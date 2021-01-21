package topology

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	topologyv1 "github.com/lyft/clutch/backend/api/topology/v1"
)

const (
	column                   = "column"
	columnIdentifer          = "column."
	metadata                 = "metadata"
	metadataIdentifer        = "metadata."
	queryDefaultLimit uint64 = 100
	maxResultLimit           = 1000
)

var filterSortFieldValidator = regexp.MustCompile(`^[a-zA-Z0-9.]*$`)

func paginatedQueryBuilder(
	filter *topologyv1.SearchRequest_Filter,
	sort *topologyv1.SearchRequest_Sort,
	pageToken string,
	limit uint64,
) (sq.SelectBuilder, uint64, error) {
	queryLimit := queryDefaultLimit
	if limit > maxResultLimit {
		return sq.SelectBuilder{}, 0, status.Error(codes.InvalidArgument, "maximum query limit is 1000")
	} else if limit > 0 {
		queryLimit = limit
	}

	// If no page is supplied default to 0
	var pageNum uint64 = 0
	var err error
	if len(pageToken) > 0 {
		pageNum, err = strconv.ParseUint(pageToken, 10, 64)
		if err != nil {
			return sq.SelectBuilder{}, 0, status.Error(codes.InvalidArgument, "unable to parse page_token")
		}
	}

	queryOffset := pageNum * limit

	query := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("id", "data", "metadata").
		From("topology_cache").
		Limit(queryLimit).
		Offset(queryOffset)

	if filter != nil {
		query, err = filterQueryBuilder(query, filter)
		if err != nil {
			return sq.SelectBuilder{}, 0, err
		}
	}

	if sort != nil {
		query, err = sortQueryBuilder(query, sort)
		if err != nil {
			return sq.SelectBuilder{}, 0, err
		}
	}

	// Blindly increment pageNum by 1 for next_page_token
	return query, pageNum + 1, nil
}

func filterQueryBuilder(query sq.SelectBuilder, f *topologyv1.SearchRequest_Filter) (sq.SelectBuilder, error) {
	if f.Search != nil && len(f.Search.Field) > 0 {
		if err := validateFilterSortField(f.Search.Field); err != nil {
			return sq.SelectBuilder{}, err
		}

		searchIdentiferExpr := sq.Expr("id")
		identifer, err := getFilterSortPrefixIdentifer(f.Search.Field)
		if err != nil {
			return sq.SelectBuilder{}, err
		}

		if identifer == column {
			searchIdentiferExpr = sq.Expr(strings.TrimPrefix(f.Search.Field, columnIdentifer))
		} else if identifer == metadata {
			mdQuery, err := convertMetadataToQuery(f.Search.Field)
			if err != nil {
				return sq.SelectBuilder{}, err
			}
			searchIdentiferExpr = sq.Expr(mdQuery)
		}

		query = query.Where(sq.Expr("? LIKE ?", searchIdentiferExpr, fmt.Sprintf("%%%s%%", f.Search.Text)))
	}

	if len(f.TypeUrl) > 0 {
		query = query.Where(sq.Eq{"resolver_type_url": f.TypeUrl})
	}

	if f.Metadata != nil {
		metadataJson, err := json.Marshal(f.Metadata)
		if err != nil {
			return sq.SelectBuilder{}, err
		}
		query = query.Where(sq.Expr("metadata @> ?::jsonb", metadataJson))
	}

	return query, nil
}

func sortQueryBuilder(query sq.SelectBuilder, s *topologyv1.SearchRequest_Sort) (sq.SelectBuilder, error) {
	if len(s.Field) > 0 {
		if err := validateFilterSortField(s.Field); err != nil {
			return sq.SelectBuilder{}, err
		}

		direction := getDirection(s.Direction)
		identifer, err := getFilterSortPrefixIdentifer(s.Field)
		if err != nil {
			return sq.SelectBuilder{}, err
		}

		if identifer == column {
			query = query.OrderByClause(fmt.Sprintf("? %s", direction), strings.TrimPrefix(s.Field, columnIdentifer))
		} else if identifer == metadata {
			mdQuery, err := convertMetadataToQuery(s.Field)
			if err != nil {
				return sq.SelectBuilder{}, err
			}
			query = query.OrderByClause(fmt.Sprintf("? %s", direction), mdQuery)
		}
	} else {
		query = query.OrderBy("ID ASC")
	}

	return query, nil
}

func validateFilterSortField(input string) error {
	if !filterSortFieldValidator.MatchString(input) {
		return fmt.Errorf("Illegal Query Operation: [%s]", input)
	}

	return nil
}

func getFilterSortPrefixIdentifer(identifer string) (string, error) {
	switch strings.Split(identifer, ".")[0] {
	case column:
		return column, nil
	case metadata:
		return metadata, nil
	default:
		return "", fmt.Errorf("Unsupported identifer: [%s]", identifer)
	}
}

func convertMetadataToQuery(input string) (string, error) {
	splitMetadata := strings.Split(strings.TrimPrefix(input, metadataIdentifer), ".")

	if splitMetadata[0] == "" {
		return "", fmt.Errorf("Incomplete metadata identifer: [%s]", metadata)
	}

	var query string
	if len(splitMetadata) == 1 {
		query = fmt.Sprintf("metadata->>'%s'", splitMetadata[0])
	} else {
		for i := range splitMetadata {
			splitMetadata[i] = fmt.Sprintf("'%s'", splitMetadata[i])
		}
		query = fmt.Sprintf("metadata->%s", strings.Join(splitMetadata, "->"))

		// Must replace the last -> with ->> to perform a text search rather than a jsonb one.
		lastIndex := strings.LastIndex(query, "->")
		query = query[:lastIndex] + strings.Replace(query[lastIndex:], "->", "->>", 1)
	}

	return query, nil
}

func getDirection(direction topologyv1.SearchRequest_Sort_Direction) string {
	switch direction {
	case topologyv1.SearchRequest_Sort_ASCENDING:
		return "ASC"
	case topologyv1.SearchRequest_Sort_DESCENDING:
		return "DESC"
	default:
		return "ASC"
	}
}
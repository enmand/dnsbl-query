package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/ent/gen/ent/dnsblquery"
	"github.com/enmand/dnsbl-query/internal/graphql/internal/gen"
)

func (r *dNSBLQueryResolver) IP(ctx context.Context, obj *ent.DNSBLQuery) (*ent.IP, error) {
	return obj.Edges.IPAddressOrErr()
}

func (r *dNSBLQueryResolver) Responses(ctx context.Context, obj *ent.DNSBLQuery, after *ent.Cursor, before *ent.Cursor, first *int, last *int) (*ent.DNSBLResponseConnection, error) {
	return obj.QueryResponses().Paginate(ctx, after, first, before, last)
}

func (r *dNSBLResponseResolver) Query(ctx context.Context, obj *ent.DNSBLResponse) (*ent.DNSBLQuery, error) {
	return obj.Edges.QueryOrErr()
}

func (r *iPResolver) ResponseCode(ctx context.Context, obj *ent.IP) (string, error) {
	rsp, err := obj.QueryQueries().
		Order(
			ent.Desc(dnsblquery.FieldCreatedAt),
		).
		QueryResponses().
		First(ctx)
	if err != nil {
		return "", fmt.Errorf("unable to find response for IP: %w", err)
	}

	return rsp.Code, nil
}

func (r *iPResolver) Queries(ctx context.Context, obj *ent.IP, after *ent.Cursor, before *ent.Cursor, first *int, last *int, orderBy *ent.DNSBLQueryOrder) (*ent.DNSBLQueryConnection, error) {
	return obj.QueryQueries().
		Paginate(ctx, after, first, before, last,
			ent.WithDNSBLQueryOrder(orderBy),
		)
}

// DNSBLQuery returns gen.DNSBLQueryResolver implementation.
func (r *Resolver) DNSBLQuery() gen.DNSBLQueryResolver { return &dNSBLQueryResolver{r} }

// DNSBLResponse returns gen.DNSBLResponseResolver implementation.
func (r *Resolver) DNSBLResponse() gen.DNSBLResponseResolver { return &dNSBLResponseResolver{r} }

// IP returns gen.IPResolver implementation.
func (r *Resolver) IP() gen.IPResolver { return &iPResolver{r} }

type dNSBLQueryResolver struct{ *Resolver }
type dNSBLResponseResolver struct{ *Resolver }
type iPResolver struct{ *Resolver }

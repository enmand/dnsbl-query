package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	"github.com/enmand/dnsbl-query/internal/graphql/internal/gen"
)

func (r *dNSBLQueryResolver) IP(ctx context.Context, obj *ent.DNSBLQuery) (*ent.IP, error) {
	return obj.Edges.IPAddressOrErr()
}

func (r *dNSBLQueryResolver) Responses(ctx context.Context, obj *ent.DNSBLQuery) ([]*ent.DNSBLResponse, error) {
	return obj.Edges.ResponsesOrErr()
}

func (r *dNSBLResponseResolver) Query(ctx context.Context, obj *ent.DNSBLResponse) (*ent.DNSBLQuery, error) {
	return obj.Edges.QueryOrErr()
}

func (r *iPResolver) ResponseCode(ctx context.Context, obj *ent.IP) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *iPResolver) Queries(ctx context.Context, obj *ent.IP) ([]*ent.DNSBLQuery, error) {
	return obj.Edges.QueriesOrErr()
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

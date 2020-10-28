package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/enmand/dnsbl-query/internal/ent/gen/ent"
	eip "github.com/enmand/dnsbl-query/internal/ent/gen/ent/ip"
	"github.com/enmand/dnsbl-query/internal/graphql/internal/gen"
	"github.com/enmand/dnsbl-query/internal/graphql/internal/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) Enque(ctx context.Context, ip []string) (*model.Operation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) GetIPDetails(ctx context.Context, ip string) (*ent.IP, error) {
	return r.client.IP.Query().Where(eip.IPAddressEQ(ip)).Only(ctx)
}

// Mutation returns gen.MutationResolver implementation.
func (r *Resolver) Mutation() gen.MutationResolver { return &mutationResolver{r} }

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

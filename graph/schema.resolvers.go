package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
)

// CreateTiger is the resolver for the createTiger field.
func (r *mutationResolver) CreateTiger(ctx context.Context, input model.CreateTigerInput) (*model.Tiger, error) {
	panic(fmt.Errorf("not implemented: CreateTiger - createTiger"))
}

// SightingOfTiger is the resolver for the sightingOfTiger field.
func (r *mutationResolver) SightingOfTiger(ctx context.Context, id string, input model.SightingOfTigerInput) (*model.Tiger, error) {
	panic(fmt.Errorf("not implemented: SightingOfTiger - sightingOfTiger"))
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Register - register"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: Login - login"))
}

// Tigers is the resolver for the tigers field.
func (r *queryResolver) Tigers(ctx context.Context) ([]*model.Tiger, error) {
	panic(fmt.Errorf("not implemented: Tigers - tigers"))
}

// Tiger is the resolver for the tiger field.
func (r *queryResolver) Tiger(ctx context.Context, id string) (*model.Tiger, error) {
	panic(fmt.Errorf("not implemented: Tiger - tiger"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
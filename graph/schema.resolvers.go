package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/vaibhavgvk08/tigerhall-kittens/graph/model"
	"github.com/vaibhavgvk08/tigerhall-kittens/services"
	"github.com/vaibhavgvk08/tigerhall-kittens/services/common"
)

// CreateTiger is the resolver for the createTiger field.
func (r *mutationResolver) CreateTiger(ctx context.Context, input model.CreateTigerInput) (*model.Tiger, error) {
	return services.FetchTigerTrackerObject().CreateTiger(input)
}

// SightingOfTiger is the resolver for the sightingOfTiger field.
func (r *mutationResolver) SightingOfTiger(ctx context.Context, id string, input model.SightingOfTigerInput) (*model.Tiger, error) {
	return services.FetchTigerTrackerObject().SightTigerLocation(id, input)
}

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input model.CreateUserInput) (*model.Response, error) {
	return services.RegisterUser(input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginUserInput) (*model.LoginResponse, error) {
	return services.LoginUser(input)
}

// ListAllTigers is the resolver for the listAllTigers field.
func (r *queryResolver) ListAllTigers(ctx context.Context, input *model.InputParams) ([]*model.Tiger, error) {
	return services.FetchTigerTrackerObject().FetchAllTigersFromDB(input)
}

// ListAllSightingsOfATiger is the resolver for the listAllSightingsOfATiger field.
func (r *queryResolver) ListAllSightingsOfATiger(ctx context.Context, id string) (*model.TigerSightings, error) {
	tigerDBObject, err := services.FetchTigerTrackerObject().FetchTigerFromDB(id)
	return common.CreateTigerSightings(tigerDBObject, err), err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

package subgraph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"

	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/products/subgraph/generated"
	"github.com/wundergraph/cosmo/demo/pkg/subgraphs/products/subgraph/model"
)

// FindConsultancyByUpc is the resolver for the findConsultancyByUpc field.
func (r *entityResolver) FindConsultancyByUpc(ctx context.Context, upc string) (*model.Consultancy, error) {
	return consultancy, nil
}

// FindCosmoByUpc is the resolver for the findCosmoByUpc field.
func (r *entityResolver) FindCosmoByUpc(ctx context.Context, upc string) (*model.Cosmo, error) {
	return cosmo, nil
}

// FindEmployeeByID is the resolver for the findEmployeeByID field.
func (r *entityResolver) FindEmployeeByID(ctx context.Context, id int) (*model.Employee, error) {
	if id < 1 {
		return nil, nil
	}
	for _, employee := range employees {
		if id == employee.ID {
			return employee, nil
		}
	}
	return nil, nil
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }

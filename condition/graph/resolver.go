package graph

import (
	"context"
	"github.com/smallretardedfish/gql-federation/condition/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ConditionStore ConditionStore
}

type ConditionFilter struct {
	IDs []int64

	Limit  *int64
	Offset *int64

	PatientID *int64
}

type ConditionStore interface {
	CreateCondition(ctx context.Context, condition model.Condition) error
	GetConditions(ctx context.Context, filter *ConditionFilter) ([]*model.Condition, error)
	GetCondition(ctx context.Context, id int64) (*model.Condition, error)
}

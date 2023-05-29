package graph

import (
	"context"
	"github.com/smallretardedfish/gql-federation/encounter/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EncounterStore EncounterStore
}

type EncounterFilter struct {
	Limit  *int64
	Offset *int64

	IDs       []int64
	PatientID *int64
}

type EncounterStore interface {
	CreateEncounter(ctx context.Context, encounter model.Encounter) error
	GetEncounters(ctx context.Context, filter *EncounterFilter) ([]*model.Encounter, error)
	GetEncounter(ctx context.Context, id int64) (*model.Encounter, error)
}

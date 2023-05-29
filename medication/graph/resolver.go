package graph

import (
	"context"
	"github.com/smallretardedfish/gql-federation/medication/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	MedicationStore MedicationStore
}

type MedicationFilter struct {
	IDs []int64

	Limit  *int64
	Offset *int64

	PatientID *string
}

type MedicationStore interface {
	CreateMedication(ctx context.Context, medication model.Medication) error
	GetMedications(ctx context.Context, filter *MedicationFilter) ([]*model.Medication, error)
	GetMedication(ctx context.Context, id int64) (*model.Medication, error)
}

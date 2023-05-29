package faker

import (
	"context"
	"github.com/smallretardedfish/gql-federation/medication/graph"
	"github.com/smallretardedfish/gql-federation/medication/graph/model"
	"math/rand"
)

type Coding struct {
	System  *string `json:"system,omitempty"`
	Code    *string `json:"code,omitempty"`
	Display *string `json:"display,omitempty"`
}

type Medication struct {
	ID           int64         `json:"id"`
	Code         *Coding       `json:"code"`
	Form         *Coding       `json:"form"`
	Manufacturer *Organization `json:"manufacturer,omitempty"`
	Patient      *Patient      `json:"patient,omitempty"`
}

type Organization struct {
	ID int64 `json:"id"`
}

func (Organization) IsEntity() {}

type Patient struct {
	ID int64 `json:"id"`
}

func createFakeMedication() model.Medication {
	medication := model.Medication{
		ID:   generateRandomID(10000),
		Code: generateRandomCoding(),
		Form: generateRandomCoding(),
	}

	manufacturerID := generateRandomID(400)
	medication.Manufacturer = &model.Organization{
		ID: manufacturerID,
	}

	patientID := generateRandomID(20000)
	medication.Patient = &model.Patient{
		ID: patientID,
	}

	return medication
}

func generateRandomID(n int) int64 {
	return int64(rand.Intn(n)) + 1
}

func generateRandomCoding() *model.Coding {
	coding := &model.Coding{
		System:  generateRandomStringOrNil([]string{"http://example.com", "http://example.org"}),
		Code:    generateRandomStringOrNil([]string{"code1", "code2", "code3"}),
		Display: generateRandomStringOrNil([]string{"Display 1", "Display 2", "Display 3"}),
	}
	return coding
}

func generateRandomStringOrNil(options []string) *string {
	if rand.Intn(2) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(options))
	value := options[randomIndex]
	return &value
}

func PopulateMedications(count int, store graph.MedicationStore) error {
	ctx := context.Background()
	for i := 0; i < count; i++ {
		patient := createFakeMedication()
		if err := store.CreateMedication(ctx, patient); err != nil {
			return err
		}
	}
	return nil
}

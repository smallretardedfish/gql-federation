package faker

import (
	"context"
	"github.com/smallretardedfish/gql-federation/condition/graph"
	"github.com/smallretardedfish/gql-federation/condition/graph/model"
	"math/rand"
)

func createFakeCondition() model.Condition {
	condition := model.Condition{
		Code:     generateRandomCoding(),
		Category: generateRandomCoding(),
		Severity: generateRandomCoding(),
		Patient:  generateRandomPatient(),
	}
	return condition
}

func generateRandomID(n int) int64 {
	return int64(rand.Intn(n) + 1)
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

func generateRandomPatient() *model.Patient {
	patientID := generateRandomID(20000)
	patient := &model.Patient{
		ID: patientID,
	}
	return patient
}

func PopulateFakeConditions(n int, storage graph.ConditionStore) error {
	ctx := context.Background()
	for i := 0; i < n; i++ {
		if err := storage.CreateCondition(ctx, createFakeCondition()); err != nil {
			return err
		}
	}
	return nil
}

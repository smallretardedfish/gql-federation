package faker

import (
	"context"
	"github.com/smallretardedfish/gql-federation/encounter/graph"
	"github.com/smallretardedfish/gql-federation/encounter/graph/model"
	"math/rand"
	"time"
)

func createFakeEncounter() model.Encounter {
	encounter := model.Encounter{
		Status:  generateRandomStatus(),
		Type:    generateRandomCoding(),
		Period:  generateRandomPeriod(),
		Patient: generateRandomPatient(),
	}
	return encounter
}

func generateRandomID(n int) int64 {
	return int64(rand.Intn(n) + 1)
}

func generateRandomStatus() string {
	statusOptions := []string{"planned", "in-progress", "completed", "on-hold", "cancelled"}
	randomIndex := rand.Intn(len(statusOptions))
	return statusOptions[randomIndex]
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

func generateRandomPeriod() *model.Period {
	startDate := generateRandomTime().String()
	endDate := generateRandomTime().String()
	return &model.Period{
		Start: &startDate,
		End:   &endDate,
	}
}

func generateRandomTime() time.Time {
	start := time.Date(2010, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Now()
	diff := end.Sub(start)
	randomDuration := time.Duration(rand.Int63n(int64(diff)))
	return start.Add(randomDuration)
}

func generateRandomPatient() *model.Patient {
	patientID := generateRandomID(20000)
	patient := &model.Patient{
		ID: patientID,
	}
	return patient
}

func PopulateFakeEncounters(n int, storage graph.EncounterStore) error {
	ctx := context.Background()
	for i := 0; i < n; i++ {
		if err := storage.CreateEncounter(ctx, createFakeEncounter()); err != nil {
			return err
		}
	}
	return nil
}

package faker

import (
	"context"
	"fmt"
	"github.com/smallretardedfish/gql-federation/patient/graph/model"
	"github.com/smallretardedfish/gql-federation/patient/storage"
	"math/rand"
	"time"
)

func createFakePatient() model.Patient {
	patient := model.Patient{
		ID: generateRandomID(),
		Name: &model.Name{
			Use:    generateRandomStringOrNil([]string{"official", "temp", "nickname"}),
			Family: generateRandomLastName(),
			Given:  generateRandomFirstNames(2),
		},
		Gender:    generateRandomGender(),
		BirthDate: generateRandomBirthDate(),
		Address:   generateRandomAddresses(3),
		Telecom:   generateRandomTelecoms(2),
	}

	return patient
}

func generateRandomID() int64 {
	return rand.Int63n(10000) + 1
}

func generateRandomStringOrNil(options []string) *string {
	if rand.Intn(2) == 0 {
		return nil
	}
	randomIndex := rand.Intn(len(options))
	value := options[randomIndex]
	return &value
}

func generateRandomLastName() *string {
	// Generate a random last name or return nil
	if rand.Intn(2) == 0 {
		return nil
	}
	lastName := "Doe" // Replace with your preferred default value
	return &lastName
}

func generateRandomFirstNames(count int) []string {
	firstNames := make([]string, count)
	for i := 0; i < count; i++ {
		firstNames[i] = fmt.Sprintf("FirstName%d", i+1) // Replace with your preferred default value
	}
	return firstNames
}

func generateRandomGender() *string {
	genders := []string{"male", "female", "other", "unknown"} // Add more genders if needed
	randomIndex := rand.Intn(len(genders))
	gender := genders[randomIndex]
	return &gender
}

func generateRandomBirthDate() *string {
	minDate := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	maxDate := time.Now()
	randomDate := randomDate(minDate, maxDate)
	birthDate := randomDate.Format("2006-01-02")
	return &birthDate
}

func randomDate(min, max time.Time) time.Time {
	delta := max.Unix() - min.Unix()
	sec := rand.Int63n(delta) + min.Unix()
	return time.Unix(sec, 0)
}

func generateRandomAddresses(count int) []*model.Address {
	addresses := make([]*model.Address, count)
	for i := 0; i < count; i++ {
		address := &model.Address{
			Use:        generateRandomStringOrNil([]string{"home", "work"}),
			Line:       []string{"123 Fake Street"},                                                       // Replace with your preferred default value
			City:       generateRandomStringOrNil([]string{"New York", "Los Angeles", "London", "Paris"}), // Add more cities if needed
			State:      generateRandomStringOrNil([]string{"NY", "CA", "TX", "FL"}),                       // Add more states if needed
			PostalCode: generateRandomStringOrNil([]string{"12345", "67890"}),                             // Add more postal codes if needed
			Country:    generateRandomStringOrNil([]string{"USA", "UK", "France"}),                        // Add more countries if needed
		}
		addresses[i] = address
	}
	return addresses
}

func generateRandomTelecoms(count int) []*model.Telecom {
	telecoms := make([]*model.Telecom, count)
	for i := 0; i < count; i++ {
		telecom := &model.Telecom{
			System: generateRandomStringOrNil([]string{"phone", "email"}),               // Add more systems if needed
			Value:  generateRandomStringOrNil([]string{"555-1234", "test@example.com"}), // Add more values if needed
		}
		telecoms[i] = telecom
	}
	return telecoms
}

func PopulatePatients(count int, store storage.PatientStore) error {
	ctx := context.Background()
	for i := 0; i < count; i++ {
		patient := createFakePatient()
		if err := store.CreatePatient(ctx, patient); err != nil {
			return err
		}
	}
	return nil
}

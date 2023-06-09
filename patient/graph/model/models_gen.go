// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Address struct {
	Use        *string  `json:"use,omitempty"`
	Line       []string `json:"line"`
	City       *string  `json:"city,omitempty"`
	State      *string  `json:"state,omitempty"`
	PostalCode *string  `json:"postalCode,omitempty"`
	Country    *string  `json:"country,omitempty"`
}

type Name struct {
	Use    *string  `json:"use,omitempty"`
	Family *string  `json:"family,omitempty"`
	Given  []string `json:"given"`
}

type Patient struct {
	ID        int64      `json:"id"`
	Name      *Name      `json:"name"`
	Gender    *string    `json:"gender,omitempty"`
	BirthDate *string    `json:"birthDate,omitempty"`
	Address   []*Address `json:"address"`
	Telecom   []*Telecom `json:"telecom"`
}

func (Patient) IsEntity() {}

type Telecom struct {
	System *string `json:"system,omitempty"`
	Value  *string `json:"value,omitempty"`
}

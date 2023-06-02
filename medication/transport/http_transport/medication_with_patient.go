package http_transport

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

type MedicationsFilter struct {
	Limit  *int `json:"limit,omitempty"`
	Offset *int `json:"offset,omitempty"`
}

type Organization struct {
	ID int64 `json:"id"`
}

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

type Telecom struct {
	System *string `json:"system,omitempty"`
	Value  *string `json:"value,omitempty"`
}

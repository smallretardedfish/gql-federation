package storage

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/smallretardedfish/gql-federation/patient/graph/model"
)

type AddressSlice []*model.Address

func (s *AddressSlice) Scan(v any) error {
	b, ok := v.([]byte)
	if b == nil && ok {
		return nil
	}

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return nil
}
func (s AddressSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type TelecomSlice []*model.Telecom

func (s *TelecomSlice) Scan(v any) error {
	b, ok := v.([]byte)
	if b == nil && ok {
		return nil
	}

	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	return nil
}

func (s TelecomSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

type NamePg model.Name

func (n *NamePg) Scan(v any) error {
	b, ok := v.([]byte)
	if b == nil && ok {
		return nil
	}

	if err := json.Unmarshal(b, &n); err != nil {
		return err
	}
	return nil
}

func (s NamePg) Value() (driver.Value, error) {
	return json.Marshal(s)
}

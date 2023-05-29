package storage

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/smallretardedfish/gql-federation/encounter/graph/model"
)

type CodingPG model.Coding

func (c *CodingPG) Scan(v any) error {
	b := v.([]byte)
	if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}

func (c CodingPG) Value() (driver.Value, error) {
	return json.Marshal(c)
}

type PeriodPG model.Period

func (c *PeriodPG) Scan(v any) error {
	b := v.([]byte)
	if err := json.Unmarshal(b, &c); err != nil {
		return err
	}
	return nil
}

func (c PeriodPG) Value() (driver.Value, error) {
	return json.Marshal(c)
}

package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/smallretardedfish/gql-federation/medication/transport/graph"
	"github.com/smallretardedfish/gql-federation/medication/transport/graph/model"
	"strings"
)

const medicationTableName = "medications"

type MedicationPostgresStore struct {
	db *sql.DB
}

func NewMedicationPostgresStore(db *sql.DB) *MedicationPostgresStore {
	return &MedicationPostgresStore{db: db}
}

func (p *MedicationPostgresStore) CreateMedication(ctx context.Context, med model.Medication) error {
	query := fmt.Sprintf(`INSERT INTO %s(code,form,manufacturer_id,patient_id)
    			VALUES($1,$2,$3,$4)`, medicationTableName)
	_, err := p.db.ExecContext(ctx, query,
		(*CodingPG)(med.Code),
		(*CodingPG)(med.Form),
		med.Manufacturer.ID,
		med.Patient.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *MedicationPostgresStore) GetMedications(ctx context.Context, filter *graph.MedicationFilter) ([]*model.Medication, error) {
	query := fmt.Sprintf("SELECT * FROM %s", medicationTableName)
	var values []any
	var placeholders []string
	var where []string

	if filter != nil {
		if len(filter.IDs) > 0 {
			for _, id := range filter.IDs {
				values = append(values, id)
				placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
			}
			where = append(where, fmt.Sprintf("id IN (%s)", strings.Join(placeholders, ",")))
		}
		if filter.PatientID != nil {
			values = append(values, filter.PatientID)
			where = append(where, fmt.Sprintf("patient_id = $%d", len(values)))
		}

		query += strings.Join(where, " AND ")

		if filter.Limit != nil {
			query += fmt.Sprintf(" LIMIT %d ", *filter.Limit)
		}
		if filter.Offset != nil {
			query += fmt.Sprintf(" OFFSET %d ", *filter.Offset)
		}
	}

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	medications, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	return medications, nil
}

func (p *MedicationPostgresStore) scan(rows *sql.Rows) ([]*model.Medication, error) {
	var res []*model.Medication
	for rows.Next() {
		med := &model.Medication{}
		var code, form CodingPG
		var orgID, patID int64
		if err := rows.Scan(
			&med.ID,
			&code,
			&form,
			&orgID,
			&patID,
		); err != nil {
			return nil, err
		}
		med.Code = (*model.Coding)(&code)
		med.Code = (*model.Coding)(&code)
		med.Manufacturer = &model.Organization{ID: orgID}
		med.Patient = &model.Patient{ID: patID}

		res = append(res, med)
	}
	return res, nil
}

func (p *MedicationPostgresStore) GetMedication(ctx context.Context, id int64) (*model.Medication, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", medicationTableName)

	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	medications, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	if len(medications) == 0 {
		return nil, errors.New("medication not found")
	}

	return medications[0], nil
}

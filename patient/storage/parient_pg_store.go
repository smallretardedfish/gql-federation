package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/smallretardedfish/gql-federation/patient/graph/model"
	"strings"
)

const patientTableName = "patients"

type PatientFilter struct {
	Limit  *int64
	Offset *int64

	IDs []int64
}

type PatientStore interface {
	CreatePatient(ctx context.Context, patient model.Patient) error
	GetPatients(ctx context.Context, filter *PatientFilter) ([]*model.Patient, error)
	GetPatient(ctx context.Context, id int64) (*model.Patient, error)
}

type PatientPostgresStore struct {
	db *sql.DB
}

func NewPatientPostgresStore(db *sql.DB) *PatientPostgresStore {
	return &PatientPostgresStore{db: db}
}

func (p *PatientPostgresStore) CreatePatient(ctx context.Context, patient model.Patient) error {
	query := fmt.Sprintf(
		`INSERT INTO %s(
				   	gender,
					birth_date,
					name,
					address,
					telecom
				)
    			VALUES($1,$2,$3,$4,$5)`, patientTableName)

	_, err := p.db.ExecContext(ctx, query,
		patient.Gender,
		patient.BirthDate,
		(*NamePg)(patient.Name),
		AddressSlice(patient.Address),
		TelecomSlice(patient.Telecom))
	if err != nil {
		return err
	}

	return nil
}

func (p *PatientPostgresStore) GetPatients(ctx context.Context, filter *PatientFilter) ([]*model.Patient, error) {
	query := fmt.Sprintf("SELECT * FROM %s", patientTableName)
	var values []any
	var placeholders []string
	if filter != nil {
		if len(filter.IDs) > 0 {
			for _, id := range filter.IDs {
				values = append(values, id)
				placeholders = append(placeholders, fmt.Sprintf("$%d", len(values)))
			}
			query += fmt.Sprintf(" WHERE id IN (%s)", strings.Join(placeholders, ","))
		}

		if filter.Limit != nil {
			query += fmt.Sprintf(" LIMIT %d ", *filter.Limit)
		}
		if filter.Offset != nil {
			query += fmt.Sprintf(" OFFSET %d ", *filter.Offset)
		}

	}
	rows, err := p.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	patients, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	return patients, nil
}

func (p *PatientPostgresStore) scan(rows *sql.Rows) ([]*model.Patient, error) {
	var res []*model.Patient
	for rows.Next() {
		pat := &model.Patient{}
		var name *NamePg
		if err := rows.Scan(
			&pat.ID,
			&pat.Gender,
			&pat.BirthDate,
			&name,
			(*AddressSlice)(&pat.Address),
			(*TelecomSlice)(&pat.Telecom),
		); err != nil {
			return nil, err
		}
		pat.Name = (*model.Name)(name)
		res = append(res, pat)
	}
	return res, nil
}

func (p *PatientPostgresStore) GetPatient(ctx context.Context, id int64) (*model.Patient, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", patientTableName)

	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	patients, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	if len(patients) == 0 {
		return nil, errors.New("patient not found")
	}

	return patients[0], nil
}

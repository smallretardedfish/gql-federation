package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/smallretardedfish/gql-federation/encounter/graph"
	"github.com/smallretardedfish/gql-federation/encounter/graph/model"
	"strings"
)

const encounterTableName = "encounters"

type EncounterPostgresStore struct {
	db *sql.DB
}

func NewEncounterPostgresStore(db *sql.DB) *EncounterPostgresStore {
	return &EncounterPostgresStore{db: db}
}

func (p *EncounterPostgresStore) CreateEncounter(ctx context.Context, encounter model.Encounter) error {
	query := fmt.Sprintf(
		`INSERT INTO %s(
               		status,
					type,
					period,
					patient_id
				)
    			VALUES($1,$2,$3,$4)`, encounterTableName)

	_, err := p.db.ExecContext(ctx,
		query,
		encounter.Status,
		(*CodingPG)(encounter.Type),
		(*PeriodPG)(encounter.Period),
		encounter.Patient.ID,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *EncounterPostgresStore) GetEncounters(ctx context.Context, filter *graph.EncounterFilter) ([]*model.Encounter, error) {
	query := fmt.Sprintf("SELECT * FROM %s", encounterTableName)
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

	rows, err := p.db.QueryContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	encounters, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	return encounters, nil
}

func (p *EncounterPostgresStore) scan(rows *sql.Rows) ([]*model.Encounter, error) {
	var res []*model.Encounter
	for rows.Next() {
		enc := &model.Encounter{}
		if err := rows.Scan(
			&enc.ID,
		); err != nil {
			return nil, err
		}
		res = append(res, enc)
	}
	return res, nil
}

func (p *EncounterPostgresStore) GetEncounter(ctx context.Context, id int64) (*model.Encounter, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", encounterTableName)

	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	encounters, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	if len(encounters) == 0 {
		return nil, errors.New("encounter not found")
	}

	return encounters[0], nil
}

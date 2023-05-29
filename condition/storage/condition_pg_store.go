package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/smallretardedfish/gql-federation/condition/graph"
	"github.com/smallretardedfish/gql-federation/condition/graph/model"
	"strings"
)

const conditionTableName = "conditions"

type ConditionPostgresStore struct {
	db *sql.DB
}

func NewConditionPostgresStore(db *sql.DB) *ConditionPostgresStore {
	return &ConditionPostgresStore{db: db}
}

func (p *ConditionPostgresStore) CreateCondition(ctx context.Context, condition model.Condition) error {
	query := fmt.Sprintf(
		`INSERT INTO %s(
					code,
					category,
					severity,
					patient_id
				)
    			VALUES($1,$2,$3,$4)`, conditionTableName)

	_, err := p.db.ExecContext(ctx, query,
		(*CodingPG)(condition.Code),
		(*CodingPG)(condition.Category),
		(*CodingPG)(condition.Severity),
		condition.Patient.ID)
	if err != nil {
		return err
	}

	return nil
}

func (p *ConditionPostgresStore) GetConditions(ctx context.Context, filter *graph.ConditionFilter) ([]*model.Condition, error) {
	query := fmt.Sprintf("SELECT * FROM %s", conditionTableName)
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

		query += fmt.Sprintf("WHERE %s", strings.Join(where, " AND "))

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

	conditions, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	return conditions, nil
}

func (p *ConditionPostgresStore) scan(rows *sql.Rows) ([]*model.Condition, error) {
	var res []*model.Condition
	for rows.Next() {
		condition := &model.Condition{}
		var patientID int64
		var code, category, severity CodingPG
		if err := rows.Scan(
			&condition.ID,
			&code,
			&category,
			&severity,
			&patientID,
		); err != nil {
			return nil, err
		}
		condition.Code = (*model.Coding)(&code)
		condition.Category = (*model.Coding)(&category)
		condition.Severity = (*model.Coding)(&severity)
		condition.Patient = &model.Patient{ID: patientID}

		res = append(res, condition)
	}
	return res, nil
}

func (p *ConditionPostgresStore) GetCondition(ctx context.Context, id int64) (*model.Condition, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", conditionTableName)

	rows, err := p.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	conditions, err := p.scan(rows)
	if err != nil {
		return nil, err
	}

	if len(conditions) == 0 {
		return nil, errors.New("condition not found")
	}

	return conditions[0], nil
}

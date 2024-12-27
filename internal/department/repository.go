package department

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IDepartmentRepository interface {
		//GetByID(ctx context.Context, id string) (*Department, error)]
		List(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*Department, int, error)
	}

	departmentRepository struct {
		db  *pgxpool.Pool
		log logger.Logger
	}
)

func NewDepartmentRepository(db *pgxpool.Pool, logger logger.Logger) IDepartmentRepository {
	return &departmentRepository{
		db:  db,
		log: logger,
	}
}

// region SQL Queries

const (
	listDeptBaseQuery = `
		SELECT 
			id, branchId, departmentId, name, code, type, specialty, 
			parentDepartmentId, status, totalBeds, availableBeds, 
			operatingRooms, weekday, weekend, timezone, holidays, metadata, 
			createdAt, updatedAt
		FROM department 
		WHERE deleted_at IS NULL`

	countDeptQuery = `
		SELECT COUNT(*) 
		FROM departments 
		WHERE deleted_at IS NULL`
)

func (r *departmentRepository) List(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*Department, int, error) {
	r.log.Debug("repository : List : begin", logger.Fields{"branchId": branchId, "filter": filter, "page": page, "limit": limit})

	// Build query with filters
	query := listDeptBaseQuery
	countQuery := countDeptQuery
	params := make([]interface{}, 0)
	paramCount := 1

	if len(filter) > 0 {
		query += " AND"
		countQuery += " AND"
		for key, value := range filter {
			if paramCount > 1 {
				query += " AND"
				countQuery += " AND"
			}
			query += fmt.Sprintf(" %s = $%d", key, paramCount)
			countQuery += fmt.Sprintf(" %s = $%d", key, paramCount)
			params = append(params, value)
			paramCount++
		}
	}

	// Add pagination
	offset := (page - 1) * limit
	query += fmt.Sprintf(" ORDER BY created_at DESC LIMIT $%d OFFSET $%d", paramCount, paramCount+1)
	params = append(params, limit, offset)

	// Get total count
	var total int
	err := r.db.QueryRow(ctx, countQuery, params[:len(params)-2]...).Scan(&total)
	if err != nil {
		return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error", err)
	}

	// Execute main query
	rows, err := r.db.Query(ctx, query, params...)
	if err != nil {
		return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error", err)
	}
	defer rows.Close()

	depts := make([]*Department, 0)
	for rows.Next() {
		dept := &Department{
			Capacity:       Capacity{},
			OperatingHours: OperatingHours{},
			Metadata:       make(map[string]interface{}),
		}
		//var createdAt, updatedAt time.Time

		err := rows.Scan(
			&dept.ID,
			&dept.BranchID,
			&dept.OrganizationID,
			&dept.Name,
			&dept.Code,
			&dept.Type,
			&dept.Specialty,
			&dept.ParentDepartmentID,
			&dept.Status,
			&dept.Capacity.TotalBeds,
			&dept.Capacity.AvailableBeds,
			&dept.Capacity.OperatingRooms,
			&dept.OperatingHours.Weekday,
			&dept.OperatingHours.Weekend,
			&dept.OperatingHours.Timezone,
			&dept.OperatingHours.Holidays,
			&dept.Metadata,
			&dept.CreatedAt,
			&dept.UpdatedAt,
		)
		if err != nil {
			return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error", err)
		}

		// dept.CreatedAt = createdAt.Format(time.RFC3339)
		// dept.UpdatedAt = updatedAt.Format(time.RFC3339)
		depts = append(depts, dept)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error", err)
	}

	r.log.Debug("repository : List : exit", nil)
	return depts, total, nil
}

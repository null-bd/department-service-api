package department

import (
	"context"
	stderr "errors"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/null-bd/department-service-api/internal/errors"
	"github.com/null-bd/logger"
)

// region Definition

type (
	IDepartmentRepository interface {
		GetByID(ctx context.Context, id string) (*Department, error)
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
			id, branch_id, organization_id, name, code, type, specialty, 
			parent_department_id, status, capacity_total_beds, capacity_available_beds, 
			capacity_operating_rooms, operating_hours_weekday, operating_hours_weekend, 
			operating_hours_timezone, operating_hours_holidays, department_head_id,
			metadata, created_at, updated_at
		FROM departments
		WHERE deleted_at IS NULL`

	getDeptByIDQuery = `
		SELECT 
			id, branch_id, organization_id, name, code, type, specialty, 
			parent_department_id, status, capacity_total_beds, capacity_available_beds, 
			capacity_operating_rooms, operating_hours_weekday, operating_hours_weekend, 
			operating_hours_timezone, operating_hours_holidays, department_head_id,
			metadata, created_at, updated_at
		FROM departments
		WHERE id = $1 AND deleted_at IS NULL`

	countDeptQuery = `
		SELECT COUNT(*) 
		FROM departments 
		WHERE deleted_at IS NULL`
)

func (r *departmentRepository) GetByID(ctx context.Context, id string) (*Department, error) {
	r.log.Debug("repository : GetByID : begin", logger.Fields{"id": id})

	dept := &Department{
		Capacity:       Capacity{},
		OperatingHours: OperatingHours{},
		Metadata:       make(map[string]interface{}),
	}

	//var createdAt, updatedAt time.Time

	err := r.db.QueryRow(ctx, getDeptByIDQuery, id).Scan(
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
		&dept.DepartmentHeadID,
		&dept.Metadata,
		&dept.CreatedAt,
		&dept.UpdatedAt,
	)

	if err != nil {
		if stderr.Is(err, pgx.ErrNoRows) {
			return nil, errors.New(errors.ErrDeptNotFound, "Department not found", err)
		}
		return nil, errors.New(errors.ErrDatabaseOperation, "database error", err)
	}

	// org.CreatedAt = createdAt.Format(time.RFC3339)
	// org.UpdatedAt = updatedAt.Format(time.RFC3339)

	r.log.Debug("repository : GetByID : exit", logger.Fields{"department": dept})
	return dept, nil
}

func (r *departmentRepository) List(ctx context.Context, branchId string, filter map[string]interface{}, page, limit int) ([]*Department, int, error) {
	r.log.Debug("repository : List : begin", logger.Fields{"branchId": branchId, "filter": filter, "page": page, "limit": limit})

	// Ensure branchId is included in the filter
	if filter == nil {
		filter = make(map[string]interface{})
	}
	filter["branch_id"] = branchId

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
		return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error1", err)
	}

	// Execute main query
	rows, err := r.db.Query(ctx, query, params...)
	if err != nil {
		return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error2", err)
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
			&dept.DepartmentHeadID,
			&dept.Metadata,
			&dept.CreatedAt,
			&dept.UpdatedAt,
		)

		if err != nil {
			return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error3", err)
		}

		// dept.CreatedAt = createdAt.Format(time.RFC3339)
		// dept.UpdatedAt = updatedAt.Format(time.RFC3339)
		depts = append(depts, dept)
	}

	if err = rows.Err(); err != nil {
		return nil, 0, errors.New(errors.ErrDatabaseOperation, "database error4", err)
	}

	r.log.Debug("repository : List : exit", nil)
	return depts, total, nil
}

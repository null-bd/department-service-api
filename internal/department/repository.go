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
		//Create(ctx context.Context, dept *Department) (*Department, error)
		//GetByID(ctx context.Context, id string) (*Department, error)
		//GetByCode(ctx context.Context, code string) (*Department, error)
		List(ctx context.Context, filter map[string]interface{}, page, limit int) ([]*Department, int, error)
		//Update(ctx context.Context, dept *Department) error
		//Delete(ctx context.Context, id string) error
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
	// createDeptQuery = `
	// 	INSERT INTO departments (
	// 		id, name, code, type, description, status,
	// 		contact_email, contact_phone, contact_address,
	// 		metadata, created_at, updated_at
	// 	) VALUES (
	// 		$1, $2, $3, $4, $5, $6,
	// 		$7, $8, $9, $10, $11, $11
	// 	) RETURNING id`

	// getDeptByIDQuery = `
	// 	SELECT
	// 		id, branchId, departmentId, name, code, type, specialty,
	// 		parentDepartmentId, status, totalBeds, availableBeds,
	// 		operatingRooms, weekday, weekend, timezone, holidays, metadata,
	// 		createdAt, updatedAt, total, page, pages, data, pagination
	// 	FROM departments
	// 	WHERE id = $1 AND deleted_at IS NULL`

	// getDeptByCodeQuery = `
	// 	SELECT
	// 		id, branchId, departmentId, name, code, type, specialty,
	// 		parentDepartmentId, status, totalBeds, availableBeds,
	// 		operatingRooms, weekday, weekend, timezone, holidays, metadata,
	// 		createdAt, updatedAt, total, page, pages, data, pagination
	// 	FROM departments
	// 	WHERE code = $1 AND deleted_at IS NULL`

	listDeptBaseQuery = `
		SELECT 
			id, branchId, departmentId, name, code, type, specialty, 
			parentDepartmentId, status, totalBeds, availableBeds, 
			operatingRooms, weekday, weekend, timezone, holidays, metadata, 
			createdAt, updatedAt, total, page, pages, data, pagination
		FROM departments 
		WHERE deleted_at IS NULL`

	// updateDeptQuery = `
	// 	UPDATE departments
	// 	SET
	// 		name = $1,
	// 		type = $2,
	// 		description = $3,
	// 		status = $4,
	// 		contact_email = $5,
	// 		contact_phone = $6,
	// 		contact_address = $7,
	// 		metadata = $8,
	// 		updated_at = $9
	// 	WHERE id = $10 AND deleted_at IS NULL`

	// softDeleteDeptQuery = `
	// 	UPDATE departments
	// 	SET
	// 		deleted_at = $1,
	// 		updated_at = $1
	// 	WHERE id = $2 AND deleted_at IS NULL`

	countDeptQuery = `
		SELECT COUNT(*) 
		FROM departments 
		WHERE deleted_at IS NULL`
)

// func (r *departmentRepository) GetByID(ctx context.Context, id string) (*Department, error) {
// 	r.log.Debug("repository : GetByID : begin", logger.Fields{"id": id})

// 	dept := &Department{
// 		Capacity:       Capacity{},
// 		OperatingHours: OperatingHours{},
// 		Metadata:       make(map[string]interface{}),
// 	}

// 	// var createdAt, updatedAt time.Time

// 	err := r.db.QueryRow(ctx, getDeptByIDQuery, id).Scan(
// 		&dept.ID,
// 		&dept.BranchID,
// 		&dept.DepartmentID,
// 		&dept.Name,
// 		&dept.Code,
// 		&dept.Type,
// 		&dept.Specialty,
// 		&dept.ParentDepartmentID,
// 		&dept.Status,
// 		&dept.Capacity.TotalBeds,
// 		&dept.Capacity.AvailableBeds,
// 		&dept.Capacity.OperatingRooms,
// 		&dept.OperatingHours.Weekday,
// 		&dept.OperatingHours.Weekend,
// 		&dept.OperatingHours.Timezone,
// 		&dept.OperatingHours.Holidays,
// 		&dept.Metadata,
// 		&dept.CreatedAt,
// 		&dept.UpdatedAt,
// 	)

// 	if err != nil {
// 		if stderr.Is(err, pgx.ErrNoRows) {
// 			return nil, errors.New(errors.ErrDeptNotFound, "department not found", err)
// 		}
// 		return nil, errors.New(errors.ErrDatabaseOperation, "database error", err)
// 	}

// 	// dept.CreatedAt = createdAt.Format(time.RFC3339)
// 	// dept.UpdatedAt = updatedAt.Format(time.RFC3339)

// 	r.log.Debug("repository : GetByID : exit", logger.Fields{"department": dept})
// 	return dept, nil
// }

// func (r *departmentRepository) GetByCode(ctx context.Context, code string) (*Department, error) {
// 	r.log.Debug("repository : GetByCode : begin", nil)

// 	dept := &Department{
// 		Capacity:       Capacity{},
// 		OperatingHours: OperatingHours{},
// 		Metadata:       make(map[string]interface{}),
// 	}

// 	//var createdAt, updatedAt time.Time

// 	err := r.db.QueryRow(ctx, getDeptByCodeQuery, code).Scan(
// 		&dept.ID,
// 		&dept.BranchID,
// 		&dept.DepartmentID,
// 		&dept.Name,
// 		&dept.Code,
// 		&dept.Type,
// 		&dept.Specialty,
// 		&dept.ParentDepartmentID,
// 		&dept.Status,
// 		&dept.Capacity.TotalBeds,
// 		&dept.Capacity.AvailableBeds,
// 		&dept.Capacity.OperatingRooms,
// 		&dept.OperatingHours.Weekday,
// 		&dept.OperatingHours.Weekend,
// 		&dept.OperatingHours.Timezone,
// 		&dept.OperatingHours.Holidays,
// 		&dept.Metadata,
// 		&dept.CreatedAt,
// 		&dept.UpdatedAt,
// 	)
// 	if err != nil {
// 		if stderr.Is(err, pgx.ErrNoRows) {
// 			return nil, nil
// 		}
// 		return nil, errors.New(errors.ErrDatabaseOperation, "database error", err)
// 	}

// 	// dept.CreatedAt = createdAt.Format(time.RFC3339)
// 	// dept.UpdatedAt = updatedAt.Format(time.RFC3339)

// 	r.log.Debug("repository : GetByCode : exit", nil)
// 	return dept, nil
// }

func (r *departmentRepository) List(ctx context.Context, filter map[string]interface{}, page, limit int) ([]*Department, int, error) {
	r.log.Debug("repository : List : begin", logger.Fields{"filter": filter, "page": page, "limit": limit})

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
			&dept.DepartmentID,
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

// func (r *departmentRepository) Update(ctx context.Context, dept *Department) error {
// 	r.log.Debug("repository : Update : begin", nil)

// 	now := time.Now().UTC()

// 	_, err := r.db.Exec(ctx, updateDeptQuery,
// 		dept.Name,
// 		dept.Type,
// 		dept.Description,
// 		dept.Status,
// 		dept.ContactInfo.Email,
// 		dept.ContactInfo.Phone,
// 		dept.ContactInfo.Address,
// 		dept.Metadata,
// 		now,
// 		dept.ID,
// 	)
// 	if err != nil {
// 		return errors.New(errors.ErrDatabaseOperation, "database error", err)
// 	}

// 	dept.UpdatedAt = now.Format(time.RFC3339)

// 	r.log.Debug("repository : Update : exit", nil)
// 	return nil
// }

// func (r *departmentRepository) Delete(ctx context.Context, id string) error {
// 	r.log.Debug("repository : Delete : begin", nil)

// 	now := time.Now().UTC()

// 	_, err := r.db.Exec(ctx, softDeleteDeptQuery, now, id)
// 	if err != nil {
// 		return errors.New(errors.ErrDatabaseOperation, "database error", err)
// 	}

// 	r.log.Debug("repository : Delete : exit", nil)
// 	return nil
// }

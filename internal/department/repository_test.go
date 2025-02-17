package department

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/null-bd/department-service-api/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	tc   *testutil.TestContainer
	repo IDepartmentRepository
}

func TestRepositorySuite(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}
	suite.Run(t, new(RepositoryTestSuite))
}

func (s *RepositoryTestSuite) SetupSuite() {
	ctx := context.Background()

	tc, err := testutil.SetupTestContainer(ctx)
	require.NoError(s.T(), err)
	s.tc = tc

	err = s.createSchema(ctx)
	require.NoError(s.T(), err)

	mockLogger := new(mockLogger)
	mockLogger.On("Debug", mock.Anything, mock.Anything).Return()
	s.repo = NewDepartmentRepository(s.tc.Pool, mockLogger)
}

func (s *RepositoryTestSuite) TearDownSuite() {
	if s.tc != nil {
		s.tc.Teardown(context.Background())
	}
}

func (s *RepositoryTestSuite) SetupTest() {
	ctx := context.Background()
	_, err := s.tc.Pool.Exec(ctx, "DELETE FROM departments")
	require.NoError(s.T(), err)
}

func (s *RepositoryTestSuite) createSchema(ctx context.Context) error {
	schema := `
        CREATE TYPE department_type AS ENUM ('medical', 'surgical', 'diagnostic', 'emergency', 'administrative', 'support');
		CREATE TYPE department_status AS ENUM ('active', 'inactive', 'maintenance', 'emergency_only');
 
		CREATE TABLE departments (
		id UUID PRIMARY KEY,
		branch_id UUID NOT NULL,
		organization_id UUID NOT NULL,
		name VARCHAR(100) NOT NULL,
		code VARCHAR(10) NOT NULL UNIQUE,
		type department_type NOT NULL,
		specialty TEXT[],
		parent_department_id UUID REFERENCES departments(id),
		status department_status NOT NULL,
		capacity_total_beds INTEGER,
		capacity_available_beds INTEGER,
		capacity_operating_rooms INTEGER,
		operating_hours_weekday VARCHAR(11) CHECK (operating_hours_weekday ~ '^([01]?[0-9]|2[0-3]):[0-5][0-9]-([01]?[0-9]|2[0-3]):[0-5][0-9]$'),
		operating_hours_weekend VARCHAR(11) CHECK (operating_hours_weekend ~ '^([01]?[0-9]|2[0-3]):[0-5][0-9]-([01]?[0-9]|2[0-3]):[0-5][0-9]$'),
		operating_hours_timezone VARCHAR(50),
		operating_hours_holidays VARCHAR(11),
		department_head_id UUID,
		min_staff_required INTEGER,
		metadata JSONB DEFAULT '{}'::JSONB,
		created_at TIMESTAMP WITH TIME ZONE NOT NULL,
		updated_at TIMESTAMP WITH TIME ZONE NOT NULL,
		deleted_at TIMESTAMP WITH TIME ZONE,
		UNIQUE(branch_id, code),
		UNIQUE(branch_id, name)
		);
		
		CREATE TYPE staff_role AS ENUM ('doctor', 'nurse', 'technician', 'administrative', 'support');
		CREATE TYPE schedule_type AS ENUM ('full_time', 'part_time', 'on_call', 'rotating');
		
		CREATE TABLE staff_assignments (
			id UUID PRIMARY KEY,
			department_id UUID NOT NULL REFERENCES departments(id),
			staff_id UUID NOT NULL,
			role staff_role NOT NULL,
			schedule_type schedule_type NOT NULL,
			primary_department BOOLEAN DEFAULT false,
			start_date DATE NOT NULL,
			end_date DATE,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL,
			UNIQUE(department_id, staff_id)
		);
		
		CREATE INDEX idx_departments_branch_id ON departments(branch_id);
		CREATE INDEX idx_departments_organization_id ON departments(organization_id);
		CREATE INDEX idx_staff_assignments_staff_id ON staff_assignments(staff_id);
    `

	_, err := s.tc.Pool.Exec(ctx, schema)
	return err
}

// func stringPtr(s string) *string {
// 	return &s
// }

func (s *RepositoryTestSuite) TestGetByID() {
	//arrange

	ctx := context.Background()
	dept := &Department{
		ID:             uuid.New().String(),
		BranchID:       uuid.New().String(),
		OrganizationID: uuid.New().String(),
		Name:           "Test Department",
		Code:           "TEST001",
		Type:           "medical",
		// ParentDepartmentID: stringPtr(uuid.New().String()),
		Status: "active",
		Capacity: Capacity{
			TotalBeds:      0,
			AvailableBeds:  0,
			OperatingRooms: 0,
		},
		OperatingHours: OperatingHours{
			Weekday:  "09:00-17:00",
			Weekend:  "10:00-14:00",
			Timezone: "UTC+0",
			Holidays: "09:00-13:00",
		},
		// DepartmentHeadID: stringPtr(uuid.New().String()),
	}

	_, err := s.repo.Create(ctx, dept)
	require.NoError(s.T(), err)

	// Act
	result, err := s.repo.GetByID(ctx, dept.ID)

	// Assert
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), dept.ID, result.ID)
	assert.Equal(s.T(), dept.Name, result.Name)
	assert.Equal(s.T(), dept.Code, result.Code)
}

func (s *RepositoryTestSuite) TestList() {
	// Arrange
	ctx := context.Background()

	// Create test data
	depts := []*Department{
		{
			ID:             uuid.New().String(),
			BranchID:       "8f822ed4-b3c8-4539-99d5-28c4f16be1ce",
			OrganizationID: uuid.New().String(),
			Name:           "Test Department1",
			Code:           "TEST001",
			Type:           "medical",
			Status:         "active",
			Capacity: Capacity{
				TotalBeds:      0,
				AvailableBeds:  0,
				OperatingRooms: 0,
			},
			OperatingHours: OperatingHours{
				Weekday:  "09:00-17:00",
				Weekend:  "10:00-14:00",
				Timezone: "UTC+0",
				Holidays: "09:00-13:00",
			},
		},
		{
			ID:             uuid.New().String(),
			BranchID:       "8f822ed4-b3c8-4539-99d5-28c4f16be1ce",
			OrganizationID: uuid.New().String(),
			Name:           "Test Department2",
			Code:           "TEST002",
			Type:           "medical",
			Status:         "active",
			Capacity: Capacity{
				TotalBeds:      0,
				AvailableBeds:  0,
				OperatingRooms: 0,
			},
			OperatingHours: OperatingHours{
				Weekday:  "09:00-17:00",
				Weekend:  "10:00-14:00",
				Timezone: "UTC+0",
				Holidays: "09:00-13:00",
			},
		},
		{
			ID:             uuid.New().String(),
			BranchID:       "8f822ed4-b3c8-4539-99d5-28c4f16be1ce",
			OrganizationID: uuid.New().String(),
			Name:           "Test Department3",
			Code:           "TEST003",
			Type:           "medical",
			Status:         "active",
			Capacity: Capacity{
				TotalBeds:      0,
				AvailableBeds:  0,
				OperatingRooms: 0,
			},
			OperatingHours: OperatingHours{
				Weekday:  "09:00-17:00",
				Weekend:  "10:00-14:00",
				Timezone: "UTC+0",
				Holidays: "09:00-13:00",
			},
		},
	}

	for _, dept := range depts {
		_, err := s.repo.Create(ctx, dept)
		require.NoError(s.T(), err)
	}

	// Act
	filter := map[string]interface{}{
		"status": "active",
		"type":   "medical",
	}

	resultsPage1, total, err := s.repo.List(ctx, "8f822ed4-b3c8-4539-99d5-28c4f16be1ce", filter, 1, 2)
	require.NoError(s.T(), err)
	resultsPage2, _, err := s.repo.List(ctx, "8f822ed4-b3c8-4539-99d5-28c4f16be1ce", filter, 2, 1)

	// Assert
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(depts), total)
	assert.Equal(s.T(), total, 3)

	assert.Equal(s.T(), 2, len(resultsPage1))
	for _, result := range resultsPage1 {
		assert.Equal(s.T(), "active", result.Status)
		assert.Equal(s.T(), "medical", result.Type)
	}

	assert.Equal(s.T(), 1, len(resultsPage2))
	for _, result := range resultsPage2 {
		assert.Equal(s.T(), "active", result.Status)
		assert.Equal(s.T(), "medical", result.Type)
	}
}

func (s *RepositoryTestSuite) TestCreate() {

	//Arrange
	ctx := context.Background()
	now := time.Now().UTC()
	dept := &Department{
		ID:             uuid.New().String(),
		BranchID:       uuid.New().String(),
		OrganizationID: uuid.New().String(),
		Name:           "Test Department",
		Code:           "TEST001",
		Type:           "medical",
		// ParentDepartmentID: stringPtr(uuid.New().String()),
		Status: "active",
		Capacity: Capacity{
			TotalBeds:      0,
			AvailableBeds:  0,
			OperatingRooms: 0,
		},
		OperatingHours: OperatingHours{
			Weekday:  "09:00-17:00",
			Weekend:  "10:00-14:00",
			Timezone: "UTC+0",
			Holidays: "09:00-13:00",
		},
		// DepartmentHeadID: stringPtr(uuid.New().String()),
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}

	// Act
	result, err := s.repo.Create(ctx, dept)

	// Assert
	assert.NoError(s.T(), err)
	assert.NotEmpty(s.T(), result.CreatedAt)
	assert.NotEmpty(s.T(), result.UpdatedAt)

	// Verify in database
	var count int
	err = s.tc.Pool.QueryRow(ctx, "SELECT COUNT(*) FROM departments WHERE id = $1", dept.ID).Scan(&count)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, count)
	_, err = s.repo.Create(ctx, dept)
	assert.Error(s.T(), err)

}

func (s *RepositoryTestSuite) TestUpdate() {
	// Arrange
	ctx := context.Background()
	dept := &Department{
		ID:     uuid.New().String(),
		Name:   "Test Department",
		Code:   "TEST005",
		Type:   "medical",
		Status: "active",
		Capacity: Capacity{
			TotalBeds:      0,
			AvailableBeds:  0,
			OperatingRooms: 0,
		},
	}

	_, err := s.repo.Create(ctx, dept)
	require.NoError(s.T(), err)

	// Update fields
	dept.Name = "Updated Department"
	dept.Status = "inactive"

	// Act
	err = s.repo.Update(ctx, dept)

	// Assert
	assert.NoError(s.T(), err)

	// Verify in database
	updated, err := s.repo.GetByID(ctx, dept.ID)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), "Updated Department", updated.Name)
	assert.Equal(s.T(), "inactive", updated.Status)
}

package department

import (
	"context"

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
	_, err := s.tc.Pool.Exec(ctx, "TRUNCATE departments")
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

func (s *RepositoryTestSuite) TestGetbyID() {
	//arrange

	ctx := context.Background()
	dept := &Department{
		ID:             "test ID",
		BranchID:       "test BranchID",
		OrganizationID: "test OrganizationID",
		Name:           "Test Department",
		Code:           "TESTDEPT001",
		Type:           "hospital",
		Status:         "active",
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

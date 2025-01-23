package department

type Capacity struct {
	TotalBeds      int
	AvailableBeds  int
	OperatingRooms int
}

type OperatingHours struct {
	Weekday  string
	Weekend  string
	Timezone string
	Holidays string
}

// type Staffing struct {
// 	DepartmentHead   string
// 	MinStaffRequired int
// }

type Department struct {
	ID                 string
	BranchID           string
	OrganizationID     string
	Name               string
	Code               string
	Type               string
	Specialty          []string
	ParentDepartmentID *string
	Status             string
	Capacity           Capacity
	OperatingHours     OperatingHours
	DepartmentHeadID   *string
	Metadata           map[string]interface{}
	CreatedAt          string
	UpdatedAt          string
	DeletedAt          string
}

type Pagination struct {
	Total int
	Page  int
	Pages int
}

type DepartmentListResponse struct {
	Data       []Department
	Pagination Pagination
}

type ContactInfo struct {
	Email   string
	Phone   string
	Address string
}

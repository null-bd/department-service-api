package rest

type CapacityDTO struct {
	TotalBeds      int `json:"totalBeds"`
	AvailableBeds  int `json:"availableBeds"`
	OperatingRooms int `json:"operatingRooms"`
}

type OperatingHoursDTO struct {
	Weekday  string `json:"weekday"`
	Weekend  string `json:"weekend"`
	Timezone string `json:"timezone"`
	Holidays string `json:"holidays"`
}

type CreateDepartmentRequest struct {
	Name               string                 `json:"name" binding:"required"`
	Code               string                 `json:"code" binding:"required,uppercase"`
	Type               string                 `json:"type" binding:"required,oneof=medical surgical diagnostic emergency administrative support"`
	Specialty          []string               `json:"specialty"`
	ParentDepartmentID string                 `json:"parentDepartmentId"`
	Capacity           CapacityDTO            `json:"capacity"`
	OperatingHours     OperatingHoursDTO      `json:"operatingHours"`
	DepartmentHeadID   string                 `json:"departmentHeadId"`
	Metadata           map[string]interface{} `json:"metadata"`
}

type ListDepartmentResponse struct {
	ID                 string                 `json:"id" binding:"required"`
	BranchID           string                 `json:"branchId"`
	OrganizationID     string                 `json:"organizationId"`
	Name               string                 `json:"name"`
	Code               string                 `json:"code"`
	Type               string                 `json:"type"`
	Specialty          []string               `json:"specialty"`
	ParentDepartmentID string                 `json:"parentDepartmentId"`
	Status             string                 `json:"status"`
	Capacity           CapacityDTO            `json:"capacity"`
	OperatingHours     OperatingHoursDTO      `json:"operatingHours"`
	DepartmentHeadID   string                 `json:"departmentheadID"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt          string                 `json:"createdAt"`
	UpdatedAt          string                 `json:"updatedAt"`
	DeletedAt          string                 `json:"deletedAt"`
}

type Pagination struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

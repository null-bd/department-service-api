package rest

import "time"

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

type ContactInfoDTO struct {
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
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
	CreatedAt          time.Time              `json:"createdAt"`
	UpdatedAt          time.Time              `json:"updatedAt"`
	DeletedAt          time.Time              `json:"deletedAt"`
}

type Pagination struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

package rest

import "time"

// type CreateDepartmentRequest struct {
// 	Name        string                 `json:"name" binding:"required"`
// 	Code        string                 `json:"code" binding:"required,uppercase"`
// 	Type        string                 `json:"type" binding:"required,oneof=hospital clinic laboratory"`
// 	Description string                 `json:"description"`
// 	ContactInfo ContactInfoDTO         `json:"contactInfo" binding:"required"`
// 	Metadata    map[string]interface{} `json:"metadata"`
// }

// type UpdateDepartmentRequest struct {
// 	Name        string                 `json:"name" binding:"required"`
// 	Type        string                 `json:"type" binding:"required,oneof=hospital clinic laboratory"`
// 	Description string                 `json:"description"`
// 	Status      string                 `json:"status" binding:"required,oneof=active inactive suspended"`
// 	ContactInfo ContactInfoDTO         `json:"contactInfo" binding:"required"`
// 	Metadata    map[string]interface{} `json:"metadata"`
// }

type Capacity struct {
	TotalBeds      int `json:"totalBeds"`
	AvailableBeds  int `json:"availableBeds"`
	OperatingRooms int `json:"operatingRooms"`
}

type OperatingHours struct {
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

type DepartmentResponse struct {
	ID                 string         `json:"id"`
	BranchID           string         `json:"branchId"`
	DepartmentID       string         `json:"departmentId"`
	Name               string         `json:"name"`
	Code               string         `json:"code"`
	Type               string         `json:"type"`
	Specialty          []string       `json:"specialty"`
	ParentDepartmentID string         `json:"parentDepartmentId"`
	Status             string         `json:"status"`
	Capacity           Capacity       `json:"capacity"`
	OperatingHours     OperatingHours `json:"operatingHours"`
	//Staffing           Staffing               `json:"staffing"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt time.Time              `json:"createdAt"`
	UpdatedAt time.Time              `json:"updatedAt"`
}

type Pagination struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Pages int `json:"pages"`
}

// type CreateBranchRequest struct {
// 	Name           string                 `json:"name" binding:"required"`
// 	Code           string                 `json:"code" binding:"required,uppercase"`
// 	Type           string                 `json:"type" binding:"required,oneof=main satellite specialized"`
// 	Description    string                 `json:"description"`
// 	ContactInfo    ContactInfoDTO         `json:"contactInfo" binding:"required"`
// 	OperatingHours OperatingHoursDTO      `json:"operatingHours"`
// 	Capacity       CapacityDTO            `json:"capacity"`
// 	Metadata       map[string]interface{} `json:"metadata,omitempty"`
// }

// type UpdateBranchRequest struct {
// 	Name           string                 `json:"name" binding:"required"`
// 	Type           string                 `json:"type" binding:"required,oneof=main satellite specialized"`
// 	Description    string                 `json:"description"`
// 	Status         string                 `json:"status" binding:"required,oneof=active inactive suspended"`
// 	ContactInfo    ContactInfoDTO         `json:"contactInfo" binding:"required"`
// 	OperatingHours OperatingHoursDTO      `json:"operatingHours"`
// 	Capacity       CapacityDTO            `json:"capacity"`
// 	Metadata       map[string]interface{} `json:"metadata,omitempty"`
// }

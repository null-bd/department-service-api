package rest

import (
	"github.com/null-bd/department-service-api/internal/department"
)

func ToDepartmentResponse(dept *department.Department) *ListDepartmentResponse {
	return &ListDepartmentResponse{
		ID:                 dept.ID,
		BranchID:           dept.BranchID,
		OrganizationID:     dept.OrganizationID,
		Name:               dept.Name,
		Code:               dept.Code,
		Type:               dept.Type,
		Specialty:          dept.Specialty,
		ParentDepartmentID: dept.ParentDepartmentID,
		Status:             dept.Status,
		Capacity: CapacityDTO{
			TotalBeds:      dept.Capacity.TotalBeds,
			AvailableBeds:  dept.Capacity.AvailableBeds,
			OperatingRooms: dept.Capacity.OperatingRooms,
		},
		OperatingHours: OperatingHoursDTO{
			Weekday:  dept.OperatingHours.Weekday,
			Weekend:  dept.OperatingHours.Weekend,
			Timezone: dept.OperatingHours.Timezone,
			Holidays: dept.OperatingHours.Holidays,
		},
		// Staffing: StaffingDTO{
		// 	DepartmentHead:   dept.Staffing.DepartmentHead,
		// 	MinStaffRequired: dept.Staffing.MinStaffRequired,
		// },
		Metadata:  dept.Metadata,
		CreatedAt: dept.CreatedAt,
		UpdatedAt: dept.UpdatedAt,
	}
}

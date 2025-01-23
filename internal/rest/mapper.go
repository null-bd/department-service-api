package rest

import (
	"github.com/null-bd/department-service-api/internal/department"
)

func ToDepartment(req *CreateDepartmentRequest) *department.Department {
	return &department.Department{
		Name:               req.Name,
		Code:               req.Code,
		Type:               req.Type,
		Specialty:          req.Specialty,
		ParentDepartmentID: req.ParentDepartmentID,
		Capacity: department.Capacity{
			TotalBeds:      req.Capacity.TotalBeds,
			AvailableBeds:  req.Capacity.AvailableBeds,
			OperatingRooms: req.Capacity.OperatingRooms,
		},
		OperatingHours: department.OperatingHours{
			Weekday:  req.OperatingHours.Weekday,
			Weekend:  req.OperatingHours.Weekend,
			Timezone: req.OperatingHours.Timezone,
			Holidays: req.OperatingHours.Holidays,
		},
		DepartmentHeadID: req.DepartmentHeadID,
		Metadata:         req.Metadata,
	}
}

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
		DepartmentHeadID: dept.DepartmentHeadID,
		// Staffing: StaffingDTO{
		// 	DepartmentHead:   dept.Staffing.DepartmentHead,
		// 	MinStaffRequired: dept.Staffing.MinStaffRequired,
		// },
		Metadata:  dept.Metadata,
		CreatedAt: dept.CreatedAt,
		UpdatedAt: dept.UpdatedAt,
		DeletedAt: dept.DeletedAt,
	}
}

package rest

import (
	"github.com/null-bd/department-service-api/internal/department"
)

// func ToDepartment(req *CreateDepartmentRequest) *department.Department {
// 	return &department.Department{
// 		Name:        req.Name,
// 		Code:        req.Code,
// 		Type:        req.Type,
// 		Description: req.Description,
// 		ContactInfo: department.ContactInfo{
// 			Email:   req.ContactInfo.Email,
// 			Phone:   req.ContactInfo.Phone,
// 			Address: req.ContactInfo.Address,
// 		},
// 		Metadata: req.Metadata,
// 	}
// }

func ToDepartmentResponse(dept *department.Department) *DepartmentResponse {
	return &DepartmentResponse{
		ID:          dept.ID,
		Name:        dept.Name,
		Code:        dept.Code,
		Type:        dept.Type,
		Description: dept.Description,
		Status:      dept.Status,
		ContactInfo: ContactInfoDTO{
			Email:   dept.ContactInfo.Email,
			Phone:   dept.ContactInfo.Phone,
			Address: dept.ContactInfo.Address,
		},
		Metadata:  dept.Metadata,
		CreatedAt: dept.CreatedAt,
		UpdatedAt: dept.UpdatedAt,
	}
}

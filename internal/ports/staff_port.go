package ports

import (
	"github.com/BreakDown-CS/api-ecommerce/internal/domain"
)

type StaffServices interface {
	CreateStaff(domain.RequestStaffInsert) (int, error)
	SearchStaffList(domain.RequestStaffList) ([]int, error)
}

type StaffRepositories interface {
	ChackStaffInsert(string, *string) (bool, error)
	GetStaffList() ([]int, error)
}

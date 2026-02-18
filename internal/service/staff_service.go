package service

import (
	"fmt"

	"github.com/BreakDown-CS/api-ecommerce/internal/domain"
	"github.com/BreakDown-CS/api-ecommerce/internal/ports"
)

type StaffService struct {
	Repositories ports.StaffRepositories
}

func NewStaffService(repositories ports.StaffRepositories) ports.StaffServices {
	return &StaffService{Repositories: repositories}
}

func (s *StaffService) CreateStaff(requestDomain domain.RequestStaffInsert) (int, error) {
	dataChackStaff, err := s.Repositories.ChackStaffInsert(requestDomain.Username, requestDomain.EMCode)
	if err != nil {
		return 0, err
	}

	fmt.Println("AAA : ", dataChackStaff)
	return 0, nil
}

func (s *StaffService) SearchStaffList(requestDomain domain.RequestStaffList) ([]int, error) {

	dataStaffList, err := s.Repositories.GetStaffList()
	if err != nil {
		return nil, err
	}

	fmt.Println("SSS : ", dataStaffList)

	return nil, nil
}

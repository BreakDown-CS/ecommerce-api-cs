package service

import (
	"fmt"

	"github.com/BreakDown-CS/api-ecommerce/internal/ports"
)

type StaffService struct {
	Repositories ports.StaffRepositories
}

func NewStaffService(repositories ports.StaffRepositories) ports.StaffServices {
	return &StaffService{Repositories: repositories}
}

func (s *StaffService) SearchStaffList() ([]int, error) {

	dataStaffList, err := s.Repositories.GetStaffList()
	if err != nil {
		return nil, err
	}

	fmt.Println("SSS : ", dataStaffList)

	return nil, nil
}

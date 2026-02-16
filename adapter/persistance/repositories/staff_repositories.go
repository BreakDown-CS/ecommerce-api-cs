package repositories

import (
	"github.com/BreakDown-CS/api-ecommerce/internal/ports"
	"gorm.io/gorm"
)

type StaffRepositories struct {
	DBPostgres *gorm.DB
}

func NewStaffRepositories(DBPostgres *gorm.DB) ports.StaffRepositories {
	return &StaffRepositories{
		DBPostgres: DBPostgres,
	}
}

func (r *StaffRepositories) GetStaffList() ([]int, error) {
	return nil, nil
}

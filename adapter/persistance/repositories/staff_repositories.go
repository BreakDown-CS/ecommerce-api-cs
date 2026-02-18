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

func (r *StaffRepositories) ChackStaffInsert(username string, em_code *string) (bool, error) {
	var chackUsername bool
	var chackEmCode bool

	qureyChackUsername := `SELECT EXISTS (SELECT 1 FROM staff WHERE username = $1)`

	errUsername := r.DBPostgres.
		Raw(qureyChackUsername, username).
		Scan(&chackUsername).Error

	if errUsername != nil {
		return false, errUsername
	}

	qureyChackEmCode := `SELECT EXISTS (SELECT 1 FROM staff WHERE em_code = $1)`

	errEmcode := r.DBPostgres.
		Raw(qureyChackEmCode, em_code).
		Scan(&chackEmCode).Error

	if errEmcode != nil {
		return false, errEmcode
	}

	chackStaff := chackUsername && chackEmCode

	return chackStaff, nil
}

func (r *StaffRepositories) GetStaffList() ([]int, error) {
	return nil, nil
}

package ports

type StaffServices interface {
	SearchStaffList() ([]int, error)
}

type StaffRepositories interface {
	GetStaffList() ([]int, error)
}

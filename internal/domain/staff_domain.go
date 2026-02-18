package domain

import "time"

type RequestStaffInsert struct {
	Username           string     `json:"username"`
	Password           string     `json:"password"`
	StaffShop          int        `json:"staff_shop"`
	StaffType          int        `json:"staff_type"`
	StaffDepartment    int        `json:"staff_department"`
	StaffMail          string     `json:"staff_mail"`
	StartDate          time.Time  `json:"start_date"`
	EMCode             *string    `json:"emcode"`
	EndDate            *time.Time `json:"end_date"`
	StaffTimeWork      *int       `json:"staff_time_work"`
	StaffStatus        *string    `json:"staff_status"`
	StaffName          *string    `json:"staff_name"`
	StaffNickname      *string    `json:"staff_nickname"`
	StaffPhone         *string    `json:"staff_phone"`
	StaffBirthday      *string    `json:"staff_birthday"`
	Barcode            *string    `json:"barcode"`
	StaffGroupPage     *int       `json:"staff_group_page"`
	StaffLine          *string    `json:"staff_line"`
	StaffDevice        *string    `json:"staff_device"`
	StaffLock          *string    `json:"staff_lock"`
	Child              *int       `json:"children"`
	StaffBankNumber    *string    `json:"staff_bank_nmuber"`
	StaffBankId        *int       `json:"staff_bank_id"`
	StaffAddress       *string    `json:"staff_address"`
	StaffRoad          *string    `json:"staff_road"`
	StaffDistrict1     *string    `json:"staff_district_1"`
	StaffDistrict2     *string    `json:"staff_district_2"`
	StaffProvince      *string    `json:"staff_province"`
	StaffZip           *string    `json:"staff_zip"`
	StaffAddressDetail *string    `json:"staff_address_detail"`
}

type RequestStaffList struct {
	Emcode string `json:"emcode"`
}

package handler

import (
	"github.com/BreakDown-CS/api-ecommerce/adapter/persistance/models"
	"github.com/BreakDown-CS/api-ecommerce/internal/domain"
	"github.com/BreakDown-CS/api-ecommerce/internal/ports"
	"github.com/gofiber/fiber/v2"
)

type StaffHandler struct {
	service ports.StaffServices
}

func NewStaffHandler(service ports.StaffServices) *StaffHandler {
	return &StaffHandler{service: service}
}

func (h *StaffHandler) CreateStaff(c *fiber.Ctx) error {

	requestModel := models.RequestStaffCreate{}

	if err := c.BodyParser(&requestModel); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "code": fiber.StatusBadRequest, "message": "กรุณากรอกข้อมูลให้ครบถ้วน"})
	}

	requestDomain := domain.RequestStaffInsert{
		Username:           requestModel.Username,
		Password:           requestModel.Password,
		StaffShop:          requestModel.StaffShop,
		StaffType:          requestModel.StaffType,
		StaffDepartment:    requestModel.StaffDepartment,
		StaffMail:          requestModel.StaffMail,
		StartDate:          requestModel.StartDate,
		EMCode:             requestModel.EMCode,
		EndDate:            requestModel.EndDate,
		StaffTimeWork:      requestModel.StaffTimeWork,
		StaffStatus:        requestModel.StaffStatus,
		StaffName:          requestModel.StaffName,
		StaffNickname:      requestModel.StaffNickname,
		StaffPhone:         requestModel.StaffPhone,
		StaffBirthday:      requestModel.StaffBirthday,
		Barcode:            requestModel.Barcode,
		StaffGroupPage:     requestModel.StaffGroupPage,
		StaffLine:          requestModel.StaffLine,
		StaffDevice:        requestModel.StaffDevice,
		StaffLock:          requestModel.StaffLock,
		Child:              requestModel.Child,
		StaffBankNumber:    requestModel.StaffBankNumber,
		StaffBankId:        requestModel.StaffBankId,
		StaffAddress:       requestModel.StaffAddress,
		StaffRoad:          requestModel.StaffRoad,
		StaffDistrict1:     requestModel.StaffDistrict1,
		StaffDistrict2:     requestModel.StaffDistrict2,
		StaffProvince:      requestModel.StaffProvince,
		StaffZip:           requestModel.StaffZip,
		StaffAddressDetail: requestModel.StaffAddressDetail,
	}

	result, err := h.service.CreateStaff(requestDomain)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

func (h *StaffHandler) GetStaffList(c *fiber.Ctx) error {

	requestModel := models.ReqStaffList{}

	if err := c.BodyParser(&requestModel); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": false, "code": fiber.StatusBadRequest, "message": "กรุณากรอกข้อมูลให้ครบถ้วน"})
	}

	requestDomain := domain.RequestStaffList{
		Emcode: requestModel.Emcode,
	}

	result, err := h.service.SearchStaffList(requestDomain)
	if err != nil {
		return err
	}
	return c.JSON(result)
}

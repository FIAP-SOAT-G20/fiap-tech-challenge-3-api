package handler

import (
	valueobject "github.com/FIAP-SOAT-G20/FIAP-TechChallenge-Fase2/internal/core/domain/value_object"
	"github.com/go-playground/validator/v10"
)

func OrderStatusValidator(fl validator.FieldLevel) bool {
	status := fl.Field().String()
	return valueobject.IsValidOrderStatus(status)
}

func StaffRoleValidator(fl validator.FieldLevel) bool {
	role := fl.Field().String()
	return valueobject.IsValidStaffRole(role)
}

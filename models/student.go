package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name string `json:"name" validate:"nonzero"`
	CPF  string `json:"cpf" validate:"len=11, regexp=^\\d{11}$"`
	RG   string `json:"rg" validate:"len=9"`
}

func ValidateStudent(student *Student) error {
	return validator.Validate(student)
}

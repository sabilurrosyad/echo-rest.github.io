package controllers

import (
	"net/http"

	validator "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type Customer struct {
	Nama   string `validate:"required"`
	Email  string `validate:"required,email"`
	Alamat string `validate:"required"`
	Umur   int    `validate:"gte=17,lte=35"`
}

func TestVariableValidation(c echo.Context) error {
	v := validator.New()

	email := "bamssss@gmail.com"

	err := v.Var(email, "required,email")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "emailnya paok salah format or not valid",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

func TestStructValidation(c echo.Context) error {
	v := validator.New()

	cust := Customer{
		Nama:   "bams",
		Email:  "bams@gmail.com",
		Alamat: "Bangka",
		Umur:   25,
	}

	err := v.Struct(cust)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "success"})
}

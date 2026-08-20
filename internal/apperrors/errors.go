package apperors

import "net/http"

type AppError struct {
	Code    int
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}

var (
	ErrCustomerNotFound = &AppError{
		Code:    http.StatusNotFound,
		Message: "customer not found",
	}
	ErrIncomeNotFound = &AppError{
		Code:    http.StatusNotFound,
		Message: "income not found",
	}
	ErrIncomeNotCreated = &AppError{
		Code:    http.StatusConflict,
		Message: "income could not be created",
	}
	ErrInvalidProductType = &AppError{
		Code:    http.StatusBadRequest,
		Message: "invalid product type",
	}
	ErrInvalidUnitPrice = &AppError{
		Code:    http.StatusBadRequest,
		Message: "invalid unit price",
	}
	ErrInvalidPaymentStatus = &AppError{
		Code:    http.StatusBadRequest,
		Message: "invalid payment status",
	}
	ErrInternalServer = &AppError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	}
)

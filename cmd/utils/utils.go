package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

func ParseJSON(r *http.Request, payload any) error {

	if r.ContentLength == 0 {
		return fmt.Errorf("missing request body(Content-Length is 0)")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func ValidateJSON(payload any) error {
	var validate = validator.New()
	validate.RegisterValidation("date_format", validateDateFormat)
	validate.RegisterValidation("time_format", validateTimeFormat)
	validate.RegisterValidation("float_format", floatValidator)
	// validate struct fields
	if err := validate.Struct(payload); err != nil {
		return err
	}

	return nil
}

func ConvertDateTimeTotal(purchaseDate string, purchaseTime string, total string) (time.Time, float64) {
	const combinedFormat = "2006-01-02 15:04"

	combined := fmt.Sprintf("%s %s", purchaseDate, purchaseTime)

	combinedTime, _ := time.Parse(combinedFormat, combined)

	floatTotal, _ := strconv.ParseFloat(total, 64)

	return combinedTime, floatTotal

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func validateDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	const dateFormat = "2006-01-02"

	_, err := time.Parse(dateFormat, dateStr)
	return err == nil
}

func validateTimeFormat(fl validator.FieldLevel) bool {

	timeStr := fl.Field().String()

	const timeFormat = "15:04"

	_, err := time.Parse(timeFormat, timeStr)
	return err == nil
}

func floatValidator(fl validator.FieldLevel) bool {
	str := fl.Field().String()
	_, err := strconv.ParseFloat(str, 64)
	return err == nil
}

package validation

import (
	"fmt"
	"strings"

	"github.com/asaskevich/govalidator"
)

// MsgValidationFailed - API response message in case of validation failure
const MsgValidationFailed = "Validation failure"

// messages contains error messages grouped by validation rule
var messages = map[string]string{
	"int":    "Value should be an integer",
	"string": "Value should be a string",
}

// Constraint describes validation constraint object ("value-to-validate" + validation rules slice)
type Constraint struct {
	Value interface{}
	Rules []string
}

// Errors describes app validation errors structure
// {
//     "fieldName1" => []string{error1, error2, ..., errorN}
//     "fieldName2" => []string{error1, error2, ..., errorN}
//     ...
//     "fieldNameN" => []string{error1, error2, ..., errorN}
// }
type Errors map[string]string

// Validatable use this interface to provide additional validation for your struct
type Validatable interface {
	Validate() (bool, Errors)
}

// Validate check all constraints inside validationMap & return Errors in case then smth invalid
func Validate(validationMap map[string]Constraint) Errors {
	var currentError string
	var validationErrors Errors = make(map[string]string, 0)

	for field, constraint := range validationMap {
		for _, rule := range constraint.Rules {

			if vf, ok := govalidator.TagMap[rule]; ok {
				if !vf(constraint.Value.(string)) {
					currentError = getErrorMessage(rule)
				}
			} else {
				currentError = fmt.Sprintf("Unable to find `%s` validation rule", rule)
			}

			if currentError != "" {
				validationErrors[field] = currentError
				currentError = ""
			}

		}
	}

	return validationErrors
}

// ValidateStruct improves usual govalidator.ValidateStruct() method with custom errors preparing
func ValidateStruct(value interface{}) Errors {
	errors := make(Errors, 0)

	if _, err := govalidator.ValidateStruct(value); err != nil {
		for _, fieldErrors := range strings.Split(err.Error(), ";") {
			if parts := strings.Split(fieldErrors, ": "); len(parts) == 2 {
				errors[parts[0]] = parts[1]
			}
		}
	}

	return errors
}

// CustomValidateStruct invokes additional validate method after ValidateStruct
func CustomValidateStruct(value Validatable) Errors {
	errors := ValidateStruct(value)
	if len(errors) != 0 {
		return errors
	}

	_, errors = value.Validate()
	return errors
}

// AddNewValidationTag provides an interface for adding new validators to govalidator
func AddNewValidationTag(tag string, v func(s string) bool) error {
	if _, exists := govalidator.TagMap[tag]; exists {
		return fmt.Errorf("Tag:%v already exists", tag)
	}

	govalidator.TagMap[tag] = v

	return nil
}

// AddNewCustomValidationTag provide an interface for adding new custom validator to govalidator
func AddNewCustomValidationTag(tag string, v func(i, o interface{}) bool) error {
	if _, exists := govalidator.CustomTypeTagMap.Get(tag); exists {
		return fmt.Errorf("Tag:%v already exists", tag)
	}

	govalidator.CustomTypeTagMap.Set(tag, v)

	return nil
}

// getErrorMessage search for validation rule message inside "messages" described above & returns it
func getErrorMessage(rule string) string {
	if message, ok := messages[rule]; ok {
		return message
	}

	return fmt.Sprintf("Value is not valid `%s`", rule)
}

package utils

import (
	"github.com/asaskevich/govalidator"
)

func IsValidURL(str string) bool {
	return govalidator.IsURL(str)
}

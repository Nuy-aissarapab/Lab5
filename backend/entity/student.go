package entity

import (
	"github.com/asaskevich/govalidator"
	
)

type Student struct {
	Fullname string  `valid:"required~Fullname is required"`
	Age      uint    `valid:"required~Age is required, range(18|200)~Age must be at least 18"`
	Email    string  `valid:"required~Email is required, email~Email is invalid"`
	GPA      float32 `valid:"float, required~GPA is required, range(0|4)~GPA must be between 0.00 and 4.00"`
}

func (s Student) Validate() (bool, error) {
	return govalidator.ValidateStruct(s)
}
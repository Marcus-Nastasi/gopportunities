package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   bool    `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (op *Opening) Update(uo *Opening) Opening {
	op.Role = uo.Role
	op.Company = uo.Company
	op.Location = uo.Location
	op.Remote = uo.Remote
	op.Link = uo.Link
	op.Salary = uo.Salary
	return *op
}

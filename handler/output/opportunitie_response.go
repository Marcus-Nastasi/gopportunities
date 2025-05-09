package output

import "time"

type OpportunitieResponse struct {
	ID			 	 uint  `json:"id"`
	CreatedAt  time.Time  `json:"createdAt"`
	UpdatedAt  time.Time	`json:"updatedAt"`
	DeletedAt  *time.Time  `json:"deletedAt"`
	Role     	 string  `json:"role"`
	Company  	 string  `json:"company"`
	Location 	 string  `json:"location"`
	Remote   	 bool    `json:"remote"`
	Link     	 string  `json:"link"`
	Salary   	 float64 `json:"salary"`
}

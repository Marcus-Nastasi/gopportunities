package mappers

import (
	"github.com/Marcus-Nastasi/gopportunities/handler/input"
	"github.com/Marcus-Nastasi/gopportunities/handler/output"
	"github.com/Marcus-Nastasi/gopportunities/schemas"
)

func MapToDomain(o *input.OpportunitieRequest) schemas.Opening {
	return schemas.Opening{
		Role: o.Role,
		Company: o.Company,
		Location: o.Location,
		Remote: o.Remote,
		Link: o.Link,
		Salary: o.Salary,
	}
}

func MapFromDomain(o *schemas.Opening) output.OpportunitieResponse {
	return output.OpportunitieResponse{
		ID: o.ID,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
		DeletedAt: &o.DeletedAt.Time,
		Role: o.Role,
		Company: o.Company,
		Location: o.Location,
		Remote: o.Remote,
		Link: o.Link,
		Salary: o.Salary,
	}
}

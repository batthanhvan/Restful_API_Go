package student

import (
	"github.com/batthanhvan/src/utils/database"
)

type Student struct {
	database.BaseModel
	FullName     *string
	DateOfBirth  *string
	StudentCode  *string `gorm:"type:varchar(255);not null;unique;"`
	Intake       *string
	Gender       *string
	PlaceOfBirth *string
	Program      *string
	Major        *string
	Degree       *string
	ClassId      *string
	ClassName    *string `gorm:"-:migration;->"`
	Gpa          *float64
}

type Search struct {
	database.DefaultSearchModel
	Student
	Query *string
}

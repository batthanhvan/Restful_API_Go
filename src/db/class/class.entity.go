package class

import (
	"github.com/batthanhvan/src/utils/database"
)

type Class struct {
	database.BaseModel
	Name     *string `gorm:"type:varchar(255);not null;unique;"`
	Duration *string
}

type Search struct {
	database.DefaultSearchModel
	Class
	Query *string
}

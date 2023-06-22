package user

import (
	"github.com/batthanhvan/src/utils/database"
)

type User struct {
	database.BaseModel
	Username string `gorm:"type:varchar(255);not null;unique;"`
	Password string `gorm:"type:varchar(255);not null"`
	Role     int32  `gorm:"type:int"`
	Name     string `gorm:"type:varchar(255)"`

	// ExpiredAt time.Time
}

type Search struct {
	database.DefaultSearchModel
	User
	Query *string
}

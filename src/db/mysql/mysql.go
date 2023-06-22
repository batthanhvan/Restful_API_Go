package mysql

import (
	"github.com/batthanhvan/src/db/class"
	"github.com/batthanhvan/src/db/student"
	"github.com/batthanhvan/src/db/user"
)

type MysqlDB struct {
	User    user.User
	Student student.Student
	Class   class.Class
}

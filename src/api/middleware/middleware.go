package middleware

import (
	"net/http"

	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/utils/token"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Mysql mysql.MysqlDB
}

var m Middleware

func (mid *Middleware) Only(role lib.ROLE) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, err := token.Decrypt(c)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		if err := m.Mysql.VerifyRole(userId, role); err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next()
	}
}

// //for customer and businesser
// func (mid *Middleware) Any() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		userId, err := token.Decrypt(c)

// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			c.Abort()
// 			return
// 		}

// 		if m.Mysql.VerifyRole(userId, lib.ROLE_CUSTOMER) != nil || m.Mysql.VerifyRole(userId, lib.ROLE_BUSINESS) != nil {
// 		} else {
// 			c.String(http.StatusUnauthorized, "Unauthorized")
// 			c.Abort()
// 			return
// 		}
// 		c.Next()
// 	}
// }

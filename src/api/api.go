package api

import (
	"github.com/batthanhvan/src/api/middleware"
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ApiServer struct {
	User    UserController
	Student StudentController
	Mid     middleware.Middleware
	Class   ClassController
	Admin   AdminController
}

var a ApiServer

func RunMain() {
	mysql.ConnectDataBase()

	r := gin.Default()
	r.SetTrustedProxies(nil)

	// gin.SetMode(gin.ReleaseMode)
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:  []string{"Authorization", "Content-Type", "User-Agent"},
		ExposeHeaders: []string{"content-disposition", "content-description"},
	}))

	r.POST("/register", a.User.HandleRegisterPost)
	r.POST("/login", a.User.HandleLoginPost)

	userGroup := r.Group("/user")
	userGroup.PUT("/change-password", a.Mid.Only(lib.ROLE_CUSTOMER), a.User.HandleChangePasswordPut)

	studentGroup := r.Group("/student")
	// studentGroup.Use(middleware.Only(lib.ROLE_CUSTOMER))
	studentGroup.GET("/:query", a.Student.HandleSearchStudentGet)
	studentGroup.GET("/", a.Student.HandleListStudentGet)

	classGroup := r.Group("/class")
	// studentGroup.Use(middleware.Only(lib.ROLE_CUSTOMER))
	classGroup.GET("/:query", a.Class.HandleSearchClassGet)
	classGroup.GET("/", a.Class.HandleListClassGet)
	classGroup.GET("/student/:query", a.Class.HandleListStudentGetByClass)

	adminGroup := r.Group("/admin")
	adminGroup.POST("/student/create", a.Mid.Only(lib.ROLE_ADMIN), a.Admin.HandleCreateStudentPost)
	adminGroup.PUT("/student/update", a.Mid.Only(lib.ROLE_ADMIN), a.Admin.HandleUpdateStudentPut)
	adminGroup.POST("/student/delete", a.Mid.Only(lib.ROLE_ADMIN), a.Admin.HandleStudentDelete)

	adminGroup.POST("/class/create", a.Mid.Only(lib.ROLE_ADMIN), a.Admin.HandleCreateClassPost)
	adminGroup.PUT("/class/update", a.Mid.Only(lib.ROLE_ADMIN), a.Admin.HandleUpdateClassPut)
	adminGroup.POST("/class/delete", a.Mid.Only(lib.ROLE_ADMIN), a.Admin.HandleClassDelete)

	r.Run("localhost:8080")

}

package api

import (
	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	S StudentService
}

func (s *StudentController) HandleSearchStudentGet(g *gin.Context) {
	req := pb.SearchStudentGetRequest{
		Query:  g.Param("query"),
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := s.S.SearchStudentGet(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *StudentController) HandleListStudentGet(g *gin.Context) {
	req := pb.SearchStudentGetRequest{
		Query:  "",
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}
	res, err := s.S.SearchStudentGet(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

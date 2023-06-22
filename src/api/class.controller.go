package api

import (
	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/gin-gonic/gin"
)

type ClassController struct {
	S ClassService
}

func (s *ClassController) HandleSearchClassGet(g *gin.Context) {
	req := pb.SearchClassGetRequest{
		Query:  g.Param("query"),
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}

	res, err := s.S.SearchClassGet(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *ClassController) HandleListClassGet(g *gin.Context) {
	req := pb.SearchClassGetRequest{
		Query:  "",
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}
	res, err := s.S.SearchClassGet(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (s *ClassController) HandleListStudentGetByClass(g *gin.Context) {
	req := pb.ListStudentGetByClassRequest{
		Query:  g.Param("query"),
		Limit:  g.DefaultQuery("limit", "10"),
		Offset: g.DefaultQuery("offset", "0"),
	}
	res, err := s.S.ListStudentGetByClass(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

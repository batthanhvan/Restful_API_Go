package api

import (
	"net/http"

	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	S AdminService
}

func (s *AdminController) HandleCreateStudentPost(g *gin.Context) {

	req := pb.CreateStudentPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.CreateStudentPost(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *AdminController) HandleUpdateStudentPut(g *gin.Context) {

	req := pb.UpdateStudentPutRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.UpdateStudentPut(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *AdminController) HandleStudentDelete(g *gin.Context) {

	req := pb.StudentDeleteRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.StudentDelete(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *AdminController) HandleCreateClassPost(g *gin.Context) {

	req := pb.CreateClassPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.CreateClassPost(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *AdminController) HandleUpdateClassPut(g *gin.Context) {

	req := pb.UpdateClassPutRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.UpdateClassPut(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *AdminController) HandleClassDelete(g *gin.Context) {

	req := pb.ClassDeleteRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.ClassDelete(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

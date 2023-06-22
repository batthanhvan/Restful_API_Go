package api

import (
	"net/http"

	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/batthanhvan/src/utils/token"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	S UserService
}

func (s *UserController) HandleRegisterPost(g *gin.Context) {

	req := pb.RegisterPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := s.S.CreateUser(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *UserController) HandleLoginPost(g *gin.Context) {

	req := pb.LoginPostRequest{}
	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := s.S.LoginCheck(&req)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

func (s *UserController) HandleChangePasswordPut(g *gin.Context) {

	req := pb.ChangePasswordPutRequest{}

	if err := g.ShouldBindJSON(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Id, err := token.Decrypt(g)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := s.S.ChangePasswordPut(&req, Id)
	if err != nil {
		lib.BadRequest(g, err)
		return
	}
	lib.Success(g, res)

}

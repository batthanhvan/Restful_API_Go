package api

import (
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/lib/validate"
	pb "github.com/batthanhvan/src/pb"
	"golang.org/x/xerrors"
)

type UserService struct {
	Mysql mysql.MysqlDB
}

func (s *UserService) CreateUser(req *pb.RegisterPostRequest) (*pb.RegisterPostResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "Username", "Password", "Role", "Name"); !ok {
		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	err := s.Mysql.Register(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.RegisterPostResponse_Data{
		Username: req.Username,
		Name:     req.Name,
	}, nil
}

func (s *UserService) LoginCheck(req *pb.LoginPostRequest) (*pb.LoginPostResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "Username", "Password"); !ok {
		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	token, err := s.Mysql.Login(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.LoginPostResponse_Data{
		Token: token,
	}, nil
}

func (s *UserService) ChangePasswordPut(req *pb.ChangePasswordPutRequest, Id string) (*pb.ChangePasswordPutResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "OldPassword", "NewPassword"); !ok {
		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	err := s.Mysql.ChangePassword(req, Id)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.ChangePasswordPutResponse_Data{}, nil
}

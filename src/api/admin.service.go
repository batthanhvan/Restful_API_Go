package api

import (
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	"github.com/batthanhvan/src/lib/validate"
	pb "github.com/batthanhvan/src/pb"
	"golang.org/x/xerrors"
)

type AdminService struct {
	Mysql mysql.MysqlDB
}

func (s *AdminService) CreateStudentPost(req *pb.CreateStudentPostRequest) (*pb.CreateStudentPostResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "FullName", "StudentCode"); !ok {

		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	err := s.Mysql.CreateStudent(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.CreateStudentPostResponse_Data{}, nil
}

func (s *AdminService) UpdateStudentPut(req *pb.UpdateStudentPutRequest) (*pb.UpdateStudentPutResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "StudentCode"); !ok {

		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	err := s.Mysql.UpdateStudent(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.UpdateStudentPutResponse_Data{}, nil
}

func (s *AdminService) StudentDelete(req *pb.StudentDeleteRequest) (*pb.StudentDeleteResponse_Data, error) {

	// if f, ok := validate.RequiredFields(req, "StudentCode"); !ok {

	// 	err := xerrors.Errorf("%w", lib.ErrMissingField(f))
	// 	return nil, err
	// }

	err := s.Mysql.StudentDelete(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.StudentDeleteResponse_Data{}, nil
}

func (s *AdminService) CreateClassPost(req *pb.CreateClassPostRequest) (*pb.CreateClassPostResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "Name"); !ok {

		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	err := s.Mysql.CreateClass(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.CreateClassPostResponse_Data{}, nil
}

func (s *AdminService) UpdateClassPut(req *pb.UpdateClassPutRequest) (*pb.UpdateClassPutResponse_Data, error) {

	if f, ok := validate.RequiredFields(req, "Name"); !ok {

		err := xerrors.Errorf("%w", lib.ErrMissingField(f))
		return nil, err
	}

	err := s.Mysql.UpdateClass(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.UpdateClassPutResponse_Data{}, nil
}

func (s *AdminService) ClassDelete(req *pb.ClassDeleteRequest) (*pb.ClassDeleteResponse_Data, error) {

	// if f, ok := validate.RequiredFields(req, "Name"); !ok {

	// 	err := xerrors.Errorf("%w", lib.ErrMissingField(f))
	// 	return nil, err
	// }

	err := s.Mysql.ClassDelete(req)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.ClassDeleteResponse_Data{}, nil
}

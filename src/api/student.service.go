package api

import (
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/db/student"
	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/batthanhvan/src/utils/database"
	"golang.org/x/xerrors"
)

type StudentService struct {
	Mysql mysql.MysqlDB
}

func (s *StudentService) SearchStudentGet(req *pb.SearchStudentGetRequest) (*pb.SearchStudentGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := s.Mysql.TotalStudent(&student.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Skip:  int(offset),
			Limit: int(limit),
		},
		Query: &req.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	p, err := s.Mysql.ListStudent(&student.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Skip:  int(offset),
			Limit: int(limit),
			Fields: []string{
				`students.full_name`,
				`students.date_of_birth`,
				`students.student_code`,
				`students.intake`,
				`students.gender`,
				`students.place_of_birth`,
				`students.program`,
				`students.major`,
				`students.degree`,
				`classes.name as class_name`,
				`students.gpa`,
			},
		},
		Query: &req.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.SearchStudentGetResponse_Data{
		Result:     s.Mysql.ConvertStudentsToProtos(p),
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}

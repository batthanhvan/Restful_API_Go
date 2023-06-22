package api

import (
	"github.com/batthanhvan/src/db/class"
	"github.com/batthanhvan/src/db/mysql"
	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/batthanhvan/src/utils/database"
	"golang.org/x/xerrors"
)

type ClassService struct {
	Mysql mysql.MysqlDB
}

func (s *ClassService) SearchClassGet(req *pb.SearchClassGetRequest) (*pb.SearchClassGetResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := s.Mysql.TotalClass(&class.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Skip:  int(offset),
			Limit: int(limit),
		},
		Query: &req.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	p, err := s.Mysql.ListClass(&class.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Skip:  int(offset),
			Limit: int(limit),
			Fields: []string{
				`classes.name`,
				`classes.duration`,
			},
		},
		Query: &req.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &pb.SearchClassGetResponse_Data{
		Result:     s.Mysql.ConvertClassesToProtos(p),
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}

func (s *ClassService) ListStudentGetByClass(req *pb.ListStudentGetByClassRequest) (*pb.ListStudentGetByClassResponse_Data, error) {

	limit := lib.ParseInt32Val(req.Limit)
	offset := lib.ParseInt32Val(req.Offset)

	total, err := s.Mysql.TotalStudentGetByClass(&class.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Skip:  int(offset),
			Limit: int(limit),
		},
		Query: &req.Query,
	})

	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	p, err := s.Mysql.ListStudentGetByClass(&class.Search{
		DefaultSearchModel: database.DefaultSearchModel{
			Skip:  int(offset),
			Limit: int(limit),
			Fields: []string{
				`classes.name as class_name`,
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

	return &pb.ListStudentGetByClassResponse_Data{
		Result:     s.Mysql.ConvertStudentsClassToProtos(p),
		Pagination: lib.Pagination(offset, limit, total),
	}, nil
}

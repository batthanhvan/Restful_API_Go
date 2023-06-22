package mysql

import (
	"github.com/batthanhvan/src/db/student"
	"github.com/batthanhvan/src/pb"
	"github.com/batthanhvan/src/utils/database"
	gouuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"

	"golang.org/x/xerrors"
)

func applySearchStudent(db *gorm.DB, search *student.Search) *gorm.DB {
	if search.ID != "" {
		db = db.Where(student.Student{
			BaseModel: database.BaseModel{
				ID: search.ID,
			},
		})
	}
	if search.Skip != 0 {
		db = db.Offset(search.Skip)
	}

	if search.Limit != 0 {
		db = db.Limit(search.Limit)
	}

	return db
}

func (p *MysqlDB) TotalStudent(search *student.Search) (*int64, error) {

	var r int64
	if err := applySearchStudent(DB, search).Model(student.Student{}).
		Where("students.full_name like ? or students.student_code = ?", "%"+*search.Query+"%", "%"+*search.Query+"%").
		Count(&r).Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}

	return &r, nil
}

func (p *MysqlDB) ListStudent(search *student.Search) ([]*student.Student, error) {

	r := make([]*student.Student, 0)

	if err := applySearchStudent(DB, search).Model(student.Student{}).
		Joins("left join classes on students.class_id = classes.id").
		Where("students.full_name like ? or students.student_code = ?", "%"+*search.Query+"%", "%"+*search.Query+"%").Debug().
		Order(`students.student_code desc, students.full_name asc`).
		Select(search.Fields).Find(&r).Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}
	return r, nil
}

func (p *MysqlDB) ConvertStudentToProto(u *student.Student) *pb.Student {
	upb := new(pb.Student)
	if u.ID != "" {
		upb.Id = u.ID
	}

	if u.FullName != nil {
		upb.FullName = *u.FullName
	}
	if u.DateOfBirth != nil {
		upb.DateOfBirth = *u.DateOfBirth
	}
	if u.StudentCode != nil {
		upb.StudentCode = *u.StudentCode
	}
	if u.Intake != nil {
		upb.Intake = *u.Intake
	}
	if u.Gender != nil {
		upb.Gender = *u.Gender
	}
	if u.PlaceOfBirth != nil {
		upb.PlaceOfBirth = *u.PlaceOfBirth
	}
	if u.Program != nil {
		upb.Program = *u.Program
	}

	if u.Major != nil {
		upb.Major = *u.Major
	}
	if u.Degree != nil {
		upb.Degree = *u.Degree
	}
	if u.ClassName != nil {
		upb.ClassName = *u.ClassName
	}
	if u.Gpa != nil {
		upb.Gpa = *u.Gpa
	}

	return upb
}

func (p *MysqlDB) ConvertStudentsToProtos(u []*student.Student) []*pb.Student {
	arr := make([]*pb.Student, 0)
	for _, a := range u {
		arr = append(arr, p.ConvertStudentToProto(a))
	}
	return arr
}

func (u *MysqlDB) CreateStudent(req *pb.CreateStudentPostRequest) error {

	var err error
	uuid, err := gouuid.NewV4()
	if err != nil {
		panic(err)
	}
	u.Student.ID = uuid.String()
	u.Student.FullName = &req.FullName
	u.Student.DateOfBirth = &req.DateOfBirth
	u.Student.StudentCode = &req.StudentCode
	u.Student.Intake = &req.Intake
	u.Student.Gender = &req.Gender
	u.Student.PlaceOfBirth = &req.PlaceOfBirth
	u.Student.Program = &req.Program
	u.Student.Major = &req.Major
	u.Student.Degree = &req.Degree
	u.Student.ClassId = &req.ClassId
	u.Student.Gpa = &req.Gpa

	err = DB.Create(&u.Student).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func (u *MysqlDB) UpdateStudent(req *pb.UpdateStudentPutRequest) error {

	var err error
	u.Student.StudentCode = &req.StudentCode
	DB.First(&u.Student)
	u.Student.FullName = &req.FullName
	u.Student.DateOfBirth = &req.DateOfBirth
	u.Student.Intake = &req.Intake
	u.Student.Gender = &req.Gender
	u.Student.PlaceOfBirth = &req.PlaceOfBirth
	u.Student.Program = &req.Program
	u.Student.Major = &req.Major
	u.Student.Degree = &req.Degree
	u.Student.ClassId = &req.ClassId
	u.Student.Gpa = &req.Gpa

	err = DB.Save(&u.Student).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func (u *MysqlDB) StudentDelete(req *pb.StudentDeleteRequest) error {

	err := DB.Where("students.student_code = ?", req.Query).Delete(&u.Student).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

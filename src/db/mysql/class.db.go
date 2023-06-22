package mysql

import (
	"github.com/batthanhvan/src/db/class"
	"github.com/batthanhvan/src/db/student"
	"github.com/batthanhvan/src/pb"
	"github.com/batthanhvan/src/utils/database"
	gouuid "github.com/nu7hatch/gouuid"
	"gorm.io/gorm"

	"golang.org/x/xerrors"
)

func applySearchClass(db *gorm.DB, search *class.Search) *gorm.DB {
	if search.ID != "" {
		db = db.Where(class.Class{
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

func (p *MysqlDB) TotalClass(search *class.Search) (*int64, error) {

	var r int64
	if err := applySearchClass(DB, search).Model(class.Class{}).
		Where("classes.name like ? ", "%"+*search.Query+"%").
		Count(&r).Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}

	return &r, nil
}

func (p *MysqlDB) ListClass(search *class.Search) ([]*class.Class, error) {

	r := make([]*class.Class, 0)

	if err := applySearchClass(DB, search).Model(class.Class{}).
		Where("classes.name like ? ", "%"+*search.Query+"%").
		Order(`classes.name asc`).
		Select(search.Fields).Find(&r).Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}
	return r, nil
}

func (p *MysqlDB) ConvertClassToProto(u *class.Class) *pb.Class {
	upb := new(pb.Class)
	if u.ID != "" {
		upb.Id = u.ID
	}

	if u.Name != nil {
		upb.Name = *u.Name
	}
	if u.Duration != nil {
		upb.Duration = *u.Duration
	}

	return upb
}

func (p *MysqlDB) ConvertClassesToProtos(u []*class.Class) []*pb.Class {
	arr := make([]*pb.Class, 0)
	for _, a := range u {
		arr = append(arr, p.ConvertClassToProto(a))
	}
	return arr
}

func (p *MysqlDB) TotalStudentGetByClass(search *class.Search) (*int64, error) {

	var r int64
	if err := applySearchClass(DB, search).Model(class.Class{}).
		Joins("left join students on students.class_id = classes.id").
		Where("classes.name like ? ", "%"+*search.Query+"%").
		Count(&r).Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}

	return &r, nil
}

func (p *MysqlDB) ListStudentGetByClass(search *class.Search) ([]*student.Student, error) {

	r := make([]*student.Student, 0)

	if err := applySearchClass(DB, search).Model(class.Class{}).Debug().
		Joins("left join students on students.class_id = classes.id").
		Where("classes.name like ? ", "%"+*search.Query+"%").
		Order(`classes.name asc`).
		Select(search.Fields).Find(&r).Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}
	return r, nil
}

func (p *MysqlDB) ConvertStudentClassToProto(u *student.Student) *pb.Student {
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

func (p *MysqlDB) ConvertStudentsClassToProtos(u []*student.Student) []*pb.Student {
	arr := make([]*pb.Student, 0)
	for _, a := range u {
		arr = append(arr, p.ConvertStudentToProto(a))
	}
	return arr
}

func (u *MysqlDB) CreateClass(req *pb.CreateClassPostRequest) error {

	var err error
	uuid, err := gouuid.NewV4()
	if err != nil {
		panic(err)
	}
	u.Class.ID = uuid.String()
	u.Class.Name = &req.Name
	u.Class.Duration = &req.Duration

	err = DB.Create(&u.Class).Debug().Error
	if err != nil {
		return err
	}
	return nil
}
func (u *MysqlDB) UpdateClass(req *pb.UpdateClassPutRequest) error {

	var err error
	u.Class.Name = &req.Name
	DB.First(&u.Class)
	u.Class.Duration = &req.Duration

	err = DB.Save(&u.Class).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func (u *MysqlDB) ClassDelete(req *pb.ClassDeleteRequest) error {

	err := DB.Where("classes.name = ?", req.Query).Delete(&u.Class).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

package mysql

import (
	"html"
	"strings"

	"github.com/batthanhvan/src/db/user"
	"github.com/batthanhvan/src/lib"
	pb "github.com/batthanhvan/src/pb"
	"github.com/batthanhvan/src/utils/token"
	gouuid "github.com/nu7hatch/gouuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/xerrors"
)

// func applySearchUser(db *gorm.DB, search *user.Search) *gorm.DB {
// 	if search.ID != "" {
// 		db = db.Where(user.User{
// 			BaseModel: database.BaseModel{
// 				ID: search.ID,
// 			},
// 		})
// 	}
// 	if search.Skip != 0 {
// 		db = db.Offset(search.Skip)
// 	}

// 	if search.Limit != 0 {
// 		db = db.Limit(search.Limit)
// 	}

// 	return db
// }

func (u *MysqlDB) Register(req *pb.RegisterPostRequest) error {

	var err error
	u.User.Password, err = u.Hash(req.Password)
	if err != nil {
		return err
	}
	//remove spaces in username
	u.User.Username = html.EscapeString(strings.TrimSpace(req.Username))

	uuid, err := gouuid.NewV4()
	if err != nil {
		panic(err)
	}
	u.User.ID = uuid.String()
	u.User.Role = req.Role
	u.User.Name = req.Name

	err = DB.Create(&u.User).Debug().Error
	if err != nil {
		return err
	}
	return nil
}

func (u *MysqlDB) Hash(password string) (string, error) {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil

}

func (u *MysqlDB) Login(req *pb.LoginPostRequest) (string, error) {

	err := DB.Model(user.User{}).Where("username = ?", req.Username).Take(&u.User).Error
	if err != nil {
		return "", err
	}

	err = u.VerifyPassword(req.Password, u.User.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.Encrypt(u.User.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *MysqlDB) VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *MysqlDB) VerifyRole(userId string, role lib.ROLE) error {

	err := DB.Model(user.User{}).Where(`"users"."id" = ? AND "users"."role" = ?"`, userId, role).Error
	return err
}

func (u *MysqlDB) ChangePassword(req *pb.ChangePasswordPutRequest, Id string) error {

	err := DB.Model(user.User{}).Where("id = ?", Id).Take(&u.User).Error
	if err != nil {
		return err
	}

	err = u.VerifyPassword(req.OldPassword, u.User.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return err
	}
	u.User.Password, err = u.Hash(req.NewPassword)
	if err != nil {
		return err
	}

	if err = DB.Updates(&u.User).Debug().Error; err != nil {
		err = xerrors.Errorf("%w", err)
		panic(err.Error())
	}
	return nil

}

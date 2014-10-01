package models

import (
	"errors"
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
)

type User struct {
	Id			int64
	Username	string  `orm:"size(100)" form:"Username"  valid:"Required"`
	Password	string  `orm:"size(100)" form:"Password"  valid:"Required"`
}

func (g *User) TableName() string {
	return "user"
}

func init(){
	orm.RegisterModel(new(User))
}

func Pwdhash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

func CheckLogin(username string, password string) (user User, err error) {
	
	user = User{ Username: username, Password: Pwdhash( password ) }
	
	o := orm.NewOrm()
	o.Read(&user, "Username")
	
	if user.Id == 0 {
		return user, errors.New("User Not Found !!!")
	}
	
	return user, nil
}

func GetLogin(user User){
	
}


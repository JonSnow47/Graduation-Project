/*
 * Revision History:
 *     Initial: 2018/04/26        Chen Yanchen
 */

package account

import (
	"os"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/JonSnow47/Graduation-Project/shop/db/mysql"
	"github.com/JonSnow47/Graduation-Project/shop/util"
)

type accountServiceProvider struct{}

var (
	AccountService *accountServiceProvider
	db             *gorm.DB
)

type Account struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(16);unique;not null"`
	Pwd       string `gorm:"type:varchar(32)"`
	Phone     string `gorm:"type:varchar(20);unique"`
	Avatar    os.File
	Male      bool `gorm:"type:bool"`
	Level     int8 `gorm:"type:int;not null"` // Account level
	Admin     bool `gorm:"type:bool;not null;default:false"`
	State     bool `gorm:"type:bool"`
	CreateAt  time.Time
	LastLogin time.Time
}

func init() {
	db := mysql.InitMysql(mysql.DatabaseShop)
	defer db.Close()
	db = db.CreateTable(db, &Account{})
	db = db.Model(Account{})
}

// WechatLogin login with wechat permission.
func (*accountServiceProvider) WechatLogin() error {
	return nil
}

// PhoneLogin login with phone validate code.
func (*accountServiceProvider) PhoneLogin(phone string) (u *Account, err error) {
	db := mysql.InitMysql(mysql.DatabaseShop)
	defer db.Close()

	err = db.Where(&Account{Phone: phone}).First(u).Error
	if err == nil {
		return
	}

	if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	a := &Account{
		Name:     phone,
		Phone:    phone,
		Level:    Level0,
		State:    true,
		CreateAt: time.Now(),
	}
	err = db.Create(a).Error
	return nil, err
}

// Register in web.
func (*accountServiceProvider) Register(name, pwd string, avatar string) error {
	db := mysql.InitMysql(mysql.DatabaseShop)
	defer db.Close()

	pwd, err := util.PwdService.Generator(pwd)
	if err != nil {
		return err
	}

	a := &Account{
		Name:     name,
		Pwd:      pwd,
		Level:    Level0,
		CreateAt: time.Now(),
	}

	err = db.Model(&Account{}).Create(a).Error
	return err
}

// Login in web.
func (*accountServiceProvider) Login(name, pwd string) bool {
	db := mysql.InitMysql(mysql.DatabaseShop)
	defer db.Close()

	var hashPwd string

	err := db.Model(&Account{}).Where(&Account{Name: name}).Find(&Account{Pwd: hashPwd}).Error
	if err != nil {
		return false
	}

	if util.PwdService.Comparator([]byte(hashPwd), pwd) {
		return true
	}

	return false
}

// Logout delete session or close JWT.
func (*accountServiceProvider) Logout() error {
	return nil
}

// 模型的行为动作
package domain

import (
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture"
	"gorm.io/gorm/clause"
)

// 广告行为
func FindAllAdvertisements(db *infrasturcture.Database) (r []Advertisement) {
	db.Session.Find(&r)
	return
}

// 产品行为
func FindAllProducts(db *infrasturcture.Database) (r []Product) {
	db.Session.Find(&r)
	return
}

func GetProduct(db *infrasturcture.Database, id int) (r Product) {
	db.Session.Preload(clause.Associations).Take(&r, id)
	return
}

// 用户行为
func GetAccountByName(db *infrasturcture.Database, username string) (r Account, num int64) {
	return r, db.Session.Where("username = ?", username).Take(&r).RowsAffected
}

func UpdateAccount(db *infrasturcture.Database, account *Account) int64 {
	return db.Session.Save(account).RowsAffected
}

func FindUsersByNameOrEmailOrTelephone(db *infrasturcture.Database, name, email, telephone string) (r []Account) {
	db.Session.Where("username = ? OR email =? OR telephone = ?", name, email, telephone).Find(&r)
	return
}

func CreateAccount(db *infrasturcture.Database, account *Account) int64 {
	return db.Session.Create(account).RowsAffected
}

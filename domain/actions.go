// 模型的行为动作
package domain

import (
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"gorm.io/gorm/clause"
)

// 广告行为
func FindAllAdvertisements(db *db.Database) (r []Advertisement, err error) {
	return r, db.Session.Find(&r).Error
}

// 产品行为
func FindAllProducts(db *db.Database) (r []Product, err error) {
	return r, db.Session.Find(&r).Error
}

func GetProduct(db *db.Database, id int) (r Product, err error) {
	return r, db.Session.Preload(clause.Associations).Take(&r, id).Error
}

func UpdateProduct(db *db.Database, product *Product) (int64, error) {
	ret := db.Session.Save(product)
	return ret.RowsAffected, ret.Error
}

func CreateProduct(db *db.Database, product *Product) (int64, error) {
	ret := db.Session.Create(product)
	return ret.RowsAffected, ret.Error
}

func DeleteProduct(db *db.Database, id int) (int64, error) {
	ret := db.Session.Delete(&Product{}, id)
	return ret.RowsAffected, ret.Error
}

// 库存行为
func GetStockpile(db *db.Database, id int) (r Stockpile, err error) {
	return r, db.Session.Preload(clause.Associations).Take(&r, id).Error
}

func UpdateStockpile(db *db.Database, stockpile *Stockpile) (int64, error) {
	ret := db.Session.Save(stockpile)
	return ret.RowsAffected, ret.Error
}

func UpdateStockpileAmount(db *db.Database, id int, amount int) (int64, error) {
	ret := db.Session.Model(&Stockpile{}).Where("id = ?", id).Update("amount", amount)
	return ret.RowsAffected, ret.Error
}

// 用户行为
func GetAccountByName(db *db.Database, username string) (r Account, err error) {
	return r, db.Session.Where("username = ?", username).Take(&r).Error
}

func UpdateAccount(db *db.Database, account *Account) (int64, error) {
	ret := db.Session.Save(account)
	return ret.RowsAffected, ret.Error
}

func FindUsersByNameOrEmailOrTelephone(db *db.Database, name, email, telephone string) (r []Account, err error) {
	return r, db.Session.Where("username = ? OR email =? OR telephone = ?", name, email, telephone).Find(&r).Error
}

func CreateAccount(db *db.Database, account *Account) (int64, error) {
	ret := db.Session.Create(account)
	return ret.RowsAffected, ret.Error
}

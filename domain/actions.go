// 模型的行为动作，可见不支持泛型写这些东西是多么的啰嗦
package domain

import (
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"time"
)

func (s *BaseModel) IsNew() bool {
	return s.ID == 0
}

// 广告行为
func NewAdvertisement(db *db.Database) *Advertisement {
	return &Advertisement{BaseModel: BaseModel{Database: db}}
}

func (s *Advertisement) FindAll() (r []Advertisement, err error) {
	return r, s.Session.Find(&r).Error
}

// 产品行为
func NewProduct(db *db.Database) *Product {
	return &Product{BaseModel: BaseModel{Database: db}}
}

func (s *Product) FindAll() (r []Product, err error) {
	return r, s.Session.Find(&r).Error
}

func (s *Product) FindByIDs(ids []uint) (r []Product, err error) {
	return r, s.Session.Find(&r, ids).Error
}

func (s *Product) Get(id uint) (*Product, error) {
	return s, s.Session.Preload(clause.Associations).Take(s, id).Error
}

func (s *Product) Update() (int64, error) {
	ret := s.Session.Save(s)
	return ret.RowsAffected, ret.Error
}

func (s *Product) Create() (int64, error) {
	ret := s.Session.Create(s)
	return ret.RowsAffected, ret.Error
}

func (s *Product) Delete(id ...uint) (int64, error) {
	if len(id) > 0 {
		ret := s.Session.Delete(&Product{}, id)
		return ret.RowsAffected, ret.Error
	} else {
		ret := s.Session.Delete(&Product{}, s.ID)
		return ret.RowsAffected, ret.Error
	}
}

// 库存行为
func NewStockpile(db *db.Database) *Stockpile {
	return &Stockpile{BaseModel: BaseModel{Database: db}}
}

func (s *Stockpile) Get(id uint) (*Stockpile, error) {
	return s, s.Session.Take(s, id).Error
}

func (s *Stockpile) GetByProductId(id uint) (*Stockpile, error) {
	return s, s.Session.Model(&Stockpile{}).Where("product_id = ?", id).Take(s).Error
}

func (s *Stockpile) Update() (int64, error) {
	ret := s.Session.Save(s)
	return ret.RowsAffected, ret.Error
}

func (s *Stockpile) UpdateAmount(id uint, amount uint) (int64, error) {
	ret := s.Session.Model(&Stockpile{}).Where("id = ?", id).Update("amount", amount)
	return ret.RowsAffected, ret.Error
}

// 用户行为
func NewAccount(db *db.Database) *Account {
	return &Account{BaseModel: BaseModel{Database: db}}
}

func (s *Account) GetByName(username string) (*Account, error) {
	return s, s.Session.Where("username = ?", username).Take(s).Error
}

func (s *Account) Update() (int64, error) {
	ret := s.Session.Save(s)
	return ret.RowsAffected, ret.Error
}

func (s *Account) FindByNameOrEmailOrTelephone(name, email, telephone string) (r []Account, err error) {
	return r, s.Session.Where("username = ? OR email =? OR telephone = ?", name, email, telephone).Find(&r).Error
}

func (s *Account) Create() (int64, error) {
	ret := s.Session.Create(s)
	return ret.RowsAffected, ret.Error
}

// 支付单行为
func NewPayment(db *db.Database) *Payment {
	return &Payment{
		BaseModel: BaseModel{Database: db},
	}
}

func NewPaymentWithInfo(db *db.Database, user string, total float64, expires int64) *Payment {
	id := uuid.New().String()
	return &Payment{
		BaseModel:   BaseModel{Database: db},
		TotalPrice:  total,
		Expires:     expires,
		CreateTime:  time.Now(),
		PayState:    Waiting,
		PayId:       id,
		PaymentLink: "/pay/modify/" + id + "?state=PAYED&accountId=" + user,
	}
}

func (s *Payment) Create() (int64, error) {
	ret := s.Session.Create(s)
	return ret.RowsAffected, ret.Error
}

func (s *Payment) Get(id uint) (*Payment, error) {
	return s, s.Session.Take(s, id).Error
}

func (s *Payment) GetByPayID(payId string) (*Payment, error) {
	return s, s.Session.Where("pay_id = ?", payId).Take(s).Error
}

func (s *Payment) Update() (int64, error) {
	ret := s.Session.Save(s)
	return ret.RowsAffected, ret.Error
}

// 用户钱包行为
func NewWallet(db *db.Database) *Wallet {
	return &Wallet{BaseModel: BaseModel{Database: db}}
}
func (s *Wallet) GetByAccountId(id uint) (*Wallet, error) {
	return s, s.Session.Where("account_id = ?", id).Take(s).Error
}

func (s *Wallet) Update() (int64, error) {
	ret := s.Session.Save(s)
	return ret.RowsAffected, ret.Error
}

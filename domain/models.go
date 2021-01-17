// 模型结构
package domain

import (
	"errors"
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture/db"
	"time"
)

type BaseModel struct {
	ID           uint `json:"id" gorm:"primaryKey;autoIncrement;notNull"`
	*db.Database `json:"-" gorm:"-" binding:"-"`
}

// 广告实体
type Advertisement struct {
	BaseModel
	ProductId uint   `json:"productId"`
	Image     string `json:"image"`
}

// 产品实例
type Product struct {
	BaseModel
	Title          string          `json:"title"`
	Price          float64         `json:"price"` // 用浮点数来表示金额是很不好的行为，这里只是方便演示
	Rate           float32         `json:"rate"`
	Description    string          `json:"description"`
	Cover          string          `json:"cover"`
	Detail         string          `json:"detail"`
	Specifications []Specification `json:"specifications" gorm:"foreignKey:ProductId"`
}

// 产品规格
type Specification struct {
	BaseModel
	Item      string `json:"item"`
	Value     string `json:"value"`
	ProductId uint   `json:"productId"`
}

// 用户
type Account struct {
	BaseModel
	Username  string `json:"username" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Password  string `json:"password,omitempty" binding:"required"`
	Avatar    string `json:"avatar"`
	Telephone string `json:"telephone" binding:"required,numeric,startswith=1,len=11"`
	Email     string `json:"email" binding:"required,email"`
	Location  string `json:"location"`
}

// 商品库存
type Stockpile struct {
	BaseModel
	Amount    int  `json:"amount"`
	Frozen    int  `json:"frozen"`
	ProductId uint `json:"productId"`
}

// 用户钱包
type Wallet struct {
	BaseModel
	Money     float64 `json:"money"`
	AccountId uint    `json:"accountId"`
}

// 支付结算单模型
type Settlement struct {
	Items      []Item            `json:"items"`
	Purchase   Purchase          `json:"purchase"`
	ProductMap map[uint]*Product `json:"productMap"`
}

// 结算单中要购买的商品
type Item struct {
	ID     uint `json:"id"`
	Amount int  `json:"amount"`
}

// 结算单中的配送信息
type Purchase struct {
	Delivery  bool   `json:"delivery"`
	Pay       string `json:"pay"`
	Name      string `json:"name" binding:"required"`
	Telephone string `json:"telephone" binding:"required,numeric,startswith=1,len=11"`
	Location  string `json:"location"`
}

type PayState int

const (
	Waiting PayState = iota
	Cancel
	Payed
	Timeout
)

var stateName = []string{"WAITING", "CANCEL", "PAYED", "TIMEOUT"}

func ConvPayStateToString(state PayState) string {
	return stateName[state]
}

func ConvStringToPayState(state string) (PayState, error) {
	for k, v := range stateName {
		if state == v {
			return PayState(k), nil
		}
	}
	return 0, errors.New("无效的PayState")
}

// 支付单模型
// 就是传到客户端让用户给扫码或者其他别的方式付钱的对象
type Payment struct {
	BaseModel
	CreateTime  time.Time `json:"createTime"`
	PayId       string    `json:"payId"`
	TotalPrice  float64   `json:"totalPrice"`
	Expires     int64     `json:"expires"`
	PaymentLink string    `json:"paymentLink" gorm:"-"`
	PayState    PayState  `json:"payState"`
}

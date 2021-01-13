// 模型结构
package domain

import "time"

type BaseModel struct {
	ID uint `json:"id" gorm:"primaryKey"`
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
	Price          float64         `json:"price"`
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
	Password  string `json:"password" binding:"required"`
	Avatar    string `json:"avatar"`
	Telephone string `json:"telephone" binding:"required,numeric,startswith=1,len=11"`
	Email     string `json:"email" binding:"required,email"`
	Location  string `json:"location"`
}

// 支付单
type Payment struct {
	BaseModel
	CreateTime  time.Time `json:"createTime"`
	PayId       string    `json:"payId"`
	TotalPrice  float64   `json:"totalPrice"`
	Expires     uint64    `json:"expires"`
	PaymentLink string    `json:"paymentLink"`
	PayState    int       `json:"payState"`
}

// 商品库存
type Stockpile struct {
	BaseModel
	Amount  uint32  `json:"amount"`
	Frozen  uint32  `json:"frozen"`
	Product Product `json:"product"`
}

// 用户钱包
type Wallet struct {
	BaseModel
	Money   float64 `json:"money"`
	Account Account `json:"account"`
}

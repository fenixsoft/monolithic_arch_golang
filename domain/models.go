package domain

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
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Rate        float32 `json:"rate"`
	Description string  `json:"description"`
	Cover       string  `json:"cover"`
	Detail      string  `json:"detail"`
}

// 产品规格
type Specification struct {
	BaseModel
	Item      string `json:"item"`
	Value     string `json:"value"`
	ProductId uint   `json:"productId"`
}

func (db *Database) FindAllAdvertisements() (r []Advertisement) {
	db.Session.Find(&r)
	return
}

func (db *Database) FindAllProducts() (r []Product) {
	db.Session.Find(&r)
	return
}

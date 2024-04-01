package models

type Fruit struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Nutrition Nutrition `json:"nutrition" gorm:"foreignKey:FruitID"`
}

type Nutrition struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	FruitID      string  `json:"fruit_id" `
	Calories     float64 `json:"calories"`
	Fat          float64 `json:"fat"`
	Sugar        float64 `json:"sugar"`
	Carbohydrate float64 `json:"carbohydrate"`
	Protein      float64 `json:"protein"`
}

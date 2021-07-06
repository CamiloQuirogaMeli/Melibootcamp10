package product

type Product struct {
	Id           int     `form:"id" json:"id"`
	Name         string  `form:"name" json:"name"`
	Color        string  `form:"color" json:"color"`
	Price        float64 `form:"price" json:"price"`
	Stock        int     `form:"stock" json:"stock"`
	Code         string  `form:"code" json:"code"`
	Published    bool    `form:"published" json:"published"`
	CreationDate string  `form:"creationDate" json:"creationDate"`
}

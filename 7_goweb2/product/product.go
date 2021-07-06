package product

type Product struct {
	ID				int		`json: "id" form: "id"`
	Name			string	`json: "name"`
	Color			string	`json: "color"`
	Price			int		`json: "price"`
	Stock			int		`json: "stock"`
	Code			string	`json: "code"`	
	Published		bool	`json: "published"`
	CreationDate	string	`json: "creationDate"`
}
package product

type SpecialProduct struct {
	size int
}

func (p SpecialProduct) TotalSize() int {
	return p.size
}

func (p SpecialProduct) IsSpecial() bool {
	return true
}
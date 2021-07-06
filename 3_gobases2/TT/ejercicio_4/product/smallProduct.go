package product

type SmallProduct struct {
	size int
}

func (p SmallProduct) TotalSize() int {
	return p.size	
}

func (p SmallProduct) IsSpecial() bool {
	return false
}
package product

type FragileProduct struct {
	size int
}

func (p FragileProduct) TotalSize() int {
	additionalSpace := (p.size * 75) / 100
	return p.size + additionalSpace
}

func (p FragileProduct) IsSpecial() bool {
	return false
}
package product

type LargeProduct struct {
	size int
}

func (p LargeProduct) TotalSize() int {
	additionalSpace := (p.size * 20) / 100
	return p.size + additionalSpace
}

func (p LargeProduct) IsSpecial() bool {
	return false
}
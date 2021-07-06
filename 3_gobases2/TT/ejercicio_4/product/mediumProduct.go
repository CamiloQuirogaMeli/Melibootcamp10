package product

type MediumProduct struct {
	size int
}

func (p MediumProduct) TotalSize() int {
	additionalSpace := (p.size * 5) / 100
	return p.size + additionalSpace
}

func (p MediumProduct) IsSpecial() bool {
	return false
}
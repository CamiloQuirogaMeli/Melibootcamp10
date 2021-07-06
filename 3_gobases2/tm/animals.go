package tm

type Animal struct {
	Name      string
	GramsFood uint
}

var (
	TARANTULA = Animal{
		Name:      "TARANTULA",
		GramsFood: 150,
	}
	HAMSTER = Animal{
		Name:      "HAMSTER",
		GramsFood: 250,
	}
	CAT = Animal{
		Name:      "CAT",
		GramsFood: 5000,
	}
	DOG = Animal{
		Name:      "DOG",
		GramsFood: 10000,
	}
)

func CalculateFood(animal Animal) func (numAnimals uint) uint {
	switch animal {
	case CAT:
		return catFood
	case DOG:
		return dogFood
	case HAMSTER:
		return hamsterFood
	case TARANTULA:
		return tarantulaFood
	default:
		return nil
	}
}

func catFood(numAnimals uint) uint {
	return numAnimals * CAT.GramsFood
}

func dogFood(numAnimals uint) uint {
	return numAnimals * DOG.GramsFood
}

func hamsterFood(numAnimals uint) uint {
	return numAnimals * HAMSTER.GramsFood
}

func tarantulaFood(numAnimals uint) uint {
	return numAnimals * TARANTULA.GramsFood
}

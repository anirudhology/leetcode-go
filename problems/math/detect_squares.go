package math

type DetectSquares struct {
	Coordinates [][]int
	Frequencies map[string]int
}

func DetectSquaresConstructor() DetectSquares {
	return DetectSquares{
		Coordinates: [][]int{},
		Frequencies: map[string]int{},
	}
}

func (this *DetectSquares) Add(point []int) {
	this.Coordinates = append(this.Coordinates, point)
	key := string(rune(point[0])) + "," + string(rune(point[1]))
	this.Frequencies[key] += 1
}

func (this *DetectSquares) Count(point []int) int {
	x, y := point[0], point[1]
	result := 0
	for _, coordinate := range this.Coordinates {
		px, py := coordinate[0], coordinate[1]
		if DetectSquaresAbs(px-x) == DetectSquaresAbs(py-y) && px != x && py != y {
			key1, key2 := string(rune(x))+","+string(rune(py)), string(rune(px))+","+string(rune(y))
			result += this.Frequencies[key1] * this.Frequencies[key2]
		}
	}
	return result
}

func DetectSquaresAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

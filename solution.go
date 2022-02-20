package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

// A = √3/4 (a²) to determine the area of an equilateral triangle.

type sidesNum int

const Pi = 3.14
const SidesTriangle = 3
const SidesSquare = 4
const SidesCircle = 0

func CalcSquare(sideLen float64, sidesNum sidesNum) float64 {
	if sidesNum == SidesTriangle {
		return math.Sqrt(3) / 4 * math.Pow(sideLen, 2)
	} else if sidesNum == SidesSquare {
		return math.Pow(sideLen, 2)
	} else if sidesNum == SidesCircle {
		return Pi * math.Pow(sideLen, 2)
	} else {
		return 0
	}

}

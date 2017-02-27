package stringguard

import "math"

func approximatelyEqual(a, b, epsilon float64) bool {
	rhs := math.Abs(a)
	if math.Abs(a) < math.Abs(b) {
		rhs = math.Abs(b)
	}
	return math.Abs(a-b) <= (rhs * epsilon)
}

func essentiallyEqual(a, b, epsilon float64) bool {
	rhs := math.Abs(a)
	if math.Abs(a) > math.Abs(b) {
		rhs = math.Abs(b)
	}
	return math.Abs(a-b) <= (rhs * epsilon)
}

func definitelyGreaterThan(a, b, epsilon float64) bool {
	rhs := math.Abs(a)
	if math.Abs(a) < math.Abs(b) {
		rhs = math.Abs(b)
	}
	return (a - b) > (rhs * epsilon)
}

func definitelyLessThan(a, b, epsilon float64) bool {
	rhs := math.Abs(a)
	if math.Abs(a) < math.Abs(b) {
		rhs = math.Abs(b)
	}
	return (b - a) > (rhs * epsilon)
}

package stringguard // import "zhihaojun.com/stringguard"

import "regexp"
import "strconv"

const epsilon = 0.00000001

type Guard struct {
	err error
}

func NewGuard() *Guard {
	return &Guard{
		err: nil,
	}
}

func (g *Guard) Err() error {
	return g.err
}

func (g *Guard) Error() string {
	if g.err != nil {
		return g.err.Error()
	}
	return ""
}

func (g *Guard) Reset() {
	g.err = nil
}

func (g *Guard) Regex(name, str string, reg *regexp.Regexp) string {
	if g.err != nil {
		return str
	}
	if ok := reg.MatchString(str); ok == false {
		// not match
		g.err = NewRegexError(name, reg)
	}
	return str
}

func (g *Guard) MaxLen(name, str string, max int) string {
	if g.err != nil {
		return str
	}
	if len(str) > max {
		g.err = NewMaxLenError(name, max)
	}
	return str
}

func (g *Guard) MinLen(name, str string, min int) string {
	if g.err != nil {
		return str
	}
	if len(str) < min {
		g.err = NewMinLenError(name, min)
	}
	return str
}

func (g *Guard) Int(name, str string) int {
	if g.err != nil {
		return 0
	}

	intValue, err := strconv.Atoi(str)
	if err != nil {
		g.err = NewIntError(name)
	}
	return intValue
}

func (g *Guard) MaxInt(name, str string, max int) int {
	if g.err != nil {
		return 0
	}

	intValue, err := strconv.Atoi(str)
	if err != nil {
		g.err = NewIntError(name)
	}
	if intValue > max {
		g.err = NewMaxIntError(name, max)
	}
	return intValue
}

func (g *Guard) MinInt(name, str string, min int) int {
	if g.err != nil {
		return 0
	}

	intValue, err := strconv.Atoi(str)
	if err != nil {
		g.err = NewIntError(name)
	}
	if intValue < min {
		g.err = NewMinIntError(name, min)
	}
	return intValue
}

func (g *Guard) Float(name, str string) float64 {
	if g.err != nil {
		return 0
	}

	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		g.err = NewFloatError(name)
	}
	return floatValue
}

func (g *Guard) MaxFloat(name, str string, max float64) float64 {
	if g.err != nil {
		return 0
	}

	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		g.err = NewFloatError(name)
	}
	if definitelyGreaterThan(floatValue, max, epsilon) {
		g.err = NewMaxFloatError(name, max)
	}
	return floatValue
}

func (g *Guard) MinFloat(name, str string, min float64) float64 {
	if g.err != nil {
		return 0
	}

	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		g.err = NewFloatError(name)
	}
	if definitelyLessThan(floatValue, min, epsilon) {
		g.err = NewMinFloatError(name, min)
	}
	return floatValue
}

func (g *Guard) Enum(name, str string, enums ...string) string {
	if g.err != nil {
		return str
	}

	for _, v := range enums {
		if v == str {
			return str
		}
	}

	// not found
	g.err = NewEnumError(name, enums)
	return str
}

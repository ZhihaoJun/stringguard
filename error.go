package stringguard

import (
	"fmt"
	"regexp"
)

type GuardError struct {
	name string
}

func NewGuardError(name string) *GuardError {
	return &GuardError{
		name: name,
	}
}

func (ge *GuardError) Name() string {
	return ge.name
}

func (ge *GuardError) Error() string {
	return fmt.Sprintf("%s:invalid", ge.name)
}

type NotEmptyError struct {
	*GuardError
}

func NewNotEmptyError(name string) *NotEmptyError {
	return &NotEmptyError{
		GuardError: NewGuardError(name),
	}
}

type RegexError struct {
	*GuardError
	regexp *regexp.Regexp
}

func NewRegexError(name string, regexp *regexp.Regexp) *RegexError {
	return &RegexError{
		GuardError: NewGuardError(name),
		regexp:     regexp,
	}
}

func (re *RegexError) Regexp() *regexp.Regexp {
	return re.regexp
}

type MaxLenError struct {
	*GuardError
	max int
}

func NewMaxLenError(name string, max int) *MaxLenError {
	return &MaxLenError{
		GuardError: NewGuardError(name),
		max:        max,
	}
}

func (gmle *MaxLenError) Max() int {
	return gmle.max
}

type MinLenError struct {
	*GuardError
	min int
}

func NewMinLenError(name string, min int) *MinLenError {
	return &MinLenError{
		GuardError: NewGuardError(name),
		min:        min,
	}
}

func (gmle *MinLenError) Min() int {
	return gmle.min
}

type IntError struct {
	*GuardError
}

func NewIntError(name string) *IntError {
	return &IntError{
		GuardError: NewGuardError(name),
	}
}

type MaxIntError struct {
	*GuardError
	max int
}

func NewMaxIntError(name string, max int) *MaxIntError {
	return &MaxIntError{
		GuardError: NewGuardError(name),
		max:        max,
	}
}

func (mie *MaxIntError) Max() int {
	return mie.max
}

type MinIntError struct {
	*GuardError
	min int
}

func NewMinIntError(name string, min int) *MinIntError {
	return &MinIntError{
		GuardError: NewGuardError(name),
		min:        min,
	}
}

func (mie *MinIntError) Min() int {
	return mie.min
}

type FloatError struct {
	*GuardError
}

func NewFloatError(name string) *FloatError {
	return &FloatError{
		GuardError: NewGuardError(name),
	}
}

type MaxFloatError struct {
	*GuardError
	max float64
}

func NewMaxFloatError(name string, max float64) *MaxFloatError {
	return &MaxFloatError{
		GuardError: NewGuardError(name),
		max:        max,
	}
}

func (mfe *MaxFloatError) Max() float64 {
	return mfe.max
}

type MinFloatError struct {
	*GuardError
	min float64
}

func NewMinFloatError(name string, min float64) *MinFloatError {
	return &MinFloatError{
		GuardError: NewGuardError(name),
		min:        min,
	}
}

func (mfe *MinFloatError) Min() float64 {
	return mfe.min
}

type EnumError struct {
	*GuardError
	enums []string
}

func NewEnumError(name string, enums []string) *EnumError {
	return &EnumError{
		GuardError: NewGuardError(name),
		enums:      enums,
	}
}

func (ee *EnumError) Enums() []string {
	return ee.enums
}

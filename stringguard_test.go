package stringguard

import "testing"

func TestInt(t *testing.T) {
	g := NewGuard()
	i := g.Int("test", "123")
	if i != 123 {
		t.Error("[int convert error]")
	}
	i = g.Int("test1", "abd")
	err := g.Err()
	if err == nil {
		t.Error("[int guard test failed]")
	}
	switch terr := err.(type) {
	case *IntError:
		_ = terr
	default:
		t.Error("[int guard error type wrong]")
	}
}

func TestMaxLen(t *testing.T) {
	g := NewGuard()
	g.MaxLen("test", "123", 3)
	if g.Err() != nil {
		t.Error("[max len check failed]")
	}
	g.MaxLen("test", "hahaha", 2)
	err := g.Err()
	if err == nil {
		t.Error("[max len error is nil]")
	}
	switch terr := err.(type) {
	case *MaxLenError:
		if terr.Max() != 2 {
			t.Error("[max len error info wrong]", terr.Max())
		}
	default:
		t.Error("[max len error type wrong]")
	}
}

func TestMinLen(t *testing.T) {
	g := NewGuard()
	g.MinLen("test", "123", 2)
	if g.Err() != nil {
		t.Error("[min len check failed]")
	}
	g.MinLen("test", "123", 10)
	err := g.Err()
	if err == nil {
		t.Error("[min len error is nil]")
	}
	switch terr := err.(type) {
	case *MinLenError:
		if terr.Min() != 10 {
			t.Error("[min len error info wrong]", terr.Min())
		}
	default:
		t.Error("[min len error type wrong]")
	}
}

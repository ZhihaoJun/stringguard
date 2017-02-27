package stringguard

import "testing"

func TestInt(t *testing.T) {
	g := New()
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

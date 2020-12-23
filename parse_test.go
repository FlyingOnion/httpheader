package header

import (
	"net/http"
	"testing"
	"time"
)

type A struct {
	I  int           `httpheader:"foo;int"`
	T  time.Time     `httpheader:"bar;time"`
	Du time.Duration `httpheader:"dag;duration"`
	IS []int         `httpheader:"baz;[]int"`
}

func TestHeader(t *testing.T) {
	h := http.Header{}
	h.Set("foo", "1")
	h.Set("bar", time.Now().Format(timeFormat))
	h.Set("dag", "1m4s")
	h.Add("baz", "2")
	h.Add("baz", "3")
	var s A
	if err := Parse(h, &s); err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(s.I, s.T.Unix(), s.Du.Seconds(), s.IS)
}

func BenchmarkHeader(b *testing.B) {
	h := http.Header{}
	h.Set("foo", "1")
	var s A
	for i := 0; i < b.N; i++ {
		_ = Parse(h, &s)
	}
}

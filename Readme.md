# httpheader
httpheader is a Go library for parsing `http.Header` to a struct.

## install
`go get -u github.com/FlyingOnion/httpheader`

or

`go mod download github.com/FlyingOnion/httpheader`

## usage
see `parse_test.go`

```
type A struct {
	I  int           `httpheader:"foo;int"`
	T  time.Time     `httpheader:"bar;time"`
	Du time.Duration `httpheader:"dag;duration"`
	IS []int         `httpheader:"baz;[]int"`
}

h := http.Header{}
h.Set("foo", "1")
h.Set("bar", time.Now().Format("2006-01-02 15:04:05"))
h.Set("dag", "1m4s")
h.Add("baz", "2")
h.Add("baz", "3")

var s A
err := Parse(h, &s)
```

## other type support?
This is a quick developing version. 

I will add them in the future (maybe not until I need them).

## LICENSE

[Anti-996 LICENSE](https://github.com/FlyingOnion/httpheader/blob/master/LICENSE)

surprised?
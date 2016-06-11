package help

import "testing"

type student struct {
	Id   int64
	Name string
}

func TestMapArray(t *testing.T) {
	var v []interface{}
	v[0] = student{1, "username"}
	v[1] = student{2, "liujx"}

	m := MapArray(v, "Id", "Name")
	t.Error(m)
}

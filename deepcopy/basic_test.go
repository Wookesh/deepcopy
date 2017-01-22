package deepcopy

import "testing"

func TestClone(t *testing.T) {
	i := new(int)
	*i = 5

	o := Clone_int(i)

	*i = 6

	if *o == *i {
		t.Fail()
	}
}

func TestListClone(t *testing.T) {
	l := make([]int, 0)
	for i := 0; i < 5; i++ {
		l = append(l, i)
	}

	z := CloneList_int(l)

	for i := 0; i < 5; i++ {
		l[i] = i + 6
	}

	for i := 0; i < 5; i++ {
		if l[i] != z[i]+6 {
			t.Fail()
		}
	}
}

func simpleClone(a int) *int {
	v := int(a)
	return &v
}

func BenchmarkSimple(b *testing.B) {

	v := new(int)
	*v = 42
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = simpleClone(*v)
	}

}

func BenchmarkPtr(b *testing.B) {

	v := new(int)
	*v = 42
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = Clone_int(v)
	}

}

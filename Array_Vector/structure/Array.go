package Array_Vector

type List interface {
	Add(int)
	Get(int) int
	Remove(int)
	Size() int
}

type Array struct {
	data []int
}
type Vector struct {
	data []int
}

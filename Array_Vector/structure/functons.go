package Array_Vector

func (al *Array) Add(element int) {
	al.data = append(al.data, element)
}

func (al *Array) Get(index int) int {
	return al.data[index]
}

func (al *Array) Remove(index int) {
	al.data = append(al.data[:index], al.data[index+1:]...)
}
func (vc *Vector) reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func (al *Array) Size() int {
	return len(al.data)
}
func (al *Array) Print() {

}
func (vc *Vector) Add(element int) {
	vc.data = append(vc.data, element)
}
func (vc *Vector) Get(index int) int {
	return vc.data[index]
}
func (vc *Vector) Remove(index int) {
	vc.data = append(vc.data[:index], vc.data[index+1:]...)
}
func (vc *Vector) size() int {
	return len(vc.data)
}

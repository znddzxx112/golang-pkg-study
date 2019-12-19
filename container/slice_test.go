package containert

import "testing"

func TestSlice(t *testing.T) {
	intSlice := make([]int, 2, 4)
	t.Log(intSlice)
	t.Log(len(intSlice))
	t.Log(cap(intSlice))

	intSlice[0] = 1

	intSlice = append(intSlice, 1)
	intSlice = append(intSlice, 2, 3)
	t.Log(intSlice)
	t.Log(len(intSlice))
	t.Log(cap(intSlice))

	int2Slice := []int{6, 7, 8}
	intSlice = append(intSlice, int2Slice...)
	t.Log(intSlice)

	int3Slice := []int{10, 11, 12}
	copy(intSlice[2:], int3Slice[:3])
	t.Log(intSlice)

	replacei := 3
	intSlice = append(intSlice, 0)
	copy(intSlice[replacei+1:], intSlice[replacei:])
	intSlice[replacei] = 5
	t.Log(intSlice)

	deletei := 2
	deleteitem := intSlice[deletei]
	t.Log(deleteitem)
	copy(intSlice[deletei:], intSlice[deletei+1:])
	intSlice = intSlice[:len(intSlice)-1]
	t.Log(intSlice)

	clearInt := make([]int, 2)
	t.Log(copy(intSlice[7:], clearInt))
	t.Log(intSlice)
}

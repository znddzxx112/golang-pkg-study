package sortt

import (
	"fmt"
	"sort"
	"testing"
)

func TestInts(t *testing.T) {
	ints := []int{
		3, 2, 5, 1,
	}

	sort.Ints(ints)
	t.Log(sort.IntsAreSorted(ints))
	t.Log(sort.SearchInts(ints, 5))
	t.Log(sort.SearchInts(ints, 7))
	t.Log(ints)
}

type Student struct {
	Name  string
	Score int
}

type StudentsSlice []Student

func (s StudentsSlice) Len() int {
	return len(s)
}

func (s StudentsSlice) Less(i, j int) bool {
	return (s)[i].Score <= (s)[j].Score
}

func (s StudentsSlice) Swap(i, j int) {
	(s)[i], (s)[j] = (s)[j], (s)[i]
}

func TestSort(t *testing.T) {
	students := StudentsSlice{
		{"foo", 4},
		{"bar", 6},
		{"car", 5},
	}
	t.Log(sort.Search(students.Len(), func(i int) bool {
		return students[i].Score == 5
	}))
	sort.Sort(students)
	t.Log(students)
	t.Log(sort.IsSorted(students))
	t.Log(sort.Search(len(students), func(j int) bool {
		fmt.Println(students[j].Score)
		return students[j].Score == 4
	}))
}

// sort multi key

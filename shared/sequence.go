package shared

import (
	"fmt"
	"sort"
)

type Sequence []int

//Methods required by sort.Interface
func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//Copy returns a copy of the Sequence
func (s Sequence) Copy() Sequence {
	seq := make(Sequence, 0, len(s))
	return append(seq, s...)
}

//Method for printing. Sorts the sequence before it.
func (s Sequence) String() string {
	s = s.Copy() //Make a copy; don't overwrite argument. ?????

	//Implementing the interfaces of Sort as above.
	//sort.Sort(s)

	//We convert Sequence to an int slice so no need to implement the interfaces.
	sort.IntSlice(s).Sort()

	return fmt.Sprint([]int(s))
}

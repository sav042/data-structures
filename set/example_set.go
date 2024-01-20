package set

import "github.com/deckarep/golang-set"

func Example() {
	s1 := mapset.NewSet(1, 2)
	s2 := mapset.NewSet(2, 3)
	_ = s1.Difference(s2) //1,3
	_ = s1.Union(s2)      //1,2,3
}

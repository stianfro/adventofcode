package main

import "fmt"

type Crate struct {
	Stack1 []string
	Stack2 []string
	Stack3 []string
}

// type Crate struct {
// 	Stack1 []string
// 	Stack2 []string
// 	Stack3 []string
// 	Stack4 []string
// 	Stack5 []string
// 	Stack6 []string
// 	Stack7 []string
// 	Stack8 []string
// 	Stack9 []string
// }

func printCrate(crate Crate) {
	fmt.Println(crate.Stack1)
	fmt.Println(crate.Stack2)
	fmt.Println(crate.Stack3)
	// fmt.Println(crate.Stack4)
	// fmt.Println(crate.Stack5)
	// fmt.Println(crate.Stack6)
	// fmt.Println(crate.Stack7)
	// fmt.Println(crate.Stack8)
	// fmt.Println(crate.Stack9)
}

func main() {
	input := Crate{
		Stack1: []string{"Z", "N"},
		Stack2: []string{"D", "C", "M"},
		Stack3: []string{"P"},
	}
	// input := Crate{
	// 	Stack1: []string{"S", "L", "W"},
	// 	Stack2: []string{"J", "T", "N", "Q"},
	// 	Stack3: []string{"S", "C", "H", "F", "J"},
	// 	Stack4: []string{"T", "R", "M", "W", "N", "G", "B"},
	// 	Stack5: []string{"T", "R", "L", "S", "D", "H", "Q", "B"},
	// 	Stack6: []string{"M", "J", "B", "V", "F", "H", "R", "L"},
	// 	Stack7: []string{"D", "W", "R", "N", "J", "M"},
	// 	Stack8: []string{"B", "Z", "T", "F", "H", "N", "D", "J"},
	// 	Stack9: []string{"H", "L", "Q", "N", "B", "F", "T"},
	// }

	printCrate(input)
}

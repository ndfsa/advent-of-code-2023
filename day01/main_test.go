package day01

import (
	"fmt"
	"testing"
)

func TestSolvePart1Example(t *testing.T) {
	res, err := SolvePart1("./example.txt", 4)

	if err != nil {
		t.Fatal(err)
	}

	if res != 142 {
		t.Fatal(fmt.Sprintf("incorrect: %d", res))
	}
	t.Logf("res: %d", res)
}

// func TestSolvePart1(t *testing.T) {
// 	res, err := SolvePart1("./input.txt", 1000)
//
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if res != 54573 {
// 		t.Fatal(fmt.Sprintf("incorrect: %d", res))
// 	}
// 	t.Logf("res: %d", res)
// }

func TestSolvePart2Example(t *testing.T) {
	res, err := SolvePart2("./example2.txt", 7)

	if err != nil {
		t.Fatal(err)
	}

	if res != 281 {
		t.Fatal(fmt.Sprintf("incorrect: %d", res))
	}
	t.Logf("res: %d", res)
}

// func TestSolvePart2(t *testing.T) {
// 	res, err := SolvePart2("./input.txt", 1000)
//
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if res != 54591 {
// 		t.Fatal(fmt.Sprintf("incorrect: %d", res))
// 	}
// 	t.Logf("res: %d", res)
// }

package dynamicprogramming

import (
	"fmt"
	"testing"
)

func TestGridTraveler(t *testing.T) {
	fmt.Println("tests")
	fmt.Print(gridTravelerMemoized(1, 1, make(map[string]int)))
	fmt.Print(" | ")
	fmt.Print(gridTravelerMemoized(2, 3, make(map[string]int)))
	fmt.Print(" | ")
	fmt.Print(gridTravelerMemoized(3, 2, make(map[string]int)))
	fmt.Print(" | ")
	fmt.Print(gridTravelerMemoized(3, 3, make(map[string]int)))
	fmt.Print(" | ")
	fmt.Print(gridTravelerMemoized(18, 18, make(map[string]int)))
	fmt.Print(" | ")

	fmt.Println()
	fmt.Print(gridTraveler(1, 1)) //1
	fmt.Print(" | ")
	fmt.Print(gridTraveler(2, 3)) //3
	fmt.Print(" | ")
	fmt.Print(gridTraveler(3, 2)) //3
	fmt.Print(" | ")
	fmt.Print(gridTraveler(3, 3)) //6
	fmt.Print(" | ")
	fmt.Print(gridTraveler(18, 18)) //2333606220
	fmt.Print(" | ")

}

/*

Traveller on a 2d grid,start at top left corner and goal is to reach bottom right corner
can move either right or down.
How many ways you can travel on a grid from the top left to the bottom right for dimension m*n
For example : gridTraveler(2,3)->3 (2 rows, 3 columns)
*/

func gridTraveler(row int, column int) int {
	if row == 1 && column == 1 {
		return 1
	}
	if row == 0 || column == 0 {
		return 0
	}
	return gridTraveler(row-1, column) + gridTraveler(row, column-1)
}

func gridTravelerMemoized(row int, column int, pathMap map[string]int) int {
	key:=fmt.Sprintf("%d:%d",row,column)
	if value, exists := pathMap[key]; exists {
		return value
	}
	if row == 1 && column == 1 {
		return 1
	}
	if row == 0 || column == 0 {
		return 0
	}
	pathMap[key] = gridTravelerMemoized(row-1, column, pathMap) + gridTravelerMemoized(row, column-1, pathMap)

	return pathMap[key]
}

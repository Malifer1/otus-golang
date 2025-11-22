package main

import (
	"fmt"
	"strings"
    "strconv"
)

func create_chessboard(size int) string {
	var board strings.Builder
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if (i+j)%2 == 0 {
				board.WriteString(" ")
			} else {
				board.WriteString("#")
			}
		}
		if i < size-1 {
			board.WriteString("\n")
		}
	}

	return board.String()
}

func main() {
    var stdin string
    fmt.Print("Enter size of chessboard: ")
	fmt.Scanln(&stdin)
	size, err:= strconv.Atoi(strings.TrimSpace(stdin))
	if err != nil {
		fmt.Println("An error:\n", err, "\noccured after entering number")
		return
	}
	if size <= 0 {
		fmt.Println("Size of chessboard must be a positive value!")
		return
	}
	board := create_chessboard(size)
    fmt.Println("Chessobard with size: ", size)
	fmt.Println(board)
}

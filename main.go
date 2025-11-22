package main

import (
	"fmt"
	"strings"
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
	size := 8
	chessBoard := create_chessboard(size)
	fmt.Println("Chessobard with size: ", size)
	fmt.Println(chessBoard)
}

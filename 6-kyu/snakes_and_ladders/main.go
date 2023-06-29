package cw

func SnakesAndLadders(board, dice []int) int {
	boardIndex := 0
	for _, roll := range dice {
		if boardIndex+roll < len(board) {
			boardIndex += roll
			if boardIndex+1 == len(board) {
				break
			}
			boardIndex += board[boardIndex]
		}
	}
	return boardIndex
}

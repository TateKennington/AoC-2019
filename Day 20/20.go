package main

import (
	"bufio"
	"fmt"
	"os"
)

func alpha(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

type state struct {
	pos   [2]int
	dist  int
	level int
}

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	var grid = make([][]rune, 0)
	for scn.Scan() {
		var line = scn.Text()
		grid = append(grid, make([]rune, 0, len(line)))
		for _, c := range line {
			grid[len(grid)-1] = append(grid[len(grid)-1], c)
		}
	}
	var warps = make(map[[2]rune][][2]int)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '.' {
				continue
			}
			if i-2 >= 0 && alpha(grid[i-1][j]) && alpha(grid[i-2][j]) {
				if _, ok := warps[[2]rune{grid[i-2][j], grid[i-1][j]}]; ok {
					warps[[2]rune{grid[i-2][j], grid[i-1][j]}] = append(warps[[2]rune{grid[i-2][j], grid[i-1][j]}], [2]int{i, j})
				} else {
					warps[[2]rune{grid[i-2][j], grid[i-1][j]}] = [][2]int{[2]int{i, j}}
				}
			}
			if j-2 >= 0 && alpha(grid[i][j-1]) && alpha(grid[i][j-2]) {
				if _, ok := warps[[2]rune{grid[i][j-2], grid[i][j-1]}]; ok {
					warps[[2]rune{grid[i][j-2], grid[i][j-1]}] = append(warps[[2]rune{grid[i][j-2], grid[i][j-1]}], [2]int{i, j})
				} else {
					warps[[2]rune{grid[i][j-2], grid[i][j-1]}] = [][2]int{[2]int{i, j}}
				}
			}
			if i+2 < len(grid) && alpha(grid[i+1][j]) && alpha(grid[i+2][j]) {
				if _, ok := warps[[2]rune{grid[i+1][j], grid[i+2][j]}]; ok {
					warps[[2]rune{grid[i+1][j], grid[i+2][j]}] = append(warps[[2]rune{grid[i+1][j], grid[i+2][j]}], [2]int{i, j})
				} else {
					warps[[2]rune{grid[i+1][j], grid[i+2][j]}] = [][2]int{[2]int{i, j}}
				}
			}
			if j+2 < len(grid[i]) && alpha(grid[i][j+1]) && alpha(grid[i][j+2]) {
				if _, ok := warps[[2]rune{grid[i][j+1], grid[i][j+2]}]; ok {
					warps[[2]rune{grid[i][j+1], grid[i][j+2]}] = append(warps[[2]rune{grid[i][j+1], grid[i][j+2]}], [2]int{i, j})
				} else {
					warps[[2]rune{grid[i][j+1], grid[i][j+2]}] = [][2]int{[2]int{i, j}}
				}
			}
		}
	}
	var queue = []state{state{warps[[2]rune{'A', 'A'}][0], 0, 0}}
	var visited = make(map[[3]int]bool)
	for len(queue) > 0 {
		var curr = queue[0]
		var pos = curr.pos
		queue = queue[1:]

		visited[[3]int{pos[0], pos[1], curr.level}] = true

		if pos == warps[[2]rune{'Z', 'Z'}][0] && curr.level == 0 {
			fmt.Println(curr.dist)
			break
		}

		if pos[0]-1 >= 0 && grid[pos[0]-1][pos[1]] == '.' && !visited[[3]int{pos[0] - 1, pos[1], curr.level}] {
			queue = append(queue, state{[2]int{pos[0] - 1, pos[1]}, curr.dist + 1, curr.level})
		}
		if pos[1]-1 >= 0 && grid[pos[0]][pos[1]-1] == '.' && !visited[[3]int{pos[0], pos[1] - 1, curr.level}] {
			queue = append(queue, state{[2]int{pos[0], pos[1] - 1}, curr.dist + 1, curr.level})
		}
		if pos[0]+1 < len(grid) && grid[pos[0]+1][pos[1]] == '.' && !visited[[3]int{pos[0] + 1, pos[1], curr.level}] {
			queue = append(queue, state{[2]int{pos[0] + 1, pos[1]}, curr.dist + 1, curr.level})
		}
		if pos[1]+1 < len(grid[pos[0]]) && grid[pos[0]][pos[1]+1] == '.' && !visited[[3]int{pos[0], pos[1] + 1, curr.level}] {
			queue = append(queue, state{[2]int{pos[0], pos[1] + 1}, curr.dist + 1, curr.level})
		}

		if pos[0]-2 >= 0 && alpha(grid[pos[0]-1][pos[1]]) && alpha(grid[pos[0]-2][pos[1]]) {
			if len(warps[[2]rune{grid[pos[0]-2][pos[1]], grid[pos[0]-1][pos[1]]}]) < 2 {
				continue
			}
			newPos := warps[[2]rune{grid[pos[0]-2][pos[1]], grid[pos[0]-1][pos[1]]}][0]
			if newPos == pos {
				newPos = warps[[2]rune{grid[pos[0]-2][pos[1]], grid[pos[0]-1][pos[1]]}][1]
			}

			newLevel := 0
			if pos[0] < len(grid)/2 {
				newLevel = curr.level - 1
			} else {
				newLevel = curr.level + 1
			}

			if newLevel < 0 {
				continue
			}

			if visited[[3]int{newPos[0], newPos[1], newLevel}] {
				continue
			}
			queue = append(queue, state{newPos, curr.dist + 1, newLevel})
		}
		if pos[1]-2 >= 0 && alpha(grid[pos[0]][pos[1]-1]) && alpha(grid[pos[0]][pos[1]-2]) {
			if len(warps[[2]rune{grid[pos[0]][pos[1]-2], grid[pos[0]][pos[1]-1]}]) < 2 {
				continue
			}
			newPos := warps[[2]rune{grid[pos[0]][pos[1]-2], grid[pos[0]][pos[1]-1]}][0]
			if newPos == pos {
				newPos = warps[[2]rune{grid[pos[0]][pos[1]-2], grid[pos[0]][pos[1]-1]}][1]
			}

			newLevel := 0
			if pos[1] < len(grid[pos[0]])/2 {
				newLevel = curr.level - 1
			} else {
				newLevel = curr.level + 1
			}

			if newLevel < 0 {
				continue
			}

			if visited[[3]int{newPos[0], newPos[1], newLevel}] {
				continue
			}

			queue = append(queue, state{newPos, curr.dist + 1, newLevel})
		}
		if pos[0]+2 < len(grid) && alpha(grid[pos[0]+1][pos[1]]) && alpha(grid[pos[0]+2][pos[1]]) {
			if len(warps[[2]rune{grid[pos[0]+1][pos[1]], grid[pos[0]+2][pos[1]]}]) < 2 {
				continue
			}
			newPos := warps[[2]rune{grid[pos[0]+1][pos[1]], grid[pos[0]+2][pos[1]]}][0]
			if newPos == pos {
				newPos = warps[[2]rune{grid[pos[0]+1][pos[1]], grid[pos[0]+2][pos[1]]}][1]
			}

			newLevel := 0
			if pos[0] < len(grid)/2 {
				newLevel = curr.level + 1
			} else {
				newLevel = curr.level - 1
			}

			if newLevel < 0 {
				continue
			}

			if visited[[3]int{newPos[0], newPos[1], newLevel}] {
				continue
			}

			queue = append(queue, state{newPos, curr.dist + 1, newLevel})
		}
		if pos[1]+2 < len(grid[pos[0]]) && alpha(grid[pos[0]][pos[1]+1]) && alpha(grid[pos[0]][pos[1]+2]) {
			if len(warps[[2]rune{grid[pos[0]][pos[1]+1], grid[pos[0]][pos[1]+2]}]) < 2 {
				continue
			}
			newPos := warps[[2]rune{grid[pos[0]][pos[1]+1], grid[pos[0]][pos[1]+2]}][0]
			if newPos == pos {
				newPos = warps[[2]rune{grid[pos[0]][pos[1]+1], grid[pos[0]][pos[1]+2]}][1]
			}

			newLevel := 0
			if pos[1] < len(grid[pos[0]])/2 {
				newLevel = curr.level + 1
			} else {
				newLevel = curr.level - 1
			}

			if newLevel < 0 {
				continue
			}

			if visited[[3]int{newPos[0], newPos[1], newLevel}] {
				continue
			}

			queue = append(queue, state{newPos, curr.dist + 1, newLevel})
		}
	}
}

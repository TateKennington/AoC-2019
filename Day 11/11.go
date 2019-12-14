package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Computer struct {
	codes        map[int]int
	relativeBase int
	curr         int
	halted       bool
}

func newComputer(codes map[int]int) *Computer {
	return &Computer{codes, 0, 0, false}
}

func (this *Computer) compute(input int) int {
	step := 0
	codes := this.codes
	relativeBase := this.relativeBase
	for curr := this.curr; !this.halted; curr += step {
		opcode := codes[curr] % 100
		param := codes[curr] / 100
		switch opcode {
		case 99:
			this.halted = true
		case 1:
			res := 0
			if param%10 == 0 {
				res = codes[codes[curr+1]]
			} else if param%10 == 1 {
				res = codes[curr+1]
			} else {
				res = codes[relativeBase+codes[curr+1]]
			}

			if (param/10)%10 == 0 {
				res += codes[codes[curr+2]]
			} else if (param/10)%10 == 1 {
				res += codes[curr+2]
			} else {
				res += codes[relativeBase+codes[curr+2]]
			}

			if (param/100)%10 == 0 {
				codes[codes[curr+3]] = res
			} else {
				codes[relativeBase+codes[curr+3]] = res
			}
			step = 4
		case 2:
			res := 0
			if param%10 == 0 {
				res = codes[codes[curr+1]]
			} else if param%10 == 1 {
				res = codes[curr+1]
			} else {
				res = codes[relativeBase+codes[curr+1]]
			}

			if (param/10)%10 == 0 {
				res *= codes[codes[curr+2]]
			} else if (param/10)%10 == 1 {
				res *= codes[curr+2]
			} else {
				res *= codes[relativeBase+codes[curr+2]]
			}

			if (param/100)%10 == 0 {
				codes[codes[curr+3]] = res
			} else {
				codes[relativeBase+codes[curr+3]] = res
			}
			step = 4
		case 3:
			if param == 0 {
				codes[codes[curr+1]] = input
			} else {
				codes[relativeBase+codes[curr+1]] = input
			}
			step = 2
		case 4:
			res := 0
			if param == 0 {
				res = codes[codes[curr+1]]
			} else if param == 1 {
				res = codes[curr+1]
			} else {
				res = codes[relativeBase+codes[curr+1]]
			}
			curr += 2
			this.curr = curr
			this.relativeBase = relativeBase
			return res
		case 5:
			var cond bool
			if param%10 == 0 {
				cond = codes[codes[curr+1]] != 0
			} else if param%10 == 1 {
				cond = codes[curr+1] != 0
			} else {
				cond = codes[relativeBase+codes[curr+1]] != 0
			}

			if cond {
				step = 0
				if (param/10)%10 == 0 {
					curr = codes[codes[curr+2]]
				} else if (param/10)%10 == 1 {
					curr = codes[curr+2]
				} else {
					curr = codes[relativeBase+codes[curr+2]]
				}
			} else {
				step = 3
			}
		case 6:
			var cond bool
			if param%10 == 0 {
				cond = codes[codes[curr+1]] == 0
			} else if param%10 == 1 {
				cond = codes[curr+1] == 0
			} else {
				cond = codes[relativeBase+codes[curr+1]] == 0
			}

			if cond {
				step = 0
				if (param/10)%10 == 0 {
					curr = codes[codes[curr+2]]
				} else if (param/10)%10 == 1 {
					curr = codes[curr+2]
				} else {
					curr = codes[relativeBase+codes[curr+2]]
				}
			} else {
				step = 3
			}
		case 7:
			first := 0
			second := 0
			step = 4
			if param%10 == 0 {
				first = codes[codes[curr+1]]
			} else if param%10 == 1 {
				first = codes[curr+1]
			} else {
				first = codes[relativeBase+codes[curr+1]]
			}

			if (param/10)%10 == 0 {
				second = codes[codes[curr+2]]
			} else if (param/10)%10 == 1 {
				second = codes[curr+2]
			} else {
				second = codes[relativeBase+codes[curr+2]]
			}

			res := 0
			if first < second {
				res = 1
			} else {
				res = 0
			}

			if (param/100)%10 == 0 {
				codes[codes[curr+3]] = res
			} else {
				codes[relativeBase+codes[curr+3]] = res
			}
		case 8:
			first := 0
			second := 0
			step = 4
			if param%10 == 0 {
				first = codes[codes[curr+1]]
			} else if param%10 == 1 {
				first = codes[curr+1]
			} else {
				first = codes[relativeBase+codes[curr+1]]
			}

			if (param/10)%10 == 0 {
				second = codes[codes[curr+2]]
			} else if (param/10)%10 == 1 {
				second = codes[curr+2]
			} else {
				second = codes[relativeBase+codes[curr+2]]
			}

			res := 0
			if first == second {
				res = 1
			} else {
				res = 0
			}

			if (param/100)%10 == 0 {
				codes[codes[curr+3]] = res
			} else {
				codes[relativeBase+codes[curr+3]] = res
			}
		case 9:
			if param == 0 {
				relativeBase += codes[codes[curr+1]]
			}
			if param == 1 {
				relativeBase += codes[curr+1]
			}
			if param == 2 {
				relativeBase += codes[relativeBase+codes[curr+1]]
			}
			step = 2
		}
	}
	return -1
}

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input = scanner.Text()
	var codesText = strings.Split(input, ",")
	var codes = make(map[int]int)
	for i, code := range codesText {
		x, _ := strconv.Atoi(code)
		codes[i] = x
	}
	var computer = newComputer(codes)
	var directions = [4][2]int{[2]int{0, 1}, [2]int{1, 0}, [2]int{0, -1}, [2]int{-1, 0}}
	var dir = 0
	var pos = [2]int{0, 5}
	var panels = make(map[[2]int]int)
	var left = 0
	var right = 0
	var up = 0
	var down = 0
	panels[pos] = 1
	for !computer.halted {
		var color = computer.compute(panels[pos])
		var turn = computer.compute(panels[pos])
		if color == -1 || turn == -1 {
			break
		}
		if pos[0] < left {
			left = pos[0]
		}
		if pos[1] < down {
			down = pos[1]
		}
		if pos[0] > right {
			right = pos[0]
		}
		if pos[1] > up {
			up = pos[1]
		}
		panels[pos] = color
		if turn == 0 {
			dir--
			if dir < 0 {
				dir += 4
			}
		} else {
			dir = (dir + 1) % 4
		}
		pos[0] += directions[dir][0]
		pos[1] += directions[dir][1]
	}
	fmt.Println(len(panels))
	fmt.Printf("%d %d %d %d\n", right, up, left, down)
	for i := up; i >= down; i-- {
		for j := left; j <= right; j++ {
			color, ok := panels[[2]int{j, i}]
			if !ok {
				fmt.Print("0")
			} else {
				fmt.Print(color)
			}
		}
		fmt.Println()
	}
}

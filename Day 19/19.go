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
	var codesCopy = make(map[int]int)
	for key, value := range codes {
		codesCopy[key] = value
	}
	return &Computer{codesCopy, 0, 0, false}
}

func (this *Computer) compute(input []int) int {
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
			if len(input) == 0 {
				curr += 2
				this.curr = curr
				this.relativeBase = relativeBase
				return -1
			}
			if param == 0 {
				codes[codes[curr+1]] = input[0]
			} else {
				codes[relativeBase+codes[curr+1]] = input[0]
			}
			input = input[1:]
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
	var ans []int
	upper := 1600
	lower := 1500
	for upper-lower > 1 {
		fmt.Printf("Upper: %d, Lower: %d\n", upper, lower)
		mid := (upper + lower) / 2
		pos := []int{mid, 0}
		for newComputer(codes).compute(pos) != 1 {
			pos[1]++
		}

		done := true
		for newComputer(codes).compute(pos) == 1 {
			done = true
			for i := pos[0]; i < pos[0]+100 && done; i++ {
				for j := pos[1]; j < pos[1]+100 && done; j++ {
					if newComputer(codes).compute([]int{i, j}) != 1 {
						done = false
					}
				}
			}
			if done {
				ans = pos
				break
			}
			pos[1]++
		}
		if done {
			upper = mid - 1
		} else {
			lower = mid
		}
	}
	fmt.Println(ans)
}

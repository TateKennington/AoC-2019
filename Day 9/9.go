package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input = scanner.Text()
	var codesText = strings.Split(input, ",")
	var codes = make(map[int]int)
	var output []int
	for i, code := range codesText {
		x, _ := strconv.Atoi(code)
		codes[i] = x
	}
	step := 0
	relativeBase := 0
	for curr := 0; curr < len(codes); curr += step {
		opcode := codes[curr] % 100
		param := codes[curr] / 100
		switch opcode {
		case 99:
			goto end
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
				codes[codes[curr+1]] = 2
			} else {
				codes[relativeBase+codes[curr+1]] = 2
			}
			step = 2
		case 4:
			if param == 0 {
				output = append(output, codes[codes[curr+1]])
			} else if param == 1 {
				output = append(output, codes[curr+1])
			} else {
				output = append(output, codes[relativeBase+codes[curr+1]])
			}
			step = 2
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
end:
	fmt.Println(output)
}

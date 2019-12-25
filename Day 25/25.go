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

func commandify(s string) []int {
	res := make([]int, 0, len(s))
	for _, c := range s {
		res = append(res, int(c))
	}
	return res
}

func main() {
	f, _ := os.Open("25.in")
	defer f.Close()
	var scanner = bufio.NewScanner(f)
	var kb = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input = scanner.Text()
	var codesText = strings.Split(input, ",")
	var codes = make(map[int]int)
	for i, code := range codesText {
		x, _ := strconv.Atoi(code)
		codes[i] = x
	}
	var computer = newComputer(codes)
	var output = make([][]rune, 1)
	var curr = 0
	var cmd []int
	var cmds = make([][]int, 0)
	var brute = false
	var items = []string{"mutex", "dark matter", "cake", "klein bottle", "tambourine", "fuel cell", "astrolabe", "monolith"}
	for !computer.halted {
		if brute {
			cmd = cmds[0]
		}
		temp := computer.compute(cmd)
		if temp == -1 {
			break
		}
		if temp == 10 {
			if curr < len(output) && len(output[curr]) == 8 {
				prompt := true
				for i, x := range commandify("Command?") {
					if rune(x) != output[curr][i] {
						prompt = false
						break
					}
				}
				if prompt && !brute {
					kb.Scan()
					if kb.Text() == "brute" {
						brute = true
						for i := 0; i < (1 << uint(len(items)+1)); i++ {
							fmt.Printf("%d/%d\n", i, (1 << uint(len(items)+1)))
							for j := 0; j < len(items); j++ {
								if i&(1<<uint(j)) != 0 {
									cmds = append(cmds, commandify(fmt.Sprintf("take %s\n", items[j])))
								}
							}
							cmds = append(cmds, commandify("north\n"))
							for j := 0; j < len(items); j++ {
								cmds = append(cmds, commandify(fmt.Sprintf("drop %s\n", items[j])))
							}
						}
					} else {
						cmd = commandify(kb.Text())
						cmd = append(cmd, '\n')
					}
				}
				if prompt && brute {
					cmds = cmds[1:]
				}
			}
			curr++
			fmt.Println()
		} else {
			fmt.Printf("%c", temp)
			for curr >= len(output) {
				output = append(output, make([]rune, 0))
			}
			output[curr] = append(output[curr], rune(temp))
		}
	}
}

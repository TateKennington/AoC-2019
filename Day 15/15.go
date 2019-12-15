package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	f, _ := os.Open("15.in")
	defer f.Close()
	var scanner = bufio.NewScanner(f)
	//var kb = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var input = scanner.Text()
	var codesText = strings.Split(input, ",")
	var codes = make(map[int]int)
	for i, code := range codesText {
		x, _ := strconv.Atoi(code)
		codes[i] = x
	}
	var computer = newComputer(codes)
	var directions = [4][2]int{[2]int{0, 1}, [2]int{0, -1}, [2]int{-1, 0}, [2]int{1, 0}}
	var pos = [2]int{0, 0}
	var panels = make(map[[2]int]int)
	var dist = make(map[[2]int]int)
	var left = 0
	var right = 0
	var up = 0
	var down = 0
	panels[pos] = 1
	dist[pos] = 0
	var airSystem [2]int
	var count = 10000000
	for !computer.halted && count > 0 {
		count--
		/*if count%1000 == 0 {
			fmt.Println(count)
		}*/
		dir := 1 + rand.Intn(4)
		status := computer.compute(dir)
		newPos := [2]int{pos[0] + directions[dir-1][0], pos[1] + directions[dir-1][1]}
		if status == 0 {
			panels[newPos] = 5
		} else if status == 1 {
			panels[pos]--
			if temp, ok := dist[newPos]; !ok || dist[pos]+1 < temp {
				dist[newPos] = dist[pos] + 1
			}
			pos = newPos
			panels[pos] = 1
		} else {
			panels[pos]--
			if temp, ok := dist[newPos]; !ok || dist[pos]+1 < temp {
				dist[newPos] = dist[pos] + 1
			}
			pos = newPos
			panels[pos] = 4
			airSystem = pos
			//done = true
		}
		if newPos[0] < left {
			left = newPos[0]
		}
		if newPos[1] < down {
			down = newPos[1]
		}
		if newPos[0] > right {
			right = newPos[0]
		}
		if newPos[1] > up {
			up = newPos[1]
		}
	}
	fmt.Println("===============================")
	for i := up; i >= down; i-- {
		for j := left; j <= right; j++ {
			temp := [2]int{j, i}
			tile, ok := panels[temp]
			if !ok {
				fmt.Print(" ")
			} else {
				fmt.Print(tile)
			}
		}
		fmt.Println()
	}
	fmt.Println("================================")
	fmt.Println(dist[airSystem])
	var ans = 0
	var q = make([][3]int, 1)
	q[0] = [3]int{airSystem[0], airSystem[1], 0}
	for len(q) > 0 {
		var curr = q[0]
		var point = [2]int{curr[0], curr[1]}
		q = q[1:]
		dist[point] = -1
		if curr[2] > ans {
			ans = curr[2]
		}
		for i := 0; i < 4; i++ {
			newPos := [2]int{point[0] + directions[i][0], point[1] + directions[i][1]}
			newDist, ok := dist[newPos]
			if ok && newDist != -1 {
				q = append(q, [3]int{newPos[0], newPos[1], curr[2] + 1})
			}
		}
	}
	fmt.Println(ans)
}

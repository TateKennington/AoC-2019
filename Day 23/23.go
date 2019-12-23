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
	input        chan int
	output       chan int
}

func newComputer(codes map[int]int) *Computer {
	var codesCopy = make(map[int]int)
	for key, value := range codes {
		codesCopy[key] = value
	}
	return &Computer{codesCopy, 0, 0, false, make(chan int, 1000), make(chan int, 1000)}
}

func (this *Computer) compute() int {
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
			res := 0
			ok := true
			select {
			case res, ok = <-(this.input):
			default:
				res = -1
			}
			if !ok {
				this.halted = true
				return -1
			}
			if param == 0 {
				codes[codes[curr+1]] = res
			} else {
				codes[relativeBase+codes[curr+1]] = res
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
			step = 2
			this.output <- res
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
	var scn = bufio.NewScanner(os.Stdin)
	scn.Scan()
	var input = scn.Text()
	var codesText = strings.Split(input, ",")
	var codes = make(map[int]int)
	for i, code := range codesText {
		x, _ := strconv.Atoi(code)
		codes[i] = x
	}
	var computers = make([]*Computer, 50)
	for i := range computers {
		computers[i] = newComputer(codes)
		computers[i].input <- i
		go computers[i].compute()
		fmt.Printf("Started %d\n", i)
	}
	defer func() {
		for _, computer := range computers {
			close(computer.input)
		}
	}()
	lastSent := -1
	nat := false
	count := 0
	natX := 0
	natY := 0
	for {
		//time.Sleep(50 * time.Millisecond)
		idle := true
		for _, computer := range computers {
			select {
			case addr := <-computer.output:
				idle = false
				x := <-computer.output
				y := <-computer.output
				if addr == 255 {
					natX = x
					natY = y
					nat = true
					//fmt.Println(natY)
				} else {
					computers[addr].input <- x
					computers[addr].input <- y
				}
			default:
				continue
			}
		}
		if idle {
			count++
		} else {
			count = 0
		}
		if count > 10000 && nat {
			count = 0
			//nat = false
			fmt.Println(natY)
			if lastSent == natY {
				fmt.Println(natY)
				goto end
			} else {
				lastSent = natY
				computers[0].input <- natX
				computers[0].input <- natY
			}
		}
	}
end:
}

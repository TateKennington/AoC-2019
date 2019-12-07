package main;

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
);

type Computer struct{
	memory []int;
	curr int;
	phase int;
	halted bool;
	first bool;
}

func newComputer(_codes[] int) *Computer{
	res := Computer{};
	res.memory = make([]int, len(_codes));
	res.first = true;
	copy(res.memory, _codes);
	return &res;
}

func (this *Computer) compute(input int) int{
	step:= 0;
	var first = this.first;
	var codes = this.memory;
	for curr:=this.curr; curr<int(len(codes)); curr+=step{
		opcode := codes[curr] % 100;
		param := codes[curr] / 100;
		switch(opcode){
			case 99:
				this.halted = true;
				goto end;
			case 1:
				res:=0;
				if param % 10 == 0{
					res = codes[codes[curr+1]];
				} else {
					res = codes[curr+1];
				}

				if (param/10) % 10 == 0{
					res += codes[codes[curr+2]];
				} else {
					res += codes[curr+2];
				}
				codes[codes[curr+3]] = res;
				step = 4;
			case 2:
				res:=0;
				if param%10 == 0{
					res = codes[codes[curr+1]];
				} else {
					res = codes[curr+1];
				}

				if (param/10) %10 == 0{
					res *= codes[codes[curr+2]];
				} else {
					res *= codes[curr+2];
				}
				codes[codes[curr+3]] = res;
				step = 4;
			case 3:
				if first {
					codes[codes[curr+1]] = this.phase;
				} else {
					codes[codes[curr+1]] = input;
				}
				this.first = false;
				first = false;
				step = 2;
			case 4:
				this.curr = curr+2;
				this.first = first;
				if param == 0 {
					return codes[codes[curr+1]];
				} else {
					return codes[curr+1];
				}
			case 5:
				var cond bool;
				if param % 10 == 0{
					cond = codes[codes[curr+1]] != 0;
				} else {
					cond = codes[curr+1] != 0;
				}

				if cond {
					step = 0;
					if (param/10) % 10 == 0{
						curr = codes[codes[curr+2]];
					} else {
						curr = codes[curr+2];
					}
				} else {
					step = 3;
				}
			case 6:
				var cond bool;
				if param % 10 == 0{
					cond = codes[codes[curr+1]] == 0;
				} else {
					cond = codes[curr+1] == 0;
				}

				if cond {
					step = 0;
					if (param/10) % 10 == 0{
						curr = codes[codes[curr+2]];
					} else {
						curr = codes[curr+2];
					}
				} else {
					step = 3;
				}
			case 7:
				first := 0;
				second := 0;
				step = 4;
				if param % 10 == 0{
					first = codes[codes[curr+1]];
				} else {
					first = codes[curr+1];
				}

				if (param/10) % 10 == 0{
					second = codes[codes[curr+2]];
				} else {
					second = codes[curr+2];
				}

				if first < second {
					codes[codes[curr+3]] = 1;
				} else {
					codes[codes[curr+3]] = 0;
				}
			case 8:
				first := 0;
				second := 0;
				step = 4;
				if param % 10 == 0{
					first = codes[codes[curr+1]];
				} else {
					first = codes[curr+1];
				}

				if (param/10) % 10 == 0{
					second = codes[codes[curr+2]];
				} else {
					second = codes[curr+2];
				}

				if first == second {
					codes[codes[curr+3]] = 1;
				} else {
					codes[codes[curr+3]] = 0;
				}
		}
	}
	end:
	return -1;
}

func process(codes []int, used [5]bool, phase [5]int, curr int, computers []*Computer) int{
	if curr == 5{
		for i, _ := range computers{
			computers[i] = newComputer(codes);
			computers[i].phase = phase[i];
		}
		input := 0;
		for i:=0; !computers[i].halted; i=(i+1)%5{
			temp:=computers[i].compute(input);
			if computers[i].halted{
				break;
			}
			input = temp;
		}
		return input;
	}
	var res = 0;
	for i:=0; i<5; i++{
		if !used[i]{
			used[i] = true;
			phase[curr] = int(i+5);
			temp := process(codes, used, phase, curr+1, computers);
			used[i] = false;
			if temp>res{
				res = temp;
			}
		}
	}
	return res;
}

func main(){
	var scanner = bufio.NewScanner(os.Stdin);
	scanner.Scan();
	var input = scanner.Text();
	var codesText = strings.Split(input, ",");
	var codes []int;
	for _, code := range codesText {
		x,  _ := strconv.Atoi(code);
		codes = append(codes, int(x));
	}
	var used [5]bool;
	var phase [5]int;
	fmt.Println(process(codes, used, phase, 0, make([]*Computer, 5)));
}
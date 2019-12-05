package main;

import "fmt";
import "bufio";
import "os";
import "strings";
import "strconv";

func main(){
	var scanner = bufio.NewScanner(os.Stdin);
	scanner.Scan();
	var input = scanner.Text();
	var codesText = strings.Split(input, ",");
	var codes []int;
	var output []int;
	for _, code := range codesText {
		x,  _ := strconv.Atoi(code);
		codes = append(codes, x);
	}
	//codes[1] = noun;
	//codes[2] = verb;
	step:= 0;
	for curr:=0; curr<len(codes); curr+=step{
		//fmt.Println(codes);
		opcode := codes[curr] % 100;
		param := codes[curr] / 100;
		//fmt.Print(codes[curr]);
		//fmt.Print(": ");
		switch(opcode){
			case 99:
				goto end;
			case 1:
				res:=0
				if param % 10 == 0{
					res = codes[codes[curr+1]];
					fmt.Printf("%d ", codes[codes[curr+1]]);
				} else {
					res = codes[curr+1];
					fmt.Printf("%d ", codes[curr+1]);
				}

				if (param/10) % 10 == 0{
					res += codes[codes[curr+2]];
					fmt.Printf("%d ", codes[codes[curr+2]]);
				} else {
					res += codes[curr+2];
					fmt.Printf("%d ", codes[curr+2]);
				}
				codes[codes[curr+3]] = res;
				step = 4;
			case 2:
				res:=0
				if param%10 == 0{
					res = codes[codes[curr+1]];
					fmt.Printf("%d ", codes[codes[curr+1]]);
				} else {
					res = codes[curr+1];
					fmt.Printf("%d ", codes[curr+1]);
				}

				if (param/10) %10 == 0{
					res *= codes[codes[curr+2]];
					fmt.Printf("%d ", codes[codes[curr+2]]);
				} else {
					res *= codes[curr+2];
					fmt.Printf("%d ", codes[curr+2]);
				}
				codes[codes[curr+3]] = res;
				step = 4;
			case 3:
				codes[codes[curr+1]] = 5;
				fmt.Printf("%d ", codes[curr+1]);
				step = 2;
			case 4:
				if param == 0 {
					output = append(output, codes[codes[curr+1]]);
					fmt.Printf("%d ", codes[codes[curr+1]]);
				} else {
					output = append(output, codes[curr+1]);
					fmt.Printf("%d ", codes[curr+1]);
				}
				step = 2;
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
		//fmt.Println();
	}
	end:
	fmt.Println(codes);
	fmt.Println(output);
}
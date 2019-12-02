package main;

import "fmt";
import "bufio";
import "os";
import "strings";
import "strconv";

func main(){
	var scanner = bufio.NewScanner(os.Stdin);
	scanner.Scan();
	done := false;
	for noun:=0; noun<99 && !done; noun++{
		for verb:=0; verb<99 && !done; verb++{
			var input = scanner.Text();
			var codesText = strings.Split(input, ",");
			var codes []int;
			for _, code := range codesText {
				x,  _ := strconv.Atoi(code);
				codes = append(codes, x);
			}
			codes[1] = noun;
			codes[2] = verb;
			for curr:=0; curr<len(codes); curr+=4{
				switch(codes[curr]){
					case 99:
						break;
					case 1:
						codes[codes[curr+3]] = codes[codes[curr+1]] + codes[codes[curr+2]];
					case 2:
						codes[codes[curr+3]] = codes[codes[curr+1]] * codes[codes[curr+2]];
				}
			}
			if codes[0] == 19690720{
				fmt.Println(noun*100+verb);
				done = true;
			}
		}
	}
}
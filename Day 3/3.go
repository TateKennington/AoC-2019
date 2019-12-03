package main;

import "fmt";
import "bufio";
import "strconv";
import "os";
import "strings";


func main(){
	var kb = bufio.NewScanner(os.Stdin);
	var points = make(map[[2]int]int);
	var curr = [2]int{0,0};
	var step = [2]int{0,0};
	var count = 0;
	var ans = -1;
	var wire1 string;
	var wire2 string;

	kb.Scan();
	wire1 = kb.Text();
	kb.Scan();
	wire2 = kb.Text();

	for _, x := range strings.Split(wire1, ","){
		dist, _ := strconv.Atoi(x[1:]);

		switch(x[0]){
			case 'R':
				step[0] = 1;
				step[1] = 0;
			case 'L':
				step[0] = -1;
				step[1] = 0;
			case 'U':
				step[0] = 0;
				step[1] = 1;
			case 'D':
				step[0] = 0;
				step[1] = -1;
		}

		for i:=0; i<dist; i++{
			count++;
			curr[0]+=step[0];
			curr[1]+=step[1];
			points[curr] = count;
		}
	}

	curr = [2]int{0,0};
	count = 0;
	for _, x := range strings.Split(wire2, ","){
		dist, _ := strconv.Atoi(x[1:]);

		switch(x[0]){
			case 'R':
				step[0] = 1;
				step[1] = 0;
			case 'L':
				step[0] = -1;
				step[1] = 0;
			case 'U':
				step[0] = 0;
				step[1] = 1;
			case 'D':
				step[0] = 0;
				step[1] = -1;
		}

		for i:=0; i<dist; i++{
			count++;
			curr[0]+=step[0];
			curr[1]+=step[1];
			if temp, ok := points[curr]; ok{
				if ans == -1 || temp + count < ans {
					ans = temp+count;
				}
			}
		}
	}
	fmt.Println(ans);
}
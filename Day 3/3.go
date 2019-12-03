package main;

import "fmt";
import "bufio";
import "strconv";
import "os";
import "strings";

func main(){
	var kb = bufio.NewScanner(os.Stdin);
	kb.Scan();
	var wire1 = kb.Text();
	var points = make(map[[2]int]int);
	var curr = [2]int{0,0};
	var count = 0;
	for _, x := range strings.Split(wire1, ","){
		var dir = x[0];
		dist, _ := strconv.Atoi(x[1:]);
		if dir == 'R'{
			for i:=0; i<dist; i++{
				count++;
				curr[0]++;
				points[curr] = count;
			}
		}
		if dir == 'L'{
			for i:=0; i<dist; i++{
				count++;
				curr[0]--;
				points[curr] = count;
			}
		}
		if dir == 'U'{
			for i:=0; i<dist; i++{
				count++;
				curr[1]++;
				points[curr] = count;
			}
		}
		if dir == 'D'{
			for i:=0; i<dist; i++{
				count++;
				curr[1]--;
				points[curr] = count;
			}
		}
	}
	kb.Scan();
	var wire2 = kb.Text();
	curr = [2]int{0,0};
	var ans = -1;
	count = 0;
	for _, x := range strings.Split(wire2, ","){
		var dir = x[0];
		dist, _ := strconv.Atoi(x[1:]);
		if dir == 'R'{
			for i:=0; i<dist; i++{
				count++;
				curr[0]++;
				if temp, ok := points[curr]; ok{
					if ans == -1 || temp+count<ans{
						ans = temp+count;
					}
				}
			}
		}
		if dir == 'L'{
			for i:=0; i<dist; i++{
				count++;
				curr[0]--;
				if temp, ok := points[curr]; ok{
					if ans == -1 || temp+count<ans{
						ans = temp+count;
					}
				}
			}
		}
		if dir == 'U'{
			for i:=0; i<dist; i++{
				count++;
				curr[1]++;
				if temp, ok := points[curr]; ok{
					if ans == -1 || temp+count<ans{
						ans = temp+count;
					}
				}
			}
		}
		if dir == 'D'{
			for i:=0; i<dist; i++{
				count++;
				curr[1]--;
				if temp, ok := points[curr]; ok{
					if ans == -1 || temp+count<ans{
						ans = temp+count;
					}
				}
			}
		}
	}
	fmt.Println(ans);
}
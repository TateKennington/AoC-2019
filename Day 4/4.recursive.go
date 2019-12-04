package main;

import "fmt";

const LOWER = 246540;
const UPPER = 787419;

func isValid(n int) bool{
	var adj = false;
	var prev = 10;
	for n > 0{
		curr := n%10;
		next := (n/10)%10;
		count := 1;
		for n > 0 && curr == next{
			count++;
			n/=10;
			curr = n%10;
			next = (n/10)%10;
		}
		if count == 2{
			adj = true;
		}
		if curr > prev{
			return false;
		}
		prev = curr;
		n/=10;
	}
	return adj;
}

func process(curr int, buffer [6]int) int{
	if curr >= 6{
		var num = 0;
		for _, x := range buffer{
			num*=10;
			num+=x;
		}
		if num>=LOWER && num<=UPPER && isValid(num){
			return 1;
		}
		return 0;
	}
	var res = 0;
	var lower = 0;
	if curr-1>=0{
		lower = buffer[curr-1];
	}
	buffer[curr] = lower
	for buffer[curr] <= 9{
		res+=process(curr+1, buffer);
		buffer[curr]++;
	}
	return res;
}

func main(){
	fmt.Println(process(0, [6]int{0,0,0,0,0,0}));
}
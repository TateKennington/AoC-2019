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

func main(){
	var ans = 0;
	for i:=LOWER; i<=UPPER; i++{
		if isValid(i){
			ans++;
		}
	}
	fmt.Println(ans);
}
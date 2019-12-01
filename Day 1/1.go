package main;

import "fmt";

func main(){
	var x = 0;
	var ans = 0;
	for{
		n, _ := fmt.Scanf("%d\n", &x);
		
		if n==0{
			break;
		}

		for x/3-2 > 0{
			ans+=x/3-2;
			x = x/3-2;
		}
	}
	fmt.Println(ans);
}
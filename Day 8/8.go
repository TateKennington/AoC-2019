package main;

import (
	"fmt"
	"os"
	"bufio"
);

const WIDTH = 25;
const HEIGHT = 6; 
const IMAGE_SIZE = WIDTH*HEIGHT

func main(){
	var scn = bufio.NewScanner(os.Stdin);
	var layerStart = 0;
	var ans = 0;
	var minZeros = IMAGE_SIZE;
	var decoded [IMAGE_SIZE]int;

	scn.Scan();
	var image = scn.Text();
	for i, _ := range decoded{
		decoded[i] = 2;
	}

	for layerStart<len(image){
		oneCount:=0;
		twoCount:=0;
		zeroCount:=0;
		for offset :=0; offset<IMAGE_SIZE; offset++{
			switch(image[layerStart+offset]){
			case '0':
				if decoded[offset] == 2{
					decoded[offset] = 0;
				}
				zeroCount++;
			case '1':
				if decoded[offset] == 2{
					decoded[offset] = 1;
				}
				oneCount++;
			case '2':
				twoCount++;
			}
		}
		if zeroCount < minZeros{
			minZeros = zeroCount;
			ans = oneCount*twoCount;
		}
		layerStart+=WIDTH*HEIGHT;
	}
	for i:=0; i<HEIGHT; i++{
		for j:=0; j<WIDTH; j++{
			fmt.Print(decoded[i*WIDTH+j]);
		}
		fmt.Println();
	}
	fmt.Println(ans);
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func next(in []int) []int {
	var res = make([]int, len(in))
	var prefix = make([]int, len(in)+1)
	var phases = [4]int{0, 1, 0, -1}
	var sum = 0
	for i := range in {
		prefix[i] = sum
		sum += in[i]
	}
	prefix[len(in)] = sum
	for i := range in {
		var curr = 0
		var phase = 1
		for j := i; j < len(in); j += (i + 1) {
			var right = sum
			if j+i+1 < len(prefix) {
				right = prefix[j+i+1]
			}
			curr += (right - prefix[j]) * phases[phase]
			phase = (phase + 1) % 4
		}
		if curr > 0 {
			res[i] = curr % 10
		} else {
			res[i] = -1 * curr % 10
		}
	}
	return res
}

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	scn.Scan()
	var input = scn.Text()
	var signal = make([]int, 0, 1000*len(input))
	for x := 0; x < 10000; x++ {
		for i := range input {
			signal = append(signal, int(input[i]-'0'))
		}
	}
	var offset = 0
	for i := 0; i < 7; i++ {
		offset *= 10
		offset += signal[i]
	}
	for i := 0; i < 100; i++ {
		fmt.Println(i)
		signal = next(signal)
	}
	fmt.Println(signal[offset : offset+8])
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func colinear(gradiants [][2]int, gradiant [2]int) bool {
	for _, vector := range gradiants {
		if (vector[0]*gradiant[0]+vector[1]*gradiant[1])*(vector[0]*gradiant[0]+vector[1]*gradiant[1]) == (vector[0]*vector[0]+vector[1]*vector[1])*(gradiant[0]*gradiant[0]+gradiant[1]*gradiant[1]) {
			if vector[0]*gradiant[0]+vector[1]*gradiant[1] >= 0 {
				return true
			}
		}
	}
	return false
}

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	var x = 0
	var y = 0
	var asteroids [][2]int
	for scn.Scan() {
		var line = scn.Text()
		fmt.Println(line)
		x = 0
		for _, c := range line {
			if c == '#' {
				asteroids = append(asteroids, [2]int{x, y})
			}
			x++
		}
		y++
	}
	var ans = 0
	var best = [2]int{-1, -1}
	var index = 0
	for i, asteroid := range asteroids {
		count := 0
		var gradiants [][2]int
		for j, other := range asteroids {
			gradiant := [2]int{asteroid[0] - other[0], asteroid[1] - other[1]}
			if i != j && !colinear(gradiants, gradiant) {
				count++
				gradiants = append(gradiants, gradiant)
			}
		}
		if count > ans {
			ans = count
			best = asteroid
			index = i
		}
	}
	asteroids[index] = asteroids[len(asteroids)-1]
	asteroids = asteroids[:len(asteroids)-1]
	sort.Slice(asteroids, func(i, j int) bool {
		abdet := (asteroids[i][0]-best[0])*(asteroids[j][1]-best[1]) - (asteroids[i][1]-best[1])*(asteroids[j][0]-best[0])
		nadet := (asteroids[i][0] - best[0])
		nbdet := (asteroids[j][0] - best[0])

		if abdet == 0 {
			return (asteroids[i][1]-best[1])*(asteroids[i][1]-best[1])+(asteroids[i][0]-best[0])*(asteroids[i][0]-best[0]) < (asteroids[j][1]-best[1])*(asteroids[j][1]-best[1])+(asteroids[j][0]-best[0])*(asteroids[j][0]-best[0])
		}

		if nadet == 0 {
			if asteroids[i][1]-best[1] < 0 {
				return true
			}
			return nbdet < 0
		}

		if nbdet == 0 {
			if asteroids[j][1]-best[1] < 0 {
				return false
			}
			return nadet > 0
		}

		if abdet > 0 && nadet > 0 && nbdet > 0 {
			return true
		}
		if abdet < 0 && nadet < 0 && nbdet < 0 {
			return false
		}
		if abdet < 0 && nadet > 0 && nbdet < 0 {
			return true
		}
		if abdet > 0 && nadet < 0 && nbdet > 0 {
			return false
		}
		if abdet < 0 && nadet > 0 && nbdet > 0 {
			return false
		}
		if abdet > 0 && nadet > 0 && nbdet > 0 {
			return true
		}
		if abdet > 0 && nadet < 0 && nbdet < 0 {
			return true
		}
		if abdet > 0 && nadet > 0 && nbdet < 0 {
			return true
		}
		if abdet < 0 && nadet < 0 && nbdet > 0 {
			return false
		}
		return true
	})
	var curr = 0
	var prev = asteroids[curr]
	var order [][2]int
	var destroyed = make([]bool, len(asteroids))
	destroyed[index] = true
	for len(order) < 200 {
		prev = asteroids[curr]
		order = append(order, prev)
		destroyed[curr] = true
		for asteroids[curr] == best || destroyed[curr] || colinear([][2]int{[2]int{prev[0] - best[0], prev[1] - best[1]}}, [2]int{asteroids[curr][0] - best[0], asteroids[curr][1] - best[1]}) {
			curr = (curr + 1) % len(asteroids)
			if asteroids[curr] == prev {
				curr = (curr + 1) % len(asteroids)
				break
			}
		}
	}
	fmt.Println(order)
	fmt.Println(ans)
}

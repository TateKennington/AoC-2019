package main

import (
	"bufio"
	"fmt"
	"os"
)

func printGrid(grid [5][5]bool) {
	for _, line := range grid {
		for _, cell := range line {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	var grids = make(map[int][5][5]bool)
	var counts = make(map[int][5][5]int)
	var grid = grids[0]
	for i := 0; scn.Scan(); i++ {
		line := scn.Text()
		for j, c := range line {
			if c == '.' {
				grid[i][j] = false
			} else {
				grid[i][j] = true
			}
		}
	}
	grids[0] = grid
	for x := 0; x < 200; x++ {
		for key, grid := range grids {
			grid[2][2] = false
			count := counts[key]
			for i := range grid {
				for j := range grid[i] {
					if grid[i][j] {
						if i-1 >= 0 {
							count[i-1][j]++
						}
						if j-1 >= 0 {
							count[i][j-1]++
						}
						if i+1 < len(grid) {
							count[i+1][j]++
						}
						if j+1 < len(grid[i]) {
							count[i][j+1]++
						}
					}
				}
			}
			counts[key] = count
		}
		for key := range grids {
			if _, ok := grids[key-1]; !ok {
				var temp [5][5]bool
				grids[key-1] = temp
			}
			if _, ok := grids[key+1]; !ok {
				var temp [5][5]bool
				grids[key+1] = temp
			}
		}
		oldGrids := make(map[int][5][5]bool)
		for key, value := range grids {
			oldGrids[key] = value
		}
		for key, grid := range grids {
			count := counts[key]
			upperGrid := oldGrids[key-1]
			lowerGrid := oldGrids[key+1]
			grid[2][2] = false
			for i := range grid {
				for j := range grid[i] {
					var adj = count[i][j]
					if i == 0 && upperGrid[1][2] {
						adj++
					}
					if i == 4 && upperGrid[3][2] {
						adj++
					}
					if j == 0 && upperGrid[2][1] {
						adj++
					}
					if j == 4 && upperGrid[2][3] {
						adj++
					}
					if i == 1 && j == 2 {
						for k := 0; k < 5; k++ {
							if lowerGrid[0][k] {
								adj++
							}
						}
					}
					if i == 3 && j == 2 {
						for k := 0; k < 5; k++ {
							if lowerGrid[4][k] {
								adj++
							}
						}
					}
					if i == 2 && j == 1 {
						for k := 0; k < 5; k++ {
							if lowerGrid[k][0] {
								adj++
							}
						}
					}
					if i == 2 && j == 3 {
						for k := 0; k < 5; k++ {
							if lowerGrid[k][4] {
								adj++
							}
						}
					}
					if grid[i][j] && adj != 1 {
						grid[i][j] = false
					} else if !grid[i][j] && (adj == 1 || adj == 2) {
						if i != 2 || j != 2 {
							grid[i][j] = true
						}
					}
				}
			}
			grids[key] = grid
		}
		for key := range counts {
			var temp [5][5]int
			counts[key] = temp
		}
	}
	var ans = 0
	for _, grid := range grids {
		for _, line := range grid {
			for _, cell := range line {
				if cell {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

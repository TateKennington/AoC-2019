package main

import (
	"fmt"
	"math/big"
)

type key struct {
	position int
	velocity int
}

func gcd(a, b int64) int64 {
	var r int64
	for b != 0 {
		r = a % b
		a = b
		b = r
	}
	return a
}

func main() {
	var moons [4][3]int
	var velocities [4][3]int
	var periods [3]int
	fmt.Scanf("<x=%d, y=%d, z=%d>\n", &moons[0][0], &moons[0][1], &moons[0][2])
	fmt.Scanf("<x=%d, y=%d, z=%d>\n", &moons[1][0], &moons[1][1], &moons[1][2])
	fmt.Scanf("<x=%d, y=%d, z=%d>\n", &moons[2][0], &moons[2][1], &moons[2][2])
	fmt.Scanf("<x=%d, y=%d, z=%d>\n", &moons[3][0], &moons[3][1], &moons[3][2])
	var init [4][3]int
	for i := range moons {
		for j := range velocities[i] {
			init[i][j] = moons[i][j]
		}
	}
	var x int
	var done = false
	for x = 0; !done; x++ {
		done = true
		if periods[0] == 0 && moons[0][0] == init[0][0] && moons[1][0] == init[1][0] && moons[2][0] == init[2][0] && moons[3][0] == init[3][0] && velocities[0][0] == 0 && velocities[1][0] == 0 && velocities[2][0] == 0 && velocities[3][0] == 0 {
			periods[0] = x
		}
		done = done && periods[0] != 0
		if periods[1] == 0 && moons[0][1] == init[0][1] && moons[1][1] == init[1][1] && moons[2][1] == init[2][1] && moons[3][1] == init[3][1] && velocities[0][1] == 0 && velocities[1][1] == 0 && velocities[2][1] == 0 && velocities[3][1] == 0 {
			periods[1] = x
		}
		done = done && periods[1] != 0
		if periods[2] == 0 && moons[0][2] == init[0][2] && moons[1][2] == init[1][2] && moons[2][2] == init[2][2] && moons[3][2] == init[3][2] && velocities[0][2] == 0 && velocities[1][2] == 0 && velocities[2][2] == 0 && velocities[3][2] == 0 {
			periods[2] = x
		}
		done = done && periods[2] != 0
		for i, _ := range moons {
			for j, _ := range moons {
				if i != j {
					if moons[i][0] < moons[j][0] {
						velocities[i][0] += 1
					} else if moons[i][0] > moons[j][0] {
						velocities[i][0] -= 1
					}

					if moons[i][1] < moons[j][1] {
						velocities[i][1] += 1
					} else if moons[i][1] > moons[j][1] {
						velocities[i][1] -= 1
					}

					if moons[i][2] < moons[j][2] {
						velocities[i][2] += 1
					} else if moons[i][2] > moons[j][2] {
						velocities[i][2] -= 1
					}
				}
			}
		}
		for i, _ := range moons {
			for j, _ := range velocities[i] {
				moons[i][j] += velocities[i][j]
			}
		}
	}
	var ans = 0
	for i := range moons {
		var pot = 0
		var kin = 0
		for j := range velocities[i] {
			if moons[i][j] > 0 {
				pot += moons[i][j]
			} else {
				pot -= moons[i][j]
			}
			if velocities[i][j] > 0 {
				kin += velocities[i][j]
			} else {
				kin -= velocities[i][j]
			}
		}
		ans += pot * kin
	}
	fmt.Println(ans)
	fmt.Println(periods)
	lcm := big.NewInt(1)
	for _, x := range periods {
		temp := big.NewInt(1).GCD(nil, nil, lcm, big.NewInt(int64(x)))
		lcm.Mul(lcm, big.NewInt(int64(x)))
		//fmt.Println(lcm)
		//fmt.Println(temp)
		lcm.Div(lcm, temp)
	}
	fmt.Println(lcm)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type reagent struct {
	name     string
	quantity int
}

var oreMemo map[reagent]int

func oreCalc(name string, quantity int, chemicals map[string][]reagent, excess map[string]int) int {
	var formula = chemicals[name]
	var times = 0
	if excess[formula[0].name] > 0 {
		if excess[formula[0].name] <= quantity {
			quantity -= excess[formula[0].name]
			excess[formula[0].name] = 0
		}
		if excess[formula[0].name] > quantity {
			excess[formula[0].name] -= quantity
			quantity = 0
		}
	}
	for formula[0].quantity*times < quantity {
		times++
	}
	var res = 0
	for _, reagent := range formula[1:] {
		if reagent.name == "ORE" {
			res += reagent.quantity * times
		} else {
			res += oreCalc(reagent.name, reagent.quantity*times, chemicals, excess)
		}
	}
	if quantity < formula[0].quantity*times {
		excess[formula[0].name] += formula[0].quantity*times - quantity
	}
	//fmt.Printf("%s x %d : %d\n", name, quantity, res)
	return res
}

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	var chemicals = make(map[string][]reagent)
	oreMemo = make(map[reagent]int)
	for scn.Scan() {
		input := strings.Split(scn.Text(), " => ")
		result := strings.Split(input[1], " ")
		amount, _ := strconv.Atoi(result[0])
		chemical := reagent{result[1], amount}
		for _, term := range strings.Split(input[0], ", ") {
			temp := strings.Split(term, " ")
			quantity, _ := strconv.Atoi(temp[0])
			if _, ok := chemicals[result[1]]; !ok {
				chemicals[result[1]] = make([]reagent, 0, 2)
				chemicals[result[1]] = append(chemicals[result[1]], chemical)
			}
			chemicals[result[1]] = append(chemicals[result[1]], reagent{temp[1], quantity})
		}
	}
	var excess = make(map[string]int)
	//orePerFuel := oreCalc("FUEL", 1, chemicals, make(map[string]int))
	ore := int64(1000000000000)
	ans := int64(0)
	for amount := 100000; amount > 0; amount /= 10 {
		for orePer := int64(oreCalc("FUEL", amount, chemicals, excess)); orePer <= ore; orePer = int64(oreCalc("FUEL", amount, chemicals, excess)) {
			ore -= orePer
			ans += int64(amount)
		}
	}
	fmt.Println(ans)

}

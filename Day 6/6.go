package main;

import (
	"fmt"
	"bufio"
	"os"
	"strings"
);

type Node struct{
	children []*Node;
	parent *Node;
	level int;
};

func lowestCommonAncestor(a, b *Node) *Node{
	var parents = make(map[*Node]bool);
	for a != nil{
		parents[a] = true;
		a = a.parent;
	}
	for !parents[b]{
		b = b.parent;
	}
	return b;
}

func orbitalDistance(a, b, c *Node) int{
	return (a.parent.level-c.level)+(b.parent.level-c.level);
}

func main(){
	var ans = 0;
	var scn = bufio.NewScanner(os.Stdin);
	var star = make(map[string]*Node);
	for scn.Scan(){
		var input = strings.Split(scn.Text(), ")");
		if star[input[0]] == nil{
			var temp Node;
			star[input[0]] = &temp;
		}
		if star[input[1]] == nil{
			var temp Node;
			star[input[1]] = &temp;
		}
		star[input[0]].children = append(star[input[0]].children, star[input[1]]);
		star[input[1]].parent = star[input[0]];
	}
	var queue = []*Node{star["COM"]};
	queue[0].level = 0;
	queue[0].parent = nil;
	for len(queue) > 0{
		var curr = queue[0];
		ans+=curr.level;
		for _, node := range curr.children{
			node.level = curr.level+1;
			queue = append(queue, node);
		}
		queue = queue[1:];
	}
	fmt.Println(ans);
	fmt.Println(orbitalDistance(star["YOU"], star["SAN"], lowestCommonAncestor(star["YOU"], star["SAN"])));
}
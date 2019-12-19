package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

// An Item is something we manage in a priority queue.
type Item struct {
	value    state // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value state, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type edge struct {
	dist int
	mask int
	dest [2]int
}

type state struct {
	pos  [2]int
	mask int
	dist int
}

func main() {
	var scn = bufio.NewScanner(os.Stdin)
	var grid = make([][]rune, 0)
	var poses = make([][2]int, 0, 4)
	var keys = make([][2]int, 0)
	for i := 0; scn.Scan(); i++ {
		line := scn.Text()
		grid = append(grid, make([]rune, 0, len(line)))
		for j, c := range line {
			if c == '@' {
				poses = append(poses, [2]int{i, j})
				c = '.'
			} else if c != '.' && c != '#' {
				if c >= 'a' && c <= 'z' {
					keys = append(keys, [2]int{i, j})
				}
			}
			grid[i] = append(grid[i], c)
		}
	}
	var adj = make([][]edge, len(keys))
	for i := range adj {
		adj[i] = make([]edge, len(keys))
	}
	for _, k := range keys {
		var visited = make(map[[2]int]bool)
		seen := 0
		//fmt.Println("bfs")
		for queue := [][4]int{[4]int{k[0], k[1], 0, 0}}; len(queue) > 0; {
			var curr = queue[0]
			queue = queue[1:]
			var pos = [2]int{curr[0], curr[1]}
			if grid[pos[0]][pos[1]] >= 'a' && grid[pos[0]][pos[1]] <= 'z' {
				adj[grid[k[0]][k[1]]-'a'][grid[pos[0]][pos[1]]-'a'] = edge{curr[2], curr[3], pos}
				seen = seen | (1 << uint(grid[pos[0]][pos[1]]-'a'))
				var done = true
				for i := 0; i < len(keys) && done; i++ {
					if seen&(1<<uint(i)) == 0 {
						done = false
					}
				}
				if done {
					break
				}
			}

			if grid[pos[0]][pos[1]] >= 'A' && grid[pos[0]][pos[1]] <= 'Z' {
				curr[3] = curr[3] | (1 << uint(grid[pos[0]][pos[1]]-'A'))
			}

			if pos[0]-1 >= 0 && grid[pos[0]-1][pos[1]] != '#' && !visited[[2]int{pos[0] - 1, pos[1]}] {
				queue = append(queue, [4]int{pos[0] - 1, pos[1], curr[2] + 1, curr[3]})
			}
			if pos[1]-1 >= 0 && grid[pos[0]][pos[1]-1] != '#' && !visited[[2]int{pos[0], pos[1] - 1}] {
				queue = append(queue, [4]int{pos[0], pos[1] - 1, curr[2] + 1, curr[3]})
			}
			if pos[0]+1 <= len(grid) && grid[pos[0]+1][pos[1]] != '#' && !visited[[2]int{pos[0] + 1, pos[1]}] {
				queue = append(queue, [4]int{pos[0] + 1, pos[1], curr[2] + 1, curr[3]})
			}
			if pos[1]+1 <= len(grid[pos[0]]) && grid[pos[0]][pos[1]+1] != '#' && !visited[[2]int{pos[0], pos[1] + 1}] {
				queue = append(queue, [4]int{pos[0], pos[1] + 1, curr[2] + 1, curr[3]})
			}
			visited[pos] = true
		}
		adj[grid[k[0]][k[1]]-'a'][grid[k[0]][k[1]]-'a'] = edge{0, 1 << uint(len(keys)+1), k}
	}
	ans := 0
	for _, pos := range poses {
		var queue = make(PriorityQueue, 0)
		heap.Init(&queue)
		var visited = make(map[[2]int]bool)
		quadrantMask := 0
		for i := range keys {
			quadrantMask = quadrantMask | (1 << uint(i))
		}
		for q := [][4]int{[4]int{pos[0], pos[1], 0, 0}}; len(q) > 0; {
			var curr = q[0]
			q = q[1:]
			var pos = [2]int{curr[0], curr[1]}
			if grid[pos[0]][pos[1]] >= 'a' && grid[pos[0]][pos[1]] <= 'z' {
				quadrantMask = quadrantMask ^ (1 << uint(grid[pos[0]][pos[1]]-'a'))
			}
			if pos[0]-1 >= 0 && grid[pos[0]-1][pos[1]] != '#' && !visited[[2]int{pos[0] - 1, pos[1]}] {
				q = append(q, [4]int{pos[0] - 1, pos[1], curr[2] + 1, curr[3]})
			}
			if pos[1]-1 >= 0 && grid[pos[0]][pos[1]-1] != '#' && !visited[[2]int{pos[0], pos[1] - 1}] {
				q = append(q, [4]int{pos[0], pos[1] - 1, curr[2] + 1, curr[3]})
			}
			if pos[0]+1 <= len(grid) && grid[pos[0]+1][pos[1]] != '#' && !visited[[2]int{pos[0] + 1, pos[1]}] {
				q = append(q, [4]int{pos[0] + 1, pos[1], curr[2] + 1, curr[3]})
			}
			if pos[1]+1 <= len(grid[pos[0]]) && grid[pos[0]][pos[1]+1] != '#' && !visited[[2]int{pos[0], pos[1] + 1}] {
				q = append(q, [4]int{pos[0], pos[1] + 1, curr[2] + 1, curr[3]})
			}
			visited[pos] = true
		}
		visited = make(map[[2]int]bool)
		for q := [][4]int{[4]int{pos[0], pos[1], 0, 0}}; len(q) > 0; {
			var curr = q[0]
			q = q[1:]
			var pos = [2]int{curr[0], curr[1]}
			if grid[pos[0]][pos[1]] >= 'a' && grid[pos[0]][pos[1]] <= 'z' {
				heap.Push(&queue, &Item{state{pos, quadrantMask | (1 << uint(grid[pos[0]][pos[1]]-'a')), curr[2]}, curr[2], 0})
			}

			if grid[pos[0]][pos[1]] >= 'A' && grid[pos[0]][pos[1]] <= 'Z' {
				continue
			}

			if pos[0]-1 >= 0 && grid[pos[0]-1][pos[1]] != '#' && !visited[[2]int{pos[0] - 1, pos[1]}] {
				q = append(q, [4]int{pos[0] - 1, pos[1], curr[2] + 1, curr[3]})
			}
			if pos[1]-1 >= 0 && grid[pos[0]][pos[1]-1] != '#' && !visited[[2]int{pos[0], pos[1] - 1}] {
				q = append(q, [4]int{pos[0], pos[1] - 1, curr[2] + 1, curr[3]})
			}
			if pos[0]+1 <= len(grid) && grid[pos[0]+1][pos[1]] != '#' && !visited[[2]int{pos[0] + 1, pos[1]}] {
				q = append(q, [4]int{pos[0] + 1, pos[1], curr[2] + 1, curr[3]})
			}
			if pos[1]+1 <= len(grid[pos[0]]) && grid[pos[0]][pos[1]+1] != '#' && !visited[[2]int{pos[0], pos[1] + 1}] {
				q = append(q, [4]int{pos[0], pos[1] + 1, curr[2] + 1, curr[3]})
			}
			visited[pos] = true
		}

		fmt.Println("Built")

		var v = make(map[state]int)
		for {
			var curr = heap.Pop(&queue).(*Item).value
			dist := curr.dist
			curr.dist = -1
			if v[curr] == 0 || v[curr] > dist {
				v[curr] = dist
			}
			//fmt.Println(curr)

			var done = true
			for i := 0; i < len(keys) && done; i++ {
				if curr.mask&(1<<uint(i)) == 0 {
					done = false
				}
			}
			if done {
				ans += dist
				fmt.Println(dist)
				break
			}

			for _, edge := range adj[grid[curr.pos[0]][curr.pos[1]]-'a'] {
				newMask := curr.mask | (1 << uint(grid[edge.dest[0]][edge.dest[1]]-'a'))
				if edge.mask&curr.mask == edge.mask && curr.mask != newMask && (v[state{edge.dest, newMask, -1}] == 0 || v[state{edge.dest, newMask, -1}] > dist+edge.dist) {
					v[state{edge.dest, newMask, -1}] = dist + edge.dist
					heap.Push(&queue, &Item{state{edge.dest, newMask, dist + edge.dist}, dist + edge.dist, 0})
				}
			}
		}
	}
	fmt.Println(ans)
}

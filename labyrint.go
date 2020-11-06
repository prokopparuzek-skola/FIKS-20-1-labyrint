package main

import "fmt"

type vertex struct {
	x     int
	y     int
	END   bool
	up    bool
	right bool
	down  bool
	left  bool
}
type steps struct {
	V      bool
	parent int
	step   rune
}

const (
	TILE     rune = '.'
	WALL     rune = '#'
	START    rune = 's'
	END      rune = 'x'
	MONSTRUM rune = 'M'
)

func solve(maze []vertex, start int, C int) (out string) {
	var Nsteps, Fsteps []int
	var work []steps
	var step int
	Nsteps = make([]int, 0)
	Nsteps = append(Nsteps, start)
	Fsteps = make([]int, 0)
	work = make([]steps, len(maze))
	for len(Nsteps) != 0 {
		for _, s := range Nsteps {
			if maze[s].END {
				var rev string
				for t := s; t != start; t = work[t].parent {
					rev += string(work[t].step)
				}
				for i := len(rev) - 1; i >= 0; i-- {
					out += string(rev[i])
				}
				return
			}
			x := s % C
			y := s / C
			if maze[s].up && !work[s-C].V { // UP
				Fsteps = append(Fsteps, (y-1)*C+x)
				work[s-C].V = true
				work[s-C].step = '^'
				work[s-C].parent = s
			}
			if maze[s].right && !work[s+1].V { // RIGHT
				Fsteps = append(Fsteps, y*C+x+1)
				work[s+1].V = true
				work[s+1].step = '>'
				work[s+1].parent = s
			}
			if maze[s].down && !work[s+C].V { // DOWN
				Fsteps = append(Fsteps, (y+1)*C+x)
				work[s+C].V = true
				work[s+C].step = 'v'
				work[s+C].parent = s
			}
			if maze[s].left && !work[s-1].V { // LEFT
				Fsteps = append(Fsteps, y*C+x-1)
				work[s-1].V = true
				work[s-1].step = '<'
				work[s-1].parent = s
			}
		}
		Nsteps = Fsteps
		Fsteps = make([]int, 0)
		step++
	}
	return "tak to je konec"
}

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ { // Načte všechny vstupy
		var R, C int
		var lab []rune
		var maze []vertex
		var start int
		fmt.Scan(&R, &C)
		lab = make([]rune, R*C)
		maze = make([]vertex, R*C)
		for j := 0; j < R; j++ { // načte řádky
			var line string
			fmt.Scanln(&line)
			for k, c := range line { // vyplní pole
				lab[j*C+k] = c
			}
		}
		//fmt.Println(lab)
		for j, c := range lab {
			x := j % C
			y := j / C
			if c != WALL {
				maze[j].x = x
				maze[j].y = y
				if y > 0 { // UP
					if lab[(y-1)*C+x] != WALL {
						maze[j].up = true
					}
				}
				if x < C-1 { // RIGHT
					if lab[y*C+x+1] != WALL {
						maze[j].right = true
					}
				}
				if y < R-1 { // DOWN
					if lab[(y+1)*C+x] != WALL {
						maze[j].down = true
					}
				}
				if x > 0 { // LEFT
					if lab[y*C+x-1] != WALL {
						maze[j].left = true
					}
				}
			}
			if c == END {
				maze[j].END = true
			} else if c == START {
				start = j
			}
		}
		//fmt.Println(maze, start)
		fmt.Println(solve(maze, start, C))
	}
}

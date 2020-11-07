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
	tile  bool
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

func bfs(maze []vertex, start int, C int) (out string) {
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
				for t := s; t != start; t = work[t].parent {
					out += string(work[t].step)
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
	return "cenok ej ot kat"
}

func MMove(maze []vertex, C, start, end int) (move int) {
	var Nsteps, Fsteps []int
	var work []steps
	Nsteps = make([]int, 0)
	Nsteps = append(Nsteps, start)
	Fsteps = make([]int, 0)
	work = make([]steps, len(maze))
	for len(Nsteps) != 0 {
		for _, s := range Nsteps {
			if s == end {
				for t := s; t != start; t = work[t].parent {
					move = t
				}
				return
			}
			x := s % C
			y := s / C
			if maze[s].up && !work[s-C].V { // UP
				Fsteps = append(Fsteps, (y-1)*C+x)
				work[s-C].V = true
				work[s-C].parent = s
			}
			if maze[s].down && !work[s+C].V { // DOWN
				Fsteps = append(Fsteps, (y+1)*C+x)
				work[s+C].V = true
				work[s+C].parent = s
			}
			if maze[s].left && !work[s-1].V { // LEFT
				Fsteps = append(Fsteps, y*C+x-1)
				work[s-1].V = true
				work[s-1].parent = s
			}
			if maze[s].right && !work[s+1].V { // RIGHT
				Fsteps = append(Fsteps, y*C+x+1)
				work[s+1].V = true
				work[s+1].parent = s
			}
		}
		Nsteps = Fsteps
		Fsteps = make([]int, 0)
	}
	return -1
}

func dfs(maze []vertex, s, C, monster int, work []int, nav [][]int) (out string, E bool) {
	defer fmt.Println(")")
	fmt.Println("(", s, maze[s], monster, work[s])
	var mnext int
	var near bool = true
	work[s]++
	if maze[s].END {
		return "", true
	}
	if monster == s {
		work[s]--
		return "m", false
	}
	if maze[s].up && work[s-C] == 0 && s-C != monster { // UP
		mnext = nav[monster][s-C]
		steps, e := dfs(maze, s-C, C, mnext, work, nav)
		if e {
			if len(out) == 0 || len(out)-1 > len(steps) {
				out = steps + "^"
				E = true
			}
		}
		near = false
	}
	if maze[s].right && work[s+1] == 0 && s+1 != monster { // RIGHT
		mnext = nav[monster][s+1]
		steps, e := dfs(maze, s+1, C, mnext, work, nav)
		if e {
			if len(out) == 0 || len(out)-1 > len(steps) {
				out = steps + ">"
				E = true
			}
		}
		near = false
	}
	if maze[s].down && work[s+C] == 0 && s+C != monster { // DOWN
		mnext = nav[monster][s+C]
		steps, e := dfs(maze, s+C, C, mnext, work, nav)
		if e {
			if len(out) == 0 || len(out)-1 > len(steps) {
				out = steps + "v"
				E = true
			}
		}
		near = false
	}
	if maze[s].left && work[s-1] == 0 && s-1 != monster { // LEFT
		mnext = nav[monster][s-1]
		steps, e := dfs(maze, s-1, C, mnext, work, nav)
		if e {
			if len(out) == 0 || len(out)-1 > len(steps) {
				out = steps + "<"
				E = true
			}
		}
		near = false
	}
	if work[s] <= 1 && !near && !E && out != "m" {
		mnext = nav[monster][s]
		steps, e := dfs(maze, s, C, mnext, work, nav)
		if e {
			if len(out) == 0 || len(out)-1 > len(steps) {
				out = steps + "o"
				E = true
			}
		}
	}
	work[s]--
	if E {
		return
	} else {
		return "cenok ej ot kat", false
	}
}

func navigateMonster(maze []vertex, C int) (nav [][]int) {
	nav = make([][]int, len(maze))
	for i, _ := range nav {
		if !maze[i].tile {
			continue
		}
		nav[i] = make([]int, len(maze))
		for j, _ := range nav[i] {
			if !maze[j].tile {
				nav[i][j] = -1
			}
			nav[i][j] = MMove(maze, C, i, j)
		}
	}
	return
}

func solve(maze []vertex, start int, C int, monster int) (out string) {
	var rev string
	if monster == -1 {
		rev = bfs(maze, start, C)
	} else {
		if t := bfs(maze, start, C); t == "cenok ej ot kat" {
			rev = "cenok ej ot kat"
		} else {
			if t := bfs(maze, monster, C); t[0] == 'c' {
				rev = bfs(maze, start, C)
			} else {
				work := make([]int, len(maze))
				nav := navigateMonster(maze, C)
				rev, _ = dfs(maze, start, C, monster, work, nav)
			}
		}
	}
	for i := len(rev) - 1; i >= 0; i-- {
		out += string(rev[i])
	}
	return
}

func multi(maze []vertex, start int, C int, monster int, outs []string, t int, end chan bool) {
	outs[t] = solve(maze, start, C, monster)
	end <- true
}

func main() {
	var T int
	var outs []string
	var end []chan bool
	fmt.Scan(&T)
	outs = make([]string, T)
	end = make([]chan bool, T)
	for i := range end {
		end[i] = make(chan bool)
	}
	for i := 0; i < T; i++ { // Načte všechny vstupy
		var R, C int
		var lab []rune
		var maze []vertex
		var start, monster int
		monster = -1
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
				maze[j].tile = true
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
			} else if c == MONSTRUM {
				monster = j
			}
		}
		//fmt.Println(maze, start)
		go multi(maze, start, C, monster, outs, i, end[i])
	}
	for i, _ := range outs {
		<-end[i]
		if outs[i] == "m" {
			fmt.Println("tak to je konec")
		} else {
			fmt.Println(outs[i])
		}
	}
}

import os
import strconv

import intcode

fn main() {
    input := os.read_file('input.txt')?
    
    mut code := []i64
    for c in input.split(",") {
        code << strconv.parse_int(c, 10, 64)
    }

	oxygen := find_oxygen_or_furthest_path(code, Point{0,0,[]int})
	println(oxygen.path.len) // star 1
	furthest := find_oxygen_or_furthest_path(code, oxygen)
	println(furthest.path.len - oxygen.path.len) // star 2
}

fn find_oxygen_or_furthest_path(code []i64, start Point) Point {
	directions := [Point {0, 1, []int}, Point {0, -1, []int}, Point {-1, 0, []int}, Point {1, 0, []int}] // const array
	
	mut open_set := [start]
	mut visited := map[string]bool
	mut longest_path := start

	for open_set.len > 0 {
		current := open_set.first()
		open_set.delete(0)

		visited[current.str()] = true

		if longest_path.path.len < current.path.len { longest_path = current }

		mut i := 0
		for i < 4 {
			mut path := current.path.clone()
			path << i + 1
			next := Point{ current.x + directions[i].x, current.y + directions[i].y, path }

			if !(next.str() in visited) {
				mut droid := Droid{}
				droid.load(code)
				status := droid.move(next.path)

				match status {
					0 { visited[next.str()] = true }
					1 { open_set << next }
					2 { return next }
					else { exit(1) }
				}
			}
			i++
		}
	}
	return longest_path
}

struct Droid {
mut:
    computer intcode.Computer
	open_set []Point
	visited map[string]bool
}
fn (d mut Droid) load(code []i64) {
    d.computer.load(code)
}
fn (d mut Droid) move(path []int) int {
	mut tile_type := i64(-1)
	for dir in path {
		_, tile_type = d.computer.run(dir)
	}

	return tile_type
}

struct Point {
    x int
    y int
	path []int
}
pub fn (p Point) str() string {
    return "($p.x, $p.y)"
}


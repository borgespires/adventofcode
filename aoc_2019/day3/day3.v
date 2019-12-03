import os
import strconv
import math

fn main() {
	input := os.read_lines("input.txt")?
	
	wire1 := get_path_points(input[0])
	wire2 := get_path_points(input[1])
	
	mut m := map[string]int
	
	for p in wire1 {
		if !(p.str() in m) || p.steps < m[p.str()] {
			m[p.str()] = p.steps
		}
	}
	
	mut nearest := 100000000
	mut shortest := 100000000
	
	for p in wire2 {
		if (p.str() in m) {
			manhatan := int(math.abs(p.x) + math.abs(p.y))
			steps := p.steps + m[p.str()]

			if nearest > manhatan { nearest = manhatan }
			if shortest > steps { shortest = steps }
		}
	}
	
	println(nearest) // star 1
	println(shortest) // star 2
}

fn get_path_points(path string) []Point {
	mut points := []Point
	mut current := Point{0, 0, 0}

	for section in path.split(",") {
		dir := section[0]
		mut steps := strconv.atoi(section[1..])
		
		for steps > 0 {
			if dir == `D` { current = Point{ current.x, current.y - 1, current.steps + 1 }}
			if dir == `U` { current = Point{ current.x, current.y + 1, current.steps + 1 }}
			if dir == `R` { current = Point{ current.x + 1, current.y, current.steps + 1 }}
			if dir == `L` { current = Point{ current.x - 1, current.y, current.steps + 1 }}
			
			points << current
			
			steps--
		}
	}

	return points
}

struct Point {
	x int
	y int
	steps int
}

pub fn (p Point) str() string {
	return "($p.x, $p.y)"
}
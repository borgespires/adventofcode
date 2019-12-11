import os
import strconv
import math

import intcode

fn main() {
    input := os.read_file('input.txt')?
    
    mut code := []i64
    for c in input.split(",") {
        code << strconv.parse_int(c, 10, 64)
    }
    
    mut r1 := new_robot(0)
    r1.load(code)
    t, _  := r1.run()
    println(t.keys().len) // star 1

    mut r2 := new_robot(1)
    r2.load(code)
    tiles, color := r2.run()
    print_registration(tiles, color) // star 2
}

struct Robot {
mut:
    computer intcode.Computer
    painted_tiles map[string]int
    position Point
    direction int = 0
}
fn (r mut Robot) load(code []i64) {
    r.computer.load(code)
}
fn (r mut Robot) run() (map[string]Point, map[string]int) {
    mut tiles := map[string]Point

    for {
        input := if r.position.str() in r.painted_tiles {
            r.painted_tiles[r.position.str()]
        } else { 0 }
        
        _, color := r.computer.run(input)
        exit_code, rotation_dir := r.computer.run(input)

        if exit_code == 0 { break }
        
        if color != input { 
            r.painted_tiles[r.position.str()] = int(color)
            tiles[r.position.str()] = r.position
        }

        r.rotate(int(rotation_dir))
        r.move()
    }

    return tiles, r.painted_tiles
}
fn (r mut Robot) rotate(dir int) {
    if dir == 0 { r.direction-- } else { r.direction++ }
    if r.direction < 0 { r.direction += 4 }
    else { r.direction %= 4 }
}
fn (r mut Robot) move() {
    directions := [Point {0, 1}, Point {1, 0}, Point {0, -1}, Point {-1, 0}] // const array
    dir := directions[r.direction]
    r.position = Point { r.position.x + dir.x, r.position.y + dir.y }
}

fn new_robot(phase int) Robot {
    return Robot { computer: intcode.Computer{ phase_setting: phase } position: Point { 0, 0 }}
}

struct Point {
    x int
    y int
}
pub fn (p Point) str() string {
    return "($p.x, $p.y)"
}

fn print_registration(tiles map[string]Point, color_map map[string]int) {
    mut min_x := math.max_i32
    mut max_x := math.min_i32
    mut min_y := math.max_i32
    mut max_y := math.min_i32

    for tile in tiles.keys() {
        x := tiles[tile].x
        y := tiles[tile].y
        
        if x > max_x { max_x = x }
        if x < min_x { min_x = x }
        if y > max_y { max_y = y }
        if y < min_y { min_y = y }
    }

    mut y := max_y
    for y >= min_y {
        mut x := min_x
        for x <= max_x {
            point := Point { x, y }
            color := color_map[point.str()]
			
            if color == 1 { print("#") } else { print(".") }
            x++
        }
        println("")
        y--
    }
}
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
    
	mut game := Game {}
	game.load(code)
	game.play()

	mut count := 0
	for pos in game.board.keys() {
		if game.board[pos] == 2 { count++ }
	}
	println(count) // star 1

	code[0] = 2 // free-play hack
	game.load(code)
    game.play()
	println(game.score) // star 2
}

struct Game {
mut:
    computer intcode.Computer
    board map[string]int
	tiles map[string]Point
	score int = 0
	player Point
	ball Point
}
fn (g mut Game) load(code []i64) {
    g.computer.load(code)
	g.player = Point {0, 0}
	g.ball = Point {0, 0}
}
fn (g mut Game) play() {
	for {
		joystick := g.get_joystick_input()
		
        _, x := g.computer.run(joystick)
		_, y := g.computer.run(0)
        exit_code, tile_id := g.computer.run(0)

		if exit_code == 0 { break }

		g.update_state(int(x), int(y), int(tile_id))
		// g.draw()
    }
}
fn (g Game) get_joystick_input() int {
	if g.player.x > g.ball.x { return -1 }
	if g.player.x < g.ball.x { return 1 }
	return 0
}
fn (g mut Game) update_state(x, y, id int) {
	if x==-1 && y==0 { 
		g.score = id
	} else {
		p := Point{ x, y }
		g.board[p.str()] = id
		g.tiles[p.str()] = p

		if id == 3 { g.player = p }
		if id == 4 { g.ball = p }
	}
}
fn (g Game) draw() {
	mut min_x := math.max_i32
    mut max_x := math.min_i32
    mut min_y := math.max_i32
    mut max_y := math.min_i32

    for tile in g.tiles.keys() {
        x := g.tiles[tile].x
        y := g.tiles[tile].y
        
        if x > max_x { max_x = x }
        if x < min_x { min_x = x }
        if y > max_y { max_y = y }
        if y < min_y { min_y = y }
    }

	println("score: $g.score")
	
    mut y := max_y
    for y >= min_y {
        mut x := min_x
        for x <= max_x {
            point := Point { x, y }
            color := g.board[point.str()]
			
			match color {
				0    { print(" ") }
				1    { print("|") }
				2    { print("#") }
				3    { print("-") }
				4    { print("O") }
				else { print(".") }
			}
            x++
        }
        println("")
        y--
    }
}

struct Point {
    x int
    y int
}
pub fn (p Point) str() string {
    return "($p.x, $p.y)"
}
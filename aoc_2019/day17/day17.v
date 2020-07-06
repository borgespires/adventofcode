import os
import strconv

import intcode

fn main() {
	test_sum_alignment()

    input := os.read_file('input.txt')?
    
    mut code := []i64
    for c in input.split(",") {
        code << strconv.parse_int(c, 10, 64)
    }

	code[0] = 2
	mut droid := Droid{}
	droid.load(code)
	// width, height, image := droid.display()

	// sum := sum_alignment_paramters(image, width, height)
	// println(sum)

	println(droid.move())
}

fn sum_alignment_paramters(image string, width, height int) int {
	mut sum := 0
	for i, px in image {
		x := i % width
		y := (i - x) / width
		// println("$x, $y, $i")
		if px == 35 && is_intersection(image, i, width) {
			sum += x * y
			// println("$x, $y, $i")
		}
	}

	return sum
}

fn is_intersection(image string, position int, width int) bool {
	top := position - width
	bottom := position + width
	left := position - 1
	right := position + 1

	if top < 0 || bottom >= image.len || left < 0 || right >= image.len { return false }


	// println("=>> $top, $bottom, $left, $right")

	return image[top] == 35 &&
		image[bottom] == 35 &&
		image[left] == 35 &&
		image[right] == 35
}


struct Droid {
mut:
    computer intcode.Computer
}
fn (d mut Droid) load(code []i64) {
    d.computer.load(code)

	main := "A,B,B,A,C,B,C,C,B,A\n"
	a := "R,10,R,8,L,10,L,10\n"
	b := "R,8,L,6,L,6\n"
	c := "L,10,R,10,L,6\n"

	input := main + a + b + c + "n\n"

	mut tmp := []i64
	for s in input { tmp << i64(s) }

	println('before')
	d.computer.input_this(tmp)
	println('after')

}
fn (d mut Droid) display() (int, int, string) {

	mut image := []byte
	mut x := 0
	mut y := 0

	for {
		exit_code, out := d.computer.run(0)
		if exit_code == 0 { break }

		match out {
			10 { y++ }
			else { 
				image << out
				x++
			}
		}
	}
	
	return (x/y) + 1, y, string(image)
}

fn (d mut Droid) move() int {

	

	mut image := []byte

	// mut i := 0
	for {
		// cmd := input[i]
		// println(cmd)
		exit_code, out := d.computer.run(0)
		if exit_code == 0 { break }
		
		// if i < input.len - 1 { i++ }
		// else {
		image << out
		// }

		// println(out)
	}

	println(image.len)
	println(string(image))
	
	return -1
}


fn test_sum_alignment() {
	image := "..#............#..........#######...####.#...#...#.##############..#...#...#....#####...^.."
	width := 13
	height := 7

	assert sum_alignment_paramters(image, width, height) == 76
}
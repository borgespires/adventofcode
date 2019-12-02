import os
import strconv

const (
	op_sum 		= 1
	op_multiply	= 2
	op_exit		= 99
)

fn main() {
	input := os.read_file('input.txt')?
	mut intcode := []int
	
	for c in input.split(",") {
		intcode << strconv.atoi(c)
	}

	mut c := Computer {}
	c.load(intcode)
	println(c.run(12, 2)) // star 1

	noun, verb := find_correct_params(19690720, intcode)
	println(100 * noun + verb) // star 2
}

fn find_correct_params(expectedOut int, intcode []int) (int, int) {
	mut c := Computer {}
	
	mut noun := 0
	for noun < 100 {
		mut verb := 0
		for verb < 100 {
			c.load(intcode)
			if (c.run(noun, verb) == expectedOut) {
				return noun, verb
			}
			verb++
		}
		noun++
	}
	// return error("Did not find any input pair that produces $expectedOut")
	return -1, -1
}

fn eval(op int, arg1 int, arg2 int) int {
	if (op == op_sum) { return arg1	+ arg2 }
	if (op == op_multiply) { return arg1 * arg2 }
	
	// return error("Invalid operation $op")
	return -1
}

struct Computer {
mut:
	memory []int = []int
}

fn (c mut Computer) load(intcode []int) {
	c.memory = []
	c.memory << intcode
}

fn (c mut Computer) run(noun int, verb int) int {
	c.memory[1] = noun
	c.memory[2] = verb
	
	mut ptr := 0
	for {
		op := c.memory[ptr]

		if op == op_exit {
			break
		}

		address1 := c.memory[ptr + 1]
		address2 := c.memory[ptr + 2]
		address3 := c.memory[ptr + 3]

		param1 := c.memory[address1]
		param2 := c.memory[address2]
		c.memory[address3] = eval(op, param1, param2)

		ptr += 4
	}
	return c.memory[0]
}
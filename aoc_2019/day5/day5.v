import os
import math
import strconv

const (
    op_sum      = 1
    op_mult     = 2
    op_in       = 3
    op_out      = 4
    op_jit      = 5
    op_jif      = 6
    op_lt       = 7
    op_eq       = 8
    op_exit     = 99
)

fn main() {
    input := os.read_file('input.txt')?
    mut intcode := []int
    for c in input.split(",") {
        intcode << strconv.atoi(c)
    }
    mut c := Computer {}
    c.load(intcode)
    println(c.run(1))

	c.load(intcode)
    println(c.run(5))
}

fn eval(op int, arg1 int, arg2 int) int {
    if (op == op_sum) { return arg1 + arg2 }
    if (op == op_mult) { return arg1 * arg2 }
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
fn (c mut Computer) run(phase int) int {
    mut ptr := 0
	mut out := -1
    for {
        modes, opcode := parse_instruction(c.memory[ptr])

        if opcode == op_exit {
            break
        }
        if opcode == op_in {
            address := c.memory[ptr + 1]
            c.memory[address] = phase
            // println("STDIN > ")
            ptr += 2
        }
        if opcode == op_out {
            address := c.memory[ptr + 1]
            out = c.memory[address]
            // println("STDOUT > $out")
            ptr += 2
        }
        if opcode == op_sum || opcode == op_mult {
            address1 := c.memory[ptr + 1]
            address2 := c.memory[ptr + 2]
            address3 := c.memory[ptr + 3]
            
			param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            
			c.memory[address3] = eval(opcode, param1, param2)
            
			ptr += 4
        }
        if opcode == op_jit {
            address1 := c.memory[ptr + 1]
            address2 := c.memory[ptr + 2]
            
			param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            
			if param1 != 0 { ptr = param2 }
            else { ptr += 3 }
        }
        if opcode == op_jif {
            address1 := c.memory[ptr + 1]
            address2 := c.memory[ptr + 2]
            
			param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            
			if param1 == 0 { ptr = param2 }
            else { ptr += 3 }
        }
        if opcode == op_lt {
            address1 := c.memory[ptr + 1]
            address2 := c.memory[ptr + 2]
            address3 := c.memory[ptr + 3]
            
			param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            
			if param1 < param2 { c.memory[address3] = 1 } 
            else { c.memory[address3] = 0 }
            
			ptr += 4
        }
        if opcode == op_eq {
            address1 := c.memory[ptr + 1]
            address2 := c.memory[ptr + 2]
            address3 := c.memory[ptr + 3]
            
			param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            
			if param1 == param2 { c.memory[address3] = 1 } 
            else { c.memory[address3] = 0 }
            
			ptr += 4
        }
    }
    return out
}

fn parse_instruction(instruction int) ([]int, int) {
    mut modes := []int
    opcode := get_digit(instruction, 1) + get_digit(instruction, 2) * 10
    mut n := 0
    for n++ < 3 { modes << get_digit(instruction, n + 2) }
    return modes, opcode
}

fn get_digit(number int, idx int) int {
    d := int(math.ceil(math.pow(10, (idx - 1))))
    k := number / d
    return k % 10
}
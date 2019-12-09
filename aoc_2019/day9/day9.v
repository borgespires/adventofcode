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
    op_base   	= 9
    op_exit     = 99
)

enum State {
    halt conf run
}

fn main() {
    input := os.read_file('input.txt')?
    
	mut intcode := []i64
    for c in input.split(",") {
        intcode << strconv.parse_int(c, 10, 64)
    }
    
	mut c1 := Computer { phase_setting: 1 }
    c1.load(intcode)
    println(c1.run(0).var_1)

    mut c2 := Computer { phase_setting: 2 }
    c2.load(intcode)
    println(c2.run(0).var_1)
}

struct Computer {
    phase_setting i64 = i64(0)
mut:
    ptr i64 = i64(0)
    base i64 = i64(0)
    memory []i64 = []i64
    state State = State.halt
}
fn (c mut Computer) load(intcode []i64) {
    c.memory = []
    c.state = .conf
    c.memory << intcode
    c.ptr = 0
}
fn (c mut Computer) run(stdin i64) (int, i64) {
    for {
        modes, opcode := parse_instruction(c.memory[c.ptr])
        
		// println("$modes, $opcode")
        
		if opcode == op_exit {
            c.state = .halt
            break
        }
        if opcode == op_in {
            mut input := if c.state == .conf { 
                c.state = .run
                c.phase_setting
            } else { stdin }
            // println("STDIN > $input")
            c.write(c.ptr + 1, modes[0], input)
            c.ptr += 2
        }
        if opcode == op_out {
            value := c.read(c.ptr + 1, modes[0])
            // println("STDOUT > $value")
            c.ptr += 2
			return 1, value
        }
        if opcode == op_sum {
            param1 := c.read(c.ptr + 1, modes[0])
            param2 := c.read(c.ptr + 2, modes[1])
            c.write(c.ptr + 3, modes[2], param1 + param2)
            c.ptr += 4
        }
		if opcode == op_mult {
            param1 := c.read(c.ptr + 1, modes[0])
            param2 := c.read(c.ptr + 2, modes[1])
            c.write(c.ptr + 3, modes[2], param1 * param2)
            c.ptr += 4
        }
        if opcode == op_jit {
            param1 := c.read(c.ptr + 1, modes[0])
            param2 := c.read(c.ptr + 2, modes[1])
            if param1 != 0 { c.ptr = param2 }
            else { c.ptr += 3 }
        }
        if opcode == op_jif {
            param1 := c.read(c.ptr + 1, modes[0])
            param2 := c.read(c.ptr + 2, modes[1])
            if param1 == 0 { c.ptr = param2 }
            else { c.ptr += 3 }
        }
        if opcode == op_lt {
            param1 := c.read(c.ptr + 1, modes[0])
            param2 := c.read(c.ptr + 2, modes[1])
            if param1 < param2 { c.write(c.ptr + 3, modes[2], 1) } 
            else { c.write(c.ptr + 3, modes[2], 0) }
            c.ptr += 4
        }
        if opcode == op_eq {
            param1 := c.read(c.ptr + 1, modes[0])
            param2 := c.read(c.ptr + 2, modes[1])
            if param1 == param2 { c.write(c.ptr + 3, modes[2], 1) } 
            else { c.write(c.ptr + 3, modes[2], 0) }
            c.ptr += 4
        }
        if opcode == op_base {
            c.base += c.read(c.ptr + 1, modes[0])
            c.ptr += 2
        }
    }
    return 0, stdin
}
fn (c mut Computer) resize(address i64) {
	size := int(address - c.memory.len + 1)
    c.memory << [i64(0)].repeat(size)
}
fn (c mut Computer) read(ptr i64, mode i64) i64 {
    address := c.memory[ptr]

    if mode == 0 {
        if address >= c.memory.len { c.resize(address) }
        return c.memory[address]
    }
    if mode == 1 { return address }
    if mode == 2 {
		ref := c.base + address
        if ref >= c.memory.len { c.resize(ref) }
        return c.memory[ref]
    }
    return -1 // errror
}
fn (c mut Computer) write(ptr i64, mode i64, value i64) {
    address := c.memory[ptr]

    if mode == 0 {
        if address >= c.memory.len { c.resize(address) }
        c.memory[address] = value
    }
    if mode == 2 { 
        ref := c.base + address
        if ref >= c.memory.len { c.resize(ref) }
        c.memory[ref] = value 
    }
}

fn parse_instruction(instruction i64) ([]int, int) {
    mut modes := []int
    opcode := get_digit(instruction, 1) + get_digit(instruction, 2) * 10
    mut n := 0
    for n++ < 3 { modes << get_digit(instruction, n + 2) }
    return modes, opcode
}
fn get_digit(number i64, idx int) int {
    d := int(math.ceil(math.pow(10, idx - 1)))
    k := number / d
    return int(k % 10)
}
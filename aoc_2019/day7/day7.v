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

enum State {
	halt conf run
}

fn main() {
    input := os.read_file('input.txt')?
    
	mut intcode := []int
    for c in input.split(",") {
        intcode << strconv.atoi(c)
    }

	println(get_highest_signal(intcode, [0,1,2,3,4])) // star 1
	println(get_highest_signal(intcode, [5,6,7,8,9])) // star 2 [7, 8, 6, 9, 5]
}

fn get_highest_signal(intcode []int, phase_settings []int) int {
	mut max := 0
	for setting in permutation(phase_settings) {
		mut system := new_system(setting)
		system.load(intcode)
		signal := system.run()
		if signal > max { max = signal }
	}

	return max
}

struct System {
mut:
	computers []Computer
}
fn (s mut System) load(intcode []int) {
	mut i:= 0
	for i < s.computers.len { 
		s.computers[i++].load(intcode)
	}
}
fn (s mut System) run() int {
	size := s.computers.len
	mut signal := 0
	mut i := 0
	mut halted := 0

	for halted < size {
		ecode, out := s.computers[i].run(signal)
		if ecode == 0 { halted++ }
		signal = out
		i = (i + 1) % size
	}

	return signal
}

fn new_system(setting []int) System {
	mut computers := []Computer
	for s in setting {
		computers << Computer { phase_setting: s }
	}
	return System { computers }
}


struct Computer {
	phase_setting int = 0
mut:
	ptr int = 0
    memory []int = []int
	state State = State.halt
}
fn (c mut Computer) load(intcode []int) {
    c.memory = []
	c.state = .conf
    c.memory << intcode
	c.ptr = 0
}
fn (c mut Computer) run(stdin int) (int, int) {
    for {
        modes, opcode := parse_instruction(c.memory[c.ptr])

        if opcode == op_exit {
            c.state = .halt
			break
        }
        if opcode == op_in {
            address := c.memory[c.ptr + 1]
			
			mut input := if c.state == .conf	 { 
				c.state = .run	
				c.phase_setting
			} else { stdin }

			// println("STDIN > $input")
			c.memory[address] = input
            c.ptr += 2
        }
        if opcode == op_out {
            address := c.memory[c.ptr + 1]
            value := c.memory[address]
            // println("STDOUT > $value")
            c.ptr += 2
			return 1, value
        }
        if opcode == op_sum || opcode == op_mult {
            address1 := c.memory[c.ptr + 1]
            address2 := c.memory[c.ptr + 2]
            address3 := c.memory[c.ptr + 3]
            param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            c.memory[address3] = eval(opcode, param1, param2)
            c.ptr += 4
        }
        if opcode == op_jit {
            address1 := c.memory[c.ptr + 1]
            address2 := c.memory[c.ptr + 2]
            param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            if param1 != 0 { c.ptr = param2 }
            else { c.ptr += 3 }
        }
        if opcode == op_jif {
            address1 := c.memory[c.ptr + 1]
            address2 := c.memory[c.ptr + 2]
            param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            if param1 == 0 { c.ptr = param2 }
            else { c.ptr += 3 }
        }
        if opcode == op_lt {
            address1 := c.memory[c.ptr + 1]
            address2 := c.memory[c.ptr + 2]
            address3 := c.memory[c.ptr + 3]
            param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            if param1 < param2 { c.memory[address3] = 1 } 
            else { c.memory[address3] = 0 }
            c.ptr += 4
        }
        if opcode == op_eq {
            address1 := c.memory[c.ptr + 1]
            address2 := c.memory[c.ptr + 2]
            address3 := c.memory[c.ptr + 3]
            param1 := if modes[0] == 0 { c.memory[address1] } else { address1 }
            param2 := if modes[1] == 0 { c.memory[address2] } else { address2 }
            if param1 == param2 { c.memory[address3] = 1 } 
            else { c.memory[address3] = 0 }
            c.ptr += 4
        }
    }
    return 0, stdin
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

fn eval(op int, arg1 int, arg2 int) int {
    if (op == op_sum) { return arg1 + arg2 }
    if (op == op_mult) { return arg1 * arg2 }
    // return error("Invalid operation $op")
    return -1
}

fn permutation(lst []int) [][]int {
    if lst.len == 0 { return [] } 
    if lst.len == 1 { return [lst] }
  
    mut l := [[]int] // cannot create an empty matrix

	mut i := 0
    for i < lst.len {
		m := lst[i] 
		
		mut rem_lst := []int
		rem_lst << lst[..i]
		rem_lst << lst[i+1..]
  
		for p in permutation(rem_lst) {
			mut tmp := [m]
			tmp << p
			l << tmp
		}
		i++
	}

    return l[1..] 
}

import os
import strconv
import math

fn main() {
    input := os.read_file('input.txt')?
	
	println(input)

	signal := input
	mut digits := []int
	for digit in signal {
		digits << toint(digit)
	}

	phase := fft_loop(digits, 100)
	println(phase[0..8])
}

fn fft_loop(signal []int, phases int) []int {
	mut i := 0
	mut current := []int
	current << signal
	
	for i < phases {
		mut pos := 0
		mut phase_temp := []int
		for pos < signal.len {
			phase_temp << next(current, pos)
			pos++
		}
		current = phase_temp
		i++
	}
	return current
}

fn next(signal []int, pos int) int {
	mut base_pattern := []int
	base_pattern << [0].repeat(pos + 1)
	base_pattern << [1].repeat(pos + 1)
	base_pattern << [0].repeat(pos + 1)
	base_pattern << [-1].repeat(pos + 1)
	offset := 1
	mut value := 0
	
	for i, digit in signal {
		pattern_idx := (i + offset) % base_pattern.len
		multiplier := base_pattern[pattern_idx]
		value += digit * multiplier
		// print("($digit * $multiplier) + ")
	}

	st := int(math.abs(value)).str()

	// println(" = $st")

	return toint(st[st.len-1])
}

fn toint(number byte) int {
	return int(number) - 48
}
import math

const(
	min = 264360
	max = 746325
)

fn main() {
	mut i := min
	mut total_with_double := 0
	mut total_with_rigid_double := 0

	for i < max {
		if has_double(i) { total_with_double++ }
		if has_rigid_double(i) { total_with_rigid_double++ }
		i++
	}

	println(total_with_double) // star 1
	println(total_with_rigid_double) // star 2
}

fn has_double(password int) bool {
	mut n := 6
	mut prev := -1
	mut found_double := false

	for n > 0 {
		digit := get_digit(password, n)
		
		if digit < prev { return false }
		if digit == prev { found_double = true }
		
		prev = digit
		n--
	}

	return found_double
}

fn has_rigid_double(password int) bool {
	mut n := 6
	mut prev := -1
	mut has_duplicates := false
	mut duplicates := 0

	for n > 0 {
		digit := get_digit(password, n)
		
		if digit < prev { return false }
		if digit == prev { duplicates++ }
		else {
			if (duplicates == 1) { has_duplicates = true } 
			duplicates = 0
		}
		
		prev = digit
		n--
	}

	return has_duplicates || duplicates == 1
}

fn get_digit(number int, idx int) int {
	d := int(math.ceil(math.pow(10, (idx - 1))))
	k := number / d
	return k % 10
}
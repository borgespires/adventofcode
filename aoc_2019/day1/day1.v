import os
import math
import strconv

fn main() {
    modules := os.read_lines("input.txt")?

	println(total_fuel(modules, required_fuel)) // star 1
	println(total_fuel(modules, rec_required_fuel)) // star 2
}

fn total_fuel(modules []string, fuel_calculator fn (int) int) int {
	mut total := 0

    for mass in modules {
        total += fuel_calculator(strconv.atoi(mass))
    }

	return total
}

fn required_fuel(mass int) int {
	return int(math.floor(mass / 3) - 2)
}

fn rec_required_fuel(mass int) int {
	fuel := required_fuel(mass)

	if fuel < 0 {
		return 0
	} else {
		return fuel + rec_required_fuel(fuel)
	}
}
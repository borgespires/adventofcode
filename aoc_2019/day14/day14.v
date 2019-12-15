import os
import strconv

fn main() {
    input := os.read_lines('input.txt')?
	mut reactions := map[string]Reaction
    for reaction_equation in input {
		out_chemical, reaction := parse_reaction(reaction_equation)
		reactions[out_chemical] = reaction
    }

	println(produce_fuel(reactions, 1)) // star 1
	println(max_fuel_production(reactions, 1000000000000)) // star 2
}

fn max_fuel_production(reactions map[string]Reaction, cargo u64) u64 {
	mut start := u64(0)
	mut end := cargo
	mut fuel_guess := u64(0)

	for {
		fuel_guess = (end - start) / u64(2) + start
		used_ore := produce_fuel(reactions, fuel_guess)

		if used_ore <= cargo && end - start <= 1 { break }
		if used_ore > cargo { end = fuel_guess } else { start = fuel_guess }
	}

	return fuel_guess
}

fn produce_fuel(reactions map[string]Reaction, fuel u64) u64 {
	mut total_ore := u64(0)
	mut obtain := [Chemical{ "FUEL", fuel }]
	mut surplus := map[string]u64

	for obtain.len > 0 {
		mut product := obtain[0]
		obtain.delete(0)

		if surplus[product.name] >= product.quantity {
			surplus[product.name] = surplus[product.name] - product.quantity
			continue
		}

		product.quantity -= surplus[product.name]
		
		reaction := reactions[product.name]
		multiplier := round_up(product.quantity, reaction.output.quantity)

		surplus[product.name] = (multiplier * reaction.output.quantity) - product.quantity

		for reagent in reaction.input {
			needed_quantity := multiplier * reagent.quantity
			if reagent.name == "ORE" {
				total_ore += needed_quantity
			} else {
				obtain << Chemical{ reagent.name, needed_quantity }
			}
		}
	}

	return total_ore
}

fn parse_reaction(equation string) (string, Reaction) {
	expressions := equation.split(" => ")
	lhs := expressions[0]
	rhs := expressions[1]

	mut input := []Chemical
	output := parse_chemical(rhs)

	for term in lhs.split(", ") {
		input << parse_chemical(term)
	}

	return output.name, Reaction { input, output }
}

fn parse_chemical(term string) Chemical {
	parts := term.split(" ")
	return Chemical{ parts[1], strconv.atoi(parts[0]) }
}

fn round_up(num u64, divisor u64) u64 {
    return  (num + divisor - u64(1)) / divisor
}

struct Reaction {
	input []Chemical
	output Chemical
}
pub fn (r Reaction) str() string {
    return "$r.input => $r.output"
}

struct Chemical {
	name string
mut:
	quantity u64
}
pub fn (c Chemical) str() string {
    return "$c.name/($c.quantity/)"
}
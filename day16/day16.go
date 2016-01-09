package day16

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

func parse(s string) map[string]int {
    r := regexp.MustCompile(`^Sue \d+: (\w+): (\d+), (\w+): (\d+), (\w+): (\d+)$`)
    match := r.FindStringSubmatch(s)

    sue := make(map[string]int)

    if match != nil {
        am1, _ := strconv.Atoi(match[2])
        am3, _ := strconv.Atoi(match[4])
        am5, _ := strconv.Atoi(match[6])
        sue[match[1]] = am1
        sue[match[3]] = am3
        sue[match[5]] = am5
    }

    return sue
}

func match(sue map[string]int, compounds map[string]int) bool {
    for thing, amount := range sue {
        if reading, _ := compounds[thing]; amount != reading { return false } 
    }

    return true
}

func _match(sue map[string]int, compounds map[string]int) bool {
    for thing, amount := range sue {
        reading, _ := compounds[thing]

        switch thing {
        case "cats", "trees":
            if amount <= reading { return false }
        case "pomeranians", "goldfish":
            if amount >= reading { return false }
        default:
            if amount != reading { return false }
        }
    }

    return true
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    var first, second int

    compounds := map[string]int{
        "children": 3,
        "cats": 7,
        "samoyeds": 2,
        "pomeranians": 3,
        "akitas": 0,
        "vizslas": 0,
        "goldfish": 5,
        "trees": 3,
        "cars": 2,
        "perfumes": 1,
    }

    for i, str := range lines {
        if match(parse(str), compounds) { first = i+1 }
        if _match(parse(str), compounds) { second = i+1 }
    }

    return first, second
}
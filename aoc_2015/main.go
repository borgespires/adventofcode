package main

import (
    "os"
    "fmt"

    "adventofcode/day1"
    "adventofcode/day2"
    "adventofcode/day3"
    "adventofcode/day4"
    "adventofcode/day5"
    "adventofcode/day6"
    "adventofcode/day7"
    "adventofcode/day8"
    "adventofcode/day9"
    "adventofcode/day10"
    "adventofcode/day11"
    "adventofcode/day12"
    "adventofcode/day13"
    "adventofcode/day14"
    "adventofcode/day15"
    "adventofcode/day16"
    "adventofcode/day17"
    "adventofcode/day18"
    "adventofcode/day19"
    "adventofcode/day20"
    "adventofcode/day21"
    "adventofcode/day22"
    "adventofcode/day23"
    "adventofcode/day24"
    "adventofcode/day25"
)

func main() {
    var s1, s2 interface{}
    puzzle := os.Args[1]

    switch puzzle {
        case "1": s1, s2 = day1.Solve(os.Args[2])
        case "2": s1, s2 = day2.Solve(os.Args[2])
        case "3": s1, s2 = day3.Solve(os.Args[2])
        case "4": s1, s2 = day4.Solve("bgvyzdsv")
        case "5": s1, s2 = day5.Solve(os.Args[2])
        case "6": s1, s2 = day6.Solve(os.Args[2])
        case "7": s1, s2 = day7.Solve(os.Args[2])
        case "8": s1, s2 = day8.Solve(os.Args[2])
        case "9": s1, s2 = day9.Solve(os.Args[2])
        case "10": s1, s2 = day10.Solve("1321131112")
        case "11": s1, s2 = day11.Solve("hepxcrrq")
        case "12": s1, s2 = day12.Solve(os.Args[2])
        case "13": s1, s2 = day13.Solve(os.Args[2])
        case "14": s1, s2 = day14.Solve(os.Args[2], 2503)
        case "15": s1, s2 = day15.Solve(os.Args[2])
        case "16": s1, s2 = day16.Solve(os.Args[2])
        case "17": s1, s2 = day17.Solve(os.Args[2])
        case "18": s1, s2 = day18.Solve(os.Args[2])
        case "19": s1, s2 = day19.Solve(os.Args[2])
        case "20": s1, s2 = day20.Solve(29000000)
        case "21": s1, s2 = day21.Solve(8, 2, 100)
        case "22": s1, s2 = day22.Solve(9, 58)
        case "23": s1, s2 = day23.Solve(os.Args[2])
        case "24": s1, s2 = day24.Solve(os.Args[2])
        case "25": s1, s2 = day25.Solve(2947, 3029)
        default: {
            fmt.Printf("Invalid day %s, choose a valid one [1-25]\n", puzzle)
            return
        }
    }

    fmt.Printf("Answers for adventofcode day %s\n", puzzle)
    fmt.Printf("star 1: %v\n", s1)
    fmt.Printf("star 2: %v\n", s2)
}
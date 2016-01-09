package day14

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

type Reindeer struct {
    speed, moving, rest int
    travelled, points int
}

func parse(s string) (name string, speed int, moving int, rest int) {
    r := regexp.MustCompile(`^(\w+).*?(\d+).*?(\d+).*?(\d+).*$`)
    match := r.FindStringSubmatch(s)

    if match != nil {
        name = match[1]
        speed, _ = strconv.Atoi(match[2])
        moving, _ = strconv.Atoi(match[3])
        rest, _ = strconv.Atoi(match[4])
    }

    return
}

func race(reindeers map[string]*Reindeer, timeLimit int) int {
    winnerDistance := 0

    for _, r := range reindeers {
        burstTime := r.moving + r.rest
        bursts := timeLimit / burstTime
        remainder := timeLimit % burstTime

        if remainder > r.moving { remainder = r.moving }

        distance := (bursts * r.moving * r.speed) + remainder * r.speed

        if distance > winnerDistance { winnerDistance = distance }
    }

    return winnerDistance
}

func burstRace(reindeers map[string]*Reindeer, timeLimit int) int {
    winnerPoints := 0

    for clk := 0; clk <= timeLimit; clk++ {
        first := []string{}
        firstDist := 0

        for name, r := range reindeers {
            burstTime := r.moving + r.rest
            if remainder := clk % burstTime; remainder < r.moving { r.travelled += r.speed }

            if r.travelled > firstDist { 
                firstDist = r.travelled
                first = append([]string{}, name)
            } else if r.travelled == firstDist {
                first = append(first, name) 
            }
        }

        for _, f := range first { reindeers[f].points++ }
    }

    for _, r := range reindeers {
        if r.points > winnerPoints { winnerPoints = r.points }
    }

    return winnerPoints
}

func Solve(filepath string, timeLimit int) (int, int) {
    lines := io.Readlines(filepath)
    reindeers := make(map[string]*Reindeer)

    for _, str := range lines {
        name, spd, tm, tr := parse(str)

        reindeers[name] = &Reindeer {
            speed: spd,
            moving: tm,
            rest: tr,
        }
    }

    return race(reindeers, timeLimit), burstRace(reindeers, timeLimit)
}
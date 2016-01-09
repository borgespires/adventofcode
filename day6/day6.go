package day6

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

type Grid struct {
    lights [1000][1000]int
}

func (g *Grid) get(x int, y int) int {
    return g.lights[x][y]
}

func (g *Grid) light(x int, y int, cmd int) {
    var light *int = &(g.lights[x][y])

    switch cmd {
    case 1:
        *light = 1
    case 2:
        *light = (*light+1)%2
    case -1:
        *light = 0
    default:
    }
}

func (g *Grid) _light(x int, y int, cmd int) {
    var light *int = &(g.lights[x][y])
    *light += cmd
    if *light < 0 { *light = 0}
}

func parse(s string) map[string]int {
    r := regexp.MustCompile(`^(?:turn )?(?P<id>on|toggle|off) (?P<xi>\d{0,3}),(?P<yi>\d{0,3}) through (?P<xf>\d{0,3}),(?P<yf>\d{0,3})$`)
    captures := make(map[string]int)

    match := r.FindStringSubmatch(s)

    if match == nil {
        return captures
    }

    for i, name := range r.SubexpNames() {
        var value int

        // Ignore the whole regexp match and unnamed groups
        if i == 0 || name == "" { continue }

        if name == "id" {
            switch match[i] {
            case "on":
                value = 1
            case "toggle":
                value = 2
            case "off":
                value = -1
            default:
            }
        } else { value, _ = strconv.Atoi(match[i]) }

        captures[name] = value
    }

    return captures
}


func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    
    gridA := new(Grid)
    gridB := new(Grid)
    
    countA := 0
    countB := 0

    for _, str := range lines {
        cmd := parse(str)

        for x:=cmd["xi"]; x <= cmd["xf"]; x++ {
            for y:=cmd["yi"]; y <= cmd["yf"]; y++ {
                gridA.light(x, y, cmd["id"])
                gridB._light(x, y, cmd["id"])
            }
        }
    }

    for x:=0; x < 1000; x++ {
        for y:=0; y < 1000; y++ {
            countA += gridA.get(x, y)
            countB += gridB.get(x, y)
        }
    }

    return countA, countB
}
package day23

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

type Instruction struct {
    cmd string
    reg string
    offset int
}

type Machine struct {
    register map[string]uint
}

func newMachine() *Machine {
    return &Machine{
        register: map[string]uint{ "a":0, "b":0 },
    }
}

func (m *Machine) run(program []Instruction) uint {
    for i := 0; i < len(program); {
        inst := program[i]

        switch inst.cmd {
            case "hlf":
                m.register[inst.reg] /= 2
            case "tpl":
                m.register[inst.reg] *= 3
            case "inc":
                m.register[inst.reg] += 1
            case "jmp":
                i += inst.offset
                continue
            case "jie":
                if m.register[inst.reg] % 2 == 0 { i += inst.offset; continue }
            case "jio":
                if m.register[inst.reg] == 1 { i += inst.offset; continue }
        }

        i++
    }

    return m.register["b"]
}

func parse(lines []string) []Instruction {
    r := regexp.MustCompile(`^(?P<cmd>\w{3})\s(?P<register>\w{1})?((?:,\s)?(?P<offset>(\+|-)?\d+))?$`)
    captures := []Instruction{}

    for _, s := range lines {
        match := r.FindStringSubmatch(s)

        if match == nil { continue }

        inst := Instruction{}

        for i, name := range r.SubexpNames() {
            var value int

            if i == 0 || name == "" { continue }

            switch name {
            case "cmd":
                inst.cmd = match[i]
            case "register":
                inst.reg = match[i]
            case "offset":
                value, _ = strconv.Atoi(match[i])
                inst.offset = value
            default:
            }
        }

        captures = append(captures, inst)
    }

    return captures
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)

    program := parse(lines)
    incRegAByOne := Instruction{cmd:"inc", reg: "a"}
    program2 := append([]Instruction{incRegAByOne}, program...)

    return int(newMachine().run(program)), int(newMachine().run(program2))
}
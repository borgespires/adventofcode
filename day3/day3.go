package day3

import (
    "fmt"

    "adventofcode/io"
)

type Robot struct {
    x, y int
}

func (bot Robot) house() string {
    return fmt.Sprintf("%d,%d", bot.x, bot.y)
}

type Santa struct {
    turn int
    robots []*Robot
    visited map[string]bool
    total int
}

func NewSanta(n int) *Santa {
    santa := &Santa {
        robots: make([]*Robot, 0, n),
        visited: map[string]bool{ "0,0": true, },
        total: 1,
    }

    for i := 0; i < n; i++ {
        santa.robots = append(santa.robots, new(Robot))
    }
    
    return santa
}

func (santa *Santa) move(r rune) {
    bot := santa.robots[santa.turn]

    switch r {
        case 94:
            bot.y++
        case 118:
            bot.y--
        case 62:
            bot.x++
        case 60:
            bot.x--
        default:
    }

    house := bot.house()

    if _, ok := santa.visited[house]; !ok {
        santa.visited[house] = true
        santa.total++
    }

    santa.turn = (santa.turn + 1) % len(santa.robots)
}

func Solve(filepath string) (int, int) {
    s := io.Readlines(filepath)[0]
    santa := NewSanta(1)
    santaWithBot := NewSanta(2)

    for _, r := range s {
        santa.move(r)
        santaWithBot.move(r)
    }

    return santa.total, santaWithBot.total
}
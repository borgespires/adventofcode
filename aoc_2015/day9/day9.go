package day9

import (
    "fmt"
    "regexp"
    "strconv"

    "adventofcode/io"
)

type Route struct {
    to *Location
    distance int
}

func (r *Route) String() string { return fmt.Sprintf("%s,%d", r.to.name, r.distance) }

type Location struct {
    routes []Route
    name string
    visited bool
}

func (l *Location) add(r Route) { l.routes = append(l.routes, r) }
func (l *Location) String() string { return fmt.Sprintf("%v,%t", l.routes, l.visited) }

type Map struct {
    locations map[string]*Location
}

func newMap() *Map {
    return &Map {
        locations: make(map[string]*Location),
    }
}

func (m *Map) getsert(name string) (local *Location) {
    var exists bool

    if local, exists = m.locations[name]; !exists {
        local = &Location{ name: name, }
        m.locations[name] = local
    }

    return local
}

func (m *Map) travel(current *Location, traveledDist int, alreadyVisited int) (bool, int, int) {
    shortestTour := int(^uint(0) >> 1)
    longestTour := traveledDist
    allVisited := false
    deadEnd := true

    current.visited = true

    for _, route := range current.routes {
        var sTour int
        var lTour int
        dest := route.to

        if dest.visited { continue } else { deadEnd = false }

        allVisited, sTour, lTour = m.travel(dest, traveledDist + route.distance, alreadyVisited+1)

        if allVisited && sTour < shortestTour { shortestTour = sTour }
        if allVisited && lTour > longestTour { longestTour = lTour }
    }

    if deadEnd { 
        shortestTour = traveledDist
        if len(m.locations) == alreadyVisited + 1 { allVisited = true }
    }

    current.visited = false

    return allVisited, shortestTour, longestTour
}

func parse(s string) (from string, to string, d int) {
    r := regexp.MustCompile(`^(\w+)\sto\s(\w+)\s=\s(\d+)$`)
    match := r.FindStringSubmatch(s)

    if match != nil {
        from = match[1]
        to = match[2]
        d, _ = strconv.Atoi(match[3])
    }

    return
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    
    shortestTour := int(^uint(0) >> 1)
    longestTour := 0
    m := newMap()

    for _, str := range lines {
        aname, bname, d := parse(str)

        a := m.getsert(aname)
        b := m.getsert(bname)

        a.add(Route{ to: b, distance: d })
        b.add(Route{ to: a, distance: d })
    }

    for startLocation, _ := range m.locations {
        allVisited, sTour, lTour := m.travel(m.locations[startLocation], 0, 0)

        if allVisited && sTour < shortestTour { shortestTour = sTour }
        if allVisited && lTour > longestTour { longestTour = lTour }
    }

    return shortestTour, longestTour
}
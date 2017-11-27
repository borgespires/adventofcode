package day7

import (
    "regexp"
    "strconv"

    "adventofcode/io"
)

type Circuit struct {
    inputs map[string]uint16
    heap map[string]uint16
    queue []map[string]string
}

func NewCircuit() *Circuit {
    return &Circuit{
        inputs: make(map[string]uint16),
        heap: make(map[string]uint16),
        queue: make([]map[string]string, 0),
    }
}

func (c *Circuit) enqueue(cmd map[string]string) { c.queue = append(c.queue, cmd) }
func (c *Circuit) dequeue() map[string]string { 
    cmd := c.queue[0]
    c.queue = c.queue[1:]
    return cmd
}
func (c *Circuit) in(wire string, signal uint16) { c.inputs[wire] = signal }
func (c *Circuit) set(wire string, signal uint16) { c.heap[wire] = signal }
func (c *Circuit) get(wire string) (signal uint16, ok bool) {
    if signal, ok = c.heap[wire]; ok { return }
    signal, ok = c.inputs[wire]
    return
}
func (c *Circuit) clearHeap() {
    c.heap = make(map[string]uint16)
}

func (c *Circuit) eval(cmd map[string]string) {
    sa, saOk := cmd["sa"] // string a
    sb, sbOk := cmd["sb"] // string b
    ia, iaOk := cmd["ia"] // int a
    ib, ibOk := cmd["ib"] // int b

    out := cmd["out"]
    gate, isGate := cmd["gate"]

    var a uint16
    var b uint16
    var aOk bool
    var bOk bool

    if saOk { a, aOk = c.get(sa) } else if iaOk {
        aOk = true
        i, _ := strconv.Atoi(ia)
        a = uint16(i)
    }
    if sbOk { b, bOk = c.get(sb) } else if ibOk {
        bOk = true 
        i, _ := strconv.Atoi(ib)
        b = uint16(i)
    }

    if !isGate {
        if iaOk { c.in(out, a); return } // input 123 -> a
        if aOk { c.set(out, a); return } // assignment b -> a

        c.enqueue(cmd)
        return
    }

    if !aOk || (gate != "NOT" && !bOk) { c.enqueue(cmd); return }
    
    switch gate {
    case "NOT":
        c.set(out, ^a)
    case "AND":
        c.set(out, a & b)
    case "OR":
        c.set(out, a | b)
    case "LSHIFT":
        c.set(out, a << b)
    case "RSHIFT":
        c.set(out, a >> b)
    default:
    }
}

func (c *Circuit) consumeQueue() {
    for {
        if len(c.queue) <= 0 { break }
        c.eval(c.dequeue())
    }
}

// `^((?:(?P<sa>[a-z]{0,2})\s|(?P<ia>\d{0,5})\s)?(?P<gate>AND|OR|LSHIFT|RSHIFT|NOT)\s)?(?:(?P<sb>[a-z]{0,2})|(?P<ib>\d{0,5}))\s->\s(?P<out>[a-z]{0,2})$`
func parse(s string) (bool, map[string]string) {
    // input regex
    inRegex := regexp.MustCompile(`^(?P<ia>\d{0,5})\s->\s(?P<out>[a-z]+)$`)
    match := inRegex.FindStringSubmatch(s)
    if match != nil { return true, capture(inRegex, match) }

    // wire regex
    wireRegex := regexp.MustCompile(`^(?P<sa>[a-z]+)\s->\s(?P<out>[a-z]+)$`)
    match = wireRegex.FindStringSubmatch(s)
    if match != nil { return false, capture(wireRegex, match) }

    // not regex
    notRegex := regexp.MustCompile(`^(?P<gate>NOT)\s(?:(?P<sa>[a-z]*)|(?P<ia>\d{0,5}))\s->\s(?P<out>[a-z]+)$`)
    match = notRegex.FindStringSubmatch(s)
    if match != nil { return false, capture(notRegex, match) }

    // other regex
    defltRegex := regexp.MustCompile(`^((?:(?P<sa>[a-z]+)\s|(?P<ia>\d{0,5})\s)?(?P<gate>AND|OR|LSHIFT|RSHIFT)\s)?(?:(?P<sb>[a-z]+)|(?P<ib>\d{0,5}))\s->\s(?P<out>[a-z]+)$`)
    match = defltRegex.FindStringSubmatch(s)
    if match != nil { return false, capture(defltRegex, match) }

    return false, make(map[string]string)
}

func capture(r *regexp.Regexp, match []string) map[string]string {
    captures := make(map[string]string)

    for i, name := range r.SubexpNames() {
        if i == 0 || name == "" || match[i] == "" { continue }
        captures[name] = match[i]
    }

    return captures
}

func Solve(filepath string) (int, int) {
    lines := io.Readlines(filepath)
    circuit := NewCircuit()

    for _, str := range lines {
        // get inputs first
        if isInput, cmd := parse(str); isInput { circuit.eval(cmd) } else { circuit.enqueue(cmd) }
    }

    nqueue := circuit.queue

    circuit.consumeQueue()
    a, _ := circuit.get("a")

    // reset circuit to get second star
    circuit.queue = nqueue
    circuit.clearHeap()
    circuit.in("b", a)
    
    circuit.consumeQueue()
    a2, _ := circuit.get("a")

    return int(a), int(a2)
}
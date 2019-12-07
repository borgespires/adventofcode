import os

fn main() {
    input := os.read_lines("input.txt")?
	mut m := map[string][]string

	for p in input {
        def := p.split(")")
        parent := def[0]
        child := def[1]
        
		if !(parent in m) { m[parent] = [] }
		
		// why this?
		mut a := m[parent]
		a << child
		m[parent] = a
    }

	nm := generate_orbits(m)
    println(total_orbits(nm)) // star 1
    println(minimal_transfers(nm, "YOU", "SAN")) // star 2
}

fn generate_orbits(m map[string][]string) map[string]Obj {
    mut queue := ["COM"] // root
	mut nm := map[string]Obj

    for queue.len > 0 {
        parent := queue[0]
		queue.delete(0)

		for child in m[parent] {
			queue << child
			nm[child] = Obj{ parent, nm[parent].orbits + 1 }
		}
    }
    return nm
}

fn total_orbits(m map[string]Obj) int {
    mut total := 0
    for obj in m.keys() { total += m[obj].orbits }
    return total
}

fn minimal_transfers(m map[string]Obj, obj1 string, obj2 string) int {
	path := path_to_root(m, obj1)
	mut curr := obj2

	for {
		if curr in path { break }
		curr = m[curr].parent
	}

	transfer1 := (m[obj1].orbits - m[curr].orbits) - 1
	transfer2 := (m[obj2].orbits - m[curr].orbits) - 1

	return  transfer1 + transfer2
}

fn path_to_root(m map[string]Obj, obj string) []string {
	mut path := []string
	mut curr := obj

	for {
		if !(curr in m) { break }
		path << curr
		curr = m[curr].parent
	}

	return path
}

struct Obj {
pub:
    parent string
    orbits int
}
pub fn (obj Obj) str() string {
    return "($obj.parent, $obj.orbits)"
}
import os
import math

fn main() {
	// tests
	test_angleof()

    input := os.read_lines('input.txt')?

    mut asteroids := []Coordinate
    mut y := 0
    for y < input.len {
        mut x := 0
        for x < input[y].len {
            if input[y][x] == `#` { asteroids << Coordinate{ x, y }}
            x++
        }
        y++
    }
	
	mut angles, asteroid_map := get_max_visibility_station(asteroids)
	println(angles.len) // star 1

	angles.sort_with_compare(compare_f64)
	destroyed := vaporize_asteroids(angles, asteroid_map)
	the_200 := destroyed[199]
	println(the_200.coordinate.x * 100 + the_200.coordinate.y) // star 2
}

fn create_target_map(coordinates []Coordinate, station Coordinate) ([]f64, map[string][]Asteroid) {
	mut asteroids := map[string][]Asteroid
	mut angles := []f64

    for target in coordinates {
        if station.str() == target.str() { continue }
        angle := angleof(station, target)
		distance := distanceof(station, target)

		if !(angle.str() in asteroids) {
			asteroids[angle.str()] = []
			angles << angle
		}
		
		// not working m[key] << value or { continue }
		mut temp_arr := asteroids[angle.str()]
		temp_arr << Asteroid{ target, distance }
		asteroids[angle.str()] = temp_arr
    }

    return angles, asteroids
}

fn get_max_visibility_station(coordinates []Coordinate) ([]f64, map[string][]Asteroid) {
	mut max_asteroids_insight := 0
	mut station := Coordinate{ -1, -1 }

    for location in coordinates {
		angles, _ := create_target_map(coordinates, location)
		if angles.len > max_asteroids_insight {
			max_asteroids_insight = angles.len
			station = location
		}
    }
	
	return create_target_map(coordinates, station)
}


fn vaporize_asteroids(angles []f64, asteroid_map map[string][]Asteroid) []Asteroid {
	mut sorted_asteroids := sort_asteroids(asteroid_map)

	mut i:= 0
	mut destroyed := []Asteroid
	for sorted_asteroids.keys().len > 0 {
		angle := angles[i].str()
		
		if angle in sorted_asteroids {
			destroyed << sorted_asteroids[angle][0]
			
			// not working m[key] << value or { continue }
			mut temp_arr := sorted_asteroids[angle]
			temp_arr.delete(0)
			sorted_asteroids[angle] = temp_arr
			
			if sorted_asteroids[angle].len == 0 { sorted_asteroids.delete(angle) }
		}
		
		i = (i + 1) % angles.len
	}

	return destroyed
}

fn sort_asteroids(asteroid_map map[string][]Asteroid) map[string][]Asteroid {
	mut sorted := map[string][]Asteroid
	
	for k in asteroid_map.keys() {
		// not working m[key] << value or { continue }
		mut temp_arr := []Asteroid
		temp_arr << asteroid_map[k]
		temp_arr.sort_with_compare(compare_asteroid)
		sorted[k] = temp_arr
	}
	
	return sorted
}

struct Coordinate {
    x int
    y int
}
pub fn (c Coordinate) str() string {
    return "($c.x, $c.y)"
}

struct Asteroid {
    coordinate Coordinate
	dst f64
}
pub fn (a Asteroid) str() string {
    return "($a.coordinate, $a.dst)"
}

fn compare_asteroid(a, b &Asteroid) int {
	if (*a).dst < (*b).dst { return -1 }
	if (*a).dst > (*b).dst { return 1 }
	return 0
}

fn compare_f64(a, b &f64) int {
	if *a < *b { return -1 }
	if *a > *b { return 1 }
	return 0
}

fn distanceof(p1 Coordinate, p2 Coordinate) f64 {
	dx := (p1.x - p2.x)
	dy := (p1.y - p2.y)
	return math.sqrt(dy * dy + dx * dx)
}

/**
 * Work out the angle from the x horizontal winding clockwise 
 * in screen space. 
 *
 * NOTE: We have Y as positive below, thus the Y values are inverted.
 * 
 * x,y -------------	x,y -------------
 *     |  1,1				|  1,1 --- 2,1 (90)
 *     |    \				|   |
 *     |     \				|   |
 *     |     2,2 (135)		|  2,1 (180)
 *
 * @return - a float from 0 to 360
 */
fn angleof(p1 Coordinate, p2 Coordinate) f64 {
    dy := f64(p1.y - p2.y)
    dx := f64(p2.x - p1.x)
    result := math.degrees(math.atan2(dx, dy))
    if result < 0 { return f64(360) + result }
    return result
}

fn test_angleof() {
    assert angleof(Coordinate{ 1, 1 }, Coordinate{ 1, 0 }) == 0
    assert angleof(Coordinate{ 1, 1 }, Coordinate{ 2, 1 }) == 90
	assert angleof(Coordinate{ 1, 1 }, Coordinate{ 2, 2 }) == 135
	assert angleof(Coordinate{ 1, 1 }, Coordinate{ 1, 2 }) == 180
    assert angleof(Coordinate{ 1, 1 }, Coordinate{ 0, 1 }) == 270
}
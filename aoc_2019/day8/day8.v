import os

const (
	width = 25 
	height = 6
)

fn main() {
	input := os.read_file('input.txt')?

	println(corruption_check(input)) // star 1
	print_image(generate_image(input)) // star 2
}

fn corruption_check(image string) int {
	size := width * height
	
	mut max := size
	mut res := -1
	
	mut i := 0
	for i < image.len {
		layer := image[i..i+size]
		zeros := count_pixel(layer, `0`)

		if max > zeros {
			res = count_pixel(layer, `1`) * count_pixel(layer, `2`)
			max = zeros
		}
		i += size
	}

	return res
}

fn count_pixel(line string, pixel byte) int {
	mut total := 0
	for p in line {
		if p == pixel { total++ }
	}

	return total
}

fn generate_image(layers string) string {
	size := width * height
	mut image := [`2`].repeat(size)
	
	mut i := 0
	for i < layers.len {
		layer := layers[i..i+size]
		
		mut pixel := 0
		for pixel < size {
			if image[pixel] == `2` { image[pixel] = layer[pixel] }
			pixel++
		}

		i += size
	}
	return string(image)
}

fn print_image(image string) {
	mut i := 0
	show := image.replace("0", " ").replace("1", "+")
	for i < show.len {
		layer := show[i..i+width]
		println(layer)
		i += width
	}
}
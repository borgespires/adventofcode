#[derive(Debug)]
struct Grid {
    cells: Vec<Vec<i32>>,
}

fn cell_power(x: usize, y: usize, serial: i32) -> i32 {
    let rack_id = x as i32 + 10;
    ((rack_id * y as i32 + serial) * rack_id / 100) % 10 - 5
}

fn summed_area_table(height: usize, width: usize, serial: i32) -> Vec<Vec<i32>> {
    let mut matrix = vec![vec![0; width]; height];

    for y in 0..height as usize {
        for x in 0..width as usize {
            let mut value = cell_power(x, y, serial);

            if y > 0 {
                value += matrix[y - 1][x];
            }
            if x > 0 {
                value += matrix[y][x - 1];
            }
            if x > 0 && y > 0 {
                value -= matrix[y - 1][x - 1];
            }
            matrix[y][x] = value;
        }
    }
    matrix
}

impl Grid {
    fn new(size: usize, serial: i32) -> Grid {
        Grid {
            cells: summed_area_table(size, size, serial),
        }
    }

    fn max_square_with_size(&self, size: usize) -> ((usize, usize), i32) {
        let mut corner = (0, 0);
        let mut max_power = std::i32::MIN;

        for y in 0..self.cells.len() - size {
            for x in 0..self.cells[y].len() - size {
                let power = self.power_of_area(x, y, size);
                if power > max_power {
                    max_power = power;
                    corner = (x, y);
                }
            }
        }
        (corner, max_power)
    }

    fn max_square(&self) -> ((usize, usize, usize), i32) {
        let mut corner = (0, 0, 0);
        let mut max_power = std::i32::MIN;

        for size in 1..self.cells.len() {
            let ((x, y), power) = self.max_square_with_size(size);
            if power > max_power {
                max_power = power;
                corner = (x, y, size);
            }
        }
        (corner, max_power)
    }

    fn power_of_area(&self, top_left_x: usize, top_left_y: usize, size: usize) -> i32 {
        let bottom_right_x = top_left_x + size - 1;
        let bottom_right_y = top_left_y + size - 1;

        let mut power = self.cells[bottom_right_y][bottom_right_x];

        if top_left_x > 0 {
            power -= self.cells[bottom_right_y][top_left_x - 1];
        }
        if top_left_y > 0 {
            power -= self.cells[top_left_y - 1][bottom_right_x];
        }
        if top_left_x > 0 && top_left_y > 0 {
            power += self.cells[top_left_y - 1][top_left_x - 1];
        }
        power
    }
}

fn main() {
    let grid = Grid::new(300, 5791);
    println!("{:?}", grid.max_square_with_size(3));
    println!("{:?}", grid.max_square());
}

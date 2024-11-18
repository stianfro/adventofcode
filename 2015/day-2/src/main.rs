use std::{fs, io};

fn read_input(file_path: &str) -> io::Result<String> {
    fs::read_to_string(file_path)
}

fn main() -> io::Result<()> {
    let file_path =
        "/Users/stianfroystein/src/github.com/stianfro/adventofcode/2015/day-2/input.txt";

    let input = read_input(file_path)?;
    let input_slice: Vec<&str> = input.split('\n').collect();

    part_one(input_slice.clone());
    // part_two(input_slice.clone());

    Ok(())
}

fn part_one(presents: Vec<&str>) {
    let mut wrapping_total = 0;

    for present in presents {
        if present.is_empty() {
            break;
        }

        let dimensions: Vec<&str> = present.split('x').collect();

        let l = dimensions[0];
        let w = dimensions[1];
        let h = dimensions[2];

        let l_int: i32 = l.parse().expect("not a valid number");
        let w_int: i32 = w.parse().expect("not a valid number");
        let h_int: i32 = h.parse().expect("not a valid number");

        // TODO: create slice of the sides
        let wrapping = 2 * l_int * w_int + 2 * w_int * h_int + 2 * h_int * l_int;
        wrapping_total += wrapping;
    }

    println!("part one: {wrapping_total}");
}

// fn part_two(dimensions: Vec<char>) {}

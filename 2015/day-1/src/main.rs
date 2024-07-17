use std::{fs, io};

fn read_input(file_path: &str) -> io::Result<String> {
    fs::read_to_string(file_path)
}

fn main() -> io::Result<()> {
    let directions = read_input("input.txt")?;
    let directions_slice: Vec<char> = directions.chars().collect();

    part_one(directions_slice.clone());
    part_two(directions_slice.clone());

    Ok(())
}

fn part_one(directions: Vec<char>) {
    let mut current_floor = 0;

    for direction in directions {
        if direction == '(' {
            current_floor += 1;
        } else if direction == ')' {
            current_floor -= 1;
        }
    }

    println!("part one: {}", current_floor);
}

fn part_two(directions: Vec<char>) {
    let mut current_floor = 0;

    for (index, direction) in directions.iter().enumerate() {
        if *direction == '(' {
            current_floor += 1;
        } else if *direction == ')' {
            current_floor -= 1;
        }

        if current_floor == -1 {
            println!("part two: {}", index + 1);
            break;
        }
    }
}

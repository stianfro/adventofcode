use std::{fs, io};

fn read_input(file_path: &str) -> io::Result<String> {
    fs::read_to_string(file_path)
}

fn main() -> io::Result<()> {
    let directions = read_input("input.txt")?;
    let directions_slice: Vec<char> = directions.chars().collect();
    let mut current_position = 0;

    for dir in directions_slice {
        if dir == '(' {
            current_position += 1;
        } else if dir == ')' {
            current_position -= 1;
        }
    }

    println!("answer: {}", current_position);

    Ok(())
}

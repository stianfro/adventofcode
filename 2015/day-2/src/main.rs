use std::{fs, io};

fn read_input(file_path: &str) -> io::Result<String> {
    fs::read_to_string(file_path)
}

fn main() -> io::Result<()> {
    let input = read_input("input.txt")?;
    let input_slice: Vec<char> = input.chars().collect();

    part_one(input_slice.clone());
    part_two(input_slice.clone());

    Ok(())
}

fn part_one(input: Vec<char>) {}

fn part_two(input: Vec<char>) {}
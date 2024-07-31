use std::{fs, io};

fn read_input(file_path: &str) -> io::Result<String> {
    fs::read_to_string(file_path)
}

fn main() -> io::Result<()> {
    let directions = read_input("input.txt")?;
    let directions_slice: Vec<char> = directions.chars().collect();

    println!("part one: {}", part_one(directions_slice.clone()));
    println!("part two: {}", part_two(directions_slice.clone()));

    Ok(())
}

fn part_one(directions: Vec<char>) -> String {
    let mut current_floor = 0;

    for direction in directions {
        if direction == '(' {
            current_floor += 1;
        } else if direction == ')' {
            current_floor -= 1;
        }
    }

    format!("{current_floor}")
}

fn part_two(directions: Vec<char>) -> String {
    let mut current_floor = 0;

    for (index, direction) in directions.iter().enumerate() {
        if *direction == '(' {
            current_floor += 1;
        } else if *direction == ')' {
            current_floor -= 1;
        }

        if current_floor == -1 {
            return format!("{}", index + 1);
        }
    }

    String::from("")
}

#[test]
fn test_part_one() {
    struct Test<'a> {
        input: &'a str,
        want: &'a str,
    }

    let tests = vec![
        Test {
            input: "(())",
            want: "0",
        },
        Test {
            input: "(((",
            want: "3",
        },
    ];

    for test in tests {
        let input_slice = test.input.chars().collect();
        let have = part_one(input_slice);

        assert_eq!(test.want, have)
    }
}

#[test]
fn test_part_two() {
    struct Test<'a> {
        input: &'a str,
        want: &'a str,
    }

    let tests = vec![
        Test {
            input: ")",
            want: "1",
        },
        Test {
            input: "()())",
            want: "5",
        },
    ];

    for test in tests {
        let input_slice = test.input.chars().collect();
        let have = part_two(input_slice);

        assert_eq!(test.want, have)
    }
}

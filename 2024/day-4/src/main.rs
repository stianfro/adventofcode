use std::fs;

fn main() {
    let input =
        fs::read_to_string("./input-example.txt").expect("Should have been able to read the file");

    println!("{}", part1(&input))
}

fn part1(input: &str) -> i32 {
    let lines = input.lines();
    for line in lines {
        println!("{line}")
    }

    1
}

#[test]
fn test_part1() {
    let input =
        fs::read_to_string("./input-example.txt").expect("Should have been able to read the file");

    assert_eq!(part1(&input), 18)
}

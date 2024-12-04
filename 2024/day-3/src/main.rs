// extern crate regex;

use regex::Regex;
use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("Should have been able to read the file");

    println!("Part 2: {}", part1(&input));
    println!("Part 2: {}", part2(&input));
}

fn part1(input: &str) -> i32 {
    let mut result = 0;

    let lines = input.lines();
    for line in lines {
        let re = Regex::new(r"mul\(\d+,\d+\)").unwrap();

        for mul in re.captures_iter(line) {
            result += multiply(String::from(&mul[0]))
        }
    }

    result
}

fn part2(input: &str) -> i32 {
    let mut result = 0;
    let mut enabled = true;

    let lines = input.lines();
    for line in lines {
        let re = Regex::new(r"(mul\(\d+,\d+\))|(don't\(\))|(do\(\))").unwrap();

        for mul in re.captures_iter(line) {
            let m = String::from(&mul[0]);

            if m == "don't()" {
                enabled = false;
                continue;
            }

            if m == "do()" {
                enabled = true;
                continue;
            }

            if enabled {
                result += multiply(m)
            }
        }
    }

    result
}

fn multiply(m: String) -> i32 {
    let m = m.replace("mul(", "");
    let m = m.replace(")", "");

    let xy = m.split(",").collect::<Vec<_>>();

    let x: i32 = xy[0].parse().unwrap();
    let y: i32 = xy[1].parse().unwrap();

    x * y
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_1() {
        let input = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))";

        assert_eq!(part1(input), 161)
    }

    #[test]
    fn test_part_2() {
        let input = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))";

        assert_eq!(part2(input), 48)
    }
}

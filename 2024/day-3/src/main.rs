// extern crate regex;

use regex::Regex;
use std::fs;

fn main() {
    let input =
        fs::read_to_string("./input-example.txt").expect("Should have been able to read the file");

    let lines = input.lines();
    for line in lines {
        let re = Regex::new(r"mul\((?P<x>\d+),(?P<y>\d+\))").unwrap();

        for mul in re.captures_iter(line) {
            let m = String::from(&mul[0]);
            println!("{}", m)
        }
    }
}

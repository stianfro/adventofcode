use std::fs;

fn main() {
    let input = fs::read_to_string("./input-example.txt").expect("no error");

    let sections = input.split("\n\n").collect::<Vec<_>>();
    // let rules = sections[0].split("\n").collect::<Vec<_>>();

    let section_one = sections[0].lines();
    // let section_two = sections[1].lines();

    for line in section_one {
        let rule = line.split("|").collect::<Vec<_>>();

        println!("{line}: Page number {} & Page Number {}", rule[0], rule[1])
    }
}

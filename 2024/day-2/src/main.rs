use std::fs;

fn main() {
    let input =
        fs::read_to_string("./input-example.txt").expect("Should have been able to read the file");

    let p1 = part_1(&input);

    println!("Part 1: {p1}");
}

// 1 2 3 4 5 <- report
// ^
// level

fn part_1(input: &str) -> i32 {
    let data = input.lines();

    for report in data {
        let levels = report.split(' ').collect::<Vec<_>>();

        for (i, &level) in levels.iter().enumerate() {
            if level.is_empty() {
                continue;
            }

            println!("{i}: {level}")
        }
    }

    1
}

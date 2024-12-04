use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("Should have been able to read the file");

    let p1 = part_1(&input);
    let p2 = part_2(&input);

    println!("Part 1: {p1}");
    println!("Part 2: {p2}");
}

fn part_1(input: &str) -> i32 {
    let lines = input.lines();

    let mut safe_reports = 0;

    for line in lines {
        let levels = line.split(' ').collect::<Vec<_>>();
        let mut levels_int: Vec<i32> = Vec::new();

        for lvl in levels {
            let lvl_int: i32 = lvl.parse().unwrap();
            levels_int.push(lvl_int);
        }

        if is_safe(levels_int) {
            safe_reports += 1;
        }
    }

    safe_reports
}

fn is_safe(report: Vec<i32>) -> bool {
    let mut report_sorted = report.clone();
    let mut report_sorted_reverse = report.clone();

    report_sorted.sort();
    report_sorted_reverse.sort_by(|a, b| b.cmp(a));

    let increasing_or_decreasing = report == report_sorted || report == report_sorted_reverse;
    let mut safe = true;

    for (i, _) in report.iter().enumerate() {
        if i + 1 == report.len() {
            continue;
        }

        let diff = (report[i] - report[i + 1]).abs();

        if !(1..=3).contains(&diff) {
            safe = false
        }
    }

    increasing_or_decreasing && safe
}

fn part_2(input: &str) -> i32 {
    let lines = input.lines();

    let mut safe_reports = 0;

    for line in lines {
        let levels = line.split(' ').collect::<Vec<_>>();
        let mut report: Vec<i32> = Vec::new();

        for lvl in levels {
            report.push(lvl.parse().unwrap());
        }

        let mut safe = false;

        for (i, _) in report.iter().enumerate() {
            let mut vec1 = report[0..i].to_vec();
            let mut vec2 = report[i + 1..report.len()].to_vec();

            vec1.append(&mut vec2);

            if is_safe(vec1) {
                safe = true
            }
        }

        if safe {
            safe_reports += 1
        }
    }

    safe_reports
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part_1() {
        let input = String::from(
            "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9",
        );
        let result = part_1(&input);

        assert_eq!(result, 2)
    }

    #[test]
    fn test_part_2() {
        let input = String::from(
            "44 47 50 51 53 54 53
70 73 75 77 80 81 84 84
1 3 4 7 10 13 16 20
47 49 52 53 55 57 60 65
69 70 71 70 71
22 23 20 21 24 27 24",
        );
        let result = part_2(&input);

        assert_eq!(result, 4)
    }
}

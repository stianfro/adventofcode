use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("Should have been able to read the file");

    let p1 = part_1(&input);
    let p2 = part_2(&input);

    println!("Part 1: {p1}");
    println!("Part 2: {p2}");
}

// 1 2 3 4 5 <- report
// ^
// level

fn part_1(input: &str) -> i32 {
    let data = input.lines();

    let mut safe_reports = 0;

    for report in data {
        let levels = report.split(' ').collect::<Vec<_>>();

        let mut safe = true;

        for (i, &level) in levels.iter().enumerate() {
            if level.is_empty() || i == 0 {
                continue;
            }

            let curr: i32 = level.parse().unwrap();
            let prev: i32 = levels[i - 1].parse().unwrap();

            let diff = curr - prev;

            let mut diff_prev: i32 = 0;

            if i >= 2 {
                let prev_prev: i32 = levels[i - 2].parse().unwrap();

                diff_prev = prev - prev_prev
            }

            if !within_limits(diff) {
                safe = false;
            }

            if !same_direction(diff, diff_prev) {
                safe = false;
            }
        }

        if !safe {
            continue;
        }

        safe_reports += 1;
    }

    safe_reports
}

fn part_2(input: &str) -> i32 {
    let data = input.lines();

    let mut safe_reports = 0;

    for report in data {
        let levels = report.split(' ').collect::<Vec<_>>();
        let mut levels_r2 = levels.clone();

        let mut unsafe_count = 0;

        for (i, &level) in levels.iter().enumerate() {
            if level.is_empty() || i == 0 {
                continue;
            }

            let curr: i32 = level.parse().unwrap();
            let prev: i32 = levels[i - 1].parse().unwrap();

            let diff = curr - prev;

            let mut diff_prev: i32 = 0;

            if i >= 2 {
                let prev_prev: i32 = levels[i - 2].parse().unwrap();

                diff_prev = prev - prev_prev
            }

            if !within_limits(diff) {
                if unsafe_count < 1 {
                    levels_r2.remove(i);
                }

                unsafe_count += 1
            }

            if !same_direction(diff, diff_prev) {
                if unsafe_count < 1 {
                    if diff >= 0 {
                        println!("Removing {prev} from '{report}'");
                        levels_r2.remove(i - 1);
                    } else {
                        println!("Removing {curr} from '{report}'");
                        levels_r2.remove(i);
                    }
                }

                unsafe_count += 1
            }
        }

        let mut safe = true;

        for (i, &level) in levels_r2.iter().enumerate() {
            if level.is_empty() || i == 0 {
                continue;
            }

            let curr: i32 = level.parse().unwrap();
            let prev: i32 = levels_r2[i - 1].parse().unwrap();

            let diff = curr - prev;

            let mut diff_prev: i32 = 0;

            if i >= 2 {
                let prev_prev: i32 = levels_r2[i - 2].parse().unwrap();

                diff_prev = prev - prev_prev
            }

            if !within_limits(diff) {
                safe = false;
            }

            if !same_direction(diff, diff_prev) {
                safe = false;
            }
        }

        if !safe {
            // println!("Unsafe: {report}");
            continue;
        }

        // println!("Safe: {report}");

        safe_reports += 1;
    }

    safe_reports
}

fn same_direction(curr: i32, prev: i32) -> bool {
    if curr <= 0 && prev <= 0 {
        return true;
    }

    if curr >= 0 && prev >= 0 {
        return true;
    }

    false
}

fn within_limits(diff: i32) -> bool {
    if diff == 0 {
        return false;
    }

    if (-3..=3).contains(&diff) {
        return true;
    }

    false
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_within_limits() {
        // ok
        assert!(within_limits(-3));
        assert!(within_limits(3));

        // not ok
        assert!(!within_limits(-5));
        assert!(!within_limits(-4));
    }

    #[test]
    fn test_same_direction() {
        // ok
        assert!(same_direction(3, 2));
        assert!(same_direction(1, 1));

        // not ok
        assert!(!same_direction(-1, 1));
        assert!(!same_direction(1, -1));
    }

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
    fn test_part_2_example() {
        let input = String::from(
            "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9",
        );
        let result = part_2(&input);

        assert_eq!(result, 4)
    }

    #[test]
    fn test_part_2_special() {
        let input = String::from("44 47 50 51 53 54 53");
        let result = part_2(&input);

        assert_eq!(result, 1)
    }
}

use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("Should have been able to read the file");

    println!("{}", part1(&input));
    println!("{}", part2(&input));
}

fn part1(input: &str) -> i32 {
    let lines = input.split("\n").collect::<Vec<_>>();

    let mut xmas = 0;

    for (y, _) in lines.iter().enumerate() {
        let lv = lines[y].split("").collect::<Vec<_>>();

        for (x, _) in lv.iter().enumerate() {
            if lv[x] != "X" {
                continue;
            }

            // forward
            if x + 3 <= lv.len() && lv[x + 1] == "M" && lv[x + 2] == "A" && lv[x + 3] == "S" {
                xmas += 1
            }

            // backward
            if x >= 3 && lv[x - 1] == "M" && lv[x - 2] == "A" && lv[x - 3] == "S" {
                xmas += 1
            }

            // up
            if y >= 3 {
                let lvup1 = lines[y - 1].split("").collect::<Vec<_>>();
                let lvup2 = lines[y - 2].split("").collect::<Vec<_>>();
                let lvup3 = lines[y - 3].split("").collect::<Vec<_>>();

                if lvup1[x] == "M" && lvup2[x] == "A" && lvup3[x] == "S" {
                    xmas += 1
                }
            }

            // down
            if y + 3 < lines.len() {
                let lvdown1 = lines[y + 1].split("").collect::<Vec<_>>();
                let lvdown2 = lines[y + 2].split("").collect::<Vec<_>>();
                let lvdown3 = lines[y + 3].split("").collect::<Vec<_>>();

                if lvdown1[x] == "M" && lvdown2[x] == "A" && lvdown3[x] == "S" {
                    xmas += 1
                }
            }

            // diag forward up
            if y >= 3 && x + 3 <= lv.len() {
                let lvup1 = lines[y - 1].split("").collect::<Vec<_>>();
                let lvup2 = lines[y - 2].split("").collect::<Vec<_>>();
                let lvup3 = lines[y - 3].split("").collect::<Vec<_>>();

                if lvup1[x + 1] == "M" && lvup2[x + 2] == "A" && lvup3[x + 3] == "S" {
                    xmas += 1
                }
            }

            // diag forward down
            if y + 3 < lines.len() && x + 3 <= lv.len() {
                let lvdown1 = lines[y + 1].split("").collect::<Vec<_>>();
                let lvdown2 = lines[y + 2].split("").collect::<Vec<_>>();
                let lvdown3 = lines[y + 3].split("").collect::<Vec<_>>();

                // TODO: might not work on all input
                if lvdown3.len() < 3 {
                    continue;
                }

                if lvdown1[x + 1] == "M" && lvdown2[x + 2] == "A" && lvdown3[x + 3] == "S" {
                    xmas += 1
                }
            }

            // diag backward up
            if x >= 3 && y >= 3 {
                let lvup1 = lines[y - 1].split("").collect::<Vec<_>>();
                let lvup2 = lines[y - 2].split("").collect::<Vec<_>>();
                let lvup3 = lines[y - 3].split("").collect::<Vec<_>>();

                if lvup3.len() < 3 {
                    continue;
                }

                if lvup1[x - 1] == "M" && lvup2[x - 2] == "A" && lvup3[x - 3] == "S" {
                    xmas += 1
                }
            }

            // diag backward down
            if y + 3 < lines.len() && x >= 3 {
                let lvdown1 = lines[y + 1].split("").collect::<Vec<_>>();
                let lvdown2 = lines[y + 2].split("").collect::<Vec<_>>();
                let lvdown3 = lines[y + 3].split("").collect::<Vec<_>>();

                if lvdown3.len() < 3 {
                    continue;
                }

                if lvdown1[x - 1] == "M" && lvdown2[x - 2] == "A" && lvdown3[x - 3] == "S" {
                    xmas += 1
                }
            }
        }
    }

    xmas
}

fn part2(input: &str) -> i32 {
    let lines = input.split("\n").collect::<Vec<_>>();

    let mut xmas = 0;

    for (y, _) in lines.iter().enumerate() {
        let lv = lines[y].split("").collect::<Vec<_>>();

        for (x, _) in lv.iter().enumerate() {
            if y + 2 >= lines.len() {
                continue;
            }

            if x + 2 >= lv.len() {
                continue;
            }

            let lvdown1 = lines[y + 1].split("").collect::<Vec<_>>();
            let lvdown2 = lines[y + 2].split("").collect::<Vec<_>>();

            // M.S
            // .A.
            // M.S
            if lv[x] == "M"
                && lv[x + 2] == "S"
                && lvdown1[x + 1] == "A"
                && lvdown2[x] == "M"
                && lvdown2[x + 2] == "S"
            {
                xmas += 1;
                continue;
            }

            // S.S
            // .A.
            // M.M
            if lv[x] == "S"
                && lv[x + 2] == "S"
                && lvdown1[x + 1] == "A"
                && lvdown2[x] == "M"
                && lvdown2[x + 2] == "M"
            {
                xmas += 1;
                continue;
            }

            // M.M
            // .A.
            // S.S
            if lv[x] == "M"
                && lv[x + 2] == "M"
                && lvdown1[x + 1] == "A"
                && lvdown2[x] == "S"
                && lvdown2[x + 2] == "S"
            {
                xmas += 1;
                continue;
            }

            // S.M
            // .A.
            // S.M
            if lv[x] == "S"
                && lv[x + 2] == "M"
                && lvdown1[x + 1] == "A"
                && lvdown2[x] == "S"
                && lvdown2[x + 2] == "M"
            {
                xmas += 1;
                continue;
            }
        }
    }

    xmas
}

#[test]
fn test_part1() {
    let input =
        fs::read_to_string("./input-example.txt").expect("Should have been able to read the file");
    let input2 = "..X...
.SAMX.
.A..A.
XMAS.S
.X....";
    assert_eq!(part1(&input), 18);
    assert_eq!(part1(input2), 4)
}

#[test]
fn test_part2() {
    let input =
        fs::read_to_string("./input-example.txt").expect("Should have been able to read the file");
    assert_eq!(part2(&input), 9)
}

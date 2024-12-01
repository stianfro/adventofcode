use std::fs;

fn main() {
    let input = fs::read_to_string("./input.txt").expect("Should have been able to read the file");

    let p1 = part_1(&input);
    let p2 = part_2(&input);

    println!("Part 1: {p1}");
    println!("Part 2: {p2}");
}

fn part_1(input: &str) -> i32 {
    let data = input.split('\n');

    let mut l_list: Vec<i32> = Vec::new();
    let mut r_list: Vec<i32> = Vec::new();

    for element in data {
        if element == "\n" || element.is_empty() {
            continue;
        }

        let line_split = element.split("   ");
        let mut line_vector: Vec<&str> = Vec::new();

        for entry in line_split {
            line_vector.push(entry);
        }

        let l = line_vector[0];
        let r = line_vector[1];

        let l_int: i32 = l.parse().unwrap();
        let r_int: i32 = r.parse().unwrap();

        l_list.push(l_int);
        r_list.push(r_int);
    }

    l_list.sort();
    r_list.sort();

    let l_len = l_list.len();
    let r_len = r_list.len();

    if l_len != r_len {
        println!("[ERROR] left and right lists do not have same length");
    }

    let mut total_distance = 0;

    for (pos, _) in l_list.iter().enumerate() {
        let l = l_list[pos];
        let r = r_list[pos];

        let mut v = [l, r];

        v.sort();

        let distance = v[1] - v[0];

        total_distance += distance
    }

    total_distance
}

fn part_2(input: &str) -> i32 {
    let data = input.split('\n');

    let mut l_list: Vec<i32> = Vec::new();
    let mut r_list: Vec<i32> = Vec::new();

    for element in data {
        if element == "\n" || element.is_empty() {
            continue;
        }

        let line_split = element.split("   ");
        let mut line_vector: Vec<&str> = Vec::new();

        for entry in line_split {
            line_vector.push(entry);
        }

        let l = line_vector[0];
        let r = line_vector[1];

        let l_int: i32 = l.parse().unwrap();
        let r_int: i32 = r.parse().unwrap();

        l_list.push(l_int);
        r_list.push(r_int);
    }

    let l_len = l_list.len();
    let r_len = r_list.len();

    if l_len != r_len {
        println!("[ERROR] left and right lists do not have same length");
    }

    let mut total_similarity_score = 0;

    for (pos, _) in l_list.iter().enumerate() {
        let l = l_list[pos];

        let mut occurences = 0;

        for r in &r_list {
            if l == *r {
                occurences += 1
            }
        }

        let similarity_score = l * occurences;

        println!("{l} * {occurences} = {similarity_score}");

        total_similarity_score += similarity_score
    }

    total_similarity_score
}

#[cfg(test)]
mod tests {
    use super::*;

    struct Test {
        input: String,
        want: i32,
    }

    #[test]
    fn test_part_one() {
        let input = fs::read_to_string("./input-example.txt")
            .expect("Should have been able to read the file");

        let tests = vec![Test { input, want: 11 }];

        for t in tests {
            let got = part_1(&t.input);
            assert_eq!(got, t.want)
        }
    }

    #[test]
    fn test_part_two() {
        let input = fs::read_to_string("./input-example.txt")
            .expect("Should have been able to read the file");

        let tests = vec![Test { input, want: 31 }];

        for t in tests {
            let got = part_2(&t.input);
            assert_eq!(got, t.want)
        }
    }
}

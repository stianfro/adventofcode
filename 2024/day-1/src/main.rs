fn main() {
    let input = "3   4
4   3
2   5
1   3
3   9
3   3";

    let data = input.split('\n');

    // let mut l_list: Vec<&str> = Vec::new();
    // let mut r_list: Vec<&str> = Vec::new();

    for element in data {
        let line_split = element.split("   ");
        let mut line_vector: Vec<String> = Vec::new();

        for entry in line_split {
            let x = entry.to_string();
            line_vector.push(x);
        }

        let l = &line_vector[0];
        let r = &line_vector[1];

        println!("l: {l}, r: {r}")
    }
}

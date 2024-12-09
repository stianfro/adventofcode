use std::fs;

pub fn read(path: &str) -> String {
    fs::read_to_string(path).expect("Should have been able to read the file")
}

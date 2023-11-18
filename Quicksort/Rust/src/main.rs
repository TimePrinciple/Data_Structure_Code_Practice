#![allow(non_snake_case)]
#![allow(dead_code)]

use std::io::BufReader;
use std::io::BufRead;
use std::fs::File;
use std::str;

fn main() {
    // Open the input file
    let input_file = File::open("../input").expect("File in \"../input\" should be present");
    // Construct BufReader
    let buf_reader = BufReader::new(input_file);
    let mut input = vec![];
    for element in buf_reader.split(b' ') {
        let element = element.unwrap();
        match str::from_utf8(&element).unwrap().parse::<i32>() {
            Ok(value) => {
                input.push(value);
            }
            Err(_) => {
                continue;
            }
        }
    }
    dbg!(input);
}

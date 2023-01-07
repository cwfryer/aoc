use std::collections::HashMap;
use std::io;
use std::io::Error;

fn main() -> Result<(), Error> {
    let mut buf = String::new();
    io::stdin().read_line(&mut buf)?;
    buf = buf.trim().to_string();
    println!("part1: {:?}", part1(&buf));
    println!("part2: {:?}", part2(&buf));
    Ok(())
}

fn part1(s: &String) -> i64 {
    let crabs = parse(s);
    let farthest = *crabs.iter().max().unwrap();
    let mut hm: HashMap<i64, i64> = HashMap::new();
    for i in 0..farthest {
        let val = crabs.iter().fold(0, |acc, &c| acc + (c - i).abs());
        hm.insert(i, val);
    }

    hm.iter()
        .filter(|(_k, v)| **v >= 0)
        .min_by(|a, b| a.1.cmp(b.1))
        .map(|(_k, &v)| v)
        .unwrap()
}

fn part2(s: &String) -> i64 {
    let crabs = parse(s);
    let farthest = *crabs.iter().max().unwrap();
    let mut hm: HashMap<i64, i64> = HashMap::new();
    for i in 0..farthest {
        let val = crabs.iter().fold(0, |acc, &c| {
            let v = (c - i).abs();
            acc + (0..=v).fold(0,|a,b| a + b)
        });
        hm.insert(i, val);
    }

    hm.iter()
        .filter(|(_k, v)| **v >= 0)
        .min_by(|a, b| a.1.cmp(b.1))
        .map(|(_k, &v)| v)
        .unwrap()
}

fn parse(s: &String) -> Vec<i64> {
    s.split(",")
        .map(|x| x.parse::<i64>().unwrap())
        .collect::<Vec<i64>>()
}

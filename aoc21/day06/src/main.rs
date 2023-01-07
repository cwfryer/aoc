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
    let fish = parse(s);
    generation(80, fish)
}
fn part2(s: &String) -> i64 {
    let fish = parse(s);
    generation(256, fish)
}

fn parse(s: &String) -> HashMap<i8, i64> {
    let fish = s
        .split(",")
        .map(|c| c.parse::<i8>().unwrap())
        .collect::<Vec<i8>>();

    let mut hm: HashMap<i8, i64> = (0..9).map(|day| (day, 0)).collect();
    for f in fish {
        *hm.entry(f).or_insert(0) += 1;
    }

    hm
}

fn generation(days: usize, fish: HashMap<i8, i64>) -> i64 {
    let mut curr_gen = fish;

    for _d in 0..days {
        let mut next_gen: HashMap<i8, i64> = (0..7)
            .filter(|x| x != &6)
            .map(|day| (day as i8, 0))
            .collect::<HashMap<i8, i64>>();

        let curr0 = *curr_gen.get(&0).unwrap_or(&0);
        next_gen.insert(6, curr0);
        next_gen.insert(8, curr0);

        for age in 1..9 {
            *next_gen.entry(age - 1).or_insert(0) += curr_gen.get(&age).unwrap_or(&0);
        }

        curr_gen = next_gen.clone();
    }

    curr_gen.values().cloned().collect::<Vec<i64>>().iter().sum()
}

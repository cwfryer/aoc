use std::io::Error;

fn main() -> Result<(), Error> {
    let data = std::fs::read_to_string("./input.txt")?;
    println!("Part 1: {:?}\n", part1(&data));
    println!("Part 2: {:?}\n", part2(&data));
    Ok(())
}

fn part1(data: &String) -> i32 {
    let numbers: Vec<i32> = data
        .trim()
        .split("\n")
        .map(|x| x.parse::<i32>().unwrap())
        .collect();

    let mut last = numbers[0];
    let mut count = 0;

    for n in &numbers[1..] {
        if *n > last {
            count += 1;
        }
        last = *n
    }

    count
}

fn part2(data: &String) -> i32 {
    let groups = data.trim()
        .split("\n")
        .map(|x| x.parse::<i32>().unwrap())
        .collect::<Vec<i32>>()
        .iter()
        .collect::<Vec<_>>()
        .windows(3)
        .map(|a| a[0] + a[1] + a[2])
        .collect::<Vec<i32>>();

    let mut last = groups[0];
    let mut count = 0;
    for d in &groups[1..] {
        if *d > last {
            count += 1;
        }
        last = *d
    }

    count
}

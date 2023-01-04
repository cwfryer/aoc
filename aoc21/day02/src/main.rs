use std::io::Error;

fn main() -> Result<(), Error> {
    let data = std::fs::read_to_string("./input.txt")?;
    println!("Part 1: {:?}\n", part1(&data));
    println!("Part 2: {:?}\n", part2(&data));
    Ok(())
}

fn part1(data: &String) -> i32 {
    let (h, v) = data
        .trim()
        .split("\n")
        .collect::<Vec<&str>>()
        .iter()
        .fold((0, 0), |acc, &s| parse_part1(acc, s));

    h * v
}

fn part2(data: &String) -> i32 {
    let (h, v, _a) = data
        .trim()
        .split("\n")
        .collect::<Vec<&str>>()
        .iter()
        .fold((0, 0, 0), |acc, &s| parse_part2(acc, s));

    h * v
}

fn parse_part1(mut hv: (i32, i32), s: &str) -> (i32, i32) {
    let s_vec = s.split(" ").collect::<Vec<&str>>();
    match s_vec[0] {
        "forward" => hv.0 += s_vec[1].parse::<i32>().unwrap(),
        "down" => hv.1 += s_vec[1].parse::<i32>().unwrap(),
        "up" => hv.1 -= s_vec[1].parse::<i32>().unwrap(),
        _ => unreachable!(),
    }

    hv
}
fn parse_part2(mut hv: (i32, i32, i32), s: &str) -> (i32, i32, i32) {
    let s_vec = s.split(" ").collect::<Vec<&str>>();
    match s_vec[0] {
        "forward" => {
            hv.0 += s_vec[1].parse::<i32>().unwrap();
            hv.1 += s_vec[1].parse::<i32>().unwrap() * hv.2
        },
        "down" => hv.2 += s_vec[1].parse::<i32>().unwrap(),
        "up" => hv.2 -= s_vec[1].parse::<i32>().unwrap(),
        _ => unreachable!(),
    }

    hv
}

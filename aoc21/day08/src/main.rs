use std::collections::{HashMap, HashSet};
use std::io::Error;
use std::io::{self, Read};

fn main() -> Result<(), Error> {
    let mut buf = String::new();
    io::stdin().read_to_string(&mut buf)?;
    buf = buf.trim().to_string();
    println!("part1: {:?}", part1(&buf));
    println!("part2: {:?}", part2(&buf));
    Ok(())
}

fn part1(s: &String) -> usize {
    let mut count = 0;
    for line in s.lines() {
        let v = line.split(" | ").skip(1).fold(0, |acc, x| {
            let c = x.split_whitespace().fold(0, |bcc, y| {
                if y.len() == 2 || y.len() == 3 || y.len() == 4 || y.len() == 7 {
                    bcc + 1
                } else {
                    bcc
                }
            });
            acc + c
        });
        count += v
    }
    count
}

fn part2(s: &String) -> i32 {
    let mut answer = 0;
    let mut num_map: HashMap<i32, HashSet<char>> = HashMap::new();
    for line in s.lines() {
        let nums = line.split_whitespace();
        num_map.insert(
            1,
            nums.clone()
                .filter(|x| x.len() == 2)
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            4,
            nums.clone()
                .filter(|x| x.len() == 4)
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            7,
            nums.clone()
                .filter(|x| x.len() == 3)
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            8,
            nums.clone()
                .filter(|x| x.len() == 7)
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            3,
            nums.clone()
                .filter(|x| x.len() == 5 && has_one(x, num_map.get(&1).unwrap()))
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            0,
            nums.clone()
                .filter(|x| x.len() == 6 && has_one(x, num_map.get(&1).unwrap()) && !has_match(x, num_map.get(&3).unwrap()))
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            6,
            nums.clone()
                .filter(|x| x.len() == 6 && !has_one(x, num_map.get(&1).unwrap()))
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            5,
            nums.clone()
                .filter(|x| x.len() == 5 && contained_in(x, num_map.get(&6).unwrap()))
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            2,
            nums.clone()
                .filter(|x| x.len() == 5 && !contained_in(x, num_map.get(&6).unwrap()) && !has_one(x, num_map.get(&1).unwrap()))
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );
        num_map.insert(
            9,
            nums.clone()
                .filter(|x| x.len() == 6 && has_match(x, num_map.get(&4).unwrap()))
                .next()
                .unwrap()
                .chars()
                .collect::<HashSet<char>>(),
        );

        let out = line.split(" | ").skip(1).map(|o| {
            o.split_whitespace().map(|d| {
                let h = d.chars().collect::<HashSet<char>>();
                num_map.iter().find_map(|(k,v)| if *v == h {Some(*k)} else {None}).unwrap_or(0)
            }).collect::<Vec<i32>>()
        }).collect::<Vec<Vec<i32>>>();
        answer += out[0][0]*1000 + out[0][1]*100 + out[0][2]*10 + out[0][3];
    }
    answer
}

fn has_one(c: &str, one: &HashSet<char>) -> bool {
    let one_vec = one.iter().collect::<Vec<&char>>();
    c.contains(*one_vec[0]) && c.contains(*one_vec[1])
}

fn has_match(c: &str, four: &HashSet<char>) -> bool {
    let four_vec = four.iter().collect::<Vec<&char>>();

    for cc in four_vec {
        if !c.contains(*cc) {
            return false
        }
    }
    true
}

fn contained_in(c: &str, container: &HashSet<char>) -> bool {
    let c_vec = c.chars().collect::<Vec<char>>();
    for cc in c_vec {
        if !container.contains(&cc) {
            return false;
        }
    }
    true
}

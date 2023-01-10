use std::{io::{Error, self, Read}, collections::BinaryHeap};

fn main() {
    let data = parse().unwrap();
    println!("Part 1: {:?}", part1(&data));
    println!("Part 2: {:?}",part2(&data));
}

enum LineStatus {
    Corrupted(u32),
    Incomplete(u64),
}
use LineStatus::*;

fn score_corrupt(closing_char: char) -> u32 {
    match closing_char {
        ')'=> 3,
        ']'=> 57, 
        '}'=> 1197,
        '>'=> 25137,
        _ => panic!("unexpected char")
    }
}

fn score_incomplete(completion_string: Vec<char>) -> u64 {
    let mut score = 0;
    for c in completion_string {
        score *= 5;
        match c {
            '(' => score += 1,
            '[' => score += 2,
            '{' => score += 3,
            '<' => score += 4,
            _ => (),
        }
    }
    score
}

fn chars_match(opening: char, closing: char) -> bool {
    match (opening,closing) {
        ('(',')') | ('{','}') | ('[',']') | ('<','>') => true,
        _ => false,
    }
}
fn parse() -> Result<Vec<Vec<char>>, Error> {
    let mut buf = String::new();
    io::stdin().read_to_string(&mut buf)?;
    buf = buf.trim().to_string();
    Ok(buf
        .split("\n")
        .map(|line| {
            line.chars()
                .collect()
        })
        .collect::<Vec<Vec<char>>>())
}

fn part1(lines: &Vec<Vec<char>>) -> u32 {
    let mut total_error_score = 0;
    let mut incomplete_scores = BinaryHeap::new();
    for line in lines {
        match check_line(line) {
            Corrupted(n) => total_error_score += n,
            Incomplete(n) => incomplete_scores.push(n),
        }
    }
    total_error_score
}

fn part2(lines: &Vec<Vec<char>>) -> u64 {
    let mut _total_error_score = 0;
    let mut incomplete_scores = BinaryHeap::new();
    for line in lines {
        match check_line(line) {
            Corrupted(n) => _total_error_score += n,
            Incomplete(n) => incomplete_scores.push(n),
        }
    }
    for _ in 0..(incomplete_scores.len()/2) {
        incomplete_scores.pop();
    }
    incomplete_scores.pop().unwrap()
}

fn check_line(l: &Vec<char>) -> LineStatus {
    let mut stack: Vec<char> = Vec::new();
    for c in l {
        if ['[','{','(','<'].contains(c) {
            stack.push(*c);
        } else {
            if let Some(opening_char) = stack.pop() {
                if chars_match(opening_char, *c) {
                    continue
                }
            }
            return Corrupted(score_corrupt(*c))
        }
    }
    let mut completion_string = Vec::new();
    while !stack.is_empty() {
        completion_string.push(stack.pop().unwrap())
    }
    Incomplete(score_incomplete(completion_string))
}

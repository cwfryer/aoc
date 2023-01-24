use std::io::{self, Read};
use std::str::FromStr;
use std::env;

#[derive(Debug, PartialEq)]
enum Mode {
    LiteralValue(i64),
    Operator(Vec<Packet>),
}

#[derive(Debug, PartialEq)]
struct Packet {
    version: i32,
    type_id: i32,
    length: u16,
    mode: Mode,
}

#[derive(Debug)]
struct ParsePacketError;

impl FromStr for Packet {
    type Err = ParsePacketError;
    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let version = i32::from_str_radix(&s[..3], 2).unwrap();
        let type_id = i32::from_str_radix(&s[3..6], 2).unwrap();
        let mode_and_length = match type_id {
            4 => parse_literal_value(&s[6..]).unwrap(),
            _ => parse_operator(&s[6..]).unwrap(),
        };

        Ok(Packet {
            version,
            type_id,
            length: mode_and_length.1 + 6,
            mode: mode_and_length.0,
        })
    }
}

fn main() {
    env::set_var("RUST_BACKTRACE", "1");
    let input = get_input();
    input.lines().map(|line| hex_to_bin(line)).for_each(|line| {
        let packet = Packet::from_str(&line).unwrap();
        println!("sum of version: {:?}", part1(&packet));
        println!("result of operations {:?}", part2(&packet));
    })
}

fn part1(packet: &Packet) -> i32 {
    match &packet.mode {
        Mode::LiteralValue(_c) => packet.version,
        Mode::Operator(children) => packet.version + children.iter().map(|p| part1(p)).sum::<i32>(),
    }
}

fn part2(packet: &Packet) -> usize {
    match &packet.mode {
        Mode::LiteralValue(lit) => *lit as usize,
        Mode::Operator(children) => {
            let evaluated_children = children.iter().map(|packet| part2(packet));
            match packet.type_id {
                0 => sum(evaluated_children),
                1 => product(evaluated_children),
                2 => minimum(evaluated_children),
                3 => maximum(evaluated_children),
                5 => greater_than(evaluated_children),
                6 => less_than(evaluated_children),
                7 => equal_to(evaluated_children),
                _ => unreachable!(),
            }
        }
    }
}

fn sum<I>(evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    evaluated_children.sum()
}

fn product<I>(evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    evaluated_children.fold(1usize, |acc, child| acc * child)
}

fn minimum<I>(evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    evaluated_children.fold(usize::MAX, |acc, child| std::cmp::min(acc, child))
}

fn maximum<I>(evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    evaluated_children.fold(0usize, |acc, child| std::cmp::max(acc, child))
}

fn greater_than<I>(mut evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    if evaluated_children.next().unwrap() > evaluated_children.next().unwrap() {
        1
    } else {
        0
    }
}

fn less_than<I>(mut evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    if evaluated_children.next().unwrap() < evaluated_children.next().unwrap() {
        1
    } else {
        0
    }
}

fn equal_to<I>(mut evaluated_children: I) -> usize
where
    I: Iterator<Item = usize>,
{
    if evaluated_children.next().unwrap() == evaluated_children.next().unwrap() {
        1
    } else {
        0
    }
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

fn parse_operator(input: &str) -> Option<(Mode, u16)> {
    let length_type_id = input.chars().next().unwrap();
    match length_type_id {
        '0' => {
            let length = u16::from_str_radix(&input[1..=15], 2).unwrap();
            let packet_string = String::from(&input[16..=15 + length as usize]);
            let packets_and_length = read_packets_by_length(&packet_string, length).unwrap();
            Some((packets_and_length.0, packets_and_length.1 + 16))
        }
        '1' => {
            let count = u16::from_str_radix(&input[1..=11], 2).unwrap();
            let packet_string = String::from(&input[12..]);
            let packets_and_length = read_packets_by_count(&packet_string, count).unwrap();
            Some((packets_and_length.0, packets_and_length.1 + 12))
        }
        _ => unreachable!(),
    }
}

fn read_packets_by_length(input: &str, length: u16) -> Option<(Mode, u16)> {
    let mut accumulated_length = 0u16;
    let mut children: Vec<Packet> = vec![];
    loop {
        let new_child_packet = Packet::from_str(&input[accumulated_length as usize..]).unwrap();
        accumulated_length += new_child_packet.length;
        children.push(new_child_packet);
        if length == accumulated_length {
            return Some((Mode::Operator(children), length));
        } else if accumulated_length > length {
            panic!("Invalid packet children. Shits too long")
        }
    }
}

fn read_packets_by_count(input: &str, count: u16) -> Option<(Mode, u16)> {
    let mut accumulated_length = 0u16;
    let mut children: Vec<Packet> = vec![];
    for _i in 0..count {
        let new_child_packet = Packet::from_str(&input[accumulated_length as usize..]).unwrap();
        accumulated_length += new_child_packet.length;
        children.push(new_child_packet);
    }
    return Some((Mode::Operator(children), accumulated_length));
}

fn parse_literal_value(input: &str) -> Option<(Mode, u16)> {
    let mut bin_str = String::new();
    let mut bit_length = 0;
    for chunk in input.chars().collect::<Vec<char>>().chunks(5) {
        if chunk.iter().next().unwrap() == &'0' {
            for c in chunk[1..].into_iter() {
                bin_str.push(*c);
            }
            bit_length += 5;
            break;
        } else {
            for c in chunk[1..].into_iter() {
                bin_str.push(*c);
            }
            bit_length += 5;
        }
    }

    Some((
        Mode::LiteralValue(i64::from_str_radix(&bin_str, 2).unwrap()),
        bit_length,
    ))
}

fn hex_to_bin(input: &str) -> String {
    input
        .chars()
        .map(|c| match c {
            '0' => "0000",
            '1' => "0001",
            '2' => "0010",
            '3' => "0011",
            '4' => "0100",
            '5' => "0101",
            '6' => "0110",
            '7' => "0111",
            '8' => "1000",
            '9' => "1001",
            'A' => "1010",
            'B' => "1011",
            'C' => "1100",
            'D' => "1101",
            'E' => "1110",
            'F' => "1111",
            _ => unreachable!(),
        })
        .fold(String::new(), |mut output, c| {
            output.push_str(c);
            output
        })
}

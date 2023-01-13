use std::collections::HashMap;
use std::io::{self, Read};
fn main() {
    let input = get_input();
    let start = input.lines().next().unwrap().trim().to_string();
    let rules = input
        .split("\n\n")
        .last()
        .unwrap()
        .lines()
        .map(|line| {
            let (a, b) = line.split_once(" -> ").unwrap();
            (a.to_string(), b.chars().next().unwrap())
        })
        .collect::<HashMap<String, char>>();

    let mut o = start
        .clone()
        .chars()
        .collect::<Vec<char>>()
        .windows(2)
        .fold(HashMap::new(), |mut acc, x| {
            let key: String = x.into_iter().collect();
            acc.entry(key).and_modify(|v| *v += 1).or_insert(1);
            acc
        });

    for _ in 0..40 {
        o = apply_rules(o, &rules);
    }

    // after applying the rules, count up the chars
    let mut count_map = o.iter().fold(HashMap::new(), |mut acc, (k,v)| {
        // take the first char of each pair, because it is a rolling window we only need the first
        let c = k.chars().next().unwrap();
        acc.entry(c).and_modify(|val| *val += *v).or_insert(*v);
        acc
    });
    // and we have to add on the starting last char
    count_map.entry(start.chars().last().unwrap()).and_modify(|v| *v += 1).or_insert(1);

    println!(
        "{}",
        count_map.values().max().unwrap() - count_map.values().min().unwrap()
    )
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

// this worked for part 1 but it is too slow for part 2
// fn apply_rules(s: String, rules: &HashMap<String,String>) -> String {
//     let mut out = String::new();
//     s.chars().collect::<Vec<char>>().windows(2).for_each(|x| {
//         let key: String = x.into_iter().collect();
//         if rules.contains_key(&key) {
//             out.push(x[0]);
//             out.push_str(rules.get(&key).as_ref().unwrap());
//         } else {
//             out.push(x[0])
//         }
//     });
//     out.push(s.chars().last().unwrap());
//
//     out
// }

fn apply_rules(m: HashMap<String, usize>, rules: &HashMap<String, char>) -> HashMap<String, usize> {
    let mut out = m.clone();
    m.iter().for_each(|(k, val)| { // for each (pair,count) in the hashmap
        if rules.contains_key(k) { // if there's a rule for that pair...
            // make a left pair of chars and a right pair of chars
            let ins: &char = rules.get(k).unwrap();
            let mut c = k.chars();
            let mut left = String::from(c.next().unwrap());
            left.push(*ins);
            let mut right = String::from(*ins);
            right.push(c.next().unwrap());

            // remove the found key and replace it with an equal amount of l and r pairs
            out.entry(k.clone()).and_modify(|v| *v -= val);
            out.entry(left).and_modify(|v| *v += val).or_insert(*val);
            out.entry(right).and_modify(|v| *v += val).or_insert(*val);
        };
    });
    out
}

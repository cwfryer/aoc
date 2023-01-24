use std::io::{self, Read};

#[derive(Debug, Clone)]
struct VecTree {
    vals: Vec<u32>,
    depths: Vec<u32>,
}

impl VecTree {
    fn parse(s: &str) -> VecTree {
        let mut vals: Vec<u32> = Vec::new();
        let mut depths: Vec<u32> = Vec::new();
        let mut depth = 0;
        s.chars().for_each(|c| match c {
            '[' => depth += 1,
            ']' => depth -= 1,
            ',' => (),
            d => {
                vals.push(d.to_digit(10).unwrap() as u32);
                depths.push(depth - 1);
            }
        });

        VecTree { vals, depths }
    }
    fn explode(&mut self) -> bool {
        for i in 0..self.depths.len() {
            let depth = self.depths[i];
            if depth != 4 {
                continue;
            }

            if i != 0 {
                self.vals[i - 1] += self.vals[i];
            }

            if i + 2 < self.vals.len() {
                self.vals[i + 2] += self.vals[i + 1];
            }

            self.vals[i] = 0;
            self.depths[i] = 3;
            self.vals.remove(i + 1);
            self.depths.remove(i + 1);

            return true;
        }
        false
    }
    fn split(&mut self) -> bool {
        for i in 0..self.vals.len() {
            let value = self.vals[i];
            if value < 10 {
                continue;
            }
            let (a, b) = if value % 2 == 0 {
                (value / 2, value / 2)
            } else {
                (value / 2, value / 2 + 1)
            };

            self.vals[i] = a;
            self.depths[i] += 1;
            self.vals.insert(i + 1, b);
            self.depths.insert(i + 1, self.depths[i]);

            return true;
        }
        false
    }
    fn reduce(&mut self) {
        loop {
            if !self.explode() && !self.split() {
                break;
            }
        }
    }
    fn add(&mut self, other: &VecTree) {
        self.vals.extend(other.vals.iter());
        self.depths.extend(other.depths.iter());
        for i in 0..self.depths.len() {
            self.depths[i] += 1;
        }
        self.reduce()
    }
    fn score(&self) -> u32 {
        let mut vals = self.vals.clone();
        let mut depths = self.depths.clone();

        while vals.len() > 1 {
            for i in 0..depths.len() {
                if depths[i] == depths[i + 1] {
                    vals[i] = 3 * vals[i] + 2 * vals[i + 1];
                    vals.remove(i + 1);
                    depths.remove(i + 1);

                    if depths[i] > 0 {
                        depths[i] -= 1;
                    }

                    break;
                }
            }
        }

        vals[0]
    }
}

fn main() {
    let input = get_input();
    let trees: Vec<VecTree> = input.lines().map(|line| VecTree::parse(line)).collect();
    let mut values: Vec<u32> = vec![];

    for i in 0..trees.len() {
        let tree = trees[i].clone();
        let mut other_trees = trees.clone();
        other_trees.remove(i);
        for ot in other_trees {
            let mut v = tree.clone();
            v.add(&ot);
            values.push(v.score());
        }
    }

    println!("{}", values.iter().fold(0,|acc,v| if *v>acc{*v}else{acc}));
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

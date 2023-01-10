use std::collections::{HashMap, HashSet, VecDeque};
use std::io::Error;
use std::io::{self, Read};

#[derive(Eq, PartialEq, Hash, Clone, Copy, Debug)]
struct Point {
    x: i64,
    y: i64,
}

impl Point {
    fn new(x: i64, y: i64) -> Point {
        Point { x, y }
    }

    fn neighbors(&self) -> Vec<Point> {
        vec![
            Point {
                x: self.x + 1,
                y: self.y,
            },
            Point {
                x: self.x - 1,
                y: self.y,
            },
            Point {
                x: self.x,
                y: self.y + 1,
            },
            Point {
                x: self.x,
                y: self.y - 1,
            },
        ]
    }
}

fn main() {
    let points = parse().unwrap();
    let p2 = points.clone();
    println!("Part 1: {:?}", part1(points));
    println!("Part 2: {:?}", part2(p2));
}

fn parse() -> Result<HashMap<Point, i64>, Error> {
    let mut buf = String::new();
    io::stdin().read_to_string(&mut buf)?;
    buf = buf.trim().to_string();
    let mut hm: HashMap<Point, i64> = HashMap::new();
    buf.split("\n").enumerate().for_each(|(r_num, row)| {
        row.chars().enumerate().for_each(|(c_num, v)| {
            hm.insert(
                Point::new(c_num as i64, r_num as i64),
                v.to_digit(10).unwrap() as i64,
            );
        });
    });
    Ok(hm)
}

fn part1(points: HashMap<Point, i64>) -> i64 {
    points
        .iter()
        .map(|(k, v)| {
            (
                k.neighbors()
                    .iter()
                    .all(|x| !in_bounds(&points, &x) || points.get(x).unwrap() > v),
                v,
            )
        })
        .filter(|(b, _)| *b)
        .fold(0, |acc, (_, v)| acc + 1 + v)
}

fn part2(points: HashMap<Point, i64>) -> i64 {
    let basins = points
        .iter()
        .map(|(k, v)| {
            (
                k,
                (
                    k.neighbors()
                        .iter()
                        .all(|x| !in_bounds(&points, &x) || points.get(x).unwrap() > v),
                    v,
                ),
            )
        })
        .filter(|(_, v)| v.0)
        .map(|(k, _)| k)
        .collect::<Vec<&Point>>();


    let mut b_size: Vec<i64> = vec![];
    for b in basins {
        b_size.push(walk_basin(&points, *b));
    }
    b_size.sort_unstable();
    b_size.reverse();

    let mut answer = 1;
    for s in &b_size[..3] {
        answer *= s;
    }
    

    answer
}

fn walk_basin(points: &HashMap<Point, i64>, start: Point) -> i64 {
    let mut queue: VecDeque<Point> = VecDeque::new();
    let mut visited: HashSet<Point> = HashSet::new();

    queue.push_front(start);
    visited.insert(start);

    while !queue.is_empty() {
        let p = queue.pop_front();

        for n in p.unwrap().neighbors() {
            if in_bounds(points, &n) && points.get(&n).unwrap() < &9 && !visited.contains(&n){
                visited.insert(n);
                queue.push_back(n)
            } else {
                continue;
            };
        }
    }

    visited.len() as i64
}

fn in_bounds(points: &HashMap<Point, i64>, p: &Point) -> bool {
    points.contains_key(p)
}

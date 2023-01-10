use std::collections::{HashMap, HashSet};
use std::io::{self, Error, Read};

#[derive(Hash, Clone, Copy, Debug, PartialEq, PartialOrd, Eq)]
struct Point {
    x: i32,
    y: i32,
}

impl Point {
    fn new(x: i32, y: i32) -> Point {
        Point { x, y }
    }
    fn touching(self) -> Vec<Point> {
        vec![
            Point::new(self.x,self.y+1),
            Point::new(self.x+1,self.y+1),
            Point::new(self.x+1,self.y),
            Point::new(self.x+1,self.y-1),
            Point::new(self.x,self.y-1),
            Point::new(self.x-1,self.y-1),
            Point::new(self.x-1,self.y),
            Point::new(self.x-1,self.y+1),
        ]
    }
}
struct Cavern {
    octos: HashMap<Point, u32>,
    height: i32,
    width: i32,
}

impl Cavern {
    fn step(&mut self) -> usize {
        
        for c in self.grid_coords() {
            self.octos.entry(c).and_modify(|e| *e += 1);
        }

        let mut flashed:HashSet<Point> = HashSet::new();
        loop {
            let flashes:HashSet<Point> = self.grid_coords()
                .into_iter()
                .filter(|c| *self.octos.get(c).unwrap() > 9)
                .filter(|c| !flashed.contains(c))
                .collect::<HashSet<Point>>();

            for c in flashes.iter().flat_map(|c| c.touching()) {
                self.octos.entry(c).and_modify(|e| *e += 1);
            }

            if flashes.is_empty() {
                break;
            }
            flashed.extend(flashes)
        }

        for c in flashed.iter() {
            self.octos.entry(*c).and_modify(|e| *e = 0);
        }

        flashed.len()
    }

    fn is_synchronized(&self) -> bool {
        self.octos.values().all(|e| *e == 0)
    }

    fn grid_coords(&self) -> HashSet<Point> {
        let mut coords = HashSet::new();
        for x in 0..self.width {
            for y in 0..self.height {
                coords.insert(Point::new(x, y));
            }
        }
        coords
    }
}

fn main() -> Result<(),Error> {
    let mut buf = String::new();
    io::stdin().read_to_string(&mut buf)?;
    buf = buf.trim().to_string();
    println!("part 1: {:?}", part1(&buf, 100));
    println!("part 2: {:?}", part2(&buf));
    Ok(())
}

fn part1(input: &String, steps: usize) -> usize {
    let mut cavern = parse(input).unwrap();
    let mut flashes = 0;
    for _ in 0..steps {
        flashes += cavern.step();
    }
    flashes
}

fn part2(input: &String) -> usize {
    let mut cavern = parse(input).unwrap();
    let mut step = 0;
    loop {
        cavern.step();
        step += 1;

        if cavern.is_synchronized() {
            break
        }
    }
    step
}

fn parse(buf: &String) -> Result<Cavern, Error> {
    let mut octos: HashMap<Point, u32> = HashMap::new();
    buf.split("\n").enumerate().for_each(|(r_num, row)| {
        row.chars().enumerate().for_each(|(c_num, v)| {
            octos.insert(
                Point::new(c_num as i32, r_num as i32),
                v.to_digit(10).unwrap() as u32,
            );
        });
    });
    let height = buf.lines().count() as i32;
    let width = buf.lines().next().unwrap().chars().count() as i32;
    Ok(Cavern {
        octos,
        height,
        width,
    })
}

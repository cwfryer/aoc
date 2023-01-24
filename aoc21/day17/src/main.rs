use itertools::Itertools;
use std::io::{self, Read};

fn try_vel(xmin: i32, xmax: i32, ymin: i32, ymax: i32, mut dx: i32, mut dy: i32) -> Option<i32> {
    let (mut x, mut y, mut maxy) = (0, 0, 0);
    loop {
        x += dx;
        y += dy;
        dx -= dx.signum();
        dy -= 1;
        if y > maxy {
            maxy = y;
        }
        match (xmin <= x && x <= xmax, ymin <= y && y <= ymax) {
            (true, true) => return Some(maxy),
            (false, _) if dx == 0 => return None,
            (_, false) if dy < 0 && y < ymin => return None,
            _ => {}
        }
    }
}

fn parse_range(s: &str) -> (i32, i32) {
    let (a, b) = s.split_once("..").unwrap();
    (a.parse().unwrap(), b.parse().unwrap())
}

fn main() {
    let input = get_input();
    let (xrange, yrange) = input[15..].split_once(", y=").unwrap();
    let (xmin, xmax) = parse_range(xrange);
    let (ymin, ymax) = parse_range(yrange);
    let maxys = (0..=xmax)
        .cartesian_product(ymin..1000)
        .filter_map(|(x, y)| try_vel(xmin, xmax, ymin, ymax, x, y))
        .collect::<Vec<_>>();

    println!("{},{}", *maxys.iter().max().unwrap(), maxys.len());
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

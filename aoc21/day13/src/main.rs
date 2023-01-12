use std::{
    collections::HashSet,
    io::{self, Read},
};

fn main() {
    let input = get_input();
    let mut input_iter = input.split("\n\n").map(|x| x.to_string());
    let coord_string = input_iter.next().unwrap().to_string();
    let instruction_string = input_iter.next().unwrap().to_string();

    let points = coord_string
        .lines()
        .map(|line| {
            let (x, y) = line.split_once(",").unwrap();
            (x.parse::<usize>().unwrap(), y.parse::<usize>().unwrap())
        })
        .collect::<HashSet<(usize, usize)>>();

    let mut max: (usize, usize) = (0, 0);
    for p in &points {
        if p.0 > max.0 {
            max.0 = p.0
        }
        if p.1 > max.1 {
            max.1 = p.1
        }
    }
    let instructions: Vec<(&str, usize)> = instruction_string
        .lines()
        .map(|line| {
            let (axis, c) = line.rsplit_once(" ").unwrap().1.split_once("=").unwrap();
            (axis, c.parse::<usize>().unwrap())
        })
        .collect();

    let mut o: HashSet<(usize, usize)> = points.clone();
    for (axis, coord) in instructions {
        o = match axis {
            "x" => {
                fold_x_axis(coord, o.clone())
            }
            "y" => {
                fold_y_axis(coord, o.clone())
            }
            _ => unreachable!(),
        };
        max = match axis {
            "x" => (coord,max.1),
            "y" => (max.0,coord),
            _ => unreachable!()
        };
    }
    for y in 0..=max.1 {
        let mut r_str = String::new();
        for x in 0..=max.0 {
            if o.contains(&(x,y)) {
                r_str.push_str("#")
            } else {
                r_str.push_str(" ")
            }
        }
        println!("{}",r_str);
    }
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

fn fold_x_axis(
    coord: usize,
    points: HashSet<(usize, usize)>,
) -> HashSet<(usize, usize)> {
    let mut output: HashSet<(usize, usize)> =
        points.iter().filter(|p| p.0 < coord).map(|p| *p).collect();

    points.iter().filter(|p| p.0 > coord).for_each(|p| {
        output.insert((coord-(p.0 - coord), p.1));
    });

    output
}

fn fold_y_axis(
    coord: usize,
    points: HashSet<(usize, usize)>,
) -> HashSet<(usize, usize)> {
    let mut output: HashSet<(usize, usize)> =
        points.iter().filter(|p| p.1 < coord).map(|p| *p).collect();

    points.iter().filter(|p| p.1 > coord).for_each(|p| {
        output.insert((p.0, coord-(p.1 - coord)));
    });

    output
}

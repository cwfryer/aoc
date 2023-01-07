use itertools::Itertools;
use std::collections::HashMap;
use std::io::Error;

fn main() -> Result<(), Error> {
    let data = std::fs::read_to_string("./input.txt")?;
    part1(&data);
    part2(&data);
    Ok(())
}

fn part1(data: &String) {
    let points = parse(data,false);
    println!("Part 1: {:?}\n",points);
}
fn part2(data: &String) {
    let points = parse(data,true);
    println!("Part 2: {:?}\n",points);
}

fn parse(data: &String, diag: bool) -> u32 {
    let mut hm: HashMap<(i32, i32), u32> = HashMap::new();
    data.lines().for_each(|line| {
        let pt = line
            .split(" -> ")
            .map(|p| {
                p.split(",")
                    .map(|x| x.parse::<i32>().unwrap())
                    .collect_tuple::<(i32, i32)>()
            })
            .map(|l| l.unwrap())
            .collect::<Vec<(i32, i32)>>();
        let p1 = pt[0];
        let p2 = pt[1];
        let slope = calculate_slope(p1, p2);

        // Part 1
        // if slope == 0 {
        //     let mut points: Vec<(i32, i32)> = vec![];
        //     if p1.0 == p2.0 {
        //         if p2.1 > p1.1 {
        //             for y in p1.1..=p2.1 {
        //                 points.push((p1.0, y))
        //             }
        //         } else {
        //             for y in p2.1..=p1.1 {
        //                 points.push((p1.0,y))
        //             }
        //         }
        //     } else {
        //         let mut y = p1.1;
        //         if p2.0 > p1.0 {
        //             for x in p1.0..=p2.0 {
        //                 points.push((x, y));
        //                 y += slope
        //             }
        //         } else {
        //             for x in p2.0..=p1.0 {
        //                 points.push((x, y));
        //                 y += slope
        //             }
        //         }
        //     }
        //     for p in points {
        //         all_points.push(p)
        //     }
        // }


        // Part 2
        if p1.0 == p2.0 {
            let small = if p1.1 > p2.1 { p2.1 } else { p1.1 };
            let large = if p1.1 > p2.1 { p1.1 } else { p2.1 };
            for y in small..=large {
                let point = (p1.0,y);
                *hm.entry(point).or_insert(0) += 1;
            }
        } else if p1.1 == p2.1 {
            let small = if p1.0 > p2.0 { p2.0 } else { p1.0 };
            let large = if p1.0 > p2.0 { p1.0 } else { p2.0 };
            for x in small..=large {
                let point = (x,p1.1);
                *hm.entry(point).or_insert(0) += 1;
            }
        } else if diag {
            let distance = (p1.0 - p2.0).abs() as i32;
            let leftmost_point = if p1.0 < p2.0 { p1 } else { p2 };
            for i in 0..=distance {
                let point = if slope > 0 {
                    (leftmost_point.0 + i, leftmost_point.1 + i)
                } else {
                    (leftmost_point.0 + i, leftmost_point.1 - i)
                };

                *hm.entry(point).or_insert(0) += 1;
            }
        }
    });

    let intersections = hm.values().fold(0, |acc, val| {
        if *val > 1 {
            return acc + 1;
        }

        acc
    });

    intersections
}

fn calculate_slope(p1: (i32, i32), p2: (i32, i32)) -> i32 {
    if p1.0 == p2.0 {
        return 0;
    }
    (p2.1 - p1.1) / (p2.0 - p1.0)
}

// fn print_map(m: &HashMap<(i32, i32), u32>) {
//     for i in 0..=9 {
//         let mut r = String::new();
//         for j in 0..=9 {
//             if m.get(&(j,i)).is_none() {
//                 r.push_str(".")
//             } else {
//                 r.push_str(m[&(j,i)].to_string().as_str())
//             }
//         }
//         println!("{}", r)
//     }
// }

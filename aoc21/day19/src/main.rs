use itertools::Itertools;
use nalgebra::{Matrix3, Vector3};
use std::collections::HashMap;
use std::io::{self, Read};

type Fingerprint = (i32, i32);
#[derive(Debug, Clone)]
struct Scanner {
    name: String,
    location: Vector3<i32>,
    beacons: Vec<Vector3<i32>>,
    fingerprints: HashMap<Fingerprint, Vec<[Vector3<i32>; 2]>>,
}

static ROTATION_MATRICES: [Matrix3<i32>; 24] = [
    Matrix3::new(1, 0, 0, 0, 1, 0, 0, 0, 1),
    Matrix3::new(1, 0, 0, 0, 0, 1, 0, -1, 0),
    Matrix3::new(1, 0, 0, 0, -1, 0, 0, 0, -1),
    Matrix3::new(1, 0, 0, 0, 0, -1, 0, 1, 0),
    Matrix3::new(0, 1, 0, 0, 0, 1, 1, 0, 0),
    Matrix3::new(0, 1, 0, 1, 0, 0, 0, 0, -1),
    Matrix3::new(0, 1, 0, 0, 0, -1, -1, 0, 0),
    Matrix3::new(0, 1, 0, -1, 0, 0, 0, 0, 1),
    Matrix3::new(0, 0, 1, 1, 0, 0, 0, 1, 0),
    Matrix3::new(0, 0, 1, 0, 1, 0, -1, 0, 0),
    Matrix3::new(0, 0, 1, -1, 0, 0, 0, -1, 0),
    Matrix3::new(0, 0, 1, 0, -1, 0, 1, 0, 0),
    Matrix3::new(-1, 0, 0, 0, -1, 0, 0, 0, 1),
    Matrix3::new(-1, 0, 0, 0, 0, 1, 0, 1, 0),
    Matrix3::new(-1, 0, 0, 0, 1, 0, 0, 0, -1),
    Matrix3::new(-1, 0, 0, 0, 0, -1, 0, -1, 0),
    Matrix3::new(0, -1, 0, 0, 0, -1, 1, 0, 0),
    Matrix3::new(0, -1, 0, 1, 0, 0, 0, 0, 1),
    Matrix3::new(0, -1, 0, 0, 0, 1, -1, 0, 0),
    Matrix3::new(0, -1, 0, -1, 0, 0, 0, 0, -1),
    Matrix3::new(0, 0, -1, -1, 0, 0, 0, 1, 0),
    Matrix3::new(0, 0, -1, 0, 1, 0, 1, 0, 0),
    Matrix3::new(0, 0, -1, 1, 0, 0, 0, -1, 0),
    Matrix3::new(0, 0, -1, 0, -1, 0, -1, 0, 0),
];

impl Scanner {
    fn from_str(s: &str) -> Scanner {
        let mut lines = s.lines();
        let name = lines.next().unwrap().to_owned();
        let beacons = lines
            .map(|line| {
                if let Some((x, y, z)) = line.split(",").collect_tuple() {
                    Vector3::from_vec(vec![
                        x.parse::<i32>().unwrap(),
                        y.parse::<i32>().unwrap(),
                        z.parse::<i32>().unwrap(),
                    ])
                } else {
                    panic!("WHAT??")
                }
            })
            .collect::<Vec<Vector3<i32>>>();
        let mut fingerprints: HashMap<Fingerprint, Vec<[Vector3<i32>; 2]>> = HashMap::new();
        for (b1, b2) in beacons.iter().tuple_combinations::<(_, _)>() {
            let d = b1 - b2;
            let fingerprint = (d[0] + d[1] + d[2], d[0].max(d[1]).max(d[2]));
            fingerprints.entry(fingerprint).or_default().push([*b1, *b2]);
        }
        Scanner {
            name,
            location: Vector3::from_vec(vec![0, 0, 0]),
            beacons,
            fingerprints,
        }
    }
}

fn main() {
    let input = get_input();
    let scanners = input
        .split("\n\n")
        .map(|chunk| Scanner::from_str(chunk))
        .collect::<Vec<Scanner>>();
    let s0 = &scanners[0];
    println!("{:?}", s0.fingerprints);
    let s1 = &scanners[1];
    println!("{:?}", s1.fingerprints);

    let mut same_dist_count = 0;
    for (f,_) in &s0.fingerprints {
        if s1.fingerprints.contains_key(f) {
            same_dist_count += 1
        }
    }

    println!("{:?}", same_dist_count);
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

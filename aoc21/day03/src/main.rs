use std::io::Error;

fn main() -> Result<(), Error> {
    let data = std::fs::read_to_string("./input.txt")?;
    println!("{}", part1(&data));
    println!("{}", part2(&data));
    Ok(())
}

fn part1(data: &String) -> u32 {
    let arr = data
        .lines()
        .map(|l| l.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut ga = String::new();
    let mut ep = String::new();
    for i in 0..arr[0].len() {
        let mut ones = 0;
        for j in 0..arr.len() {
            if arr[j][i] == '1' {
                ones += 1
            }
        }
        let zeroes = arr.len() - ones;
        if ones > zeroes {
            ga.push('1');
            ep.push('0')
        } else {
            ga.push('0');
            ep.push('1');
        }
    }

    bstr_to_int(&ga) * bstr_to_int(&ep)
}

fn part2(data: &String) -> u32 {
    let arr = data
        .lines()
        .map(|l| l.chars().collect::<Vec<char>>())
        .collect::<Vec<Vec<char>>>();

    let mut ox_arr = arr.clone();
    let mut co_arr = arr.clone();
    let mut ox = String::new();
    let mut co = String::new();

    for i in 0..=arr[0].len() {
        if ox_arr.len() == 1 {
            ox = ox_arr[0].iter().collect::<String>();
            break
        }
        let mut ones = 0;
        for j in 0..ox_arr.len() {
            if ox_arr[j][i] == '1' {
                ones += 1
            }
        }
        let zeroes = ox_arr.len() - ones;
        if ones > zeroes {
            ox_arr.retain(|e| e[i] == '1');
        } else if zeroes > ones {
            ox_arr.retain(|e| e[i] == '0');
        } else {
            ox_arr.retain(|e| e[i] == '1');
        }
    }
    for i in 0..=arr[0].len() {
        if co_arr.len() == 1 {
            co = co_arr[0].iter().collect::<String>();
            break
        }
        let mut ones = 0;
        for j in 0..co_arr.len() {
            if co_arr[j][i] == '1' {
                ones += 1
            }
        }
        let zeroes = co_arr.len() - ones;
        if ones > zeroes {
            co_arr.retain(|e| e[i] == '0')
        } else if zeroes > ones {
            co_arr.retain(|e| e[i] == '1')
        } else {
            co_arr.retain(|e| e[i] == '0')

        }
    }

    println!("ox:  {}\nco2: {}",&ox,&co);
    bstr_to_int(&ox) * bstr_to_int(&co)
}

fn bstr_to_int(s: &String) -> u32 {
    let mut out: u32 = 0;
    for (i,c) in s.chars().rev().enumerate() {
        out += c.to_digit(10).unwrap() * (2_u32.pow(i as u32))
    }
    out
}

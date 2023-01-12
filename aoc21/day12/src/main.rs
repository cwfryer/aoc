use std::{
    collections::{HashMap, VecDeque},
    io::{self, Read},
};

fn main() {
    let mut neighbors_of: HashMap<String, Vec<String>> = HashMap::new();
    let input = get_input();
    for line in input.lines() {
        let mut chunks = line.split("-");
        let a = chunks.next().unwrap().to_string();
        let b = chunks.next().unwrap().to_string();
        neighbors_of
            .entry(a.clone())
            .and_modify(|neighbors| neighbors.push(b.clone()))
            .or_insert(vec![b.clone()]);
        neighbors_of
            .entry(b)
            .and_modify(|neighbors| neighbors.push(a.clone()))
            .or_insert(vec![a]);
    }

    let mut q: VecDeque<Vec<String>> = VecDeque::new();
    q.push_back(vec![String::from("start")]);

    let mut num_paths = 0;
    while let Some(path) = q.pop_front() {
        let last = path.last().unwrap();
        let sc: HashMap<&String, i32> = path
            .iter()
            .filter(|c| c.len() < 3)
            .filter(|c| c.chars().next().unwrap().is_ascii_lowercase())
            .fold(HashMap::new(), |mut acc, item| {
                acc.entry(item).and_modify(|val| *val += 1).or_insert(1);
                acc
            });

        for n in neighbors_of.get(last).unwrap() {
            if n == "start" {
                continue;
            }
            if n == "end" {
                num_paths += 1;
                let mut p_path = path.clone();
                p_path.push(n.clone());
                continue;
            }

            if n.chars().next().unwrap().is_ascii_lowercase()
                && sc.values().any(|v| v >= &2)
                && path.contains(n)
            {
                continue;
            }

            let mut new_path = path.clone();
            new_path.push(n.clone());
            q.push_back(new_path);
        }
    }

    println!("part 1: {:?}", num_paths)
}

fn get_input() -> String {
    let mut input: String = String::new();
    io::stdin().lock().read_to_string(&mut input).unwrap();
    return input.trim().to_string();
}

// two ways to represent a graph,
// 1. point adjacency list
// 2. node w/ neighbor nodes
//
// this can be done as a BFS or DFS
// create a partial path from start to neighbors, add to a stack, and go through
// each of those
// we need to know the full length of path so we know which lower caves we've
// visited
//
// we probably need to keep a set of partial paths in the queue
//
// both dfs and bfs seem to guarantee that duplicate paths aren't made
//
// hashmap point: vec<neighbors>

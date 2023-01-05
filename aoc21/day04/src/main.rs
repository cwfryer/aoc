use nalgebra::Matrix5;
use std::io::Error;

type Board = (Matrix5<(u32, bool)>,bool);

fn main() -> Result<(), Error> {
    let data = std::fs::read_to_string("./input.txt")?;
    println!("Part 1: {:?}\n", part1(&data));
    println!("Part 2: {:?}\n", part2(&data));
    Ok(())
}

fn part1(data: &String) -> u32 {
    let (nums, mut boards) = parse(data);
    for num in nums {
        for board in boards.iter_mut() {
            update_board(board, num);
            if check_win(board) {
                board.1 = true;
                return num * unmarked_sum(board)
            }
        }
    }

    0
}

fn part2(data: &String) -> u32 {
    let (nums, mut boards) = parse(data);
    for num in nums {
        for board in boards.iter_mut() {
            if board.1 == true {
                continue
            }
            update_board(board, num);
            if check_win(board) {
                board.1 = true;
                println!("board won with score: {}",num * unmarked_sum(board));
            }
        }
    }

    0
}

fn parse(data: &String) -> (Vec<u32>, Vec<Board>) {
    let mut lines = data.split("\n\n");
    let nums = lines
        .next()
        .unwrap()
        .split(",")
        .filter_map(|x| x.parse::<u32>().ok())
        .collect::<Vec<_>>();

    let boards = lines
        .map(|board| {
            (Matrix5::from_iterator(board.lines().flat_map(|line| {
                line.split_whitespace()
                    .filter_map(|x| x.parse::<u32>().ok())
                    .map(|x| (x, false))
            })),false)
        })
        .collect::<Vec<Board>>();

    (nums, boards)
}

fn update_board(board: &mut Board, num: u32) {
    board.0.iter_mut().for_each(|x| {
        if x.0 == num {
            x.1 = true
        }
    })
}

fn check_win(board: &Board) -> bool {
    let col_win = board.0
        .column_iter()
        .any(|col| col.iter().all(|(_, marked)| *marked));

    let row_win = board.0
        .row_iter()
        .any(|row| row.iter().all(|(_, marked)| *marked));

    row_win || col_win
}

fn unmarked_sum(board: &Board) -> u32 {
    board.0
        .iter()
        .filter(|(_, marked)| *marked == false)
        .map(|(x, _)| *x)
        .sum()
}

use std::fs::File;
use std::io::prelude::*;
use std::vec::IntoIter;

struct Node {
    children: Vec<Node>,
    metadata: Vec<usize>,
    value: usize,
}

fn parse_tree(numbers: Vec<i32>) -> Node {
    fn parse_node(it: &mut IntoIter<i32>) -> Node {
        let n_children = it.next().unwrap();
        let n_metadata = it.next().unwrap();

        let children = parse_children(it, n_children);
        let metadata = parse_metadata(it, n_metadata);
        let value = calculate_node_value(&children, &metadata);

        Node {
            children,
            metadata,
            value,
        }
    }

    fn parse_children(it: &mut IntoIter<i32>, n_children: i32) -> Vec<Node> {
        let mut children = vec![];
        for _ in 0..n_children {
            children.push(parse_node(it));
        }
        children
    }

    fn parse_metadata(it: &mut IntoIter<i32>, n_metadata: i32) -> Vec<usize> {
        let mut metadata = vec![];
        for _ in 0..n_metadata {
            metadata.push(it.next().unwrap() as usize);
        }
        metadata
    }

    fn calculate_node_value(children: &Vec<Node>, metadata: &Vec<usize>) -> usize {
        if children.len() > 0 {
            let mut value = 0;
            for &child_i in metadata.iter() {
                if child_i <= children.len() {
                    value += children[child_i - 1].value;
                }
            }
            value
        } else {
            metadata.iter().sum()
        }
    }

    parse_node(&mut numbers.into_iter())
}

fn sum_metadata(root: &Node) -> usize {
    let mut stack = vec![root.clone()];
    let mut total = 0;

    while let Some(node) = stack.pop() {
        stack.extend(node.children.iter().clone());
        total += node.metadata.iter().sum::<usize>();
    }

    total
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let numbers: Vec<i32> = contents
        .split_whitespace()
        .map(|s| s.parse().unwrap())
        .collect();

    let root = parse_tree(numbers);

    println!("{:?}", sum_metadata(&root));
    println!("{:?}", root.value);
}

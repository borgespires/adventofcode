extern crate regex;

use regex::Regex;
use std::collections::BTreeMap;
use std::collections::HashSet;
use std::fs::File;
use std::io::prelude::*;

struct TaskRunner {
    completed: Vec<char>,
    in_progress: Vec<(char, u32)>,
    elapsed_time: u32,
    idle_workers: u32,
    task_eta: u32,
}

impl TaskRunner {
    fn new(workers: u32, task_eta: u32) -> TaskRunner {
        TaskRunner {
            completed: vec![],
            in_progress: vec![],
            elapsed_time: 0,
            idle_workers: workers,
            task_eta: task_eta,
        }
    }

    fn task_delay(&self, id: char) -> u32 {
        (id as u32) - b'A' as u32 + 1
    }

    fn ready_to_play(&self, (task_id, required): &(&char, &Vec<char>)) -> bool {
        let in_progress_tasks = self
            .in_progress
            .iter()
            .map(|(id, _)| id)
            .collect::<Vec<&char>>();

        !self.completed.contains(task_id)
            && !in_progress_tasks.contains(task_id)
            && required
                .iter()
                .filter(|s| !self.completed.contains(s))
                .count()
                == 0
    }

    fn assign_tasks_to_idle_workers(&mut self, tasks: &BTreeMap<char, Vec<char>>) {
        loop {
            match tasks.iter().find(|task| self.ready_to_play(task)) {
                Some((&task_id, _)) if self.idle_workers > 0 => {
                    let delay = self.task_delay(task_id);
                    self.in_progress.push((task_id, self.task_eta + delay));
                    self.idle_workers -= 1;
                }
                _ => break,
            }
        }
    }

    fn drain_finished_tasks(&mut self) -> Vec<char> {
        let (finished, other): (Vec<_>, Vec<_>) = self
            .in_progress
            .clone()
            .into_iter()
            .partition(|(_, time_left)| *time_left == 0);

        self.in_progress = other;

        let mut tasks = finished
            .into_iter()
            .map(|(id, _)| id)
            .collect::<Vec<char>>();

        tasks.sort();
        tasks
    }

    fn update_in_progress_tasks(&mut self) {
        for (_id, time_left) in self.in_progress.iter_mut() {
            *time_left -= 1;
        }
    }

    fn update_tasks_state(&mut self) {
        let mut finished_tasks = self.drain_finished_tasks();

        self.idle_workers += finished_tasks.iter().count() as u32;
        self.completed.append(&mut finished_tasks);
    }

    fn run(&mut self, tasks: &BTreeMap<char, Vec<char>>) -> (String, u32) {
        while self.completed.len() != tasks.len() {
            self.assign_tasks_to_idle_workers(&tasks);
            self.update_in_progress_tasks();
            self.update_tasks_state();
            self.elapsed_time += 1;
        }

        (self.completed.iter().collect(), self.elapsed_time)
    }
}

fn parse(s: &str) -> (char, char) {
    let re = Regex::new(r"Step (\w) must be finished before step (\w) can begin.").unwrap();
    let cap = re.captures(&s).unwrap();
    (
        cap[1].chars().next().unwrap(),
        cap[2].chars().next().unwrap(),
    )
}

fn create_backlog(instructions: Vec<(char, char)>) -> BTreeMap<char, Vec<char>> {
    let mut backlog = BTreeMap::new();
    let mut blocked_tasks = HashSet::new();
    let mut blocking_tasks = HashSet::new();

    for (requirement, task_id) in instructions {
        blocked_tasks.insert(task_id);
        blocking_tasks.insert(requirement);
        backlog
            .entry(task_id)
            .or_insert_with(Vec::new)
            .push(requirement);
    }

    let tasks_without_dependencies = blocking_tasks
        .difference(&blocked_tasks)
        .collect::<Vec<&char>>();

    for &task_id in tasks_without_dependencies {
        backlog.insert(task_id, vec![]);
    }

    backlog
}

fn main() {
    let mut f = File::open("./resources/input.txt").expect("file not found");

    let mut contents = String::new();
    f.read_to_string(&mut contents)
        .expect("something went wrong reading the file");

    let instructions = contents
        .lines()
        .map(|s| parse(s))
        .collect::<Vec<(char, char)>>();

    let backlog = create_backlog(instructions);

    println!("{:?}", TaskRunner::new(1, 0).run(&backlog));
    println!("{:?}", TaskRunner::new(5, 60).run(&backlog));
}

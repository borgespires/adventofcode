use regex::Regex;
use std::num::ParseIntError;
use std::str::FromStr;

use datetime::DateTime;

#[derive(Debug)]
pub enum Action {
    ShiftBegin { guard: u32 },
    Asleep,
    WakeUp,
    None,
}

#[derive(Debug)]
pub struct Event {
    pub datetime: DateTime,
    pub action: Action,
}

impl FromStr for Event {
    // TODO: how to define and use our own customized errors? ex: ParseEventError
    // https://doc.rust-lang.org/rust-by-example/error/multiple_error_types/define_error_type.html
    type Err = ParseIntError;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        lazy_static! {
            static ref RE: Regex = Regex::new(
                r"(?x)
            \[(?P<year>[0-9]{4})-(?P<month>[0-9]{2})-(?P<day>[0-9]{2})
            \s+(?P<hour>[0-9]{2}):(?P<minute>[0-9]{2})\]
            \s+(?:Guard\ \#(?P<id>[0-9]+)\ begins\ shift|(?P<action>.+))
            "
            ).unwrap();
        }

        let caps = RE.captures(&s).unwrap();
        let datetime = DateTime::new(
            caps["year"].parse()?,
            caps["month"].parse()?,
            caps["day"].parse()?,
            caps["hour"].parse()?,
            caps["minute"].parse()?,
        );
        let action = if let Some(id) = caps.name("id") {
            Action::ShiftBegin {
                guard: id.as_str().parse()?,
            }
        } else if caps["action"] == *"falls asleep" {
            Action::Asleep
        } else if caps["action"] == *"wakes up" {
            Action::WakeUp
        } else {
            Action::None
        };

        Ok(Event { datetime, action })
    }
}

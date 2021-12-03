use std::fs;
use std::error::Error;

fn part_one(input: &String) -> Result<(i64, i64), Box<dyn Error>> {
    // Horizontal position, depth
    let (mut x, mut y) = (0, 0);
    for (i, l) in input.split('\n').enumerate() {
        if l == "" { break };
        let mut line = l.split(' ');
        let cmd = line
            .next()
            .ok_or(format!("expected a command on line {}", i))?;
        let num = line
            .next()
            .ok_or(format!("expected a number on line {}", i))?
            .parse::<i64>()?;
        match cmd {
            "forward" => x += num,
            "down"    => y += num,
            "up"      => y -= num,
            _         => return Err(format!("expected one of `forward`, \
                                    `down` or `up` for a command on line {}", i).into())
        }
    }
    Ok((x, y))
}

fn part_two(input: &String) -> Result<(i64, i64), Box<dyn Error>> {
    // Horizontal position, depth, aim
    let (mut x, mut y, mut a) = (0, 0, 0);
    for (i, l) in input.split('\n').enumerate() {
        if l == "" { break };
        let mut line = l.split(' ');
        let cmd = line
            .next()
            .ok_or(format!("expected a command on line {}", i))?;
        let num = line
            .next()
            .ok_or(format!("expected a number on line {}", i))?
            .parse::<i64>()?;
        match cmd {
            "forward" => { x += num; y += a * num },
            "down"    => { a += num },
            "up"      => { a -= num },
            _         => return Err(format!("expected one of `forward`, \
                                    `down` or `up` for a command on line {}", i).into())
        }
    }
    Ok((x, y))
}

fn main() -> Result<(), Box<dyn Error>> {
    let input = fs::read_to_string("input")?;
    let (x, y) = part_one(&input)?;
    println!("(PI) Multiplying the horizontal position {} \
             by the final depth {} gives: {}.", x, y, x * y);
    let (x, y) = part_two(&input)?;
    println!("(PII) Multiplying the horizontal position {} \
             by the final depth {} gives: {}.", x, y, x * y);
    Ok(())
}

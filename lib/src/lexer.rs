#[derive(Debug)]
pub enum Token {
    Dice,
    LParan,
    RParan,
    Lowest,
    Highest,
    Plus,
    Minus,
    Num(usize),
}

pub fn lex_row(row: String) -> Vec<Token> {
    let mut tokens = Vec::new();
    for field in row.split_whitespace() {
        match field {
            "d" => tokens.push(Token::Dice),
            "(" => tokens.push(Token::LParan),
            ")" => tokens.push(Token::RParan),
            "l" => tokens.push(Token::Lowest),
            "h" => tokens.push(Token::Highest),
            "+" => tokens.push(Token::Plus),
            "-" => tokens.push(Token::Minus),
            tok => {
                if let Ok(num) = tok.parse::<usize>() {
                    tokens.push(Token::Num(num))
                } else {
                    todo!("Add proper error handling. Field {tok} could not be parsed");
                }
            },
        }
    }
    tokens
}

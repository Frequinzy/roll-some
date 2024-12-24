mod lexer;

use lexer::lex_row;
use lexer::Token as Token;

pub fn do_stuff() {
    let toks = lex_row("d ( ) l h + - 1 23 456".to_string());
    println!("{toks:?}");
    println!("Hello, world")
}

// row = expr | expr, row | expr, operator, row ;
// expr = modifier, action | action | int ;
// action = "d", int | "(", row, ")" ;
// modifier = low | high | int ;
// low = "l", int | int, "l", int ;
// high = "h", int | int, "h", int ;
// operator = "+", | "-" ;
struct Row {
    exprs: Vec<Expr>,
}

struct Expr {
    action: Action,
    modifier: Modifier,
}

enum Action {
    Die(usize),
    Parans(Row),
}

enum Modifier {
    Lowest{ pick_n: usize, from_n: usize},
    Highest{ pick_n: usize, from_n: usize},
    Multiplier(usize),
}

fn parse_toks(toks: Vec<Token>) -> Row { 

    Row { exprs: vec![] }
}

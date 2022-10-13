# Design

A doc to keep track of potential design decisions. Not necessarily planned implementations.

## Lexing

```enbf

```

## Parsing

```ebnf

program: statement+;

statement: expression ;

expression: literal ;



literal: string_literal
       | integer_literal
       | decimal_literal
       ;

string_literal: \"[^"]*\" ; (* Any quoted characters *)
integer_literal: [0..9]+ ;

decimal_literal: [0..9]+\.[0..9]
               | \.[0..9]+
               ;

```

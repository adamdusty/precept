# Design

A doc to keep track of potential design decisions. Not necessarily planned implementations.

## Lexing

prefix c: unicode category  
prefic p: unicode property

```ebnf
source: UTF-8 unicode text

line_break: cZl
white_space: pPattern_White_Space

tokens: token | tokens token
tokens: +
      | -
      | *
      | /
      | .
      | ,
      | (
      | )
      | {
      | }
      | [
      | ]
      | ?
      | ;
      | :
      | =
      | ==
      | !
      | !=
      | >
      | >=
      | <
      | <=
      | identifier
      | string_literal
      | character_literal
      | integer_literal
      | decimal_literal

identifier: 
```

## Parsing

```ebnf



```

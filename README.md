# Monkey Interpreter in Go

This repository follows the implementation steps from **_Writing an Interpreter in Go_** by **Thorsten Ball**, a guide to understanding how a programming language works from the ground up. The goal of this project is to implement the **Monkey programming language** in Go, practicing key concepts of lexing, parsing, abstract syntax trees (ASTs), and evaluation along the way.

## Purpose
The repository serves as a learning project and reference implementation of a simple, expressive, dynamically typed language.

By the end of this project, the interpreter will be able to:
- Tokenize Monkey source code into lexemes (Lexer)
- Parse expressions and statements into an AST (Parser)
- Evaluate programs interactively (Evaluator + REPL)
- Support functions, conditionals, arrays, hashes, and closures

## Project Structure
```
monkey-interpreter/
├── token/         # Token definitions and keyword mappings
├── lexer/         # Converts source code into tokens
├── ast/           # Abstract Syntax Tree definitions
├── parser/        # Builds AST from tokens
├── object/        # Runtime value types (Integer, Boolean, etc.)
├── evaluator/     # Executes the AST to produce results
├── repl/          # Interactive Read-Eval-Print Loop
└── main.go        # Entry point for REPL
```

## Credits
The design, flow, and learning material are based on:
> **Thorsten Ball (2016).** *Writing an Interpreter in Go.* [https://interpreterbook.com](https://interpreterbook.com)

All credit for the concept and educational material belongs to Thorsten Ball. This repository simply reproduces and tracks implementation progress for educational purposes.

## License
This project’s code is MIT licensed. The book content and its examples are copyright Thorsten Ball and used here under fair use for personal study.

---
**Author of this Implementation:** Matthew Grayson  
**Source Book:** Writing an Interpreter in Go, Thorsten Ball  
**Purpose:** Educational / Learning Compiler and Interpreter Design


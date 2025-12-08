# Monkey Interpreter

This repository contains a *Monkey* programming language interpreter written in Go.  
The project follows and expands upon the structure presented in **_Writing an Interpreter in Go_** by *Thorsten Ball* (2016), a widely respected introductory text for building interpreters from scratch.

## What Monkey Is

Monkey is a small, expression-oriented programming language designed specifically for learning how interpreters work.  
It includes:

- **Variable bindings**: `let x = 5;`
- **Integers** and **arithmetic**: `5 + 10 * 2;`
- **Booleans** and comparisons: `true == false;`
- **First-class functions**:
  ```
  let add = fn(a, b) { a + b };
  add(2, 3);
  ```
- **Lexical scoping**
- **Closures**
- **REPL support**

Monkey is intentionally minimal. Its simplicity makes it an ideal tool for understanding core concepts behind programming languages without being burdened by large specifications or complex compiler infrastructure.

## What This Interpreter Does

This interpreter is being built step-by-step:

### 1. Lexer (Tokenizer)
The lexer converts raw characters into a stream of tokens—basic syntactic units such as:

- identifiers
- numbers
- operators (`+`, `-`, `*`, `/`, `==`, `!=`)
- delimiters (`(`, `)`, `{`, `}`)
- keywords (`let`, `fn`, `true`, `false`, `return`)

The lexer is responsible for:
- ignoring whitespace
- reading identifiers and numbers
- classifying keywords
- reporting illegal characters

### 2. Parser (Pratt Parser)
A Pratt parser is used to build an abstract syntax tree (AST).  
It enables:
- operator precedence
- prefix/infix parsing
- flexible extension of grammar rules

Example parsed expression:

```
5 + 10 * 2
```

becomes an AST where `*` binds tighter than `+`.

### 3. AST Evaluation
The evaluator walks the AST and executes Monkey programs.  
It handles:

- arithmetic and boolean expressions
- variable environments (scopes)
- function calls
- return values
- blocks and conditionals

### 4. REPL
A read-eval-print loop lets users type Monkey programs interactively:

```
$ monkey
>> let x = 10;
>> x + 5;
15
```

## Project Goals

- Fully implement the Monkey language described in the book
- Add additional features not in the book (Unicode, extended operators, enhanced REPL, etc.)
- Maintain high code quality using Go idioms and tooling (`gofmt`, `pre-commit`, tests`)
- Provide a clean reference implementation for others learning interpreters

## Reference

Thorsten Ball.  
**_Writing an Interpreter in Go_**.  
(2016).  
https://interpreterbook.com/

This repository is an educational implementation inspired by the structure and exercises in the book, with modifications and extensions introduced along the way.

## Pre-commit Setup

This repo uses [pre-commit](https://pre-commit.com) to enforce formatting and hygiene.

### 1) Install

#### macOS
```bash
brew install pre-commit    # or: pipx install pre-commit
```

#### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y python3-pip
pip3 install --user pre-commit    # or: pipx install pre-commit
```

#### Fedora
```bash
sudo dnf install -y python3-pip
pip3 install --user pre-commit
```

#### Windows
PowerShell:
```powershell
winget install Python.Python.3.12
python -m pip install --user pre-commit   # or: pipx install pre-commit
```
WSL Ubuntu: follow the Ubuntu steps inside WSL.

> Ensure Go is installed so `gofmt` is on PATH.

### 2) Enable locally
```bash
pre-commit install           # installs Git hook
pre-commit run --all-files   # optional: run on entire repo
```

### 3) Hooks in this repo

```yaml
---
repos:
  - repo: local
    hooks:
      - id: gofmt
        name: gofmt
        entry: gofmt -s -w
        language: system
        types: [go]

  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v5.0.0
    hooks:
      - id: trailing-whitespace
      - id: end-of-file-fixer
      - id: mixed-line-ending
      - id: check-merge-conflict
      - id: check-yaml
      - id: check-toml

  - repo: https://github.com/igorshubovych/markdownlint-cli
    rev: v0.41.0
    hooks:
      - id: markdownlint
        args: ["--disable", "MD013"]  # allow long lines in tables/code blocks
```

### 4) Update hooks
```bash
pre-commit autoupdate
git add .pre-commit-config.yaml
```

### 5) Windows tips
```bash
git config core.autocrlf input   # avoid CRLF/LF churn
```

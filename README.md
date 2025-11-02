# monkey-interpreter

# Pre-commit Setup

This repo uses [pre-commit](https://pre-commit.com) to enforce formatting and hygiene.

## 1) Install

### macOS
```bash
brew install pre-commit    # or: pipx install pre-commit
```

### Ubuntu/Debian
```bash
sudo apt-get update
sudo apt-get install -y python3-pip
pip3 install --user pre-commit    # or: pipx install pre-commit
```

### Fedora
```bash
sudo dnf install -y python3-pip
pip3 install --user pre-commit
```

### Windows
PowerShell:
```powershell
winget install Python.Python.3.12
python -m pip install --user pre-commit   # or: pipx install pre-commit
```
WSL Ubuntu: follow the Ubuntu steps inside WSL.

> Ensure Go is installed so `gofmt` is on PATH.

## 2) Enable locally
```bash
pre-commit install           # installs Git hook
pre-commit run --all-files   # optional: run on entire repo
```

## 3) Hooks in this repo

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

## 4) Update hooks
```bash
pre-commit autoupdate
git add .pre-commit-config.yaml
```

## 5) Windows tips
```bash
git config core.autocrlf input   # avoid CRLF/LF churn
```

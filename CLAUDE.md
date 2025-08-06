# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

`tsk` is a task runner project written in Rust (edition 2024, rust-version 1.88+) that aims to execute various tasks.

## Development Commands

### Build
```bash
cargo build          # Debug build
cargo build --release # Release build
```

### Run
```bash
cargo run           # Run in debug mode
cargo run --release # Run in release mode
```

### Test
```bash
cargo test         # Run all tests
cargo test [test_name] # Run specific test
cargo test -- --nocapture # Show println! output during tests
```

### Lint & Format
```bash
cargo fmt           # Format code
cargo fmt --check   # Check formatting without applying changes
cargo clippy        # Run Rust linter
cargo clippy -- -D warnings # Treat warnings as errors
```

### Documentation
```bash
cargo doc --open   # Generate and open documentation
```

## Code Architecture

The project is currently in its initial stage with a simple structure:
- `src/main.rs`: Entry point containing the main function
- Future task runner functionality will be implemented to execute various tasks as described in the README

## Git Commit Guidelines

Follow the conventions defined in CONTRIBUTING.md:
- **feat**: New feature
- **fix**: Bug fix
- **refactor**: Code change that neither fixes a bug nor adds a feature
- **test**: Adding or correcting tests
- **style**: Formatting changes
- **chore**: Build process or auxiliary tools changes
- **docs**: Documentation changes

## Text Linting

The project uses textlint for documentation with Japanese technical writing presets configured in `.textlintrc.json`.
# CLAUDE.md

**Note**: This project uses [bd (beads)](https://github.com/steveyegge/beads)
for issue tracking. Use `bd` commands instead of markdown TODOs.
See AGENTS.md for workflow details.

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

hashii is a Go CLI tool that displays the HashiCorp "H" logo in ANSI colors. It supports multiple sizes (small, medium, large), color options (blue, cyan, green, magenta, red, yellow, white, mix, random), plain output mode, and a "dazzle mode" that rapidly cycles through colors.

## Build Commands

```bash
# Build the binary
make build

# Quick build (just compile)
make

# Install (builds, generates docs, installs)
make install

# Format code
make fmt

# Lint code (requires golint)
make lint

# Vet code
make vet

# Clean build artifacts
make clean
```

## Running

```bash
# Direct execution after build
./hashii -size=medium -color=random

# Or via go run
go run main.go -size=large -color=mix

# Dazzle mode (Ctrl+C to exit)
./hashii -dazzle -size=large

# Plain output (no ANSI colors)
./hashii -plain
```

## Architecture

Single-file application (`main.go`) with:
- Base64-encoded logo art in three sizes (`haLg`, `haMd`, `haSm` constants)
- ANSI color codes defined as `ColorCode` type with a `colorMap` for lookups
- Flag-based CLI using standard library `flag` package
- Dazzle mode runs in an infinite loop with signal handling for clean exit

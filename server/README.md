# Language Server

This small language server parses Jest diffs and converts them to valid JSON.

The parsing is based on the original JSON grammar and we're using ANTLR4 to generate a parser in Go.

## Requirements

Install [ANTLR v4](https://github.com/antlr/antlr4) with Homebrew: `brew install antlr`

## Generating the parser

You can generate the parser by running:

```
antlr -Dlanguage=Go -o internal/parser JSON.g4
```

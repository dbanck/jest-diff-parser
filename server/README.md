# Language Server

This small language server parses and validates Jest diffs.

The parsing is based on the original JSON grammar with modifications for Jest output. We're using ANTLR4 to generate a parser in Go.

## Requirements

Install [ANTLR v4](https://github.com/antlr/antlr4) with Homebrew: `brew install antlr`

## Generating the parser

You can generate the parser by running:

```
antlr -Dlanguage=Go -o internal/parser JestDiff.g4
```

## Notes

It's currently possible to parse the diff from `examples/exampleV1.diff`. For real diff parsing, like `examples/exampleV2.diff`, the grammar needs to be improved to handle trailing commas and allow objects on the outer most level.
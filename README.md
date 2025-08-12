# JSON-Parser

## Overview

A simple JSON parser in Go that resolves JSON strings into functions to retrieve data. It supports interpreting JSON expressions and can be extended with nested Lisp-like conditional syntax.

## Usage

Run the parser with:

```bash
go run main.go
```
Make sure the JSON files use the predefined functions. You can use existing JSON files and maps as references to add your own new functions.

## Features
- Parses JSON strings into callable functions.
- Supports custom functions defined in JSON.
- Enables parsing of nested Lisp-like syntax conditions via templating.

## Future Work
- Enhance the parser with a templating mechanism to handle complex nested conditions.
- Support additional operators and more advanced expression parsing.


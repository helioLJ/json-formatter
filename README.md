# JSON Formatter

A simple CLI tool to format and pretty-print JSON data.

## Features

- Read JSON data from a file or standard input.
- Format and pretty-print the JSON data.
- Output the formatted JSON to a file or standard output.

## Usage

### Format JSON from a file and output to another file:
```sh
./json-formatter -input input.json -output output.json
```

### Format JSON from a file and print to standard output:

```sh
./json-formatter -input input.json
```

### Format JSON from standard input and print to standard output:

```sh
cat input.json | ./json-formatter
```

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/json-formatter.git
cd json-formatter
```

2. Build the project:

```sh
go build -o json-formatter
```
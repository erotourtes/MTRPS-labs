# Markdown Parser

**Basic Markdown to HTML converter**  

## Usage
```bash
Usage: $0 [--format <ansi|html>] [--out <file>] <input>
  --format <ansi|html>: Specify the output format as 'ansi' or 'html'
  --out <file>: Specify the output file (default: stdout if not specified)
  <input>: Path to the input file
```
### Example
```bash
./mp --out output.html input.md
```

## Installation
```bash
git clone --depth 1 https://github.com/erotourtes/MTRPS-labs.git
cd MTRPS-labs/
go build -o mp main.go
```
## Run
```bash
./mp README.md
# or 
go run main.go README.md
```

## Run tests
```bash
go test -v ./...
```

## Revert commit 
`4f0398e79b987a75ff38cd74ba137acfa10cd4f1`  
[Github](https://github.com/erotourtes/MTRPS-labs/commit/4f0398e79b987a75ff38cd74ba137acfa10cd4f1)

## Falling tests
`aab835548d5d5fc9d77097327147342d5dad62e3`  
[Github](https://github.com/erotourtes/MTRPS-labs/commit/aab835548d5d5fc9d77097327147342d5dad62e3)

## Conclusion
I noticed that in such projects as parsers, or converters, it is much easier to write tests.
Sometimes, I even wrote tests before writing the main code, which was very helpful.

I didn't know `go` before, therefore didn't know the ways to write tests.
The `table-driven` tests are simple and easy to write, however in this project, I couldn't apply
them to every module, for example [parser_test.go](./lib/parser/parser_test.go), and as a result,
I had to write a lot of similar tests.

In my other [project](https://github.com/erotourtes/harpooner), I wrote a plugin for JetBrains, and there I couldn't write tests at all, because
of the architecture, missing knowledge, and the fact that I don't own some classes.

To sum it up, tests are very helpful and give one confidence in the code changes, merges, and refactoring.

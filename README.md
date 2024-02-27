# Markdown Parser

**Basic Markdown to HTML converter**  

## Usage
```bash
Usage:
  -out string
        Output file
        Default stdout
  <input file>       

Example:
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
./mp --out <output file> <input file>
# or 
go run main.go --out <output file> <input file>
```

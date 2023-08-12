# Golang Calculator Project Example

This is a simple calculator featuring the golang project structure.

It is written following the *"For the Love of Go"* Book by Kevin Cunningham.

The only difference is the updated command to run the calculator. It has added help message and option to run all calculator commands.

## TODO

[ ] Refactor the run command to be more clean
[ ] Include at least two more functions to the calculator e.g. log, power, cos, sin
[ ] Refactor the Sqrt function to not rely on `math` package


## Build and run

To build the project

1. First clone the repository

```shell 
git clone https://github.com/vkolev/calculator
```

2. Change the directory to the new created directory
```shell 
cd calculator
```

3. Run the tests
```shell
go test
```

4. Build an executable
```shell
go build -o calculator ./cmd/calculator
```

## Usage

```shell
./calculator help
```
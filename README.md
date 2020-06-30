# GO_algos
>Calculate algorithms throw your terminal

GO-algos is a program that calculates famous algorithms with passed user's inputs and returns calculated results all in your terminal.

## Installation
Linux:
```sh
clone https://github.com/MaxPolarfox/algo.git
make build
```

OS X 
```sh
clone https://github.com/MaxPolarfox/algo.git
make build-mac
```

Windows:

```sh
clone https://github.com/MaxPolarfox/algo.git
make build-windows
```

## Usage example

Algos that can be run in terminal:
1. "fibonacci"
2. "longest_peak"
3. "permutations"

...Many more are coming soon.

Input:
```sh
make build-mac
./Go_algos -algo=permutations 1 2 3
```

Result:
```sh
Result: [[1 2 3] [1 3 2] [2 1 3] [2 3 1] [3 1 2] [3 2 1]]%
```

Testing:
```sh
make test
```

## Development setup

```sh
go mod tidy
```

## Release History

* 0.0.1
    * Work in progress

## Meta

Maksim Pesetski  – maksim.pesetski@gmail.com


[https://github.com/MaxPolarfox](https://github.com/MaxPolarfox)

## Contributing

1. Fork it (<https://github.com/MaxPolarfox/algo.git>)
2. Create your feature branch (`git checkout -b feature/fooBar`)
3. Commit your changes (`git commit -am 'Add some fooBar'`)
4. Push to the branch (`git push origin feature/fooBar`)
5. Create a new Pull Request
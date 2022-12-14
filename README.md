# parkerwords-go

Go implementation of the problem originally discussed at https://www.youtube.com/watch?v=_-AfhLQfb6w.

This solution is based on the https://github.com/oisyn/parkerwords/ and utilizes
bit sets and go-routines to make things happen in parallel.

There is also Rust implementation of the same problem available at https://github.com/yarcat/parkerwords-rs.

## Benchmarks Overview

Processor	Intel(R) Core(TM) i7-7700K CPU @ 4.20GHz, 4200 Mhz, 4 Core(s), 8 Logical Processor(s)

Please note that I'm considering only processing times. 

Averages of 3 runs:

1) Rust: 8.437ms
2) Go: 9.136ms
3) C++: 14.638ms

## Go Benchmarks

```
$ go run .
538 solutions written to solutions.txt
Total time: 31093µs
Read:       21097µs
Process:     9749µs
Write:        247µs

$ go run .
538 solutions written to solutions.txt
Total time: 31171µs
Read:       22084µs
Process:     9087µs
Write:          0µs

$ go run .
538 solutions written to solutions.txt
Total time: 27300µs
Read:       17714µs
Process:     8571µs
Write:       1014µs
```

## Rust Benchmarks

```
$ cargo build --release

$ ./target/release/parkerwords-rs.exe
538 solutions written to solutions.txt
Total time:    15778µs
Read:           6810µs
Process:        8432µs
Write:           535µs

$ ./target/release/parkerwords-rs.exe
538 solutions written to solutions.txt
Total time:    15211µs
Read:           6318µs
Process:        8411µs
Write:           481µs

$ ./target/release/parkerwords-rs.exe
538 solutions written to solutions.txt
Total time:    15672µs
Read:           6714µs
Process:        8468µs
Write:           488µs
```

## C++ Benchmarks

```
$ g++ -o parkerwords -O3 parkerwords.cpp -pthread -std=c++20

$ ./parkerwords
538 solutions written to solutions.txt.
Total time: 147251us (0.147251s)
Read:      133279us
Process:    13972us
Write:          0us

$ ./parkerwords
538 solutions written to solutions.txt.
Total time: 27001us (0.027001s)
Read:       11030us
Process:    15971us
Write:          0us

$ ./parkerwords
538 solutions written to solutions.txt.
Total time: 26000us (0.026s)
Read:       11028us
Process:    13972us
Write:       1000us
```

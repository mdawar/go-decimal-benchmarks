# Go Decimal Packages Benchmarks

### Benchmark Results

[Benchmark Workflow](https://github.com/mdawar/go-decimal-benchmarks/actions/workflows/benchmark.yml)

### Run Benchmarks

```sh
just benchmark
# Or
go test -bench . -benchmem
```

### Run Package Benchmarks

```sh
just package shopspring
just package alpaca
just package eric
just package apd
just package govalues
just package udecimal
```

Without `just`:

```sh
go test -bench Benchmark.*/.*/shopspring -benchmem
go test -bench Benchmark.*/.*/alpaca -benchmem
go test -bench Benchmark.*/.*/eric -benchmem
go test -bench Benchmark.*/.*/apd -benchmem
go test -bench Benchmark.*/.*/govalues -benchmem
go test -bench Benchmark.*/.*/udecimal -benchmem
```

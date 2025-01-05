# Fast Fourier Transform Implementation
[![Rust](https://img.shields.io/badge/Rust-1.75%2B-orange.svg)](https://www.rust-lang.org)
[![Go](https://img.shields.io/badge/Go-1.21%2B-blue.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![FFT](https://img.shields.io/badge/Algorithm-FFT-green.svg)](https://en.wikipedia.org/wiki/Fast_Fourier_transform)

This repository contains implementations of the Fast Fourier Transform (FFT) algorithm in both Rust and Go, demonstrating different approaches to parallel computation in these languages.

## Implementation Differences

### Parallelization Approach

#### Go Implementation
- Uses goroutines and `sync.WaitGroup` for parallelization
- Direct recursive parallelization with goroutines at each level
- Lightweight concurrency model with built-in primitives
- Example:
```go
var wg sync.WaitGroup
wg.Add(2)
go func() {
    defer wg.Done()
    evenFFT = ParallelFFT(even)
}()
```

#### Rust Implementation
- Uses Rayon for parallel computation
- Two-tiered approach with parallel top level and sequential recursion
- Thread pool-based parallelization using `rayon::join`
- Example:
```rust
let (even_fft, odd_fft) = rayon::join(
    || fft(&even),
    || fft(&odd)
);
```

### Memory Management

#### Go Implementation
- Garbage collected memory management
- Simpler memory handling with slices
- Direct array access and manipulation
- Example:
```go
even := make([]complex128, n/2)
odd := make([]complex128, n/2)
```

#### Rust Implementation
- Zero-cost abstractions with ownership system
- Explicit memory management through borrowing
- Vec capacity pre-allocation for performance
- Example:
```rust
let mut even = Vec::with_capacity(n / 2);
let mut odd = Vec::with_capacity(n / 2);
```

### Complex Numbers

#### Go Implementation
- Uses built-in `complex128` type
- Native complex arithmetic operations
- Example:
```go
t := cmplx.Exp(complex(0, -2*math.Pi*float64(k)/float64(n))) * oddFFT[k]
```

#### Rust Implementation
- Uses `num_complex::Complex64` from external crate
- More explicit complex number construction
- Example:
```rust
let t = Complex64::from_polar(1.0, -2.0 * PI * (k as f64) / (n as f64)) * odd_fft[k]
```

## Performance Considerations

### Go
- Faster startup time
- Lower memory overhead
- Better for smaller datasets
- Simple concurrency model

### Rust
- Better CPU utilization with thread pool
- More efficient memory usage
- Better for larger datasets
- Zero-cost abstractions

## Dependencies

### Go
```go
import (
    "math"
    "math/cmplx"
    "sync"
)
```

### Rust
```toml
[dependencies]
num-complex = "0.4"
rayon = "1.8"
```

## Usage

### Go
```bash
go run *.go
```

### Rust
```bash
cargo run --release
```

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

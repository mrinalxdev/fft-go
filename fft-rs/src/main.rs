use num_complex::Complex64;
use std::f64::consts::PI;
use std::sync::Arc;
use rayon::prelude::*;

fn fft(input: &[Complex64]) -> Vec<Complex64> {
    let n = input.len();
    if n <= 1 {
        return input.to_vec();
    }

    let mut even = Vec::with_capacity(n / 2);
    let mut odd = Vec::with_capacity(n / 2);

    for i in 0..(n / 2) {
        even.push(input[2 * i]);
        odd.push(input[2 * i + 1]);
    }

    let even_fft = fft(&even);
    let odd_fft = fft(&odd);

    // combinig the results
    // with the twiddle factor
    let mut output = vec![Complex64::new(0.0, 0.0); n];

    for k in 0..(n / 2) {
        let t = Complex64::from_polar(
            1.0,
            -2.0 * PI * (k as f64) / (n as f64)
        ) * odd_fft[k];

        output[k] = even_fft[k] + t;
        output[k + n/2] = even_fft[k] - t;
    }

    output
}


pub fn parallel_fft(input: Vec<Complex64>) -> Vec<Complex64> {
    let n = input.len();
    if n <= 1 {
        return input;
    }

    let mut even = Vec::with_capacity(n / 2);
    let mut odd = Vec::with_capacity(n / 2);

    for i in 0..(n / 2) {
        even.push(input[2 * i]);
        odd.push(input[2 * i + 1]);
    }

    let (even_fft, odd_fft) = rayon::join(
        || fft(&even),
        || fft(&odd)
    );

    let mut output = vec![Complex64::new(0.0, 0.0); n];

    for k in 0..(n / 2) {
        let t = Complex64::from_polar(
            1.0,
            -2.0 * PI * (k as f64) / (n as f64)
        ) * odd_fft[k];

        output[k] = even_fft[k] + t;
        output[k + n/2] = even_fft[k] - t;
    }

    output
}

// function to generate a test signal
pub fn generate_signal(size: usize) -> Vec<Complex64> {
    let mut signal = Vec::with_capacity(size);
    for i in 0..size {
        signal.push(Complex64::new(
            (2.0 * PI * i as f64 / size as f64).sin(),
            0.0
        ));
    }
    signal
}

pub fn magnitude(c: Complex64) -> f64 {
    c.norm()
}

fn main() {
    let n = 1024;
    let signal = generate_signal(n);

    let start = std::time::Instant::now();
    let result = parallel_fft(signal);
    let elapsed = start.elapsed();

    println!("FFT completed in {:?}", elapsed);
    println!("First 10 FFT results:");
    for i in 0..10 {
        println!("Result[{}]: {}", i, result[i]);
    }
}

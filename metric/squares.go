// Copyright (c) 2019, The eTable Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metric

import (
	"math"

	"github.com/chewxy/math32"
	"github.com/emer/etable/norm"
)

///////////////////////////////////////////
//  SumSquares

// SumSquares32 computes the sum-of-squares distance between two vectors.
// Skips NaN's and panics if lengths are not equal.
func SumSquares32(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float32(0)
	for i, av := range a {
		bv := b[i]
		if math32.IsNaN(av) || math32.IsNaN(bv) {
			continue
		}
		dv := av - bv
		ss += dv * dv
	}
	return ss
}

// SumSquares64 computes the sum-of-squares distance between two vectors.
// Skips NaN's and panics if lengths are not equal.
func SumSquares64(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float64(0)
	for i, av := range a {
		bv := b[i]
		if math.IsNaN(av) || math.IsNaN(bv) {
			continue
		}
		dv := av - bv
		ss += dv * dv
	}
	return ss
}

///////////////////////////////////////////
//  Euclidean

// Euclidean32 computes the square-root of sum-of-squares distance
// between two vectors (aka the L2 norm).
// Skips NaN's and panics if lengths are not equal.
func Euclidean32(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := SumSquares32(a, b)
	return math32.Sqrt(ss)
}

// Euclidean64 computes the square-root of sum-of-squares distance
// between two vectors (aka the L2 norm).
// Skips NaN's and panics if lengths are not equal.
func Euclidean64(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := SumSquares64(a, b)
	return math.Sqrt(ss)
}

///////////////////////////////////////////
//  Covariance

// Covariance32 computes the mean of the co-product of each vector element minus
// the mean of that vector: cov(A,B) = E[(A - E(A))(B - E(B))]
// Skips NaN's and panics if lengths are not equal.
func Covariance32(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float32(0)
	am := norm.Mean32(a)
	bm := norm.Mean32(b)
	n := 0
	for i, av := range a {
		bv := b[i]
		if math32.IsNaN(av) || math32.IsNaN(bv) {
			continue
		}
		ss += (av - am) * (bv - bm)
		n++
	}
	if n > 0 {
		ss /= float32(n)
	}
	return ss
}

// Covariance64 computes the mean of the co-product of each vector element minus
// the mean of that vector: cov(A,B) = E[(A - E(A))(B - E(B))]
// Skips NaN's and panics if lengths are not equal.
func Covariance64(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float64(0)
	am := norm.Mean64(a)
	bm := norm.Mean64(b)
	n := 0
	for i, av := range a {
		bv := b[i]
		if math.IsNaN(av) || math.IsNaN(bv) {
			continue
		}
		ss += (av - am) * (bv - bm)
		n++
	}
	if n > 0 {
		ss /= float64(n)
	}
	return ss
}

///////////////////////////////////////////
//  Correlation

// Correlation32 computes the vector similarity in range (-1..1) as the
// mean of the co-product of each vector element minus the mean of that vector,
// normalized by the product of their standard deviations:
// cor(A,B) = E[(A - E(A))(B - E(B))] / sigma(A) sigma(B).
// (i.e., the standardized covariance) -- equivalent to the cosine of mean-normalized
// vectors.
// Skips NaN's and panics if lengths are not equal.
func Correlation32(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float32(0)
	am := norm.Mean32(a)
	bm := norm.Mean32(b)
	var avar, bvar float32
	for i, av := range a {
		bv := b[i]
		if math32.IsNaN(av) || math32.IsNaN(bv) {
			continue
		}
		ad := av - am
		bd := bv - bm
		ss += ad * bd   // between
		avar += ad * ad // within
		bvar += bd * bd
	}
	vp := math32.Sqrt(avar * bvar)
	if vp > 0 {
		ss /= vp
	}
	return ss
}

// Correlation64 computes the vector similarity in range (-1..1) as the
// mean of the co-product of each vector element minus the mean of that vector,
// normalized by the product of their standard deviations:
// cor(A,B) = E[(A - E(A))(B - E(B))] / sigma(A) sigma(B).
// (i.e., the standardized covariance) -- equivalent to the cosine of mean-normalized
// vectors.
// Skips NaN's and panics if lengths are not equal.
func Correlation64(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float64(0)
	am := norm.Mean64(a)
	bm := norm.Mean64(b)
	var avar, bvar float64
	for i, av := range a {
		bv := b[i]
		if math.IsNaN(av) || math.IsNaN(bv) {
			continue
		}
		ad := av - am
		bd := bv - bm
		ss += ad * bd   // between
		avar += ad * ad // within
		bvar += bd * bd
	}
	vp := math.Sqrt(avar * bvar)
	if vp > 0 {
		ss /= vp
	}
	return ss
}

///////////////////////////////////////////
//  InnerProduct

// InnerProduct32 computes the sum of the element-wise product of the two vectors.
// Skips NaN's and panics if lengths are not equal.
func InnerProduct32(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float32(0)
	for i, av := range a {
		bv := b[i]
		if math32.IsNaN(av) || math32.IsNaN(bv) {
			continue
		}
		ss += av * bv
	}
	return ss
}

// InnerProduct64 computes the mean of the co-product of each vector element minus
// the mean of that vector, normalized by the product of their standard deviations:
// cor(A,B) = E[(A - E(A))(B - E(B))] / sigma(A) sigma(B).
// (i.e., the standardized covariance) -- equivalent to the cosine of mean-normalized
// vectors.
// Skips NaN's and panics if lengths are not equal.
func InnerProduct64(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float64(0)
	for i, av := range a {
		bv := b[i]
		if math.IsNaN(av) || math.IsNaN(bv) {
			continue
		}
		ss += av * bv
	}
	return ss
}

///////////////////////////////////////////
//  Cosine

// Cosine32 computes the cosine of the angle between two vectors (-1..1),
// as the normalized inner product: inner product / sqrt(ssA * ssB).
// If vectors are mean-normalized = Correlation.
// Skips NaN's and panics if lengths are not equal.
func Cosine32(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float32(0)
	var ass, bss float32
	for i, av := range a {
		bv := b[i]
		if math32.IsNaN(av) || math32.IsNaN(bv) {
			continue
		}
		ss += av * bv  // between
		ass += av * av // within
		bss += bv * bv
	}
	vp := math32.Sqrt(ass * bss)
	if vp > 0 {
		ss /= vp
	}
	return ss
}

// Cosine32 computes the cosine of the angle between two vectors (-1..1),
// as the normalized inner product: inner product / sqrt(ssA * ssB).
// If vectors are mean-normalized = Correlation.
// Skips NaN's and panics if lengths are not equal.
func Cosine64(a, b []float64) float64 {
	if len(a) != len(b) {
		panic("metric: slice lengths do not match")
	}
	ss := float64(0)
	var ass, bss float64
	for i, av := range a {
		bv := b[i]
		if math.IsNaN(av) || math.IsNaN(bv) {
			continue
		}
		ss += av * bv  // between
		ass += av * av // within
		bss += bv * bv
	}
	vp := math.Sqrt(ass * bss)
	if vp > 0 {
		ss /= vp
	}
	return ss
}

///////////////////////////////////////////
//  InvCosine

// InvCosine32 computes 1 - cosine of the angle between two vectors (-1..1),
// as the normalized inner product: inner product / sqrt(ssA * ssB).
// If vectors are mean-normalized = Correlation.
// Skips NaN's and panics if lengths are not equal.
func InvCosine32(a, b []float32) float32 {
	return 1 - Cosine32(a, b)
}

// InvCosine32 computes 1 - cosine of the angle between two vectors (-1..1),
// as the normalized inner product: inner product / sqrt(ssA * ssB).
// If vectors are mean-normalized = Correlation.
// Skips NaN's and panics if lengths are not equal.
func InvCosine64(a, b []float64) float64 {
	return 1 - Cosine64(a, b)
}

///////////////////////////////////////////
//  InvCorrelation

// InvCorrelation32 computes 1 - the vector similarity in range (-1..1) as the
// mean of the co-product of each vector element minus the mean of that vector,
// normalized by the product of their standard deviations:
// cor(A,B) = E[(A - E(A))(B - E(B))] / sigma(A) sigma(B).
// (i.e., the standardized covariance) -- equivalent to the cosine of mean-normalized
// vectors.
// Skips NaN's and panics if lengths are not equal.
func InvCorrelation32(a, b []float32) float32 {
	return 1 - Correlation32(a, b)
}

// InvCorrelation64 computes 1 - the vector similarity in range (-1..1) as the
// mean of the co-product of each vector element minus the mean of that vector,
// normalized by the product of their standard deviations:
// cor(A,B) = E[(A - E(A))(B - E(B))] / sigma(A) sigma(B).
// (i.e., the standardized covariance) -- equivalent to the cosine of mean-normalized
// vectors.
// Skips NaN's and panics if lengths are not equal.
func InvCorrelation64(a, b []float64) float64 {
	return 1 - Correlation64(a, b)
}
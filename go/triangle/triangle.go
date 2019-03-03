// Package triangle implements has constants and methods associated with decyphering the input triangle
package triangle

import (
	"math"
)

// Kind is the status code for the type of triangle
type Kind int

const (
	// Deg is a Degenerate triangle
	Deg Kind = iota
	// NaT is Not a Triangle
	NaT
	// Equ is a equilateral triangle
	Equ
	// Iso is a isosceles triangle
	Iso
	// Sca is a scalene triangle
	Sca
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		return NaT
	}

	if math.IsInf(a, 1) || math.IsInf(b, 1) || math.IsInf(c, 1) {
		return NaT
	}

	if math.IsInf(a, -1) || math.IsInf(b, -1) || math.IsInf(c, -1) {
		return NaT
	}

	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	if (a == b && a+b == c) || (a == c && a+c == b) || (b == c && b+c == a) {
		return Deg
	}

	if a == b && a == c {
		return Equ
	}

	if a == b || b == c || a == c {
		return Iso
	}

	if a != b && a != c && b != c {
		return Sca
	}

	return NaT
}

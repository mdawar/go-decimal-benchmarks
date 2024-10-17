package main_test

import (
	"fmt"
	"testing"

	"github.com/alpacahq/alpacadecimal"
	"github.com/cockroachdb/apd/v3"
	ericlagergren "github.com/ericlagergren/decimal"
	govalues "github.com/govalues/decimal"
	"github.com/quagmt/udecimal"
	shopspring "github.com/shopspring/decimal"
)

func BenchmarkMultiplication(b *testing.B) {
	cases := []struct {
		x, y string
	}{
		{"1", "1"},
		{"1", "100"},
		{"1000", "0.1"},
		{"1000", "0.01"},
		{"1000", "0.001"},
	}

	for _, tc := range cases {
		b.Run(fmt.Sprintf("%sx%s", tc.x, tc.y), func(b *testing.B) {
			b.Run("shopspring", func(b *testing.B) {
				x := shopspring.RequireFromString(tc.x)
				y := shopspring.RequireFromString(tc.y)

				b.ResetTimer()
				for range b.N {
					_ = x.Mul(y)
				}
			})

			b.Run("alpacadecimal", func(b *testing.B) {
				x := alpacadecimal.RequireFromString(tc.x)
				y := alpacadecimal.RequireFromString(tc.y)

				b.ResetTimer()
				for range b.N {
					_ = x.Mul(y)
				}
			})

			b.Run("apd", func(b *testing.B) {
				x, _, err := apd.NewFromString(tc.x)
				if err != nil {
					b.Fatal(err)
				}

				y, _, err := apd.NewFromString(tc.y)
				if err != nil {
					b.Fatal(err)
				}

				b.ResetTimer()
				for range b.N {
					result := new(apd.Decimal)
					_, err := apd.BaseContext.Mul(result, x, y)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("ericlagergren", func(b *testing.B) {
				x, ok := new(ericlagergren.Big).SetString(tc.x)
				if !ok {
					b.Fatal("failed to parse x")
				}

				y, ok := new(ericlagergren.Big).SetString(tc.y)
				if !ok {
					b.Fatal("failed to parse y")
				}

				b.ResetTimer()
				for range b.N {
					_ = new(ericlagergren.Big).Mul(x, y)
				}
			})

			b.Run("govalues", func(b *testing.B) {
				x := govalues.MustParse(tc.x)
				y := govalues.MustParse(tc.y)

				b.ResetTimer()
				for range b.N {
					_, err := x.Mul(y)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("udecimal", func(b *testing.B) {
				x := udecimal.MustParse(tc.x)
				y := udecimal.MustParse(tc.y)

				b.ResetTimer()
				for range b.N {
					_ = x.Mul(y)
				}
			})
		})
	}
}

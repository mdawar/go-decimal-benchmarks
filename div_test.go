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

func BenchmarkDivision(b *testing.B) {
	cases := []struct {
		x, y string
	}{
		{"1", "1"},
		{"1", "100"},
		{"1000", "0.1"},
		{"1000", "0.01"},
		{"1000", "0.001"},
		{"123456789.12345678912", "0.001"},
	}

	for _, tc := range cases {
		b.Run(fmt.Sprintf("%s÷%s", tc.x, tc.y), func(b *testing.B) {
			b.Run("shopspring", func(b *testing.B) {
				x := shopspring.RequireFromString(tc.x)
				y := shopspring.RequireFromString(tc.y)

				b.ResetTimer()
				b.Run("Div", func(b *testing.B) {
					for range b.N {
						_ = x.Div(y)
					}
				})

				b.Run("DivRound", func(b *testing.B) {
					for range b.N {
						_ = x.DivRound(y, 0)
					}
				})
			})

			b.Run("alpacadecimal", func(b *testing.B) {
				x := alpacadecimal.RequireFromString(tc.x)
				y := alpacadecimal.RequireFromString(tc.y)

				b.ResetTimer()
				b.Run("Div", func(b *testing.B) {
					for range b.N {
						_ = x.Div(y)
					}
				})

				b.Run("DivRound", func(b *testing.B) {
					for range b.N {
						_ = x.DivRound(y, 0)
					}
				})
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

				ctx := apd.BaseContext.WithPrecision(19)

				b.ResetTimer()
				b.Run("Quo", func(b *testing.B) {
					for range b.N {
						result := new(apd.Decimal)
						_, err := ctx.Quo(result, x, y)
						if err != nil {
							b.Fatal(err)
						}
					}
				})

				b.Run("QuoInteger", func(b *testing.B) {
					for range b.N {
						result := new(apd.Decimal)
						_, err := ctx.QuoInteger(result, x, y)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
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
				b.Run("Quo", func(b *testing.B) {
					for range b.N {
						_ = new(ericlagergren.Big).Quo(x, y)
					}
				})

				b.Run("QuoRem", func(b *testing.B) {
					for range b.N {
						_, _ = new(ericlagergren.Big).QuoRem(x, y, new(ericlagergren.Big))
					}
				})

				b.Run("QuoInt", func(b *testing.B) {
					for range b.N {
						_ = new(ericlagergren.Big).QuoInt(x, y)
					}
				})
			})

			b.Run("govalues", func(b *testing.B) {
				x := govalues.MustParse(tc.x)
				y := govalues.MustParse(tc.y)

				b.ResetTimer()
				b.Run("Quo", func(b *testing.B) {
					for range b.N {
						_, err := x.Quo(y)
						if err != nil {
							b.Fatal(err)
						}
					}
				})

				b.Run("QuoRem", func(b *testing.B) {
					for range b.N {
						_, _, err := x.QuoRem(y)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})

			b.Run("udecimal", func(b *testing.B) {
				x := udecimal.MustParse(tc.x)
				y := udecimal.MustParse(tc.y)

				b.ResetTimer()
				b.Run("Div", func(b *testing.B) {
					for range b.N {
						_, err := x.Div(y)
						if err != nil {
							b.Fatal(err)
						}
					}
				})
			})
		})
	}
}

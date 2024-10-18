package main_test

import (
	"strconv"
	"testing"

	"github.com/alpacahq/alpacadecimal"
	"github.com/cockroachdb/apd/v3"
	ericlagergren "github.com/ericlagergren/decimal"
	govalues "github.com/govalues/decimal"
	"github.com/quagmt/udecimal"
	shopspring "github.com/shopspring/decimal"
)

func BenchmarkParsing(b *testing.B) {
	cases := []string{
		"0",
		"1",
		"0.31032345678911",
		"1.23",
		"123456.78910111213",
		"12345678910.123456789101112131",
	}

	for _, value := range cases {
		b.Run(value, func(b *testing.B) {
			b.Run("strconv", func(b *testing.B) {
				for range b.N {
					_, err := strconv.ParseFloat(value, 64)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("shopspring", func(b *testing.B) {
				for range b.N {
					_, err := shopspring.NewFromString(value)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("alpacadecimal", func(b *testing.B) {
				for range b.N {
					_, err := alpacadecimal.NewFromString(value)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("apd", func(b *testing.B) {
				for range b.N {
					_, _, err := apd.NewFromString(value)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("ericlagergren", func(b *testing.B) {
				for range b.N {
					_, ok := new(ericlagergren.Big).SetString(value)
					if !ok {
						b.Fatal("parse failed")
					}
				}
			})

			b.Run("govalues", func(b *testing.B) {
				for range b.N {
					_, err := govalues.Parse(value)
					if err != nil {
						b.Fatal(err)
					}
				}
			})

			b.Run("udecimal", func(b *testing.B) {
				for range b.N {
					_, err := udecimal.Parse(value)
					if err != nil {
						b.Fatal(err)
					}
				}
			})
		})
	}
}

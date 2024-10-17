package main_test

import (
	"testing"

	"github.com/alpacahq/alpacadecimal"
	"github.com/cockroachdb/apd/v3"
	ericlagergren "github.com/ericlagergren/decimal"
	govalues "github.com/govalues/decimal"
	"github.com/quagmt/udecimal"
	shopspring "github.com/shopspring/decimal"
)

func BenchmarkParsing(b *testing.B) {
	const value = "123.456"

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
}

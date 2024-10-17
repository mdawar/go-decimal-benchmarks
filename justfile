_default:
	@just -l -u

# Run benchmarks.
benchmark *args:
  go test -bench . -benchmem {{args}}

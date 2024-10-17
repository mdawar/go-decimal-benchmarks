_default:
	@just -l -u

# Run benchmarks.
benchmark:
  go test -bench . -benchmem

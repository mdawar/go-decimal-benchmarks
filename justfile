_default:
	@just -l -u

# Run benchmarks.
benchmark regexp='.' *args='':
  go test -bench {{regexp}} -benchmem {{args}}

_default:
	@just -l -u

# Run benchmarks.
benchmark regexp='.' *args='':
  go test -bench {{regexp}} -benchmem {{args}}

# Run benchmarks for a specific package.
package name *args:
  go test -bench Benchmark.*/.*/{{name}} -benchmem {{args}}

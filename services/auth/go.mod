module example.com/services/auth

go 1.23.1

require (
	example.com/core v0.0.0
	github.com/google/uuid v1.6.0
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.29.0
)

replace example.com/core v0.0.0 => ../../core

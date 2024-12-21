module example.com/services/posts

go 1.23.1

require (
	example.com/core v0.0.0
	github.com/lib/pq v1.10.9
)

require github.com/google/uuid v1.6.0 // indirect

replace example.com/core v0.0.0 => ../../core

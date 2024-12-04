module example.com/services/posts

go 1.23.1

require (
	example.com/core v0.0.0
	github.com/lib/pq v1.10.9
)

replace example.com/core v0.0.0 => ../../core

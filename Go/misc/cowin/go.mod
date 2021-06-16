module track

go 1.16

replace example.com/types => ./internal/types

replace example.com/logger => ./internal/logger

replace example.com/tracker => ./internal/tracker

require (
	example.com/logger v0.0.0-00010101000000-000000000000
	example.com/tracker v0.0.0-00010101000000-000000000000
)

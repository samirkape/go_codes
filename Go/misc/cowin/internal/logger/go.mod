module logger

go 1.16

require (
	example.com/types v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.7.7
)

replace example.com/types => ../internal/types

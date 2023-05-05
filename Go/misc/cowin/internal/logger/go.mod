module logger

go 1.16

require (
	example.com/types v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.9.0
	github.com/ugorji/go v1.1.7 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
)

replace example.com/types => ../internal/types

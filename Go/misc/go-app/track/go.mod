module example.com/track

go 1.16

require (
	example.com/types v0.0.0-00010101000000-000000000000
	example.com/logger v0.0.0-00010101000000-000000000000
	github.com/go-telegram-bot-api/telegram-bot-api v4.6.4+incompatible
	github.com/technoweenie/multipartstreamer v1.0.1 // indirect
)

replace example.com/types => ../internal/types

replace example.com/logger => ../internal/logger

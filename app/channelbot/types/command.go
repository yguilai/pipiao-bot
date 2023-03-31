package types

type (
	BotCommand struct {
		At      string
		MainCmd string
		SubCmd  string
	}

	CommandHandler interface {
		Handle(cmd *BotCommand) error
	}
)

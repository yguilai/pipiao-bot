package types

type EventHandler interface {
	Register() interface{}
}

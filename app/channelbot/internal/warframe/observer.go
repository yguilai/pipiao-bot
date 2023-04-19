package warframe

type StateObserver interface {
	onChange(str string)
}

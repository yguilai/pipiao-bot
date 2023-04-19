package warframe

type WfState string

const (
	StateAlert      WfState = "Alerts"
	StateSortie             = "Sorties"
	StateInvasion           = "Invasions"
	StateVoidTrader         = "VoidTrades"
	StateNightwave          = ""
	StateCetusCycle         = ""
)

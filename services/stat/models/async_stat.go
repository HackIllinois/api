package models

type AsyncStat struct {
	Service string
	Stat    *Stat
	Error   error
}

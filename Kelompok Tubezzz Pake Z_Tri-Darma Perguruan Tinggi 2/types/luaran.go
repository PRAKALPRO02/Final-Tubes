package types

type Luaran struct {
	Id           int
	IdTridarma   int
	BentukLuaran string
	Pelaksanaan  string
}

type DataLuaran struct {
	LastId int
	Length int
	Data   [100]Luaran
}

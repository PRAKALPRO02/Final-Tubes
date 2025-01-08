package types

type Anggota struct {
	Id         int
	IdTridarma int
	Nama       string
	Role       string
}

type DataAnggota struct {
	LastId int
	Length int
	Data   [100]Anggota
}

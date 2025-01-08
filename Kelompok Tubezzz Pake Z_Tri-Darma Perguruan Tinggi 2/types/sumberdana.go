package types

type Dana struct {
	Id         int
	IdTridarma int
	Sumber     string
	Keterangan string
	Nominal    int
}

type DataDana struct {
	LastId int
	Length int
	Data   [100]Dana
}

package services

import (
	"errors"
	constants "tri_darma/constants"
	database "tri_darma/database"
	types "tri_darma/types"
)

func AddDana(danaData types.Dana) (bool, error) {
	var dataDana = database.ReadJSON[types.DataDana](constants.FullPathDB(constants.DB_DANA_FILE))
	found, datafounded := GetTridarById(danaData.IdTridarma)
	if found {
		danaData.Id = dataDana.LastId + 1
		dataDana.LastId = danaData.Id
		dataDana.Data[dataDana.Length] = danaData
		dataDana.Length += 1
		database.WriteJSON(
			constants.FullPathDB(constants.DB_DANA_FILE),
			dataDana,
		)
		datafounded.SumDana += danaData.Nominal
		ChangeTriDarmaById(datafounded.Id, datafounded)
		return true, nil
	}
	return false, errors.New("ID tridar not found")

}

func RemoveDanaById(id int) (bool, error) {
	var dataDana = database.ReadJSON[types.DataDana](constants.FullPathDB(constants.DB_DANA_FILE))
	var idxToDelete int = -1
	for i := 0; i < dataDana.Length; i++ {
		if dataDana.Data[i].Id == id {
			idxToDelete = i
		}
	}
	if idxToDelete != -1 {
		_, datafounded := GetTridarById(dataDana.Data[idxToDelete].IdTridarma)
		datafounded.SumDana -= dataDana.Data[idxToDelete].Nominal
		for i := idxToDelete; i < dataDana.Length-1; i++ {
			dataDana.Data[i] = dataDana.Data[i+1]
		}
		dataDana.Data[dataDana.Length-1] = types.Dana{}
		dataDana.Length--
		database.WriteJSON(
			constants.FullPathDB(constants.DB_DANA_FILE),
			dataDana,
		)
		ChangeTriDarmaById(datafounded.Id, datafounded)
		return true, nil
	}
	return false, errors.New("DanaId Not Found")
}

func ChangeDanaById(id int, dataDanach types.Dana) (bool, error) {
	var dataDana = database.ReadJSON[types.DataDana](constants.FullPathDB(constants.DB_DANA_FILE))
	var idxToChange int = -1
	for i := 0; i < dataDana.Length; i++ {
		if dataDana.Data[i].Id == id {
			idxToChange = i
		}
	}
	if idxToChange != -1 {
		dataDana.Data[idxToChange] = dataDanach
		database.WriteJSON(
			constants.FullPathDB(constants.DB_DANA_FILE),
			dataDana,
		)
		return true, nil
	}
	return false, errors.New("LuaranId Not Found")
}

func ListDana() types.DataDana {
	return database.ReadJSON[types.DataDana](constants.FullPathDB(constants.DB_DANA_FILE))
}

func GetDanaById(id int) (bool, types.Dana) {
	var dataDana = database.ReadJSON[types.DataDana](constants.FullPathDB(constants.DB_DANA_FILE))
	for i := 0; i < dataDana.Length; i++ {
		if dataDana.Data[i].Id == id {
			return true, dataDana.Data[i]
		}
	}
	return false, types.Dana{}
}

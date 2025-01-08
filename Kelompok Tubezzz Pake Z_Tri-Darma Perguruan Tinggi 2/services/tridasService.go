package services

import (
	"errors"
	constants "tri_darma/constants"
	database "tri_darma/database"
	types "tri_darma/types"
)

func Add3Darma(dataDarma types.TriDarma) (bool, error) {
	var dataTRI = database.ReadJSON[types.DataTriDarma](constants.FullPathDB(constants.DB_PENELITIAN))
	dataDarma.Id = dataTRI.LastId + 1
	dataTRI.LastId = dataDarma.Id
	dataTRI.Data[dataTRI.Length] = dataDarma
	dataTRI.Length += 1
	database.WriteJSON(
		constants.FullPathDB(constants.DB_PENELITIAN),
		dataTRI,
	)
	return true, nil
}

func RemoveTriDarmaById(id int) (bool, error) {
	var dataTRI = database.ReadJSON[types.DataTriDarma](constants.FullPathDB(constants.DB_PENELITIAN))
	var idxToDelete int = -1
	for i := 0; i < dataTRI.Length; i++ {
		if dataTRI.Data[i].Id == id {
			idxToDelete = i
		}
	}
	if idxToDelete != -1 {
		for i := idxToDelete; i < dataTRI.Length-1; i++ {
			dataTRI.Data[i] = dataTRI.Data[i+1]
		}

		dataTRI.Data[dataTRI.Length-1] = types.TriDarma{}
		dataTRI.Length--
		database.WriteJSON(
			constants.FullPathDB(constants.DB_PENELITIAN),
			dataTRI,
		)
		// REMOVE ANGGOTA ALSO
		dataAnggota := ListAnggota()
		for i := 0; i < dataAnggota.Length; i++ {
			if dataAnggota.Data[i].IdTridarma == id {
				RemoveAnggotaById(dataAnggota.Data[i].Id)
			}
		}
		return true, nil
	}
	return false, errors.New("CategoryId Not Found")
}

func ChangeTriDarmaById(id int, dataTrich types.TriDarma) (bool, error) {
	var dataTRI = database.ReadJSON[types.DataTriDarma](constants.FullPathDB(constants.DB_PENELITIAN))
	var idxToChange int = -1
	for i := 0; i < dataTRI.Length; i++ {
		if dataTRI.Data[i].Id == id {
			idxToChange = i
		}
	}
	if idxToChange != -1 {
		dataTRI.Data[idxToChange] = dataTrich
		database.WriteJSON(
			constants.FullPathDB(constants.DB_PENELITIAN),
			dataTRI,
		)
		return true, nil
	}
	return false, errors.New("CategoryId Not Found")
}

func GetTridarById(id int) (bool, types.TriDarma) {
	var dataTRI = database.ReadJSON[types.DataTriDarma](constants.FullPathDB(constants.DB_PENELITIAN))
	for i := 0; i < dataTRI.Length; i++ {
		if dataTRI.Data[i].Id == id {
			return true, dataTRI.Data[i]
		}
	}
	return false, types.TriDarma{}
}

func ListTridar() types.DataTriDarma {
	return database.ReadJSON[types.DataTriDarma](constants.FullPathDB(constants.DB_PENELITIAN))
}

// Used for sorting
func UpdateAll(data types.DataTriDarma) {
	database.WriteJSON(constants.FullPathDB(constants.DB_PENELITIAN), data)
}

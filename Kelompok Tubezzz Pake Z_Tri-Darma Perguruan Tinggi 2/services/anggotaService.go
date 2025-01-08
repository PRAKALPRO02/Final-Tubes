package services

import (
	"errors"
	constants "tri_darma/constants"
	database "tri_darma/database"
	types "tri_darma/types"
)

func AddAnggota(dataAngg types.Anggota) (bool, error) {
	var dataUser = database.ReadJSON[types.DataAnggota](constants.FullPathDB(constants.DB_ANGGOTA))

	for i := 0; i < dataUser.Length; i++ {
		if dataUser.Data[i].Nama == dataAngg.Nama {
			return false, errors.New("Duplicate Name")
		}
	}
	found, datafounded := GetTridarById(dataAngg.IdTridarma)
	if found {
		dataAngg.Id = dataUser.LastId + 1
		dataUser.LastId = dataAngg.Id
		dataUser.Data[dataUser.Length] = dataAngg
		dataUser.Length += 1
		database.WriteJSON(
			constants.FullPathDB(constants.DB_ANGGOTA),
			dataUser,
		)
		datafounded.CountAnggota++
		ChangeTriDarmaById(datafounded.Id, datafounded)
		return true, nil
	}
	return false, errors.New("ID tridar not found")
}

func RemoveAnggotaById(id int) (bool, error) {
	var dataUser = database.ReadJSON[types.DataAnggota](constants.FullPathDB(constants.DB_ANGGOTA))
	var idxToDelete int = -1
	for i := 0; i < dataUser.Length; i++ {
		if dataUser.Data[i].Id == id {
			idxToDelete = i
		}
	}

	if idxToDelete != -1 {
		_, datafounded := GetTridarById(dataUser.Data[idxToDelete].IdTridarma)

		for i := idxToDelete; i < dataUser.Length-1; i++ {
			dataUser.Data[i] = dataUser.Data[i+1]
		}
		datafounded.CountAnggota--
		dataUser.Data[dataUser.Length-1] = types.Anggota{}
		dataUser.Length--
		database.WriteJSON(
			constants.FullPathDB(constants.DB_ANGGOTA),
			dataUser,
		)
		ChangeTriDarmaById(datafounded.Id, datafounded)
		return true, nil
	}
	return false, errors.New("UserId Not Found")
}

func ChangeAnggotaById(id int, datausch types.Anggota) (bool, error) {
	var dataUser = database.ReadJSON[types.DataAnggota](constants.FullPathDB(constants.DB_ANGGOTA))
	var idxToChange int = -1
	for i := 0; i < dataUser.Length; i++ {
		if dataUser.Data[i].Id == id {
			idxToChange = i
		}
	}
	if idxToChange != -1 {
		dataUser.Data[idxToChange] = datausch
		database.WriteJSON(
			constants.FullPathDB(constants.DB_ANGGOTA),
			dataUser,
		)
		return true, nil
	}
	return false, errors.New("CategoryId Not Found")
}

func ListAnggota() types.DataAnggota {
	return database.ReadJSON[types.DataAnggota](constants.FullPathDB(constants.DB_ANGGOTA))
}

func GetUserById(id int) (bool, types.Anggota) {
	var dataUser = database.ReadJSON[types.DataAnggota](constants.FullPathDB(constants.DB_ANGGOTA))
	for i := 0; i < dataUser.Length; i++ {
		if dataUser.Data[i].Id == id {
			return true, dataUser.Data[i]
		}
	}
	return false, types.Anggota{}
}

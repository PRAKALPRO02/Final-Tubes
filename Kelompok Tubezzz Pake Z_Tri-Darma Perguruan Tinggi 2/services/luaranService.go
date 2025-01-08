package services

import (
	"errors"
	constants "tri_darma/constants"
	database "tri_darma/database"
	types "tri_darma/types"
)

func AddLuaran(luaranData types.Luaran) (bool, error) {

	var dataLuaran = database.ReadJSON[types.DataLuaran](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	found, datafounded := GetTridarById(luaranData.IdTridarma)
	if found {
		luaranData.Id = dataLuaran.LastId + 1
		dataLuaran.LastId = luaranData.Id
		dataLuaran.Data[dataLuaran.Length] = luaranData
		dataLuaran.Length += 1
		database.WriteJSON(
			constants.FullPathDB(constants.DB_PRODUCTS_FILE),
			dataLuaran,
		)
		datafounded.CountLuaran++
		ChangeTriDarmaById(datafounded.Id, datafounded)
		return true, nil
	}
	return false, errors.New("ID tridar not found")

}

func RemoveProductById(id int) (bool, error) {
	var dataLuaran = database.ReadJSON[types.DataLuaran](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	var idxToDelete int = -1
	for i := 0; i < dataLuaran.Length; i++ {
		if dataLuaran.Data[i].Id == id {
			idxToDelete = i
		}
	}
	if idxToDelete != -1 {
		_, datafounded := GetTridarById(dataLuaran.Data[idxToDelete].IdTridarma)

		for i := idxToDelete; i < dataLuaran.Length-1; i++ {
			dataLuaran.Data[i] = dataLuaran.Data[i+1]
		}
		datafounded.CountLuaran--
		dataLuaran.Data[dataLuaran.Length-1] = types.Luaran{}
		dataLuaran.Length--
		database.WriteJSON(
			constants.FullPathDB(constants.DB_PRODUCTS_FILE),
			dataLuaran,
		)
		ChangeTriDarmaById(datafounded.Id, datafounded)
		return true, nil
	}
	return false, errors.New("ProductId Not Found")
}

func ChangeLuaranById(id int, datauluarch types.Luaran) (bool, error) {
	var dataLuaran = database.ReadJSON[types.DataLuaran](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	var idxToChange int = -1
	for i := 0; i < dataLuaran.Length; i++ {
		if dataLuaran.Data[i].Id == id {
			idxToChange = i
		}
	}
	if idxToChange != -1 {
		dataLuaran.Data[idxToChange] = datauluarch
		database.WriteJSON(
			constants.FullPathDB(constants.DB_PRODUCTS_FILE),
			dataLuaran,
		)
		return true, nil
	}
	return false, errors.New("LuaranId Not Found")
}

func ListLuaran() types.DataLuaran {
	return database.ReadJSON[types.DataLuaran](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
}

func GetLuaranById(id int) (bool, types.Luaran) {
	var dataLuaran = database.ReadJSON[types.DataLuaran](constants.FullPathDB(constants.DB_PRODUCTS_FILE))
	for i := 0; i < dataLuaran.Length; i++ {
		if dataLuaran.Data[i].Id == id {
			return true, dataLuaran.Data[i]
		}
	}
	return false, types.Luaran{}
}

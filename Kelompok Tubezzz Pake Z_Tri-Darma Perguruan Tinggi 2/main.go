package main

import (
	database "tri_darma/database"
	types "tri_darma/types"
	view "tri_darma/view"
)

/*
	TODO: Tidy up
*/

func main() {
	database.InitDb()
	var setDataTriDarma types.TriDarma
	view.MainMenu(&setDataTriDarma)
	if setDataTriDarma.Id != 0 {
		view.PenelitianMenu(&setDataTriDarma)
		main()
	}
}

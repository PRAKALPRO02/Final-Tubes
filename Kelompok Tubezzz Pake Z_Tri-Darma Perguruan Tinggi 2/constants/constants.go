package constants

import "path/filepath"

const (
	PATH_DATA        = "data"
	PATH_DB          = "db"
	DB_PENELITIAN    = "penelitian.json"
	DB_ANGGOTA       = "anggota.json"
	DB_PRODUCTS_FILE = "keluaran.json"
	DB_DANA_FILE     = "sumberdana.json"
)

func FullPathDB(file string) string {
	return filepath.Join(PATH_DATA, PATH_DB, file)
}

package view

import (
	"fmt"
	"os"
	"strconv"
	services "tri_darma/services"
	types "tri_darma/types"

	tablewriter "github.com/olekukonko/tablewriter"
)

func choosePeran() string {
	var choice int
	Clrscr()
	fmt.Println(border("-", "Pilih Jabatan", 50))
	fmt.Println("1. Anggota (Mahasiswa)")
	fmt.Println("2. Anggota (Dosen)")
	fmt.Println(border("-", "", 50))
	for {
		fmt.Print("Pilih : ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			return "Anggota (Mahasiswa)"
		case 2:
			return "Anggota (Dosen)"
		default:
			fmt.Println("inputan salah! Ulangi")
			delay(2)
		}
	}
}

func InputAnggota(penelitianManage *types.TriDarma) types.Anggota {
	Clrscr()
	fmt.Println(border("-", "Data Anggota", 50))
	var anggotaTemp types.Anggota
	fmt.Print("Nama : ")
	HandleLongInput(&anggotaTemp.Nama)
	anggotaTemp.Role = choosePeran()
	anggotaTemp.IdTridarma = penelitianManage.Id
	return anggotaTemp
}

func InputPendanaan(penelitianManage *types.TriDarma) types.Dana {
	Clrscr()
	var dataPay types.Dana
	var choice int
	fmt.Println(border("-", "Data Pendanaan", 50))
	fmt.Println("1. Internal")
	fmt.Println("2. External")
	fmt.Println(border("-", "", 50))
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)
	for choice > 2 || choice < 1 {
		fmt.Println("Pilihan tidak valid!, ulang!")
		fmt.Print("Pilih : ")
		fmt.Scanln(&choice)
	}
	if choice == 1 {
		dataPay.Sumber = "Internal"
	} else {
		dataPay.Sumber = "External"
	}
	fmt.Print("Keterangan : ")
	HandleLongInput(&dataPay.Keterangan)
	fmt.Print("Nominal : ")
	fmt.Scanln(&dataPay.Nominal)
	dataPay.IdTridarma = penelitianManage.Id
	return dataPay
}

func InputLuaran(penelitianManage *types.TriDarma) types.Luaran {
	Clrscr()
	var tempLuaran types.Luaran
	var choice int
	fmt.Println(border("-", "Tambah Luaran", 50))
	fmt.Println("1. Publikasi")
	fmt.Println("2. Produk")
	if penelitianManage.Tipe == "Abdimas" {
		fmt.Println("3. Seminar")
		fmt.Println("4. Pelatihan")
	}
	fmt.Println(border("-", "", 50))
	fmt.Print("Pilih : ")
	fmt.Scan(&choice)
	for (choice > 2 && penelitianManage.Tipe == "Penelitian") || (choice > 4 && penelitianManage.Tipe == "Abdimas") {
		fmt.Println("Pilihan tidak valid!, ulang!")
		fmt.Print("Pilih : ")
		fmt.Scan(&choice)
	}
	switch choice {
	case 1:
		tempLuaran.BentukLuaran = "Publikasi"
	case 2:
		tempLuaran.BentukLuaran = "Produk"
	case 3:
		tempLuaran.BentukLuaran = "Seminar"
	case 4:
		tempLuaran.BentukLuaran = "Pelatihan"
	}
	fmt.Print("Tanggal Pelaksanaan (dd/mm/yyyy) : ")
	HandleLongInput(&tempLuaran.Pelaksanaan)
	tempLuaran.IdTridarma = penelitianManage.Id
	return tempLuaran
}

func MenuTemplate(choice *int, label string) {
	Clrscr()
	fmt.Println(border("-", "Manage "+label, 50))
	fmt.Println("1. Tambah", label)
	fmt.Println("2. Lihat", label)
	fmt.Println("3. Ubah", label)
	fmt.Println("4. Hapus", label)
	fmt.Println("0. Exit")
	fmt.Println(border("-", "", 50))
	fmt.Print("Pilih : ")

	// Parameter refuses to be used in scan()
	var temp int
	fmt.Scan(&temp)
	*choice = temp
}

func AnggotaMenu(penelitianManage *types.TriDarma) {
	var pilih int
	loop := true
	for loop {
		_, user := services.GetTridarById(penelitianManage.Id)
		*penelitianManage = user
		MenuTemplate(&pilih, "Anggota")
		switch pilih {
		case 1:
			if penelitianManage.CountAnggota < 4 {
				services.AddAnggota(InputAnggota(penelitianManage))
				fmt.Println("Data berhasil disimpan, akan dialihkan dalam 2 dtk")
			} else {
				fmt.Println("Error!, Max. 4 Orang")
			}
			delay(2)
		case 2:
			Clrscr()
			fmt.Println(border("-", "Data Anggota", 50))
			var dataAnggota = services.ListAnggota()
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"No", "Nama", "Role"})
			count := 0
			for i := 0; i < dataAnggota.Length; i++ {
				if dataAnggota.Data[i].IdTridarma == penelitianManage.Id {
					table.Append([]string{strconv.Itoa(count + 1), dataAnggota.Data[i].Nama, dataAnggota.Data[i].Role})
					count++
				}
			}
			table.Render()

			fmt.Scanln()
			fmt.Println("Klik [enter] untuk lanjut")
			fmt.Scanln()
		case 3:
			var sementara [99]int
			var countTmp int
			var dataAnggota = services.ListAnggota()
			for i := 0; i < dataAnggota.Length; i++ {
				if dataAnggota.Data[i].IdTridarma == penelitianManage.Id {
					sementara[countTmp] = i
					countTmp++
				}
			}
			if countTmp == 0 {
				fmt.Println("Tidak Ada Anggota disini!")
			} else {
				Clrscr()
				fmt.Println(border("-", "Ubah Anggota", 50))
				for i := 0; i < countTmp; i++ {
					fmt.Println(i+1, ". ", dataAnggota.Data[sementara[i]].Nama)
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scan(&pilih)
				for pilih > countTmp || pilih < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scan(&pilih)
				}

				services.ChangeAnggotaById(dataAnggota.Data[sementara[pilih-1]].Id, InputAnggota(penelitianManage))
				fmt.Println("Data berhasil diubah, akan dialihkan dalam 2 dtk")
			}
			delay(2)
		case 4:
			var sementara [99]int
			var countTmp int
			var dataAnggota = services.ListAnggota()
			for i := 0; i < dataAnggota.Length; i++ {
				if dataAnggota.Data[i].IdTridarma == penelitianManage.Id {
					sementara[countTmp] = i
					countTmp++
				}
			}
			if countTmp == 0 {
				fmt.Println("Tidak Ada Anggota disini!")
			} else {
				Clrscr()
				fmt.Println(border("-", "Hapus Anggota", 50))
				for i := 0; i < countTmp; i++ {
					fmt.Println(i+1, ". ", dataAnggota.Data[sementara[i]].Nama)
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scan(&pilih)
				for pilih > countTmp || pilih < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scan(&pilih)
				}

				services.RemoveAnggotaById(dataAnggota.Data[sementara[pilih-1]].Id)

				fmt.Println("Data Berhasil dihapus!, dialihkan dalam 2 detik")
			}
			delay(2)
		case 0:
			loop = false
		default:
			fmt.Println("Inputan salah. Mengulang dalam 2 detik...")
			delay(2)
		}
	}
}

func PendanaanMenu(penelitianManage *types.TriDarma) {
	var choice int
	loop := true
	for loop {
		MenuTemplate(&choice, "Pendanaan")
		switch choice {
		case 1:
			services.AddDana(InputPendanaan(penelitianManage))
			fmt.Println("Data berhasil disimpan, akan dialihkan dalam 2 dtk")
			delay(2)
		case 2:
			Clrscr()
			fmt.Println(border("-", "Data Pendanaan", 50))
			var dataDana = services.ListDana()
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"No", "Sumber Dana", "Asal Dana", "Nomimal"})
			count := 0
			for i := 0; i < dataDana.Length; i++ {
				if dataDana.Data[i].IdTridarma == penelitianManage.Id {
					table.Append([]string{strconv.Itoa(count + 1), dataDana.Data[i].Sumber, dataDana.Data[i].Keterangan, strconv.Itoa(dataDana.Data[i].Nominal)})
					count++
				}
			}
			table.Render()

			fmt.Scanln()
			fmt.Println("Klik [enter] untuk lanjut")
			fmt.Scanln()
		case 3:
			var sementara [99]int
			var countTmp int
			var dataDana = services.ListDana()
			for i := 0; i < dataDana.Length; i++ {
				if dataDana.Data[i].IdTridarma == penelitianManage.Id {
					sementara[countTmp] = i
					countTmp++
				}
			}
			if countTmp == 0 {
				fmt.Println("Tidak Ada Dana, Kembali ke menu!")
			} else {
				Clrscr()
				fmt.Println(border("-", "Hapus Pendanaan", 50))
				for i := 0; i < countTmp; i++ {
					fmt.Println(i+1, ". ", dataDana.Data[sementara[i]].Sumber, "(", dataDana.Data[sementara[i]].Keterangan, ")")
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scan(&choice)
				for choice > countTmp || choice < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scan(&choice)
				}
				services.ChangeDanaById(dataDana.Data[sementara[choice-1]].Id, InputPendanaan(penelitianManage))
				fmt.Println("Data Berhasil diubah!, dialihkan dalam 2 detik")
			}
			delay(2)
		case 4:
			var sementara [99]int
			var countTmp int
			var dataDana = services.ListDana()
			for i := 0; i < dataDana.Length; i++ {
				if dataDana.Data[i].IdTridarma == penelitianManage.Id {
					sementara[countTmp] = i
					countTmp++
				}
			}
			if countTmp == 0 {
				fmt.Println("Tidak Ada Dana, Kembali ke menu!")
			} else {
				Clrscr()
				fmt.Println(border("-", "Hapus Pendanaan", 50))
				for i := 0; i < countTmp; i++ {
					fmt.Println(i+1, ". ", dataDana.Data[sementara[i]].Sumber, "(", dataDana.Data[sementara[i]].Keterangan, ")")
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scan(&choice)
				for choice > countTmp || choice < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scan(&choice)
				}
				services.RemoveDanaById(dataDana.Data[sementara[choice-1]].Id)
				fmt.Println("Data Berhasil dihapus!, dialihkan dalam 2 detik")
			}
			delay(2)
		case 0:
			loop = false
		default:
			fmt.Println("Inputan salah. Mengulang dalam 2 detik...")
			delay(2)
		}
	}
}

func LuaranMenu(penelitianManage *types.TriDarma) {
	var choice int
	loop := true
	for loop {
		MenuTemplate(&choice, "Luaran")
		switch choice {
		case 1:
			services.AddLuaran(InputLuaran(penelitianManage))
			fmt.Println("Data berhasil disimpan, akan dialihkan dalam 2 dtk")
			delay(2)
		case 2:
			Clrscr()
			fmt.Println(border("-", "Data Pendanaan", 50))
			var dataLuaran = services.ListLuaran()
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"No", "Bentuk", "Pelaksanaan (dd/mm/yyyy)"})
			count := 0
			for i := 0; i < dataLuaran.Length; i++ {
				if dataLuaran.Data[i].IdTridarma == penelitianManage.Id {
					table.Append([]string{strconv.Itoa(count + 1), dataLuaran.Data[i].BentukLuaran, dataLuaran.Data[i].Pelaksanaan})
					count++
				}
			}

			fmt.Scanln()
			fmt.Println("Klik [enter] untuk lanjut")
			fmt.Scanln()
		case 3:
			var sementara [99]int
			var countTmp int
			var dataLuaran = services.ListLuaran()
			for i := 0; i < dataLuaran.Length; i++ {
				if dataLuaran.Data[i].IdTridarma == penelitianManage.Id {
					sementara[countTmp] = i
					countTmp++
				}
			}
			if countTmp == 0 {
				fmt.Println("Tidak Ada Luaran, Kembali ke menu!")
			} else {
				Clrscr()
				fmt.Println(border("-", "Hapus Pendanaan", 50))
				for i := 0; i < countTmp; i++ {
					fmt.Println(i+1, ". ", dataLuaran.Data[sementara[i]].BentukLuaran, "(", dataLuaran.Data[sementara[i]].Pelaksanaan, ")")
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scan(&choice)
				for choice > countTmp || choice < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scan(&choice)
				}
				services.RemoveProductById(dataLuaran.Data[sementara[choice-1]].Id)
				fmt.Println("Data Berhasil diubah!, dialihkan dalam 2 detik")
			}
			delay(2)
		case 4:
			var sementara [99]int
			var countTmp int
			var dataLuaran = services.ListLuaran()
			for i := 0; i < dataLuaran.Length; i++ {
				if dataLuaran.Data[i].IdTridarma == penelitianManage.Id {
					sementara[countTmp] = i
					countTmp++
				}
			}
			if countTmp == 0 {
				fmt.Println("Tidak Ada Luaran, Kembali ke menu!")
			} else {
				Clrscr()
				fmt.Println(border("-", "Hapus Pendanaan", 50))
				for i := 0; i < countTmp; i++ {
					fmt.Println(i+1, ". ", dataLuaran.Data[sementara[i]].BentukLuaran, "(", dataLuaran.Data[sementara[i]].Pelaksanaan, ")")
				}
				fmt.Println(border("-", "", 50))
				fmt.Print("Pilih : ")
				fmt.Scan(&choice)
				for choice > countTmp || choice < 1 {
					fmt.Println("Pilihan tidak valid!, ulang!")
					fmt.Print("Pilih : ")
					fmt.Scan(&choice)
				}
				services.RemoveProductById(dataLuaran.Data[sementara[choice-1]].Id)
				fmt.Println("Data Berhasil dihapus!, dialihkan dalam 2 detik")
			}
			delay(2)
		case 0:
			loop = false
		default:
			fmt.Println("Inputan salah. Mengulang dalam 2 detik...")
			delay(2)
		}
	}
}

func PenelitianMenu(penelitianManage *types.TriDarma) {
	var choice int
	var anggota types.Anggota
	for {
		Clrscr()
		if penelitianManage.CountAnggota == 0 {
			fmt.Println(border("-", "Data Ketua", 50))
			fmt.Print("Nama Ketua : ")
			HandleLongInput(&anggota.Nama)
			anggota.IdTridarma = penelitianManage.Id
			anggota.Role = "Ketua"
			services.AddAnggota(anggota)
			_, user := services.GetTridarById(penelitianManage.Id)
			*penelitianManage = user
		} else {
			Clrscr()
			_, user := services.GetTridarById(penelitianManage.Id)
			*penelitianManage = user
			fmt.Println(border("-", "Detail "+penelitianManage.Tipe, 50))
			formatPrint("Judul", penelitianManage.Nama)
			formatPrint("Prodi", penelitianManage.Prodi)
			formatPrint("Tahun", penelitianManage.Tahun)
			formatPrint("Banyak Anggota", penelitianManage.CountAnggota)
			formatPrint("Banyak Luaran", penelitianManage.CountLuaran)
			formatPrint("Total Pendanaan", penelitianManage.SumDana)
			fmt.Println(border("-", "", 50))
			fmt.Println("1. Manage Anggota")
			fmt.Println("2. Manage Pendanaan")
			fmt.Println("3. Manage Luaran")
			fmt.Println("4. Ubah Detail " + penelitianManage.Tipe)
			fmt.Println("5. [DANGER] Hapus " + penelitianManage.Tipe + " ini")
			fmt.Println("0. Exit")
			fmt.Println(border("-", "", 50))
			fmt.Print("Pilih : ")
			fmt.Scan(&choice)
			switch choice {
			case 1:
				AnggotaMenu(penelitianManage)
			case 2:
				PendanaanMenu(penelitianManage)
			case 3:
				LuaranMenu(penelitianManage)
			case 4:
				inputRegTriDarma(penelitianManage, penelitianManage.Tipe)
				services.ChangeTriDarmaById(penelitianManage.Id, *penelitianManage)
				fmt.Println("Data berhasil diubah. Melanjutkan dalam 2 detik...")
				delay(2)
			case 5:
				Clrscr()
				fmt.Println(border("-", "Konfirmasi Penghapusan", 50))
				fmt.Println("Apakah anda yakin untuk menghapus " + penelitianManage.Tipe + " " + penelitianManage.Nama + "?")
				loop := true
				for loop {
					fmt.Print("[1. Ya / 2. Tidak] : ")
					fmt.Scan(&choice)
					if choice == 1 {
						services.RemoveTriDarmaById(penelitianManage.Id)
						fmt.Println("Data berhasil dihapus. Kembali ke menu awal dalam 2 detik...")
						delay(2)
						loop = false
						return
					} else if choice == 2 {
						loop = false
					} else {
						fmt.Println("Inputan Salah! ulangi")
					}
				}
			case 0:
				return
			default:
				fmt.Println("Inputan salah. Mengulang dalam 2 detik...")
				delay(2)
			}
		}

	}
}

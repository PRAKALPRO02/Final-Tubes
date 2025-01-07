package main

import (
	"fmt"
	"sort"
)

type Matakuliah struct {
	Nama  string
	SKS   int
	UTS   float64
	UAS   float64
	Quiz  float64
	Total float64
	Grade string
}

type Mahasiswa struct {
	NIM        string
	Nama       string
	Matakuliah map[string]Matakuliah
	TotalSKS   int
	TotalNilai float64
	IPK        float64
}

var mahasiswaList = make(map[string]Mahasiswa)

func hitungTotalDanGrade(mk *Matakuliah) {
	mk.Total = (mk.UTS * 0.3) + (mk.UAS * 0.4) + (mk.Quiz * 0.3)
	if mk.Total >= 80 {
		mk.Grade = "A"
	} else if mk.Total >= 70 {
		mk.Grade = "B"
	} else if mk.Total >= 60 {
		mk.Grade = "C"
	} else if mk.Total >= 50 {
		mk.Grade = "D"
	} else {
		mk.Grade = "E"
	}
}

func tambahMahasiswa() {
	var nim, nama string
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&nama)
	mahasiswaList[nim] = Mahasiswa{NIM: nim, Nama: nama, Matakuliah: make(map[string]Matakuliah)}
	fmt.Println("Mahasiswa berhasil ditambahkan!")
}

func editMahasiswa() {
	var nim, nama string
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	mhs, exists := mahasiswaList[nim]
	if !exists {
		fmt.Println("Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nama Baru: ")
	fmt.Scan(&nama)
	mhs.Nama = nama
	mahasiswaList[nim] = mhs
	fmt.Println("Mahasiswa berhasil diperbarui!")
}

func hapusMahasiswa() {
	var nim string
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	delete(mahasiswaList, nim)
	fmt.Println("Mahasiswa berhasil dihapus!")
}

func tambahMatakuliah() {
	var nim, namaMK string
	var sks int
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	mhs, exists := mahasiswaList[nim]
	if !exists {
		fmt.Println("Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nama Matakuliah: ")
	fmt.Scan(&namaMK)
	fmt.Print("Masukkan SKS: ")
	fmt.Scan(&sks)
	mk := Matakuliah{Nama: namaMK, SKS: sks}
	mhs.Matakuliah[namaMK] = mk
	mahasiswaList[nim] = mhs
	fmt.Println("Matakuliah berhasil ditambahkan!")
}

func inputNilai() {
	var nim, namaMK string
	var uts, uas, quiz float64
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	mhs, exists := mahasiswaList[nim]
	if !exists {
		fmt.Println("Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nama Matakuliah: ")
	fmt.Scan(&namaMK)
	mk, exists := mhs.Matakuliah[namaMK]
	if !exists {
		fmt.Println("Matakuliah tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nilai UTS: ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan Nilai UAS: ")
	fmt.Scan(&uas)
	fmt.Print("Masukkan Nilai Quiz: ")
	fmt.Scan(&quiz)
	mk.UTS = uts
	mk.UAS = uas
	mk.Quiz = quiz
	hitungTotalDanGrade(&mk)
	mhs.Matakuliah[namaMK] = mk
	mahasiswaList[nim] = mhs

	// Hitung IPK setelah input nilai
	hitungIPK(&mhs)

	fmt.Println("Nilai berhasil dimasukkan!")
}

func editNilai() {
	var nim, namaMK string
	var uts, uas, quiz float64
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	mhs, exists := mahasiswaList[nim]
	if !exists {
		fmt.Println("Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nama Matakuliah: ")
	fmt.Scan(&namaMK)
	mk, exists := mhs.Matakuliah[namaMK]
	if !exists {
		fmt.Println("Matakuliah tidak ditemukan!")
		return
	}
	fmt.Println("Data Matakuliah ditemukan:")
	fmt.Println("UTS:", mk.UTS, "UAS:", mk.UAS, "Quiz:", mk.Quiz)

	// Input Nilai Baru
	fmt.Print("Masukkan Nilai UTS baru: ")
	fmt.Scan(&uts)
	fmt.Print("Masukkan Nilai UAS baru: ")
	fmt.Scan(&uas)
	fmt.Print("Masukkan Nilai Quiz baru: ")
	fmt.Scan(&quiz)

	// Update Nilai dan Hitung Ulang Total dan Grade
	mk.UTS = uts
	mk.UAS = uas
	mk.Quiz = quiz
	hitungTotalDanGrade(&mk) // Recalculate the Total and Grade after editing

	mhs.Matakuliah[namaMK] = mk
	mahasiswaList[nim] = mhs

	// Update IPK setelah nilai diubah
	hitungIPK(&mhs)

	fmt.Println("Nilai berhasil diperbarui!")
}

func hapusNilai() {
	var nim, namaMK string
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	mhs, exists := mahasiswaList[nim]
	if !exists {
		fmt.Println("Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Print("Masukkan Nama Matakuliah: ")
	fmt.Scan(&namaMK)
	_, exists = mhs.Matakuliah[namaMK]
	if !exists {
		fmt.Println("Matakuliah tidak ditemukan!")
		return
	}

	// Remove Nilai Matakuliah
	delete(mhs.Matakuliah, namaMK)
	mahasiswaList[nim] = mhs

	// Update IPK setelah nilai dihapus
	hitungIPK(&mhs)

	fmt.Println("Nilai untuk matakuliah", namaMK, "telah dihapus!")
}

func hitungIPK(mhs *Mahasiswa) {
	var totalNilai float64
	var totalSKS int
	for _, mk := range mhs.Matakuliah {
		totalNilai += mk.Total * float64(mk.SKS)
		totalSKS += mk.SKS
	}
	if totalSKS > 0 {
		mhs.TotalNilai = totalNilai
		mhs.TotalSKS = totalSKS
		mhs.IPK = totalNilai / float64(totalSKS)
	}
}

func tampilkanMahasiswaByMatakuliah() {
	var namaMK string
	fmt.Print("Masukkan Nama Matakuliah: ")
	fmt.Scan(&namaMK)
	fmt.Println("Mahasiswa yang mengambil", namaMK, ":")
	for _, mhs := range mahasiswaList {
		if _, exists := mhs.Matakuliah[namaMK]; exists {
			fmt.Println(mhs.NIM, mhs.Nama)
		}
	}
}

func tampilkanMatakuliahByMahasiswa() {
	var nim string
	fmt.Print("Masukkan NIM: ")
	fmt.Scan(&nim)
	mhs, exists := mahasiswaList[nim]
	if !exists {
		fmt.Println("Mahasiswa tidak ditemukan!")
		return
	}
	fmt.Println("Matakuliah yang diambil oleh", mhs.Nama, ":")
	for _, mk := range mhs.Matakuliah {
		fmt.Println(mk.Nama, "SKS:", mk.SKS, "Total:", mk.Total, "Grade:", mk.Grade)
	}
}

func tampilkanMahasiswaTerurut() {
	var sortedMahasiswa []Mahasiswa
	for _, mhs := range mahasiswaList {
		// Hitung IPK mahasiswa
		var totalNilai float64
		var totalSKS int
		for _, mk := range mhs.Matakuliah {
			totalNilai += mk.Total * float64(mk.SKS)
			totalSKS += mk.SKS
		}
		if totalSKS > 0 {
			mhs.IPK = totalNilai / float64(totalSKS)
		}

		// Tambahkan mahasiswa yang sudah dihitung IPK-nya ke dalam list
		sortedMahasiswa = append(sortedMahasiswa, mhs)
	}

	// Urutkan mahasiswa berdasarkan IPK
	sort.Slice(sortedMahasiswa, func(i, j int) bool {
		return sortedMahasiswa[i].IPK > sortedMahasiswa[j].IPK
	})

	// Tampilkan mahasiswa terurut berdasarkan IPK
	fmt.Println("Mahasiswa terurut berdasarkan IPK:")
	for _, mhs := range sortedMahasiswa {
		fmt.Println(mhs.NIM, mhs.Nama, "IPK:", mhs.IPK)
	}
}

func main() {
	for {
		fmt.Println("\nMenu: ")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Edit Mahasiswa")
		fmt.Println("3. Hapus Mahasiswa")
		fmt.Println("4. Tambah Matakuliah")
		fmt.Println("5. Input Nilai")
		fmt.Println("6. Edit Nilai")
		fmt.Println("7. Hapus Nilai")
		fmt.Println("8. Tampilkan Mahasiswa berdasarkan Matakuliah")
		fmt.Println("9. Tampilkan Matakuliah berdasarkan Mahasiswa")
		fmt.Println("10. Tampilkan Mahasiswa Terurut")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")
		var pilihan int
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahMahasiswa()
		case 2:
			editMahasiswa()
		case 3:
			hapusMahasiswa()
		case 4:
			tambahMatakuliah()
		case 5:
			inputNilai()
		case 6:
			editNilai()
		case 7:
			hapusNilai()
		case 8:
			tampilkanMahasiswaByMatakuliah()
		case 9:
			tampilkanMatakuliahByMahasiswa()
		case 10:
			tampilkanMahasiswaTerurut()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

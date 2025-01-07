package main

import (
    "fmt"
    "os"
    "os/exec"
    "strings"
)

type PaketMCU struct {
    IDPaket   int
    NamaPaket string
    Harga     float64
}

type PemeriksaanFisik struct {
    TinggiBadan   int
    BeratBadan    int
    TekananDarah  string
    DenyutNadi    int
    SuhuTubuh     float64
    Mata          string
    Hemoglobin    string
    Trombosit     string
    GulaDarah     string
    Kolestrol     string
    RontgenDada   string
    FungsiHatiSGOT string
    FungsiHatiSGPT string
    EKG           string
}

type Pasien struct {
    Nama          string
    JenisKelamin  string
    Umur          int
    Alamat        string
    PaketMCU      PaketMCU
    TanggalMasuk string
    Fisik     PemeriksaanFisik
}

var (
	dataPasien    []Pasien
	daftarPaketMCU = []PaketMCU{
		{IDPaket: 1, NamaPaket: "Paket A", Harga: 250000},
		{IDPaket: 2, NamaPaket: "Paket B", Harga: 300000},
		{IDPaket: 3, NamaPaket: "Paket C", Harga: 350000},
	}
)

func clearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func menu() {
    for {
        clearScreen()
        fmt.Println("\n=== Aplikasi Medical Check-Up ===")
        fmt.Println("1. Paket MCU")
        fmt.Println("2. Tambah Data Pasien MCU")
        fmt.Println("3. Data Pasien MCU")
        fmt.Println("4. Masukkan MCU Pasien") 
        fmt.Println("5. Hasil MCU Pasien")    
        fmt.Println("6. Laporan Pemasukan")  
        fmt.Println("0. Keluar")
        fmt.Print("Pilih menu: ")

        var pilihan int
        fmt.Scanln(&pilihan)

        switch pilihan {
        case 1:
            listPaketMCU()
        case 2:
            tambahPasien()
        case 3:
            DataPasien() 
        case 4:
            masukkanMCUPasien() 
        case 5:
            tampilkanHasilMCUPasien() 
        case 6:
            laporanPemasukan()
        case 0:
            fmt.Println("Terima kasih! Sampai jumpa.")
            return
        default:
            fmt.Println("Pilihan tidak valid, silakan coba lagi.")
        }
    }
}

func listPaketMCU() {
    clearScreen()
    fmt.Println("\nDaftar Paket MCU:")
    for _, paket := range daftarPaketMCU {
        fmt.Printf("%d. %s (Harga: %.2f)\n", paket.IDPaket, paket.NamaPaket, paket.Harga)
        switch paket.IDPaket {
        case 1:
            fmt.Println("   - Tinggi Badan")
            fmt.Println("   - Berat Badan")
            fmt.Println("   - Tekanan Darah")
            fmt.Println("   - Denyut Nadi")
            fmt.Println("   - Suhu Tubuh")
            fmt.Println("   - Mata")
            fmt.Println("   - Hemoglobin")
            fmt.Println("   - Trombosit")
        case 2:
            fmt.Println("   - Tinggi Badan")
            fmt.Println("   - Berat Badan")
            fmt.Println("   - Tekanan Darah")
            fmt.Println("   - Denyut Nadi")
            fmt.Println("   - Suhu Tubuh")
            fmt.Println("   - Mata")
            fmt.Println("   - Hemoglobin")
            fmt.Println("   - Trombosit")
            fmt.Println("   - Gula Darah")
            fmt.Println("   - Kolestrol")
            fmt.Println("   - Rontgen Dada")
        case 3:
            fmt.Println("   - Tinggi Badan")
            fmt.Println("   - Berat Badan")
            fmt.Println("   - Tekanan Darah")
            fmt.Println("   - Denyut Nadi")
            fmt.Println("   - Suhu Tubuh")
            fmt.Println("   - Mata")
            fmt.Println("   - Hemoglobin")
            fmt.Println("   - Trombosit")
            fmt.Println("   - Gula Darah")
            fmt.Println("   - Kolestrol")
            fmt.Println("   - Rontgen Dada")
            fmt.Println("   - Fungsi Hati SGOT")
            fmt.Println("   - Fungsi Hati SGPT")
            fmt.Println("   - EKG")
        }
    }
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
    clearScreen()
}

func tambahPasien() {
    clearScreen()
    var pasien Pasien
    fmt.Println("Masukkan data pasien:")
    fmt.Print("Nama          : ")
    fmt.Scanln(&pasien.Nama)
    fmt.Print("Jenis Kelamin : ")
    fmt.Scanln(&pasien.JenisKelamin)
    fmt.Print("Umur          : ")
    fmt.Scanln(&pasien.Umur)
    fmt.Print("Alamat        : ")
    fmt.Scanln(&pasien.Alamat)
    fmt.Print("(DD-MM-YYYY)  : ")
    fmt.Scanln(&pasien.TanggalMasuk)

    // Memilih Paket MCU
    listPaketMCU()
    var pilihanPaket int
    fmt.Print("Pilih ID Paket MCU : ")
    fmt.Scanln(&pilihanPaket)

    paketDitemukan := false
    for _, paket := range daftarPaketMCU {
        if paket.IDPaket == pilihanPaket {
            pasien.PaketMCU = paket
            paketDitemukan = true
            break
        }
    }

    if !paketDitemukan {
        fmt.Println("ID Paket MCU tidak valid.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        clearScreen()
        return
    }

    dataPasien = append(dataPasien, pasien)
    fmt.Println("Data pasien berhasil ditambahkan.")
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
    clearScreen()
}

func DataPasien() {
    clearScreen()
    if len(dataPasien) == 0 {
        fmt.Println("Tidak ada data pasien saat ini.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    fmt.Println("Data Pasien MCU:")
    for i, pasien := range dataPasien {
        fmt.Printf("%d. Nama         : %s\n", i+1, pasien.Nama)
        fmt.Printf("   Jenis Kelamin: %s\n", pasien.JenisKelamin)
        fmt.Printf("   Umur         : %d\n", pasien.Umur)
        fmt.Printf("   Alamat       : %s\n", pasien.Alamat)
        fmt.Printf("   Paket MCU    : %s (Harga: %.2f)\n", pasien.PaketMCU.NamaPaket, pasien.PaketMCU.Harga)
        fmt.Printf("   Tanggal Masuk: %s\n", pasien.TanggalMasuk)
        fmt.Println("  --------------------")
    }

    fmt.Println("\nPilihan:")
    fmt.Println("1. Cari data pasien")
    fmt.Println("0. Kembali ke menu utama")
    fmt.Print("Pilih: ")
    var pilihan int
    fmt.Scanln(&pilihan)

    if pilihan == 0 {
        return
    } else if pilihan == 1 {
        clearScreen()
        var keyword string
        fmt.Print("Masukkan keyword pencarian (nama/usia/jenis kelamin/paket/tanggal): ")
        fmt.Scanln(&keyword)
        clearScreen()

        var hasil []Pasien
        for _, pasien := range dataPasien {
            if strings.Contains(strings.ToLower(pasien.Nama), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.JenisKelamin), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(fmt.Sprint(pasien.Umur)), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.PaketMCU.NamaPaket), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.TanggalMasuk), strings.ToLower(keyword)) ||
                strings.Contains(strings.ToLower(pasien.Alamat), strings.ToLower(keyword)) {
                hasil = append(hasil, pasien)
            }
        }

        if len(hasil) == 0 {
            fmt.Println("Data pasien tidak ditemukan.")
            fmt.Println("\nTekan Enter untuk melanjutkan...")
            fmt.Scanln()
            return
        }

        fmt.Println("Data pasien ditemukan:")
        for i, pasien := range hasil {
            fmt.Printf("%d. Nama          : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Jenis Kelamin : %s\n", pasien.JenisKelamin)
            fmt.Printf("   Umur          : %d\n", pasien.Umur)
            fmt.Printf("   Alamat        : %s\n", pasien.Alamat)
            fmt.Printf("   Paket MCU     : %s (Harga: %.2f)\n", pasien.PaketMCU.NamaPaket, pasien.PaketMCU.Harga)
            fmt.Printf("   Tanggal Masuk : %s\n", pasien.TanggalMasuk)
            fmt.Println("  --------------------")
        }

        var pilihanPasien int
        fmt.Print("Pilih nomor pasien: ")
        fmt.Scanln(&pilihanPasien)

        if pilihanPasien > 0 && pilihanPasien <= len(hasil) {
            // Tampilkan submenu untuk pasien yang dipilih
            for {
                fmt.Println("\nPilihan:")
                fmt.Println("1. Perbarui data pasien")
                fmt.Println("2. Hapus data pasien")
                fmt.Println("0. Kembali ke menu pencarian")
                fmt.Print("Pilih: ")
                var subPilihan int
                fmt.Scanln(&subPilihan)

                switch subPilihan {
                case 1:
                    updatePasien(hasil[pilihanPasien-1])
                case 2:
                    deletePasien(pilihanPasien - 1)
                    return
                case 0:
                    return
                default:
                    fmt.Println("Pilihan tidak valid.")
                }
            }
        }
    } else {
        fmt.Println("Pilihan tidak valid.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
    }
}

func updatePasien(pasienLama Pasien) {
    clearScreen()
    var pasienBaru Pasien
    fmt.Println("Masukkan data baru untuk pasien:")
    fmt.Print("Nama          : ")
    fmt.Scanln(&pasienBaru.Nama)
    fmt.Print("Jenis Kelamin : ")
    fmt.Scanln(&pasienBaru.JenisKelamin)
    fmt.Print("Umur          : ")
    fmt.Scanln(&pasienBaru.Umur)
    fmt.Print("Alamat        : ")
    fmt.Scanln(&pasienBaru.Alamat)
    fmt.Print("(DD-MM-YYYY)  : ")
    fmt.Scanln(&pasienBaru.TanggalMasuk)

    // Pilih Paket MCU
    listPaketMCU()
    var pilihanPaket int
    fmt.Print("Pilih ID Paket MCU: ")
    fmt.Scanln(&pilihanPaket)

    // Cari paket MCU yang sesuai
    paketDitemukan := false
    for _, paket := range daftarPaketMCU {
        if paket.IDPaket == pilihanPaket {
            pasienBaru.PaketMCU = paket
            paketDitemukan = true
            break
        }
    }

    if !paketDitemukan {
        fmt.Println("ID Paket MCU tidak valid.")
        return
    }

    // Perbarui data pasien di slice
    for i, pasien := range dataPasien {
        if pasien == pasienLama {
            dataPasien[i] = pasienBaru
            fmt.Println("Data pasien berhasil diperbarui.")
            return
        }
    }

    fmt.Println("Data pasien tidak ditemukan.")
}

func deletePasien(index int) {
    clearScreen()
    if index < 0 || index >= len(dataPasien) {
        fmt.Println("Nomor pasien tidak valid.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
         fmt.Scanln()
        return
    } 
    dataPasien = append(dataPasien[:index], dataPasien[index+1:]...)
    fmt.Println("Data pasien berhasil dihapus.")
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
    clearScreen()
}

func masukkanMCUPasien() {
    clearScreen()
    var keyword string
    fmt.Print("Masukkan nama pasien untuk memasukkan data MCU: ")
    fmt.Scanln(&keyword)

    pasienDitemukan := false
    var pasien *Pasien
    for i := range dataPasien {
        if dataPasien[i].Nama == keyword {
            pasien = &dataPasien[i]
            pasienDitemukan = true
            break
        }
    }

    if !pasienDitemukan {
        fmt.Println("Data pasien tidak ditemukan.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        return
    }

    fmt.Println("\nMasukkan data MCU untuk pasien", pasien.Nama)

    switch pasien.PaketMCU.IDPaket {
    case 1:
        fmt.Print("Tinggi Badan (cm)    : ")
        fmt.Scanln(&pasien.Fisik.TinggiBadan)
        fmt.Print("Berat Badan (kg)     : ")
        fmt.Scanln(&pasien.Fisik.BeratBadan)
        fmt.Print("Tekanan Darah (mmHg) : ")
        fmt.Scanln(&pasien.Fisik.TekananDarah)
        fmt.Print("Denyut Nadi (bpm)    : ")
        fmt.Scanln(&pasien.Fisik.DenyutNadi)
        fmt.Print("Suhu Tubuh (°C)      : ")
        fmt.Scanln(&pasien.Fisik.SuhuTubuh)
        fmt.Print("Mata                 : ")
        fmt.Scanln(&pasien.Fisik.Mata)
        fmt.Print("Hemoglobin           : ")
        fmt.Scanln(&pasien.Fisik.Hemoglobin)
        fmt.Print("Trombosit            : ")
        fmt.Scanln(&pasien.Fisik.Trombosit)
    case 2:
        fmt.Print("Tinggi Badan (cm)    : ")
        fmt.Scanln(&pasien.Fisik.TinggiBadan)
        fmt.Print("Berat Badan (kg)     : ")
        fmt.Scanln(&pasien.Fisik.BeratBadan)
        fmt.Print("Tekanan Darah (mmHg) : ")
        fmt.Scanln(&pasien.Fisik.TekananDarah)
        fmt.Print("Denyut Nadi (bpm)    : ")
        fmt.Scanln(&pasien.Fisik.DenyutNadi)
        fmt.Print("Suhu Tubuh (°C)      : ")
        fmt.Scanln(&pasien.Fisik.SuhuTubuh)
        fmt.Print("Mata                 : ")
        fmt.Scanln(&pasien.Fisik.Mata)
        fmt.Print("Hemoglobin           : ")
        fmt.Scanln(&pasien.Fisik.Hemoglobin)
        fmt.Print("Trombosit            : ")
        fmt.Scanln(&pasien.Fisik.Trombosit)
        fmt.Print("Gula Darah           : ")
        fmt.Scanln(&pasien.Fisik.GulaDarah)
        fmt.Print("Kolestrol            : ")
        fmt.Scanln(&pasien.Fisik.Kolestrol)
        fmt.Print("Rontgen Dada         : ")
        fmt.Scanln(&pasien.Fisik.RontgenDada)
    case 3:
        fmt.Print("Tinggi Badan (cm)    : ")
        fmt.Scanln(&pasien.Fisik.TinggiBadan)
        fmt.Print("Berat Badan (kg)     : ")
        fmt.Scanln(&pasien.Fisik.BeratBadan)
        fmt.Print("Tekanan Darah (mmHg) : ")
        fmt.Scanln(&pasien.Fisik.TekananDarah)
        fmt.Print("Denyut Nadi (bpm)    : ")
        fmt.Scanln(&pasien.Fisik.DenyutNadi)
        fmt.Print("Suhu Tubuh (°C)  : ")
        fmt.Scanln(&pasien.Fisik.SuhuTubuh)
        fmt.Print("Mata             : ")
        fmt.Scanln(&pasien.Fisik.Mata)
        fmt.Print("Hemoglobin       : ")
        fmt.Scanln(&pasien.Fisik.Hemoglobin)
        fmt.Print("Trombosit        : ")
        fmt.Scanln(&pasien.Fisik.Trombosit)
        fmt.Print("Gula Darah       : ")
        fmt.Scanln(&pasien.Fisik.GulaDarah)
        fmt.Print("Kolestrol        : ")
        fmt.Scanln(&pasien.Fisik.Kolestrol)
        fmt.Print("Rontgen Dada     : ")
        fmt.Scanln(&pasien.Fisik.RontgenDada)
        fmt.Print("Fungsi Hati SGOT : ")
        fmt.Scanln(&pasien.Fisik.FungsiHatiSGOT)
        fmt.Print("Fungsi Hati SGPT : ")
        fmt.Scanln(&pasien.Fisik.FungsiHatiSGPT)
        fmt.Print("EKG              : ")
        fmt.Scanln(&pasien.Fisik.EKG)
    }

    fmt.Println("Data MCU pasien berhasil ditambahkan.")
    fmt.Println("\nTekan Enter untuk melanjutkan...")
    fmt.Scanln()
}

func tampilkanHasilMCUPasien() {
    fmt.Println("\nData Hasil MCU Pasien:")
    if len(dataPasien) == 0 {
        fmt.Println("Tidak ada data pasien.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        clearScreen()
        return
    }

    for i, pasien := range dataPasien {
        fmt.Printf("%d. Nama         : %s\n", i+1, pasien.Nama)
        switch pasien.PaketMCU.IDPaket {
        case 1:
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
        case 2:
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
            fmt.Printf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah)
            fmt.Printf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol)
            fmt.Printf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada)
        case 3:
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
            fmt.Printf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah)
            fmt.Printf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol)
            fmt.Printf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada)
            fmt.Printf("   F-Hati SGOT  : %s\n", pasien.Fisik.FungsiHatiSGOT)
            fmt.Printf("   F-Hati SGPT  : %s\n", pasien.Fisik.FungsiHatiSGPT)
            fmt.Printf("   EKG          : %s\n", pasien.Fisik.EKG)
        }
        fmt.Println("   --------------------")
    }

    //...............................................................................................................

    var keyword string
    fmt.Print("\nMasukkan keyword untuk mencari hasil MCU pasien: ")
    fmt.Scanln(&keyword)

    var hasilPencarian []Pasien
    for _, pasien := range dataPasien {
        if strings.Contains(strings.ToLower(pasien.Nama), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(pasien.JenisKelamin), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(fmt.Sprint(pasien.Umur)), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(pasien.Alamat), strings.ToLower(keyword)) ||
            strings.Contains(strings.ToLower(pasien.TanggalMasuk), strings.ToLower(keyword)) {
            hasilPencarian = append(hasilPencarian, pasien)
        }
    }

    if len(hasilPencarian) == 0 {
        fmt.Println("Data pasien tidak ditemukan.")
        fmt.Println("\nTekan Enter untuk melanjutkan...")
        fmt.Scanln()
        clearScreen()
        return
    }

    fmt.Println("\nHasil pencarian:")
    for i, pasien := range hasilPencarian {
        if pasien.PaketMCU.IDPaket == 1 {
            fmt.Printf("%d. Nama        : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
        } else if pasien.PaketMCU.IDPaket == 2 {
            fmt.Printf("%d. Nama        : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
            fmt.Printf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah)
            fmt.Printf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol)
            fmt.Printf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada)
        } else if pasien.PaketMCU.IDPaket == 3 {
            fmt.Printf("%d. Nama        : %s\n", i+1, pasien.Nama)
            fmt.Printf("   Tinggi Badan : %d cm\n", pasien.Fisik.TinggiBadan)
            fmt.Printf("   Berat Badan  : %d kg\n", pasien.Fisik.BeratBadan)
            fmt.Printf("   Tekanan Darah: %s mmHg\n", pasien.Fisik.TekananDarah)
            fmt.Printf("   Denyut Nadi  : %d bpm\n", pasien.Fisik.DenyutNadi)
            fmt.Printf("   Suhu Tubuh   : %.2f °C\n", pasien.Fisik.SuhuTubuh)
            fmt.Printf("   Mata         : %s\n", pasien.Fisik.Mata)
            fmt.Printf("   Hemoglobin   : %s\n", pasien.Fisik.Hemoglobin)
            fmt.Printf("   Trombosit    : %s\n", pasien.Fisik.Trombosit)
            fmt.Printf("   Gula Darah   : %s\n", pasien.Fisik.GulaDarah)
            fmt.Printf("   Kolestrol    : %s\n", pasien.Fisik.Kolestrol)
            fmt.Printf("   Rontgen Dada : %s\n", pasien.Fisik.RontgenDada)
            fmt.Printf("   F-Hati SGOT  : %s\n", pasien.Fisik.FungsiHatiSGOT)
            fmt.Printf("   F-Hati SGPT  : %s\n", pasien.Fisik.FungsiHatiSGPT)
            fmt.Printf("   EKG          : %s\n", pasien.Fisik.EKG)
        }
        fmt.Println("   --------------------")
    }
    fmt.Println("\nTekan Enter untuk kembali ke menu utama...")
    fmt.Scanln()
    clearScreen()
}

func laporanPemasukan() {
    clearScreen()
    totalPemasukan := 0.0
    for _, pasien := range dataPasien {
        totalPemasukan += pasien.PaketMCU.Harga
    }
    fmt.Printf("Total Pemasukan: Rp %.2f\n", totalPemasukan)
    fmt.Println("\nTekan Enter untuk kembali ke menu...")
    fmt.Scanln()
}

func main() {
    menu()
}
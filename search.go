package main

import "fmt"

func mainSearch(warga [NMAX]dataWarga, jumlahWarga int, setoran [NMAX]dataSetoran, jumlahSetoran int) {
	var opsi int = -1
	for opsi != 0 {
		opsi = -1
		fmt.Print("\n===== Cari Warga =====\n")
		fmt.Print("\n1. Cari Warga Berdasarkan ID\n")
		fmt.Print("2. Cari Warga Berdasarkan Nama\n")
		fmt.Print("0. Back to Main Menu\n")
		fmt.Print("\nPilih Opsi: ")
		fmt.Scan(&opsi)
		switch opsi {
		case 1:
			cariWargaByID(warga, jumlahWarga, setoran, jumlahSetoran)
		case 2:
			cariWargaByNama(warga, jumlahWarga, setoran, jumlahSetoran)
		case 0:
			return
		default:
			fmt.Print("Input tidak valid.")
		}
	}
}

func hitungTotalSetoran(setoran [NMAX]dataSetoran, jumlahSetoran int, idWarga int) int {
	total := 0
	for i := 0; i < jumlahSetoran; i++ {
		if setoran[i].idWarga == idWarga {
			total += setoran[i].berat
		}
	}
	return total
}

func cariWargaByID(warga [NMAX]dataWarga, jumlahWarga int, setoran [NMAX]dataSetoran, jumlahSetoran int) {
	var targetID int
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&targetID)
	find := searchBinaryID(warga, jumlahWarga, targetID)
	if find == -1 {
		fmt.Print("Data tidak ditemukan.\n")
	} else {
		total := hitungTotalSetoran(setoran, jumlahSetoran, warga[find].id)
		fmt.Printf("ID: %d, Nama: %s, No HP: %s, Total Setoran: %d kg\n",
			warga[find].id, warga[find].nama, warga[find].noHP, total)
	}
}

func searchBinaryID(warga [NMAX]dataWarga, jumlahWarga int, targetID int) int {
	var left, right, mid int
	left = 0
	right = jumlahWarga - 1

	for left <= right {
		mid = (left + right) / 2
		if warga[mid].id == targetID {
			return mid
		} else if warga[mid].id < targetID {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func cariWargaByNama(warga [NMAX]dataWarga, jumlahWarga int, setoran [NMAX]dataSetoran, jumlahSetoran int) {
	var targetNama string
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&targetNama)
	find := searchSequentialName(warga, jumlahWarga, targetNama)
	if find == -1 {
		fmt.Print("Data tidak ditemukan.\n")
	} else {
		total := hitungTotalSetoran(setoran, jumlahSetoran, warga[find].id)
		fmt.Printf("ID: %d, Nama: %s, No HP: %s, Total Setoran: %d kg\n",
			warga[find].id, warga[find].nama, warga[find].noHP, total)
	}
}

func searchSequentialName(warga [NMAX]dataWarga, jumlahWarga int, targetNama string) int {
	for i := 0; i < jumlahWarga; i++ {
		if warga[i].nama == targetNama {
			return i
		}
	}
	return -1
}

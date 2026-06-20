package main

import (
	"fmt"
)

const NMAX = 999

type dataWarga struct {
	id   int
	nama string
	noHP string
}
type dataSetoran struct {
	idWarga   int
	noSetoran int
	jenis     string
	berat     int
	tanggal   int
	mingguKe  int
}

func main() {
	var Warga [NMAX]dataWarga
	var jumlahWarga int
	var Setoran [NMAX]dataSetoran
	var jumlahSetoran int

	var opsi int = -1

	for opsi != 0 {
		fmt.Print("\n===== Waste Track Mingguan =====\n")
		fmt.Print("\nMain Menu:\n")
		fmt.Print("1. Ubah Data Warga\n")
		fmt.Print("2. Ubah Data Setoran Sampah Warga\n")
		fmt.Print("3. Cari Warga\n")
		fmt.Print("4. Leaderboard Setoran\n")
		fmt.Print("5. Lihat Akumulasi Setoran\n")
		fmt.Print("0. Exit\n")
		for opsi < 0 || opsi > 5 {
			fmt.Print("\nPilih Opsi (input hanya angka 1 sampai 5): ")
			fmt.Scan(&opsi)
			if opsi != 1 && opsi != 2 && opsi != 3 && opsi != 4 && opsi != 5 && opsi != 0 {
				fmt.Print("Input tidak valid.")
			}
		}
		switch opsi {
		case 1:
			mainWarga(&Warga, &jumlahWarga)
			opsi = -1
		case 2:
			mainSetoran(Warga, &Setoran, &jumlahSetoran, jumlahWarga)
			opsi = -1
		case 3:
			mainSearch(Warga, jumlahWarga, Setoran, jumlahSetoran)
			opsi = -1
		case 4:
			mainLeaderboard(Warga, Setoran, jumlahWarga, jumlahSetoran)
			opsi = -1
		case 5:
			mainAkumulasi(Warga, Setoran, jumlahWarga, jumlahSetoran)
			opsi = -1
		}
	}
}

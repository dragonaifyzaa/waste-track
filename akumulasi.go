package main

import "fmt"

func mainAkumulasi(Warga [NMAX]dataWarga, Setoran [NMAX]dataSetoran, jumlahWarga int, jumlahSetoran int) {
	fmt.Print("\n===== Akumulasi Setoran Sampah Warga =====\n")

	var opsi int = -1

	if jumlahWarga == 0 {
		fmt.Print("Belum ada data warga.\n")
		return

	}
	fmt.Print("Data akumulasi setoran sampah mingguan warga:\n")
	totalMingguan := totalPerMinggu(Setoran, jumlahSetoran)
	for i := 0; i <= 3; i++ {
		fmt.Printf("Minggu ke-%d: %d kg\n", i+1, totalMingguan[i])
	}

	fmt.Print("\nAkumulasi Mingguan per Warga:\n")
	akumulasiPerWargaPerMinggu(Warga, Setoran, jumlahWarga, jumlahSetoran)

	for opsi != 0 {
		fmt.Print("\n0. Back to Main Menu: ")
		fmt.Scan(&opsi)
		if opsi != 0 {
			fmt.Print("Input tidak valid.")
		}
		if opsi == 0 {
			return
		}
	}
}

func totalPerMinggu(Setoran [NMAX]dataSetoran, jumlahSetoran int) [4]int {
	var totalMingguan [4]int
	for i := 0; i < jumlahSetoran; i++ {
		minggu := Setoran[i].mingguKe
		if minggu >= 1 && minggu <= 4 {
			totalMingguan[minggu-1] = totalMingguan[minggu-1] + Setoran[i].berat
		}
	}
	return totalMingguan
}

func akumulasiPerWargaPerMinggu(Warga [NMAX]dataWarga, Setoran [NMAX]dataSetoran, jumlahWarga int, jumlahSetoran int) {
	for minggu := 1; minggu <= 4; minggu++ {
		fmt.Printf("Minggu ke-%d:\n", minggu)
		for i := 0; i < jumlahWarga; i++ {
			total := 0
			for j := 0; j < jumlahSetoran; j++ {
				if Setoran[j].idWarga == Warga[i].id && Setoran[j].mingguKe == minggu {
					total += Setoran[j].berat
				}
			}
			fmt.Printf("  [%d] %s : %d kg\n", Warga[i].id, Warga[i].nama, total)
		}
	}
}

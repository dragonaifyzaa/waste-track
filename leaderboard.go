package main

import "fmt"

type dataLeaderboard struct {
	idWarga    int
	namaWarga  string
	totalBerat int
}

func mainLeaderboard(warga [NMAX]dataWarga, setoran [NMAX]dataSetoran, jumlahWarga int, jumlahSetoran int) {
	var opsi int = -1
	fmt.Print("\n===== Leaderboard Setoran =====\n\n")

	var leaderboard [NMAX]dataLeaderboard
	leaderboard = kalkulasiBeratWarga(warga, setoran, jumlahWarga, jumlahSetoran)

	fmt.Print("Selection Sort:\n")
	selectionSortDescend(&leaderboard, jumlahWarga)

	fmt.Print("\nInsertion Sort:\n")
	InsertionSortDescend(&leaderboard, jumlahWarga)

	fmt.Print("\nLeaderboard Setoran Sampah Warga:\n")
	leaderboardTable(leaderboard, jumlahWarga)

	fmt.Print("\n")

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

func kalkulasiBeratWarga(warga [NMAX]dataWarga, setoran [NMAX]dataSetoran, jumlahWarga int, jumlahSetoran int) [NMAX]dataLeaderboard {
	var leaderboard [NMAX]dataLeaderboard
	for i := 0; i < jumlahWarga; i++ {
		leaderboard[i].idWarga = warga[i].id
		leaderboard[i].namaWarga = warga[i].nama
		leaderboard[i].totalBerat = 0
		for j := 0; j < jumlahSetoran; j++ {
			if setoran[j].idWarga == warga[i].id {
				leaderboard[i].totalBerat += setoran[j].berat
			}
		}
	}
	return leaderboard
}

func selectionSortDescend(leaderboard *[NMAX]dataLeaderboard, jumlahWarga int) {
	for i := 0; i < jumlahWarga-1; i++ {
		maxIDX := i
		for j := i + 1; j < jumlahWarga; j++ {
			if leaderboard[j].totalBerat > leaderboard[maxIDX].totalBerat {
				maxIDX = j
			}
		}
		leaderboard[i], leaderboard[maxIDX] = leaderboard[maxIDX], leaderboard[i]
		for i := 0; i < jumlahWarga; i++ {
			fmt.Print(leaderboard[i].totalBerat, " ")
		}
		fmt.Print("\n")
	}
}

func InsertionSortDescend(leaderboard *[NMAX]dataLeaderboard, jumlahWarga int) {
	for i := 1; i < jumlahWarga; i++ {
		key := leaderboard[i]
		j := i - 1
		for j >= 0 && leaderboard[j].totalBerat < key.totalBerat {
			leaderboard[j+1] = leaderboard[j]
			j--
		}
		leaderboard[j+1] = key
		for i := 0; i < jumlahWarga; i++ {
			fmt.Print(leaderboard[i].totalBerat, " ")
		}
		fmt.Print("\n")
	}
}

func leaderboardTable(leaderboard [NMAX]dataLeaderboard, jumlahWarga int) {
	fmt.Print("\n| ID | Nama Warga | Total Berat Setoran (kg) |\n")
	fmt.Print("|----|------------|-------------------|\n")
	for i := 0; i < jumlahWarga; i++ {
		fmt.Printf("| %d | %s | %d |\n", leaderboard[i].idWarga, leaderboard[i].namaWarga, leaderboard[i].totalBerat)
	}
}

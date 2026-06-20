package main

import "fmt"

func mainWarga(warga *[NMAX]dataWarga, jumlahWarga *int) {
	var opsi int = -1
	for opsi != 0 {
		opsi = -1
		fmt.Print("\n===== Ubah Data Warga =====\n")
		fmt.Print("\nPilih Opsi:\n")
		fmt.Print("1. Input Data Warga Baru\n")
		fmt.Print("2. Ganti Data Warga yang Sudah Ada\n")
		fmt.Print("3. Hapus Data Warga\n")
		fmt.Print("4. Lihat dan Cari Data Warga\n")
		fmt.Print("0. Exit\n")
		for opsi < 0 || opsi > 4 {
			fmt.Print("\nPilih Opsi (input hanya angka 0 sampai 4): ")
			fmt.Scan(&opsi)
			if opsi != 0 && opsi != 1 && opsi != 2 && opsi != 3 && opsi != 4 {
				fmt.Print("Input tidak valid.")
			}
		}
		switch opsi {
		case 1:
			inputWarga(warga, jumlahWarga)
		case 2:
			updateWarga(warga, jumlahWarga)
		case 3:
			deleteWarga(warga, jumlahWarga)
		case 4:
			lihatWarga(warga, jumlahWarga)
			cariWarga(warga, jumlahWarga)
		case 0:
			return
		}
	}
}

func autoIDWarga(warga *[NMAX]dataWarga, jumlahWarga int) int {
	if jumlahWarga == 0 {
		return 1
	} else {
		return warga[jumlahWarga-1].id + 1
	}
}

func lihatWarga(warga *[NMAX]dataWarga, jumlahWarga *int) {
	for i := 0; i < *jumlahWarga; i++ {
		fmt.Printf("ID: %d, Nama: %s, No HP: %s\n", warga[i].id, warga[i].nama, warga[i].noHP)
	}
}

func cariWarga(warga *[NMAX]dataWarga, jumlahWarga *int) {
	var opsi int = -1
	for opsi != 1 && opsi != 2 && opsi != 0 {
		fmt.Print("\nCari berdasarkan:\n")
		fmt.Print("1. ID\n")
		fmt.Print("2. Nama\n")
		fmt.Print("0. Skip Cari\n")
		fmt.Print("Pilih: ")
		fmt.Scan(&opsi)
		if opsi != 1 && opsi != 2 && opsi != 0 {
			fmt.Print("Opsi tidak valid.\n")
		}
	}
	switch opsi {
	case 1:
		searchByID(warga, jumlahWarga)
	case 2:
		searchByNama(warga, jumlahWarga)
	case 0:
		return
	}
}

func searchByID(warga *[NMAX]dataWarga, jumlahWarga *int) {
	var targetID int
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&targetID)
	find := cariIDWarga(warga, jumlahWarga, targetID)
	if find == -1 {
		fmt.Print("Data tidak ditemukan.\n")
	} else {
		fmt.Printf("ID: %d, Nama: %s, No HP: %s\n",
			warga[find].id, warga[find].nama, warga[find].noHP)
	}
}

func cariIDWarga(warga *[NMAX]dataWarga, jumlahWarga *int, targetID int) int {
	var left, right, mid int
	left = 0
	right = *jumlahWarga - 1

	//binary search
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

func searchByNama(warga *[NMAX]dataWarga, jumlahWarga *int) {
	var input string
	var i int
	fmt.Print("Masukkan Nama: ")
	fmt.Scan(&input)
	var ketemu bool = false

	//sequential search
	for i < *jumlahWarga {
		if warga[i].nama == input {
			fmt.Printf("ID: %d, Nama: %s, No HP: %s\n",
				warga[i].id, warga[i].nama, warga[i].noHP)
			ketemu = true
		}
		i++
	}
	if !ketemu {
		fmt.Print("Data tidak ditemukan.\n")
	}
}

func inputWarga(warga *[NMAX]dataWarga, jumlahWarga *int) {
	fmt.Print("\nKetik \"SELESAI\" untuk mengakhiri input\n")
	var i int = *jumlahWarga
	for i < NMAX {
		fmt.Print("\nWarga [", i+1, "]:\n")
		fmt.Print("Masukkan nama warga (tanpa spasi atau gunakan \"_\"): ")
		fmt.Scanln(&warga[i].nama)
		if warga[i].nama == "SELESAI" {
			warga[i] = dataWarga{}
			i = NMAX
		} else {
			fmt.Print("Masukkan nomor HP: ")
			fmt.Scan(&warga[i].noHP)
			if warga[i].noHP == "SELESAI" {
				warga[i] = dataWarga{}
				i = NMAX
			} else {
				warga[i].id = autoIDWarga(warga, *jumlahWarga)
				*jumlahWarga++
				i++
			}
		}
	}
	fmt.Print("\nData warga berhasil ditambahkan.\n")
}

func deleteWarga(warga *[NMAX]dataWarga, jumlahWarga *int) {
	var targetID int
	fmt.Print("Masukkan ID warga yang ingin dihapus: ")
	fmt.Scan(&targetID)

	find := cariIDWarga(warga, jumlahWarga, targetID)
	if find == -1 {
		fmt.Print("ID tidak ditemukan.\n")
		return
	}

	var i int
	for i = find; i < *jumlahWarga-1; i++ {
		warga[i] = warga[i+1]
	}
	warga[*jumlahWarga-1] = dataWarga{}
	*jumlahWarga--
	fmt.Print("Data berhasil dihapus.\n")
}

func updateWarga(warga *[NMAX]dataWarga, jumlahWarga *int) {
	var targetID int
	fmt.Print("Masukkan ID warga yang ingin diubah: ")
	fmt.Scan(&targetID)

	find := cariIDWarga(warga, jumlahWarga, targetID)
	if find == -1 {
		fmt.Print("ID tidak ditemukan.\n")
		return
	}

	var opsi int
	fmt.Printf("Data ditemukan: ID: %d Nama: %s NoHP: %s\n",
		warga[find].id, warga[find].nama, warga[find].noHP)
	fmt.Print("1. Ubah Nama\n2. Ubah No HP\n")
	fmt.Print("Pilih: ")
	fmt.Scan(&opsi)

	switch opsi {
	case 1:
		fmt.Print("Nama baru: ")
		fmt.Scan(&warga[find].nama)
	case 2:
		fmt.Print("No HP baru: ")
		fmt.Scan(&warga[find].noHP)
	default:
		fmt.Print("Opsi tidak valid.\n")
		return
	}
	fmt.Print("Data berhasil diubah.\n")
}

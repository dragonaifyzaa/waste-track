package main

import "fmt"

func mainSetoran(warga [NMAX]dataWarga, setoran *[NMAX]dataSetoran, jumlahSetoran *int, jumlahWarga int) {
	var opsi int = -1
	for opsi != 0 {
		opsi = -1
		fmt.Print("\n===== Ubah Data Setoran =====\n")
		fmt.Print("\nPilih Opsi:\n")
		fmt.Print("1. Tambah Setoran\n")
		fmt.Print("2. Cari dan Ganti Setoran\n")
		fmt.Print("3. Cari dan Hapus Setoran\n")
		fmt.Print("4. Lihat ID Warga\n")
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
			inputSetoran(warga, setoran, jumlahSetoran, jumlahWarga)
		case 2:
			updateSetoran(setoran, jumlahSetoran)
		case 3:
			deleteSetoran(setoran, jumlahSetoran)
		case 4:
			lihatIDWarga(warga, jumlahWarga)
		}
	}
}

func autoNoSetoran(setoran *[NMAX]dataSetoran, jumlahSetoran *int, idWarga int) int {
	var count int
	for i := 0; i < *jumlahSetoran; i++ {
		if setoran[i].idWarga == idWarga {
			count++
		}
	}
	return count + 1
}

func cariWargaAda(warga [NMAX]dataWarga, jumlahWarga int, idWarga int) bool {
	for i := 0; i < jumlahWarga; i++ {
		if warga[i].id == idWarga {
			return true
		}
	}
	return false
}

func inputSetoran(warga [NMAX]dataWarga, setoran *[NMAX]dataSetoran, jumlahSetoran *int, jumlahWarga int) {
	var idWarga int
	fmt.Print("\nMasukkan ID Warga: ")
	fmt.Scan(&idWarga)
	if !cariWargaAda(warga, jumlahWarga, idWarga) {
		fmt.Print("Warga tidak ditemukan.\n")
		return
	}

	lihatSetoran(setoran, jumlahSetoran, idWarga)

	fmt.Print("\nKetik \"SELESAI\" untuk mengakhiri input\n")
	fmt.Print("Khusus \"Berat\", \"Tanggal\", atau \"Minggu Ke-\", Ketik \"-1\" untuk mengakhiri input\n")
	var i int = *jumlahSetoran
	for i < NMAX {
		no := autoNoSetoran(setoran, jumlahSetoran, idWarga)
		fmt.Print("\nSetoran [", no, "]:\n")
		fmt.Print("Masukkan Jenis Sampah: ")
		fmt.Scan(&setoran[i].jenis)
		if setoran[i].jenis == "SELESAI" {
			setoran[i] = dataSetoran{}
			i = NMAX
		} else {
			fmt.Print("Masukkan Berat Sampah (kg): ")
			fmt.Scan(&setoran[i].berat)
			if setoran[i].berat == -1 {
				setoran[i] = dataSetoran{}
				i = NMAX
			} else {
				fmt.Print("Masukkan Tanggal Setoran (1-31): ")
				fmt.Scan(&setoran[i].tanggal)
				if setoran[i].tanggal == -1 {
					setoran[i] = dataSetoran{}
					i = NMAX
				} else {
					fmt.Print("Masukkan Minggu Ke: ")
					fmt.Scan(&setoran[i].mingguKe)
					if setoran[i].mingguKe == -1 {
						setoran[i] = dataSetoran{}
						i = NMAX
					} else {
						setoran[i].idWarga = idWarga
						setoran[i].noSetoran = no
						*jumlahSetoran++
						i++
					}
				}
			}
		}
	}
	fmt.Print("\nData setoran berhasil ditambahkan.\n")
}

func updateSetoran(setoran *[NMAX]dataSetoran, jumlahSetoran *int) {
	var idWarga, noSetoran int
	fmt.Print("Masukkan ID Warga: ")
	fmt.Scan(&idWarga)

	lihatSetoran(setoran, jumlahSetoran, idWarga)

	fmt.Print("Masukkan No Setoran yang ingin diubah: ")
	fmt.Scan(&noSetoran)
	find := cariNoSetoran(setoran, jumlahSetoran, idWarga, noSetoran)
	if find == -1 {
		fmt.Print("Setoran tidak ditemukan.\n")
		return
	}

	var opsi int = -1
	for opsi < 0 || opsi > 3 {
		fmt.Print("\nPilih Opsi:\n")
		fmt.Print("1. Ganti Jenis Sampah\n")
		fmt.Print("2. Ganti Berat Sampah\n")
		fmt.Print("3. Ganti Tanggal Setoran\n")
		fmt.Print("0. Skip Ganti Data Setoran\n")
		fmt.Print("Pilih: ")
		fmt.Scan(&opsi)
		if opsi < 0 || opsi > 3 {
			fmt.Print("Input tidak valid.\n")
		}
	}
	switch opsi {
	case 1:
		fmt.Print("Jenis Sampah baru: ")
		fmt.Scan(&setoran[find].jenis)
	case 2:
		fmt.Print("Berat Sampah baru (kg): ")
		fmt.Scan(&setoran[find].berat)
	case 3:
		fmt.Print("Tanggal Setoran baru (1-31): ")
		fmt.Scan(&setoran[find].tanggal)
	case 0:
		return
	}
	fmt.Print("Data setoran berhasil diubah.\n")
}

func deleteSetoran(setoran *[NMAX]dataSetoran, jumlahSetoran *int) {
	var idWarga, noSetoran int
	fmt.Print("Masukkan ID Warga: ")
	fmt.Scan(&idWarga)

	lihatSetoran(setoran, jumlahSetoran, idWarga)

	fmt.Print("Masukkan No Setoran yang ingin dihapus: ")
	fmt.Scan(&noSetoran)
	find := cariNoSetoran(setoran, jumlahSetoran, idWarga, noSetoran)
	if find == -1 {
		fmt.Print("Setoran tidak ditemukan.\n")
		return
	}

	var i int
	for i = find; i < *jumlahSetoran-1; i++ {
		setoran[i] = setoran[i+1]
	}
	setoran[*jumlahSetoran-1] = dataSetoran{}
	*jumlahSetoran--
	fmt.Print("Data setoran berhasil dihapus.\n")
}

func lihatSetoran(setoran *[NMAX]dataSetoran, jumlahSetoran *int, idWarga int) {
	var ketemu bool
	fmt.Print("\nData Setoran Warga:\n")
	for i := 0; i < *jumlahSetoran; i++ {
		if setoran[i].idWarga == idWarga {
			fmt.Printf("Setoran [%d]: %s | %d kg | Tgl %d | Minggu ke-%d\n",
				setoran[i].noSetoran, setoran[i].jenis, setoran[i].berat, setoran[i].tanggal, setoran[i].mingguKe)
			ketemu = true
		}
	}
	if !ketemu {
		fmt.Print("Warga belum memiliki setoran.\n")
	}
}

func cariNoSetoran(setoran *[NMAX]dataSetoran, jumlahSetoran *int, idWarga int, noSetoran int) int {
	for i := 0; i < *jumlahSetoran; i++ {
		if setoran[i].idWarga == idWarga && setoran[i].noSetoran == noSetoran {
			return i
		}
	}
	return -1
}

func lihatIDWarga(warga [NMAX]dataWarga, jumlahWarga int) {
	for i := 0; i < jumlahWarga; i++ {
		fmt.Printf("ID: %d, Nama: %s, No HP: %s\n", warga[i].id, warga[i].nama, warga[i].noHP)
	}
}

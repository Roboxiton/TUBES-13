package main

import "fmt"

const max = 1000

type mindstone struct {
	namaTugas string
	dTanggal  int
	dBulan    int
	dTahun    int
	prioritas int
	progress  float64
	mood      int
	stress    float64
}

type tabMindstone [max]mindstone

func main() {
	var n string
	var j int
	var d tabMindstone
	for n != "8" {
		fmt.Println("--------------Selamat Datang di Mindstone--------------")
		fmt.Printf("%38s\n", "Masukkan pilihan Anda")
		fmt.Printf("%s\n", "1. Tambah Data")
		fmt.Printf("%s\n", "2. Lihat Data")
		fmt.Printf("%s\n", "3. Cari Data")
		fmt.Printf("%s\n", "4. Update Data")
		fmt.Printf("%s\n", "5. Hapus Data")
		fmt.Printf("%s\n", "6. Urutkan Data")
		fmt.Printf("%s\n", "7. Tampilkan Statistik Data")
		fmt.Printf("%s\n", "8. Keluar")
		fmt.Scan(&n)
		if n == "1" {
			tambahData(&d, &j)
		} else if n == "2" {
			lihatData(&d, &j)
		} else if n == "3" {
			cariData(&d, &j)
		} else if n == "4" {
			updateData(&d, &j)
		} else if n == "5" {
			hapusData(&d, &j)
		} else if n == "6" {
			urutkanData(&d, &j)
		} else if n == "7" {
			statistikData(&d, &j)
		} else if n == "8" {
			fmt.Println("Terima kasih telah menggunakan Mindstone!")
		} else {
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}

func tambahData(d *tabMindstone, j *int) {
	// User menambahkan jumlah tugas dan mengisi identitas tugas
	var i, n int
	fmt.Print("Silahkan isi jumlah data: ")
	fmt.Scan(&n)
	for i < n {
		for *j >= max {
			fmt.Println("Kapasitas data sudah penuh. Tidak dapat menambahkan data baru.")
			return
		}
		fmt.Print("Nama Tugas: ")
		fmt.Scan(&d[*j].namaTugas)
		fmt.Print("Deadline Tanggal: ")
		fmt.Scan(&d[*j].dTanggal)
		for d[*j].dTanggal < 1 || d[*j].dTanggal > 31 {
			fmt.Println("Tanggal tidak valid, silakan masukkan tanggal antara 1 dan 31.")
			fmt.Print("Deadline Tanggal: ")
			fmt.Scan(&d[*j].dTanggal)
		}
		fmt.Print("Deadline Bulan: ")
		fmt.Scan(&d[*j].dBulan)
		for d[*j].dBulan < 1 || d[*j].dBulan > 12 {
			fmt.Println("Bulan tidak valid, silakan masukkan bulan antara 1 dan 12.")
			fmt.Print("Deadline Bulan: ")
			fmt.Scan(&d[*j].dBulan)
		}
		fmt.Print("Deadline Tahun: ")
		fmt.Scan(&d[*j].dTahun)
		for d[*j].dTahun < 1 {
			fmt.Println("Tahun tidak valid, silakan masukkan tahun yang valid.")
			fmt.Print("Masukkan Tahun: ")
			fmt.Scan(&d[*j].dTahun)
		}
		fmt.Print("Prioritas (1-10): ")
		fmt.Scan(&d[*j].prioritas)
		for d[*j].prioritas < 1 || d[*j].prioritas > 10 {
			fmt.Println("Prioritas tidak valid, silakan masukkan prioritas antara 1 dan 10.")
			fmt.Print("Prioritas (1-10): ")
			fmt.Scan(&d[*j].prioritas)
		}
		fmt.Print("Progress (%): ")
		fmt.Scan(&d[*j].progress)
		for d[*j].progress < 0 || d[*j].progress > 100 {
			fmt.Println("Progress tidak valid, silakan masukkan progress antara 0 dan 100.")
			fmt.Print("Progress (%): ")
			fmt.Scan(&d[*j].progress)
		}
		fmt.Print("Mood (1-10): ")
		fmt.Scan(&d[*j].mood)
		for d[*j].mood < 1 || d[*j].mood > 10 {
			fmt.Println("Mood tidak valid, silakan masukkan mood antara 1 dan 10.")
			fmt.Print("Mood (1-10): ")
			fmt.Scan(&d[*j].mood)
		}
		fmt.Print("Stress (1-10): ")
		fmt.Scan(&d[*j].stress)
		for d[*j].stress < 1 || d[*j].stress > 10 {
			fmt.Println("Stress tidak valid, silakan masukkan stress antara 1 dan 10.")
			fmt.Print("Stress (1-10): ")
			fmt.Scan(&d[*j].stress)
		}
		i++
		(*j)++
	}
}

func lihatData(d *tabMindstone, j *int) {
	// User bisa melihat list tugas setelah mengisi data
	var i int
	var enter string
	if *j == 0 {
		fmt.Println("Belum ada data yang dimasukkan.")
	} else {
		for enter != "X" && enter != "x" {
			fmt.Printf("%-15s %-15s %-10s %-10s %-10s %-10s\n", "Nama Tugas", "Deadline", "Prioritas", "Progress", "Mood", "Stress")
			for i = 0; i < *j; i++ {
				fmt.Printf("%-15s %d/%d/%-11d %-10d %-10.2f %-10d %-10.2f\n", d[i].namaTugas, d[i].dTanggal, d[i].dBulan, d[i].dTahun, d[i].prioritas, d[i].progress, d[i].mood, d[i].stress)

			}
			fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk melihat data lagi.")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" && enter != "Y" && enter != "y" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama atau tekan Y untuk melihat data lagi.")
				fmt.Scan(&enter)
			}
		}
	}
}

func cariData(d *tabMindstone, j *int) {
	// User bisa mencari data tugas setelah diisi data dengan 2 metode pencarian yaitu Sequential Search untuk mencari berdasarkan nama tugas dan Binary Search untuk mencari berdasarkan deadline tugas
	var i, idx int
	var namaTugas string
	var tanggal, bulan, tahun int
	var enter string
	fmt.Println("Pilih metode yang ingin Anda cari: ")
	fmt.Print("1. Cari berdasarkan nama tugas (SequentialSearch)\n")
	fmt.Print("2. Cari berdasarkan deadline (BinarySearch)\n")
	fmt.Print("3. Kembali ke menu utama\n")
	fmt.Scan(&i)
	for i < 1 || i > 3 {
		fmt.Println("Pilihan tidak valid, silakan pilih antara 1, 2, atau 3.")
		fmt.Scan(&i)
	}
	switch i {
	case 1:
		for enter != "X" && enter != "x" {
			fmt.Print("Masukkan nama tugas yang ingin dicari: ")
			fmt.Scan(&namaTugas)
			idx = SequentialNamaTugas(*d, *j, namaTugas)
			if idx == -1 {
				fmt.Printf("Data tidak ditemukan untuk tugas: %s\n", namaTugas)
			} else {
				fmt.Printf("Data ditemukan untuk tugas: %s\n", namaTugas)
				fmt.Printf("Deadline: %d/%d/%d\n", d[idx].dTanggal, d[idx].dBulan, d[idx].dTahun)
				fmt.Printf("Prioritas: %d\n", d[idx].prioritas)
				fmt.Printf("Progress: %.2f\n", d[idx].progress)
				fmt.Printf("Mood: %d\n", d[idx].mood)
				fmt.Printf("Stress: %.2f\n", d[idx].stress)
			}
			fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk mencari data lagi.")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" && enter != "Y" && enter != "y" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama atau tekan Y untuk mencari data lagi.")
				fmt.Scan(&enter)
			}
		}
	case 2:
		for enter != "X" && enter != "x" {
			fmt.Print("Masukkan deadline yang ingin dicari (format: tanggal bulan tahun): ")
			fmt.Scan(&tanggal, &bulan, &tahun)
			idx = BinaryTanggal(*d, *j, tanggal, bulan, tahun)
			if idx == -1 {
				fmt.Printf("Data tidak ditemukan untuk deadline: %d/%d/%d\n", tanggal, bulan, tahun)
			} else {
				fmt.Printf("Data ditemukan untuk deadline: %d/%d/%d\n", tanggal, bulan, tahun)
				fmt.Printf("Nama Tugas: %s\n", d[idx].namaTugas)
				fmt.Printf("Prioritas: %d\n", d[idx].prioritas)
				fmt.Printf("Progress: %.2f\n", d[idx].progress)
				fmt.Printf("Mood: %d\n", d[idx].mood)
				fmt.Printf("Stress: %.2f\n", d[idx].stress)
			}
			fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk mencari data lagi.")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" && enter != "Y" && enter != "y" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama atau tekan Y untuk mencari data lagi.")
				fmt.Scan(&enter)
			}
		}
	case 3:
		return
	}
}

func updateData(d *tabMindstone, j *int) {
	// USer bisa melakukan update data untuk mengubah progress tugas atau mengubah deadline tugas
	var namaTugas string
	var enter string
	var i int
	for enter != "X" && enter != "x" {
		fmt.Print("Masukkan nama tugas yang ingin diupdate: ")
		fmt.Scan(&namaTugas)
		i = 0
		for i < *j && d[i].namaTugas != namaTugas {
			i++
		}
		if i < *j {
			fmt.Printf("Data ditemukan untuk tugas: %s\n", namaTugas)
			fmt.Print("Silakan isi data yang anda mau ubah.\n")
			fmt.Print("Deadline Tanggal: ")
			fmt.Scan(&d[i].dTanggal)
			for d[i].dTanggal < 1 || d[i].dTanggal > 31 {
				fmt.Println("Tanggal tidak valid, silakan masukkan tanggal antara 1 dan 31.")
				fmt.Print("Deadline Tanggal: ")
				fmt.Scan(&d[i].dTanggal)
			}
			fmt.Print("Deadline Bulan: ")
			fmt.Scan(&d[i].dBulan)
			for d[i].dBulan < 1 || d[i].dBulan > 12 {
				fmt.Println("Bulan tidak valid, silakan masukkan bulan antara 1 dan 12.")
				fmt.Print("Deadline Bulan: ")
				fmt.Scan(&d[i].dBulan)
			}
			fmt.Print("Deadline Tahun: ")
			fmt.Scan(&d[i].dTahun)
			for d[i].dTahun < 1 {
				fmt.Println("Tahun tidak valid, silakan masukkan tahun yang valid.")
				fmt.Print("Masukkan Tahun: ")
				fmt.Scan(&d[i].dTahun)
			}
			fmt.Print("Prioritas (1-10): ")
			fmt.Scan(&d[i].prioritas)
			for d[i].prioritas < 1 || d[i].prioritas > 10 {
				fmt.Println("Prioritas tidak valid, silakan masukkan prioritas antara 1 dan 10.")
				fmt.Print("Prioritas (1-10): ")
				fmt.Scan(&d[i].prioritas)
			}
			fmt.Print("Progress (%): ")
			fmt.Scan(&d[i].progress)
			for d[i].progress < 0 || d[i].progress > 100 {
				fmt.Println("Progress tidak valid, silakan masukkan progress antara 0 dan 100.")
				fmt.Print("Progress (%): ")
				fmt.Scan(&d[i].progress)
			}
			fmt.Print("Mood (1-10): ")
			fmt.Scan(&d[i].mood)
			for d[i].mood < 1 || d[i].mood > 10 {
				fmt.Println("Mood tidak valid, silakan masukkan mood antara 1 dan 10.")
				fmt.Print("Mood (1-10): ")
				fmt.Scan(&d[i].mood)
			}
			fmt.Print("Stress (1-10): ")
			fmt.Scan(&d[i].stress)
			for d[i].stress < 1 || d[i].stress > 10 {
				fmt.Println("Stress tidak valid, silakan masukkan stress antara 1 dan 10.")
				fmt.Print("Stress (1-10): ")
				fmt.Scan(&d[i].stress)
			}
			fmt.Println("Data berhasil diupdate.")
			fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk mengupdate data lagi.")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" && enter != "Y" && enter != "y" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama atau tekan Y untuk mengupdate data lagi.")
				fmt.Scan(&enter)
			}
		} else {
			fmt.Printf("Data tidak ditemukan untuk tugas: %s\n", namaTugas)
			fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk mencari data lagi.")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" && enter != "Y" && enter != "y" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama atau tekan Y untuk mencari data lagi.")
				fmt.Scan(&enter)
			}
		}
	}
}

func hapusData(d *tabMindstone, j *int) {
	// User bisa menghapus data jika tugas tidak diperlukan lagi
	var i, idx int
	var enter string
	for enter != "X" && enter != "x" {
		fmt.Print("Masukkan index yang ingin dihapus: ")
		fmt.Scan(&idx)
		if idx < 0 || idx >= *j {
			fmt.Println("Index tidak valid.")
			return
		}
		if idx < *j {
			fmt.Println("Data yang akan dihapus: ")
			fmt.Printf("%-15s %d/%d/%-11d %-10d %-10.2f %-10d %-10.2f\n", d[idx].namaTugas, d[idx].dTanggal, d[idx].dBulan, d[idx].dTahun, d[idx].prioritas, d[idx].progress, d[idx].mood, d[idx].stress)
			fmt.Print("Apakah Anda yakin ingin menghapus data ini? (Y/N): ")
			fmt.Scan(&enter)
			for enter != "Y" && enter != "y" && enter != "N" && enter != "n" {
				fmt.Println("Pilihan tidak valid, silakan tekan Y untuk menghapus data atau N untuk membatalkan.")
				fmt.Scan(&enter)
			}
			if enter == "Y" || enter == "y" {
				for i = idx; i < *j-1; i++ {
					d[i] = d[i+1]
				}
				(*j)--
				fmt.Println("Data berhasil dihapus.")
				fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk menghapus data lagi.")
				fmt.Scan(&enter)
			} else {
				fmt.Println("Penghapusan data dibatalkan.")
				fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk menghapus data lagi.")
				fmt.Scan(&enter)
			}
		}
	}
}

func urutkanData(d *tabMindstone, j *int) {
	// User bisa mengurutkan data berdasarkan prioritas atau mood dengan menggunakan metode pengurutan Selection Sort untuk mengurutkan berdasarkan prioritas dan Insertion Sort untuk mengurutkan berdasarkan mood
	var idx int
	var enter string
	for enter != "X" && enter != "x" {
		fmt.Println("Pilih metode pengurutan: ")
		fmt.Print("1. Urutkan berdasarkan prioritas (Selection Sort)\n")
		fmt.Print("2. Urutkan berdasarkan mood (Insertion Sort)\n")
		fmt.Print("3. Kembali ke menu utama\n")
		fmt.Scan(&idx)
		for idx < 1 || idx > 3 {
			fmt.Println("Pilihan tidak valid, silakan pilih antara 1, 2, atau 3.")
			fmt.Scan(&idx)
		}
		switch idx {
		case 1:
			SelectionPrioritas(d, j)
			fmt.Println("Data berhasil diurutkan berdasarkan prioritas.")
			fmt.Println("Tekan X untuk kembali ke menu utama")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama.")
				fmt.Scan(&enter)
			}
		case 2:
			insertionMood(d, j)
			fmt.Println("Data berhasil diurutkan berdasarkan mood.")
			fmt.Println("Tekan X untuk kembali ke menu utama")
			fmt.Scan(&enter)
			for enter != "X" && enter != "x" {
				fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama.")
				fmt.Scan(&enter)
			}
		case 3:
			return
		}
	}
}

func statistikData(d *tabMindstone, j *int) {
	// User bisa melihat statistik data seperti jumlah tugas yang sudah selesai, jumlah tugas yang belum selesai, persentase tugas yang sudah selesai, dan rata-rata stress level dari semua tugas
	var TugasSelesai, TotalTugas, i int
	var AvgStress float64
	var enter string
	if *j == 0 {
		fmt.Println("Belum ada data yang dimasukkan.")
		return
	}
	for i = 0; i < *j; i++ {
		if d[i].progress == 100 {
			TugasSelesai++
		}
		TotalTugas++
		AvgStress += d[i].stress
	}
	AvgStress = AvgStress / float64(TotalTugas)
	for enter != "X" && enter != "x" {
		fmt.Printf("Jumlah Tugas Selesai: %d\n", TugasSelesai)
		fmt.Printf("Jumlah Total Tugas: %d\n", TotalTugas)
		fmt.Printf("Persentase selesai: %.2f\n", (float64(TugasSelesai)/float64(TotalTugas))*100)
		fmt.Printf("Rata-rata Stress: %.2f\n", AvgStress)
		fmt.Println("Tekan X untuk kembali ke menu utama atau tekan Y untuk melihat data lagi.")
		fmt.Scan(&enter)
		for enter != "X" && enter != "x" && enter != "Y" && enter != "y" {
			fmt.Println("Pilihan tidak valid, silakan tekan X untuk kembali ke menu utama atau tekan Y untuk melihat data lagi.")
			fmt.Scan(&enter)
		}
	}
}

func SequentialNamaTugas(d tabMindstone, j int, namaTugas string) int {
	var i int
	for i < j && d[i].namaTugas != namaTugas {
		i++
	}
	if i < j {
		return i
	} else {
		return -1
	}
}

func BinaryTanggal(d tabMindstone, j int, tanggal, bulan, tahun int) int {
	var left, right, mid int
	left = 0
	right = j - 1
	mid = (left + right) / 2
	for left <= right {
		if d[mid].dTahun == tahun && d[mid].dBulan == bulan && d[mid].dTanggal == tanggal {
			return mid
		} else if d[mid].dTahun < tahun || (d[mid].dTahun == tahun && d[mid].dBulan < bulan) || (d[mid].dTahun == tahun && d[mid].dBulan == bulan && d[mid].dTanggal < tanggal) {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left + right) / 2
	}
	return -1
}

func SelectionPrioritas(d *tabMindstone, j *int) {
	var i, idx, pass int
	var temp mindstone
	pass = 1
	for pass < *j {
		idx = pass - 1
		i = pass
		for i < *j {
			if d[idx].prioritas < d[i].prioritas {
				idx = i
			} else if d[idx].prioritas == d[i].prioritas {
				if d[idx].progress < d[i].progress {
					idx = i
				}
			}
			i++
		}
		temp = d[pass-1]
		d[pass-1] = d[idx]
		d[idx] = temp
		pass++
	}
}

func insertionMood(d *tabMindstone, j *int) {
	var pass, i int
	var temp mindstone
	pass = 1
	for pass < *j {
		i = pass
		temp = d[pass]
		for i > 0 && ((temp.mood < d[i-1].mood) || (temp.mood == d[i-1].mood && temp.prioritas < d[i-1].prioritas)) {
			d[i] = d[i-1]
			i--
		}
		d[i] = temp
		pass++
	}
}

package main

import "fmt"

const NMAX = 9999

type data struct {
	Judul           string
	Penulis         string
	Pembimbing      string
	topikPenelitian string
	statusKelulusan string
	tahunLulus      int
}
type tabData [NMAX]data

func nambahDokumen(A *tabData, n *int) {
	if *n < NMAX {
		fmt.Println("--- Tambah Dokumen Skripsi ---")

		fmt.Print("Masukkan Judul: ")
		fmt.Scan(&A[*n].Judul)

		fmt.Print("Masukkan Penulis: ")
		fmt.Scan(&A[*n].Penulis)

		fmt.Print("Masukkan Pembimbing: ")
		fmt.Scan(&A[*n].Pembimbing)

		fmt.Print("Masukkan Topik Penelitian: ")
		fmt.Scan(&A[*n].topikPenelitian)

		fmt.Print("Masukkan Status Kelulusan: ")
		fmt.Scan(&A[*n].statusKelulusan)

		fmt.Print("Masukkan Tahun Lulus: ")
		fmt.Scan(&A[*n].tahunLulus)

		*n = *n + 1
		fmt.Printf("Data berhasil disimpan!\n")
	} else {
		fmt.Println("Jumlah dokumen sudah mencapai batas maksimum.")
	}
}

func hapusDokumen(A *tabData, n *int) {
	var index int
	fmt.Println("--- Hapus Dokumen Skripsi ---")
	fmt.Printf("Masukkan nomor dokumen yang ingin dihapus (1-%d): ", *n)
	fmt.Scan(&index)
	if index >= 1 && index <= *n {
		for i := index - 1; i < *n-1; i++ {
			A[i] = A[i+1]
		}
		*n = *n - 1
		fmt.Println("Dokumen berhasil dihapus!")
	} else {
		fmt.Println("Nomor dokumen tidak valid.")
	}
}

func lihatDokumen(A *tabData, n int) {
	fmt.Println("--- Daftar Dokumen Skripsi ---")
	for i := 0; i < n; i++ {
		fmt.Printf("Dokumen %d:\n", i+1)
		fmt.Printf("  Judul: %s\n", A[i].Judul)
		fmt.Printf("  Penulis: %s\n", A[i].Penulis)
		fmt.Printf("  Pembimbing: %s\n", A[i].Pembimbing)
		fmt.Printf("  Topik Penelitian: %s\n", A[i].topikPenelitian)
		fmt.Printf("  Status Kelulusan: %s\n", A[i].statusKelulusan)
		fmt.Printf("  Tahun Lulus: %d\n", A[i].tahunLulus)
	}
}

func sequentialSearch(A *tabData, n int, judul string) int {
	for i := 0; i < n; i++ {
		if A[i].Judul == judul {
			return i
		}
	}
	return -1
}

func binarySearch(A *tabData, n int, penulisDicari string) int {
	insertionSortByPenulis(A, n)
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if A[mid].Penulis == penulisDicari {
			return mid
		} else if A[mid].Penulis < penulisDicari {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func selectionSort(A *tabData, BorS int, n int) {
	var temp data
	if BorS == 1 {
		for i := 0; i < n-1; i++ {
			minIndex := i
			for j := i + 1; j < n; j++ {
				if A[j].tahunLulus < A[minIndex].tahunLulus {
					minIndex = j
				}
			}
			temp = A[i]
			A[i] = A[minIndex]
			A[minIndex] = temp
		}
	} else if BorS == 2 {
		for i := 0; i < n-1; i++ {
			maxIndex := i
			for j := i + 1; j < n; j++ {
				if A[j].tahunLulus > A[maxIndex].tahunLulus {
					maxIndex = j
				}
			}
			temp = A[i]
			A[i] = A[maxIndex]
			A[maxIndex] = temp
		}
	}
}

func insertionSort(A *tabData, BorS int, n int) {
	var temp data
	if BorS == 1 {
		for i := 1; i < n; i++ {
			temp = A[i]
			j := i - 1
			for j >= 0 && A[j].tahunLulus > temp.tahunLulus {
				A[j+1] = A[j]
				j = j - 1
			}
			A[j+1] = temp
		}
	} else if BorS == 2 {
		for i := 1; i < n; i++ {
			temp = A[i]
			j := i - 1
			for j >= 0 && A[j].tahunLulus < temp.tahunLulus {
				A[j+1] = A[j]
				j = j - 1
			}
			A[j+1] = temp
		}
	}
}

func insertionSortByPenulis(A *tabData, n int) {
	var temp data
	for i := 1; i < n; i++ {
		temp = A[i]
		j := i - 1
		for j >= 0 && A[j].Penulis > temp.Penulis {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = temp
	}
}

func tampilkanStatistik(A *tabData, n int) {
	fmt.Println("\n================ STATISTIK DATA SKRIPSI ================")
	fmt.Printf("Total dokumen yang tersimpan saat ini: %d dokumen\n", n)
	fmt.Println("--------------------------------------------------------")

	if n == 0 {
		fmt.Println("silahkan tambahkan dokumen")
		fmt.Println("========================================================")
		return
	}
	insertionSort(A, 1, n)
	fmt.Println("Jumlah judul skripsi per tahun:")
	tahunAwal := A[0].tahunLulus
	jumlahPerTahun := 1

	for i := 1; i < n; i++ {
		if A[i].tahunLulus == tahunAwal {
			jumlahPerTahun = jumlahPerTahun + 1
		} else {
			fmt.Printf("  - Tahun %d: %d judul\n", tahunAwal, jumlahPerTahun)
			tahunAwal = A[i].tahunLulus
			jumlahPerTahun = 1
		}
	}
	fmt.Printf("  - Tahun %d: %d judul\n", tahunAwal, jumlahPerTahun)
	fmt.Println("========================================================")
}

func main() {
	var jumlah int
	var menu int
	var kembaliMenu bool = true

	var A tabData = tabData{
		{"Sistem_Informasi_Akademik_Berbasis_Web", "Andi", "Dr.Budi", "Sistem_Informasi", "Lulus", 2020},
		{"Sistem_Pakar_Diagnosa_Penyakit", "Budi", "Prof.Hermawan", "Kecerdasan_Buatan", "Lulus", 2021},
		{"Analisis_Sentimen_Twitter", "Siti", "Dr.Ahmad", "Data_Science", "Lulus", 2022},
		{"Aplikasi_E-Commerce_Microservices", "Rian", "Siti_Nurjanah", "RPL", "Lulus", 2023},
		{"Implementasi_IoT_Penyiraman", "Dewi", "Andi_Wijaya", "IoT", "Belum_Lulus", 2024},
	}

	jumlah = 5

	for kembaliMenu {
		fmt.Printf("==============MENU===============\n")
		fmt.Println("1. Tambah Dokumen")
		fmt.Println("2. Lihat Dokumen")
		fmt.Println("3. Hapus Dokumen")
		fmt.Println("4. Cari Dokumen")
		fmt.Println("5. Urutkan Dokumen")
		fmt.Println("6. Statistik Dokumen")
		fmt.Println("7. Keluar")
		fmt.Printf("================================\n")
		fmt.Scan(&menu)
		if menu == 1 {
			nambahDokumen(&A, &jumlah)
		} else if menu == 2 {
			lihatDokumen(&A, jumlah)
		} else if menu == 3 {
			hapusDokumen(&A, &jumlah)
		} else if menu == 4 {
			fmt.Println("Pilih metode pencarian:")
			fmt.Println("1. Sequential Search (Berdasarkan Judul)")
			fmt.Println("2. Binary Search (Berdasarkan Penulis)")
			var pilihanPencarian, index int
			fmt.Scan(&pilihanPencarian)
			if pilihanPencarian == 1 {
				fmt.Print("Masukkan judul dokumen yang ingin dicari: ")
				var judul string
				fmt.Scan(&judul)
				index = sequentialSearch(&A, jumlah, judul)
			} else if pilihanPencarian == 2 {
				fmt.Print("Masukkan nama penulis dokumen yang ingin dicari: ")
				var penulis string
				fmt.Scan(&penulis)
				index = binarySearch(&A, jumlah, penulis)
			}
			if index != -1 {
				fmt.Printf("Dokumen ditemukan pada posisi %d:\n", index+1)
				fmt.Printf("  Judul: %s\n", A[index].Judul)
				fmt.Printf("  Penulis: %s\n", A[index].Penulis)
				fmt.Printf("  Pembimbing: %s\n", A[index].Pembimbing)
				fmt.Printf("  Topik Penelitian: %s\n", A[index].topikPenelitian)
				fmt.Printf("  Status Kelulusan: %s\n", A[index].statusKelulusan)
				fmt.Printf("  Tahun Lulus: %d\n", A[index].tahunLulus)
			} else {
				fmt.Println("Dokumen tidak ditemukan.")
			}
		} else if menu == 5 {
			fmt.Println("Pilih metode pengurutan berdasarkan tahun lulus:")
			fmt.Println("1. Selection Sort")
			fmt.Println("2. Insertion Sort")
			var pilihan int
			fmt.Scan(&pilihan)
			fmt.Println("1. Ascending")
			fmt.Println("2. Descending")
			var urutan int
			fmt.Scan(&urutan)
			if pilihan == 1 {
				selectionSort(&A, urutan, jumlah)
				fmt.Println("Dokumen berhasil diurutkan menggunakan Selection Sort.")
			} else if pilihan == 2 {
				insertionSort(&A, urutan, jumlah)
				fmt.Println("Dokumen berhasil diurutkan menggunakan Insertion Sort.")
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		} else if menu == 6 {
			tampilkanStatistik(&A, jumlah)
		} else if menu == 7 {
			fmt.Println("Terima kasih telah menggunakan program ini!")
			kembaliMenu = false
		} else {
			fmt.Println("MASUKKAN ANGKA YANG BENAR WOI!!!")
		}
	}
}

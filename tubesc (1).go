package main

import (
	"fmt"
	"time"
)

type Langganan struct {
	Nama             string
	Biaya            int
	JatuhTempo       string
	Kategori         string
	metodePembayaran string
}

var daftarLangganan []Langganan

func main() {
	daftarLangganan = append(daftarLangganan,
		Langganan{"Netflix", 200000, "01-06-2025", "Entertaiment", "Ovo"},
		Langganan{"Spotify", 50000, "14-06-2025", "Music", "Gopay"},
		Langganan{"YoutubePremium", 100000, "20-06-2025", "Entertaiment", "Gopay"},
		Langganan{"Biznet", 1000000, "10-06-2025", "Internet", "ShopeePay"},
		Langganan{"AlightMotion", 50000, "14-07-2025", "Entertaiment", "Gopay"},
	)
	var inputan int

	for {
		fmt.Println("Daftar Menu Aplikasi Langganan")
		fmt.Println("1. Tambah Daftar Langganan")
		fmt.Println("2. Ubah Daftar Langganan")
		fmt.Println("3. Hapus Daftar Langganan")
		fmt.Println("4. Total Pengeluaran Perbulan")
		fmt.Println("5. Cari Daftar Langganan Berdasarkan Nama")
		fmt.Println("6. Cari Daftar Langganan Berdasarkan Kategori")
		fmt.Println("7. Rekomendasi Penghentian Langganan")
		fmt.Println("8. Sortir Melalui Biaya Langganan Terbesar")
		fmt.Println("9. Sortir Melalui Jatuh Tempo Tercepat")
		fmt.Println("10. Tampilkan Semua Daftar Langganan")
		fmt.Println("11. Keluar")
		fmt.Print("Masukkan pilihan: ")

		fmt.Scanln(&inputan)

		switch inputan {
		case 1:
			tambahLangganan()
		case 2:
			ubahLangganan()
		case 3:
			hapusLangganan()
		case 4:
			totalPengeluaran()
		case 5:
			cariNamaLangganan()
		case 6:
			cariBerdasarkanKategori()
		case 7:
			rekomendasiPenghentian()
		case 8:
			sortirLanggananBiaya()
		case 9:
			sortirJatuhTempoTercepat()
		case 10:
			tampilkanLangganan()
		case 11:
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Pilihan Tidak Tersedia.")
		}

		fmt.Println()
	}
}
func tambahLangganan() {
	var nama, jatuhTempo, kategori, metode string
	var biaya int

	fmt.Print("Masukkan Nama Langganan: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Biaya Langganan (Rp): ")
	fmt.Scanln(&biaya)
	fmt.Print("Masukkan Jatuh Tempo: ")
	fmt.Scanln(&jatuhTempo)
	fmt.Print("Masukkan Kategori: ")
	fmt.Scanln(&kategori)
	fmt.Print("Masukkan Metode Pembayaran: ")
	fmt.Scanln(&metode)

	daftarLangganan = append(daftarLangganan, Langganan{
		Nama: nama, Biaya: biaya, JatuhTempo: jatuhTempo, Kategori: kategori, metodePembayaran: metode})
	fmt.Println("Daftar langganan berhasil ditambahkan.")
}
func ubahLangganan() {
	var input string
	fmt.Print("Masukkan Nama Langganan yang ingin diubah: ")
	fmt.Scanln(&input)

	for i := 0; i < len(daftarLangganan); i++ {
		if daftarLangganan[i].Nama == input {
			var nama, jatuhTempo, kategori, metode string
			var biaya int

			fmt.Print("Masukkan nama langganan yang baru: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan biaya langganan yang baru: ")
			fmt.Scanln(&biaya)
			fmt.Print("Masukkan jatuh tempo langganan yang baru: ")
			fmt.Scanln(&jatuhTempo)
			fmt.Print("Masukkan kategori langganan yang baru: ")
			fmt.Scanln(&kategori)
			fmt.Print("Masukkan metode pembayaran langganan yang baru: ")
			fmt.Scanln(&metode)

			daftarLangganan[i] = Langganan{
				Nama: nama, Biaya: biaya, JatuhTempo: jatuhTempo, Kategori: kategori, metodePembayaran: metode,
			}
			fmt.Println("Langganan berhasil diubah.")
			return
		}
	}

	fmt.Println("Nama langganan tidak ditemukan.")
}
func hapusLangganan() {
	var hapus string
	fmt.Print("Masukkan Nama Langganan yang ingin dihapus: ")
	fmt.Scanln(&hapus)

	for i := 0; i < len(daftarLangganan); i++ {
		if daftarLangganan[i].Nama == hapus {
			daftarLangganan = append(daftarLangganan[:i], daftarLangganan[i+1:]...)
			fmt.Println("Langganan berhasil dihapus.")
			return
		}
	}

	fmt.Println("Nama langganan tidak ditemukan.")
}
func totalPengeluaran() {
	var total int
	for i := 0; i < len(daftarLangganan); i++ {
		total += daftarLangganan[i].Biaya
	}
	fmt.Printf("Total pengeluaran per bulan: %d\n", total)
}
func cariNamaLangganan() {
	var input string
	fmt.Print("Cari nama: ")
	fmt.Scanln(&input)

	var i int
	var d Langganan

	for i = 0; i < len(daftarLangganan); i++ {
		d = daftarLangganan[i]
		if d.Nama == input {
			fmt.Printf("Data ditemukan: Nama: %s\n", d.Nama)
			fmt.Printf("Biaya: %d\n", d.Biaya)
			fmt.Printf("Jatuh Tempo: %s\n", d.JatuhTempo)
			fmt.Printf("Kategori: %s\n", d.Kategori)
			fmt.Printf("Metode Pembayaran: %s\n", d.metodePembayaran)
			return
		}
	}
	fmt.Println("Tidak ditemukan.")
}
func cariBerdasarkanKategori() {
	if len(daftarLangganan) == 0 {
		fmt.Println("Tidak ada data langganan.")
		return
	}
	for i := 0; i < len(daftarLangganan)-1; i++ {
		min := i
		for j := i + 1; j < len(daftarLangganan); j++ {
			if daftarLangganan[j].Kategori < daftarLangganan[min].Kategori {
				min = j
			}
		}
		daftarLangganan[i], daftarLangganan[min] = daftarLangganan[min], daftarLangganan[i]
	}

	fmt.Print("Masukkan kategori yang dicari: ")
	var cariKategori string
	fmt.Scanln(&cariKategori)
	kiri, kanan := 0, len(daftarLangganan)-1
	idx := -1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarLangganan[tengah].Kategori == cariKategori {
			idx = tengah
			break
		} else if daftarLangganan[tengah].Kategori < cariKategori {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if idx == -1 {
		fmt.Println("Kategori tidak ditemukan.")
		return
	}
	start, end := idx, idx
	for start > 0 && daftarLangganan[start-1].Kategori == cariKategori {
		start--
	}
	for end < len(daftarLangganan)-1 && daftarLangganan[end+1].Kategori == cariKategori {
		end++
	}

	fmt.Printf("Ditemukan %d data dengan kategori \"%s\":\n", end-start+1, cariKategori)
	for i := start; i <= end; i++ {
		d := daftarLangganan[i]
		fmt.Printf("- %s (Rp%d), Jatuh Tempo: %s, Metode: %s\n", d.Nama, d.Biaya, d.JatuhTempo, d.metodePembayaran)
	}
}

func sortirLanggananBiaya() {
	if len(daftarLangganan) == 0 {
		fmt.Println("tidak ada data langganan")
		return
	}
	for i := 0; i < len(daftarLangganan)-1; i++ {
		BiayaTerbesar := i
		for j := i + 1; j < len(daftarLangganan); j++ {
			if daftarLangganan[j].Biaya > daftarLangganan[BiayaTerbesar].Biaya {
				BiayaTerbesar = j
			}
		}
		daftarLangganan[i], daftarLangganan[BiayaTerbesar] = daftarLangganan[BiayaTerbesar], daftarLangganan[i]
	}

	fmt.Println("Daftar langganan berdasarkan biaya terbesar:")
	tampilkanLangganan()
}
func sortirJatuhTempoTercepat() {
	if len(daftarLangganan) == 0 {
		fmt.Println("Tidak ada data langganan.")
		return
	}

	for i := 1; i < len(daftarLangganan); i++ {
		j := i
		for j > 0 && daftarLangganan[j-1].JatuhTempo > daftarLangganan[j].JatuhTempo {
			daftarLangganan[j], daftarLangganan[j-1] = daftarLangganan[j-1], daftarLangganan[j]
			j--
		}
	}

	fmt.Println("Daftar langganan berdasarkan jatuh tempo tercepat:")
	tampilkanLangganan()
}

func rekomendasiPenghentian() {
	if len(daftarLangganan) == 0 {
		fmt.Println("tidak ada data langganan")
		return
	}

	var rekomendasi Langganan
	for i := 0; i < len(daftarLangganan); i++ {
		if daftarLangganan[i].Biaya > rekomendasi.Biaya {
			rekomendasi = daftarLangganan[i]
		}
	}
	fmt.Printf("Rekomendasi penghentian langganan: %s (Biaya: %d, Kategori: %s)\n",
		rekomendasi.Nama, rekomendasi.Biaya, rekomendasi.Kategori)
}
func tampilkanLangganan() {
	if len(daftarLangganan) == 0 {
		fmt.Println("tidak ada data langganan")
		return
	}

	fmt.Println("=====================================================================================")
	fmt.Println("| Nama Langganan  |  Biaya  | Jatuh Tempo |   Kategori          | Metode Pembayaran |")
	fmt.Println("|-----------------|--------|------------|--------------------|----------------------|")

	for i := 0; i < len(daftarLangganan); i++ {
		langganan := daftarLangganan[i]

		fmt.Printf("| %-15s | %-7d | %-12s | %-20s | %-18s |\n",
			langganan.Nama,
			langganan.Biaya,
			langganan.JatuhTempo,
			langganan.Kategori,
			langganan.metodePembayaran)
	}

	fmt.Println("=======================================================================================")

	layout := "02-01-2006"
	now := time.Now()

	for i := 0; i < len(daftarLangganan); i++ {
		langganan := daftarLangganan[i]

		if jatuhTempo, err := time.Parse(layout, langganan.JatuhTempo); err == nil {
			selisih := int(jatuhTempo.Sub(now).Hours() / 24)

			if selisih < 0 {
				fmt.Printf("Langganan %s sudah jatuh tempo!\n", langganan.Nama)
			} else if selisih == 0 {
				fmt.Printf("Langganan %s akan jatuh tempo hari ini!\n", langganan.Nama)
			} else if selisih <= 3 {
				fmt.Printf("Langganan %s akan jatuh tempo dalam %d hari lagi\n", langganan.Nama, selisih)
			}
		}
	}
}

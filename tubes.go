package main
import "fmt"

const NMAX int = 100

type Film struct {
	judul     string
	genre     string
	tahun     int
	deskripsi string
	rating    float64
}

type Koleksi struct {
	data [NMAX]Film
	n    int
}

func BacaInput() string {
	var res string
	var c byte
	var done bool

	done = false
	for !done {
		c = 0
		fmt.Scanf("%c", &c)
		if c == '\n' {
			if res != "" {
				done = true
			}
		} else {
			if c != '\r' && c != 0 {
				res += string(c)
			}
		}
	}
	return res
}

func tampilkanFilm(f Film) {
	fmt.Println("------------------------------------------------")
	fmt.Printf("Judul     : %s\n", f.judul)
	fmt.Printf("Genre     : %s\n", f.genre)
	fmt.Printf("Tahun     : %d\n", f.tahun)
	fmt.Printf("Rating    : %.1f/10\n", f.rating)
	fmt.Printf("Deskripsi : %s\n", f.deskripsi)
}

func cetakKoleksi(k Koleksi) {
	if k.n == 0 {
		fmt.Println("Koleksi film kosong.")
	} else {
		var i int
		for i = 0; i < k.n; i++ {
			tampilkanFilm(k.data[i])
		}
		fmt.Println("------------------------------------------------")
	}
}

func tambahFilm(k *Koleksi) {
	if k.n >= NMAX {
		fmt.Println("Gagal menambahkan film, memori penuh!")
		return
	}

	var f Film
	fmt.Print("Masukkan Judul Film        : ")
	f.judul = BacaInput()
	fmt.Print("Masukkan Genre Film        : ")
	f.genre = BacaInput()
	fmt.Print("Masukkan Tahun Rilis       : ")
	fmt.Scan(&f.tahun)
	fmt.Print("Masukkan Skor Rating (0-10): ")
	fmt.Scan(&f.rating)
	fmt.Print("Masukkan Deskripsi Singkat : ")
	f.deskripsi = BacaInput()

	k.data[k.n] = f
	k.n++
	fmt.Println("\nFilm berhasil ditambahkan ke koleksi!")
}

func cariJudulSequential(k Koleksi, judul string) int {
	var idx int
	var i int
	idx = -1
	for i = 0; i < k.n && idx == -1; i++ {
		if k.data[i].judul == judul {
			idx = i
		}
	}
	return idx
}

func ubahFilm(k *Koleksi) {
	var judul string
	var idx int
	fmt.Print("Masukkan Judul Film yang ingin diubah: ")
	judul = BacaInput()

	idx = cariJudulSequential(*k, judul)
	if idx == -1 {
		fmt.Println("Film tidak ditemukan!")
	} else {
		fmt.Println("\nFilm ditemukan! Silakan masukkan data baru:")
		fmt.Print("Masukkan Judul Baru        : ")
		k.data[idx].judul = BacaInput()
		fmt.Print("Masukkan Genre Baru        : ")
		k.data[idx].genre = BacaInput()
		fmt.Print("Masukkan Tahun Rilis Baru  : ")
		fmt.Scan(&k.data[idx].tahun)
		fmt.Print("Masukkan Skor Rating Baru  : ")
		fmt.Scan(&k.data[idx].rating)
		fmt.Print("Masukkan Deskripsi Baru    : ")
		k.data[idx].deskripsi = BacaInput()
		fmt.Println("\nData film berhasil diperbarui!")
	}
}

func hapusFilm(k *Koleksi) {
	var judul string
	var idx int
	fmt.Print("Masukkan Judul Film yang ingin dihapus: ")
	judul = BacaInput()

	idx = cariJudulSequential(*k, judul)
	if idx == -1 {
		fmt.Println("Film tidak ditemukan!")
	} else {
		var i int
		for i = idx; i < k.n-1; i++ {
			k.data[i] = k.data[i+1]
		}
		k.n--
		fmt.Println("\nFilm berhasil dihapus dari koleksi!")
	}
}

func urutGenreAscending(k *Koleksi) {
	var i int
	var j int
	var key Film
	for i = 1; i < k.n; i++ {
		key = k.data[i]
		j = i - 1
		for j >= 0 && k.data[j].genre > key.genre {
			k.data[j+1] = k.data[j]
			j--
		}
		k.data[j+1] = key
	}
}

func cariGenreBinaryRecursive(data [NMAX]Film, low, high int, genre string) int {
	if low > high {
		return -1
	}
	var mid int
	mid = (low + high) / 2
	if data[mid].genre == genre {
		return mid
	} else if data[mid].genre > genre {
		return cariGenreBinaryRecursive(data, low, mid-1, genre)
	} else {
		return cariGenreBinaryRecursive(data, mid+1, high, genre)
	}
}

func urutRatingSelection(k *Koleksi) {
	var i int
	var j int
	var maxIdx int
	var temp Film
	for i = 0; i < k.n-1; i++ {
		maxIdx = i
		for j = i + 1; j < k.n; j++ {
			if k.data[j].rating > k.data[maxIdx].rating {
				maxIdx = j
			}
		}
		temp = k.data[i]
		k.data[i] = k.data[maxIdx]
		k.data[maxIdx] = temp
	}
}

func urutTahunInsertion(k *Koleksi) {
	var i int
	var j int
	var key Film
	for i = 1; i < k.n; i++ {
		key = k.data[i]
		j = i - 1
		for j >= 0 && k.data[j].tahun > key.tahun {
			k.data[j+1] = k.data[j]
			j--
		}
		k.data[j+1] = key
	}
}

func tampilkanStatistik(k Koleksi) {
	if k.n == 0 {
		fmt.Println("Koleksi film kosong. Tidak ada statistik tersedia.")
		return
	}

	var genres [NMAX]string
	var counts [NMAX]int
	var R int
	var totalRating float64
	var i int
	var j int
	var film Film
	var found bool
	var idx int
	var rataRata float64

	R = 0
	totalRating = 0.0

	for i = 0; i < k.n; i++ {
		film = k.data[i]
		totalRating += film.rating

		found = false
		idx = 0
		for j = 0; j < R && !found; j++ {
			if genres[j] == film.genre {
				found = true
				idx = j
			}
		}

		if found {
			counts[idx]++
		} else {
			genres[R] = film.genre
			counts[R] = 1
			R++
		}
	}

	fmt.Println("\n================================================")
	fmt.Println("             STATISTIK KOLEKSI FILM             ")
	fmt.Println("================================================")
	fmt.Println("Jumlah film berdasarkan kategori genre:")
	for i = 0; i < R; i++ {
		fmt.Printf("- %s: %d film\n", genres[i], counts[i])
	}
	rataRata = totalRating / float64(k.n)
	fmt.Printf("\nRata-rata skor rating dari seluruh film: %.2f/10\n", rataRata)
	fmt.Println("================================================")
}

func main() {
	var k Koleksi
	var running bool
	var pilihan int
	k.n = 0

	k.data[k.n].judul = "Spiderman"
	k.data[k.n].genre = "Action"
	k.data[k.n].tahun = 2026
	k.data[k.n].deskripsi = "Seorang remaja berkekuatan super melawan kejahatan."
	k.data[k.n].rating = 9.0
	k.n++

	k.data[k.n].judul = "Barbie"
	k.data[k.n].genre = "Fantasi"
	k.data[k.n].tahun = 2023
	k.data[k.n].deskripsi = "Barbie mencari jati diri di dunia nyata."
	k.data[k.n].rating = 7.0
	k.n++

	k.data[k.n].judul = "Mean Girls"
	k.data[k.n].genre = "Drama"
	k.data[k.n].tahun = 2024
	k.data[k.n].deskripsi = "Drama persahabatan dan popularitas di sekolah menengah."
	k.data[k.n].rating = 8.0
	k.n++

	running = true

	for running {
		fmt.Println("\n====================================")
		fmt.Println("         MENU UTAMA CINEREVIEW      ")
		fmt.Println("====================================")
		fmt.Println("1. Tambah Koleksi Film")
		fmt.Println("2. Ubah Data Film")
		fmt.Println("3. Hapus Data Film")
		fmt.Println("4. Lihat Semua Koleksi Film")
		fmt.Println("5. Pencarian Film (Judul/Genre)")
		fmt.Println("6. Urutkan Koleksi Film")
		fmt.Println("7. Lihat Statistik Koleksi")
		fmt.Println("0. Keluar Aplikasi")
		fmt.Println("====================================")
		fmt.Print("Pilih Opsi Menu (0-7): ")

		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahFilm(&k)
		} else if pilihan == 2 {
			ubahFilm(&k)
		} else if pilihan == 3 {
			hapusFilm(&k)
		} else if pilihan == 4 {
			fmt.Println("\n=== DAFTAR SELURUH KOLEKSI FILM ===")
			cetakKoleksi(k)
		} else if pilihan == 5 {
			var opt int
			fmt.Println("\n--- METODE PENCARIAN ---")
			fmt.Println("1. Cari Berdasarkan Judul ")
			fmt.Println("2. Cari Berdasarkan Genre ")
			fmt.Print("Pilih metode (1-2): ")
			fmt.Scan(&opt)

			if opt == 1 {
				var judul string
				var idx int
				fmt.Print("Masukkan Judul Film yang dicari: ")
				judul = BacaInput()
				idx = cariJudulSequential(k, judul)
				if idx == -1 {
					fmt.Println("Film tidak ditemukan.")
				} else {
					fmt.Println("\nFilm Berhasil Ditemukan:")
					tampilkanFilm(k.data[idx])
				}
			} else if opt == 2 {
				var genre string
				var idx int
				fmt.Print("Masukkan Kategori Genre yang dicari: ")
				genre = BacaInput()

				urutGenreAscending(&k)

				idx = cariGenreBinaryRecursive(k.data, 0, k.n-1, genre)
				if idx == -1 {
					fmt.Println("Film dengan genre tersebut tidak ditemukan.")
				} else {
					var left int
					var right int
					var i int
					fmt.Println("\nFilm dengan genre tersebut ditemukan:")
					left = idx
					for left > 0 && k.data[left-1].genre == genre {
						left--
					}
					right = idx
					for right < k.n-1 && k.data[right+1].genre == genre {
						right++
					}
					for i = left; i <= right; i++ {
						tampilkanFilm(k.data[i])
					}
				}
			} else {
				fmt.Println("Opsi pilihan salah.")
			}
		} else if pilihan == 6 {
			var opt int
			fmt.Println("\n--- OPSI PENGURUTAN ---")
			fmt.Println("1. Urutkan Rating: Tertinggi ke Terendah ")
			fmt.Println("2. Urutkan Tahun Rilis: Terlama ke Terbaru ")
			fmt.Print("Pilih opsi pengurutan (1-2): ")
			fmt.Scan(&opt)

			if opt == 1 {
				urutRatingSelection(&k)
				fmt.Println("\nKoleksi film diurutkan berdasarkan rating tertinggi!")
				cetakKoleksi(k)
			} else if opt == 2 {
				urutTahunInsertion(&k)
				fmt.Println("\nKoleksi film diurutkan berdasarkan tahun rilis terlama!")
				cetakKoleksi(k)
			} else {
				fmt.Println("Opsi pilihan salah.")
			}
		} else if pilihan == 7 {
			tampilkanStatistik(k)
		} else if pilihan == 0 {
			fmt.Println("\nTerima kasih telah menggunakan aplikasi CineReview!")
			running = false
		} else {
			fmt.Println("Pilihan menu tidak valid, silakan ulangi.")
		}
	}
}
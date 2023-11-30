package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Bandara struct untuk menyimpan informasi bandara
type Bandara struct {
	Nama  string
	Tarif int
}

// Pesawat struct untuk menyimpan informasi pemesanan tiket
type Pesawat struct {
	NamaPemesan      string
	AsalBandara      string
	TujuanBandara    string
	PulangPergi      string
	Tanggal          string
	TanggalPulang    string
	JumlahPenumpang  int
	JenisTiket       string
	MetodePembayaran string
}

func main() {
	// Data referensi bandara
	bandaraCGK := Bandara{"Bandara CGK", 700000}
	bandaraJOG := Bandara{"Bandara JOG", 800000}
	bandaraSRG := Bandara{"Bandara SRG", 900000}

	bandaraList := []Bandara{bandaraCGK, bandaraJOG, bandaraSRG}

	// Input informasi dari pengguna
	fmt.Print("Masukkan Nama Pemesan: ")
	var namaPemesan string
	fmt.Scanln(&namaPemesan)

	for i, bandara := range bandaraList {
		fmt.Printf("%d. %s\n", i+1, bandara.Nama)
	}
	fmt.Print("Pilih Asal Bandara: ")

	var asalBandaraIndex int
	fmt.Scanln(&asalBandaraIndex)
	asalBandara := bandaraList[asalBandaraIndex-1].Nama

	for i, bandara := range bandaraList {
		fmt.Printf("%d. %s\n", i+1, bandara.Nama)
	}
	fmt.Print("Pilih Tujuan Bandara: ")

	var tujuanBandaraIndex int
	fmt.Scanln(&tujuanBandaraIndex)
	tujuanBandara := bandaraList[tujuanBandaraIndex-1].Nama

	fmt.Print("Pilih Pergi (P), Pulang (PL), atau Pulang Pergi (PP): ")
	var pulangPergi string
	fmt.Scanln(&pulangPergi)

	var tanggal, tanggalPulang string
	if pulangPergi == "PL" || pulangPergi == "PP" {
		fmt.Print("Masukkan Tanggal Pulang (Hari/Bulan/Tahun): ")
		fmt.Scanln(&tanggalPulang)
	}
	fmt.Print("Masukkan Tanggal Pergi (Hari/Bulan/Tahun): ")
	fmt.Scanln(&tanggal)

	fmt.Print("Masukkan Jumlah Penumpang: ")
	var jumlahPenumpang int
	fmt.Scanln(&jumlahPenumpang)

	fmt.Println("1. Ekonomi")
	fmt.Println("2. Bisnis")
	fmt.Println("3. First Class")
	fmt.Print("Pilih Jenis Tiket: ")
	var jenisTiketIndex int
	fmt.Scanln(&jenisTiketIndex)
	jenisTiket := map[int]string{1: "Ekonomi", 2: "Bisnis", 3: "First Class"}[jenisTiketIndex]

	fmt.Print("Pilih Metode Pembayaran (Transfer/Kredit): ")
	var metodePembayaran string
	fmt.Scanln(&metodePembayaran)

	// Hitung total biaya berdasarkan jenis penerbangan
	tarifAsal := bandaraList[asalBandaraIndex-1].Tarif
	tarifTujuan := bandaraList[tujuanBandaraIndex-1].Tarif
	totalBiaya := 0

	switch pulangPergi {
	case "P":
		totalBiaya = tarifTujuan * jumlahPenumpang
	case "PL":
		totalBiaya = tarifAsal * jumlahPenumpang
	case "PP":
		totalBiaya = (tarifAsal + tarifTujuan) * jumlahPenumpang
	}

	// Diskon 2% untuk metode pembayaran kredit
	if metodePembayaran == "Kredit" {
		diskon := float64(totalBiaya) * 0.02
		totalBiaya -= int(diskon)
	}

	// Generate kode unik 3 digit
	rand.Seed(time.Now().UnixNano())
	kodeUnik := rand.Intn(1000)

	// Tampilkan informasi pemesanan
	fmt.Println("\nInformasi Pemesanan:")
	fmt.Printf("Nama Pemesan: %s\n", namaPemesan)
	fmt.Printf("Asal Bandara: %s\n", asalBandara)
	fmt.Printf("Tujuan Bandara: %s\n", tujuanBandara)
	fmt.Printf("Pulang/Pergi: %s\n", map[string]string{"P": "Pergi", "PL": "Pulang", "PP": "Pulang Pergi"}[pulangPergi])
	fmt.Printf("Tanggal Pergi: %s\n", tanggal)
	fmt.Printf("Tanggal Pulang: %s\n", tanggalPulang)
	fmt.Printf("Jumlah Penumpang: %d\n", jumlahPenumpang)
	fmt.Printf("Jenis Tiket: %s\n", jenisTiket)
	fmt.Printf("Metode Pembayaran: %s\n", metodePembayaran)
	fmt.Printf("Total Biaya: Rp. %d%d\n", totalBiaya, kodeUnik)
}

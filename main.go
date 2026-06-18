package main

import "fmt"

const NMAX = 1000

type Skripsi struct {
	NIM         int
	Nama       string
	Judul      string
	Pembimbing string
	Topik      string
	Status     string
	Tahun      int
	NoUrut		int 
}

type daftar [NMAX]Skripsi

var dataSkripsi daftar

func main() {
	var pilihan, jumlahData int

	jumlahData = 0
	pilihan = 0

	for pilihan != 8 {

		fmt.Println("+==========================================+")
		fmt.Println("|             SkripsiLn                    |")
		fmt.Println("|      SISTEM INVENTARIS SKRIPSI           |")
		fmt.Println("+==========================================+")
		fmt.Println("| 1. Tambah Data Skripsi                   |")
		fmt.Println("| 2. Lihat Data Skripsi                    |")
		fmt.Println("| 3. Edit Data Skripsi                     |")
		fmt.Println("| 4. Hapus Data Skripsi                    |")
		fmt.Println("| 5. Cari Data                             |")
		fmt.Println("| 6. Urutkan Data                          |")
		fmt.Println("| 7. Statistik Skripsi                     |")
		fmt.Println("| 8. Exit                                  |")
		fmt.Println("+==========================================+")

		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			dataSkripsi, jumlahData = tambahData(jumlahData)

		case 2:
			tampilData(jumlahData)

		case 3:
			dataSkripsi = editData(jumlahData)

		case 4:
			dataSkripsi, jumlahData = hapusData(jumlahData)

		case 5:
    		cariData(jumlahData)

		case 6:
			urutData (jumlahData)

		case 7:
			statistik(jumlahData)

		case 8:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tambahData(jumlahData int) (daftar, int) {

	fmt.Println("\n+==========================================+")
	fmt.Println("|           TAMBAH DATA SKRIPSI            |")
	fmt.Println("+==========================================+")

	if jumlahData >= NMAX {
		fmt.Println("Penyimpanan Penuh!")
		return dataSkripsi, jumlahData
	}

	dataSkripsi[jumlahData].NoUrut = jumlahData + 1
	
	fmt.Print("penggunaan spasi diggantikan dengan '_'\n")

	fmt.Print("NIM Mahasiswa : ")
	fmt.Scan(&dataSkripsi[jumlahData].NIM)

	fmt.Print("Nama Mahasiswa : ")
	fmt.Scan(&dataSkripsi[jumlahData].Nama)

	fmt.Print("Judul Skripsi  : ")
	fmt.Scan(&dataSkripsi[jumlahData].Judul)

	fmt.Print("Pembimbing     : ")
	fmt.Scan(&dataSkripsi[jumlahData].Pembimbing)

	fmt.Print("Topik          : ")
	fmt.Scan(&dataSkripsi[jumlahData].Topik)

	fmt.Print("Status         : ")
	fmt.Scan(&dataSkripsi[jumlahData].Status)

	fmt.Print("Tahun          : ")
	fmt.Scan(&dataSkripsi[jumlahData].Tahun)

	jumlahData++

	fmt.Println("Data berhasil ditambahkan.")
	return dataSkripsi, jumlahData
}

func tampilData(jumlahData int) {
	var i, idx, pass, j int
	var tempDaftar daftar
	var temp Skripsi
	
	for i = 0; i < jumlahData; i++{
		tempDaftar[i] = dataSkripsi[i]
	}
	
	for pass = 1; pass < jumlahData; pass++{
		idx = pass-1
		for j = pass; j < jumlahData; j++{
			if tempDaftar[idx].NoUrut > tempDaftar[j].NoUrut{
				idx = j
			}
		}
		temp = tempDaftar[pass-1]
		tempDaftar[pass-1] = tempDaftar[idx]
		tempDaftar[idx] = temp
	}
	
	
	fmt.Println("\n+===============================================================+")
	fmt.Println("|                     DATA SKRIPSI                              |")
	fmt.Println("+===============================================================+")

	if jumlahData == 0 {
		fmt.Println("Belum ada data.")
		fmt.Println("+---------------------------------------------------------------+")
	}else {
		for i = 0; i < jumlahData; i++ {

		fmt.Println("NIM         :", tempDaftar[i].NIM)
		fmt.Println("Nama        :", tempDaftar[i].Nama)
		fmt.Println("Judul       :", tempDaftar[i].Judul)
		fmt.Println("Pembimbing  :", tempDaftar[i].Pembimbing)
		fmt.Println("Topik       :", tempDaftar[i].Topik)
		fmt.Println("Status      :", tempDaftar[i].Status)
		fmt.Println("Tahun       :", tempDaftar[i].Tahun)

		fmt.Println("+---------------------------------------------------------------+")
		}
	}
}

func lihatData(jumlahData int) {
	var i int
	
	fmt.Println("\n+===============================================================+")
	fmt.Println("|                     DATA SKRIPSI                              |")
	fmt.Println("+===============================================================+")

	if jumlahData == 0 {
		fmt.Println("Belum ada data.")
		fmt.Println("+---------------------------------------------------------------+")
	}else {
		for i = 0; i < jumlahData; i++ {

		fmt.Println("NIM         :", dataSkripsi[i].NIM)
		fmt.Println("Nama        :", dataSkripsi[i].Nama)
		fmt.Println("Judul       :", dataSkripsi[i].Judul)
		fmt.Println("Pembimbing  :", dataSkripsi[i].Pembimbing)
		fmt.Println("Topik       :", dataSkripsi[i].Topik)
		fmt.Println("Status      :", dataSkripsi[i].Status)
		fmt.Println("Tahun       :", dataSkripsi[i].Tahun)

		fmt.Println("+---------------------------------------------------------------+")
		}
	}
}

func editData(jumlahData int) daftar {

	var nomor, index, i int

	tampilData(jumlahData)

	if jumlahData == 0 {
		return dataSkripsi
	}

	fmt.Print("Pilih NIM data yang ingin diubah: ")
	fmt.Scan(&nomor)

	index = -1
	for i = 0; i < jumlahData ; i++{
		if dataSkripsi[i].NIM == nomor{
			index = i
		}
	}

	if index == -1{
		fmt.Println("NIM Tidak Ditemukan")
		return dataSkripsi
	}
	

	fmt.Println("\n+==========================================+")
	fmt.Println("|           MASUKKAN DATA BARU             |")
	fmt.Println("+==========================================+")

	fmt.Print("Nama Mahasiswa : ")
	fmt.Scan(&dataSkripsi[index].Nama)

	fmt.Print("Judul Skripsi  : ")
	fmt.Scan(&dataSkripsi[index].Judul)

	fmt.Print("Pembimbing     : ")
	fmt.Scan(&dataSkripsi[index].Pembimbing)

	fmt.Print("Topik          : ")
	fmt.Scan(&dataSkripsi[index].Topik)

	fmt.Print("Status         : ")
	fmt.Scan(&dataSkripsi[index].Status)

	fmt.Print("Tahun          : ")
	fmt.Scan(&dataSkripsi[index].Tahun)
	
	lihatData(jumlahData)

	fmt.Println("Data Berhasil Diperbarui")

	

	return dataSkripsi
}

func hapusData(jumlahData int) (daftar, int) {

	var nomor, index, i int

	tampilData(jumlahData)

	if jumlahData == 0 {
		return dataSkripsi, jumlahData
	}

	fmt.Print("Pilih NIM Data Yang Ingin Dihapus: ")
	fmt.Scan(&nomor)

	index = -1
	for i = 0; i < jumlahData ; i++{
		if dataSkripsi[i].NIM == nomor{
			index = i
		}
	}

	if index == -1{
		fmt.Println("NIM Tidak Ditemukan")
		return dataSkripsi, jumlahData
	}

	for i = index; i < jumlahData-1; i++ {
		dataSkripsi[i] = dataSkripsi[i+1]
	}

	jumlahData--
	lihatData(jumlahData)

	fmt.Println("Data Berhasil Dihapus")

	return dataSkripsi, jumlahData

	}
func cariData(jumlahData int) {

	var keyword string
	var pilihan, pilihmetode int

	if jumlahData == 0 {
		fmt.Println("\n+==========================================+")
		fmt.Println("| Belum ada data.                          |")
		fmt.Println("+==========================================+")
	}else{

		fmt.Println("\n+==========================================+")
		fmt.Println("|               CARI DATA                  |")
		fmt.Println("+==========================================+")
		fmt.Println("| 1. Cari Berdasarkan Nama                 |")
		fmt.Println("| 2. Cari Berdasarkan Judul                |")
		fmt.Println("+==========================================+")

		fmt.Print("Pilih metode pencarian: ")
		fmt.Scan(&pilihan)
		
		fmt.Println("\n+==========================================+")
		fmt.Println("|               CARI DATA                  |")
		fmt.Println("+==========================================+")
		fmt.Println("| 1. Cari Menggunakan Sequential Search    |")
		fmt.Println("| 2. Cari Menggunaka Binary Search         |")
		fmt.Println("+==========================================+")

		fmt.Print("Pilih metode pencarian: ")
		fmt.Scan(&pilihmetode)
		
		if pilihan == 1 {
			fmt.Print("Masukkan Nama Mahasiswa: ")
			fmt.Scan(&keyword)
		}else if pilihan == 2{
			fmt.Print("Masukkan Judul Skripsi: ")
			fmt.Scan(&keyword)
		}else{
			fmt.Println("Pilihan Tidak valid.")
		}
		
		if pilihan == 1{
			if pilihmetode == 1{
				sequentialNama(jumlahData, keyword)
				
			}else if pilihmetode == 2{
				binaryNama(jumlahData, keyword)
			}else{
				fmt.Println("Pilihan Tidak valid.")
			}
		}else if pilihan == 2{
			if pilihmetode == 1{
				sequentialJudul(jumlahData, keyword)
			}else if pilihmetode == 2{
				binaryJudul(jumlahData, keyword)
			}else{
				fmt.Println("Pilihan Tidak valid.")
			} 
		}	
	}
}

func sequentialNama(jumlahData int, keyword string){
	var i int
	var ditemukan bool
	
	ditemukan = false
	for i = 0; i < jumlahData; i++ {
		if dataSkripsi[i].Nama == keyword {
			fmt.Println("\nData Ditemukan")
			fmt.Println("NIM          :", dataSkripsi[i].NIM)
			fmt.Println("Nama        :", dataSkripsi[i].Nama)
			fmt.Println("Judul       :", dataSkripsi[i].Judul)
			fmt.Println("Pembimbing  :", dataSkripsi[i].Pembimbing)
			fmt.Println("Topik       :", dataSkripsi[i].Topik)
			fmt.Println("Status      :", dataSkripsi[i].Status)
			fmt.Println("Tahun       :", dataSkripsi[i].Tahun)
			ditemukan = true
		}
	}
	if ditemukan == false{
		fmt.Println("Data Tidak Ditemukan")
	}
	
}

func sequentialJudul(jumlahData int, keyword string){
	var i int
	var ditemukan bool
	
	ditemukan = false
	for i = 0; i < jumlahData; i++ {
		if dataSkripsi[i].Judul == keyword {
			fmt.Println("\nData Ditemukan")
			fmt.Println("NIM          :", dataSkripsi[i].NIM)
			fmt.Println("Nama        :", dataSkripsi[i].Nama)
			fmt.Println("Judul       :", dataSkripsi[i].Judul)
			fmt.Println("Pembimbing  :", dataSkripsi[i].Pembimbing)
			fmt.Println("Topik       :", dataSkripsi[i].Topik)
			fmt.Println("Status      :", dataSkripsi[i].Status)
			fmt.Println("Tahun       :", dataSkripsi[i].Tahun)
			ditemukan = true
		}
	}
	if ditemukan == false{
		fmt.Println("Data Tidak Ditemukan")
	}
	
}

func binaryNama(jumlahData int, keyword string){
	var left, right, mid, found, i, start, end int
	
	selectionNamaAscending(jumlahData)
	left = 0
	right = jumlahData-1
	found = -1
			
	for left<= right && found == -1{
		mid = (left + right)/ 2
		if keyword < dataSkripsi[mid].Nama{
			right = mid -1
		}else if keyword > dataSkripsi[mid].Nama{
			left = mid + 1
		}else{
			found = mid
		}
	}
			
	if found != -1{
		start = found
		for start > 0 && dataSkripsi[start-1].Nama == keyword{
			start--
		}
		end = found
		for end < jumlahData-1 && dataSkripsi[end+1].Nama == keyword{
			end++
		}
		for i = start; i <= end; i++{
			fmt.Println("\nData Ditemukan")
			fmt.Println("NIM          :", dataSkripsi[i].NIM)
			fmt.Println("Nama        :", dataSkripsi[i].Nama)
			fmt.Println("Judul       :", dataSkripsi[i].Judul)
			fmt.Println("Pembimbing  :", dataSkripsi[i].Pembimbing)
			fmt.Println("Topik       :", dataSkripsi[i].Topik)
			fmt.Println("Status      :", dataSkripsi[i].Status)
			fmt.Println("Tahun       :", dataSkripsi[i].Tahun)	
		}
	}else{
		fmt.Println("Data Tidak Ditemukan")
	}
}

func binaryJudul(jumlahData int, keyword string){
	var left, right, mid, found, i, start, end int
	
	selectionJudulAscending(jumlahData)
	left = 0
	right = jumlahData-1
	found = -1
			
	for left<= right && found == -1{
		mid = (left + right)/ 2
		if keyword < dataSkripsi[mid].Judul{
			right = mid -1
		}else if keyword > dataSkripsi[mid].Judul{
			left = mid + 1
		}else{
			found = mid
		}
	}
			
	if found != -1{
		start = found
		for start > 0 && dataSkripsi[start-1].Judul == keyword{
			start--
		}
		end = found
		for end < jumlahData-1 && dataSkripsi[end+1].Judul == keyword{
			end++
		}
		for i = start; i <= end; i++{
			fmt.Println("\nData Ditemukan")
			fmt.Println("NIM          :", dataSkripsi[i].NIM)
			fmt.Println("Nama        :", dataSkripsi[i].Nama)
			fmt.Println("Judul       :", dataSkripsi[i].Judul)
			fmt.Println("Pembimbing  :", dataSkripsi[i].Pembimbing)
			fmt.Println("Topik       :", dataSkripsi[i].Topik)
			fmt.Println("Status      :", dataSkripsi[i].Status)
			fmt.Println("Tahun       :", dataSkripsi[i].Tahun)	
		}	
	}else{
		fmt.Println("Data Tidak Ditemukan")
	}
}

func urutData(jumlahData int){

	var pilihan, pilihmetode, pilihurutan int

	if jumlahData == 0 {
		fmt.Println("Belum ada data.")
	}else{

		fmt.Println("\n+==========================================+")
		fmt.Println("|            URUTKAN DATA                  |")
		fmt.Println("+==========================================+")
		fmt.Println("| 1. Urutkan Berdasarkan Nama              |")
		fmt.Println("| 2. Urutkan Berdasarkan Tahun             |")
		fmt.Println("+==========================================+")

		fmt.Print("Pilih metode pengurutan: ")
		fmt.Scan(&pilihan)
		
		fmt.Println("\n+==========================================+")
		fmt.Println("|            URUTKAN DATA                  |")
		fmt.Println("+==========================================+")
		fmt.Println("| 1. Urutkan Menggunakan Selection Sort    |")
		fmt.Println("| 2. Urutkan Menggunakan Insertion Sort    |")
		fmt.Println("+==========================================+")

		fmt.Print("Pilih metode pengurutan: ")
		fmt.Scan(&pilihmetode)

		fmt.Println("\n+==========================================+")
		fmt.Println("|            URUTKAN DATA                  |")
		fmt.Println("+==========================================+")
		fmt.Println("| 1. Urutkan Secara Ascending              |")
		fmt.Println("| 2. Urutkan Secara Descending             |")
		fmt.Println("+==========================================+")

		fmt.Print("Pilih metode pengurutan: ")
		fmt.Scan(&pilihurutan)
		
		if pilihan == 1 {
			if pilihmetode == 1{
				if pilihurutan == 1{
					selectionNamaAscending(jumlahData)
					
					fmt.Println("Data berhasil diurutkan berdasarkan Nama Menggunakan Selection Sort Secara Ascending.")
				}else if pilihurutan == 2{
					selectionNamaDescending(jumlahData)
					fmt.Println("Data berhasil diurutkan berdasarkan Nama Menggunakan Selection Sort Secara Descending.")
				}else{
					fmt.Println("Pilihan Tidak valid.")
				}
			}else if pilihmetode == 2{
				if pilihurutan == 1{
					insertionNamaAscending(jumlahData)
					fmt.Println("Data berhasil diurutkan berdasarkan Nama Menggunakan Insertion Sort Secara Ascending.")
				}else if pilihurutan == 2{
					insertionNamaDescending(jumlahData)
					fmt.Println("Data berhasil diurutkan berdasarkan Nama Menggunakan Insertion Sort Secara Descending.")
				}else{
					fmt.Println("Pilihan Tidak valid.")
				}
			}else{
				fmt.Println("Pilihan Tidak valid.")
			}
		}else if pilihan == 2{
			if pilihmetode == 1{
				if pilihurutan == 1{
					selectionTahunAscending(jumlahData)
					fmt.Println("Data berhasil diurutkan berdasarkan Tahun Menggunakan Selection Sort Secara Ascending.")
				}else if pilihurutan == 2{
					selectionTahunDescending(jumlahData)
					fmt.Println("Data berhasil diurutkan berdasarkan Tahun Menggunakan Selection Sort Secara Descending.")
				}else{
					fmt.Println("Pilihan Tidak valid.")
				}
			}else if pilihmetode == 2{
				if pilihurutan == 1{
					insertionTahunAscending(jumlahData)
					fmt.Println("Data berhasil diurutkan berdasarkan Tahun Menggunakan Insertion Sort Secara Ascending.")
				}else if pilihurutan == 2{
					insertionTahunDescending(jumlahData)
					fmt.Println("Data 2berhasil diurutkan berdasarkan Tahun Menggunakan Insertion Sort Secara Descending.")
				}else{
					fmt.Println("Pilihan Tidak valid.")
				}
			}else{
				fmt.Println("Pilihan Tidak valid.")
			}
		}else{
			fmt.Println("Pilihan Tidak valid.")
		}
		
		lihatData(jumlahData)
	}
}

func selectionJudulAscending(jumlahData int){
	var pass, i, idx int
	var temp Skripsi
	
	pass = 1
	for pass <= jumlahData-1{
		idx = pass-1
		i = pass
		for i < jumlahData{
			if dataSkripsi[idx].Judul > dataSkripsi[i].Judul{
				idx = i
			}
			i++
		}
		temp = dataSkripsi[pass-1]
		dataSkripsi[pass-1]= dataSkripsi[idx]
		dataSkripsi[idx]= temp
		pass++
	}
}

func selectionNamaAscending(jumlahData int){
	var pass, i, idx int
	var temp Skripsi
	
	pass = 1
	for pass <= jumlahData-1{
		idx = pass-1
		i = pass
		for i < jumlahData{
			if dataSkripsi[idx].Nama > dataSkripsi[i].Nama{
				idx = i
			}
			i++
		}
		temp = dataSkripsi[pass-1]
		dataSkripsi[pass-1]= dataSkripsi[idx]
		dataSkripsi[idx]= temp
		pass++
	}
}

func selectionNamaDescending(jumlahData int){
	var pass, i, idx int
	var temp Skripsi
	
	pass = 1
	for pass <= jumlahData-1{
		idx = pass-1
		i = pass
		for i < jumlahData{
			if dataSkripsi[idx].Nama < dataSkripsi[i].Nama{
				idx = i
			}
			i++
		}
		temp = dataSkripsi[pass-1]
		dataSkripsi[pass-1]= dataSkripsi[idx]
		dataSkripsi[idx]= temp
		pass++
	}
}

func selectionTahunAscending(jumlahData int){
	var pass, i, idx int
	var temp Skripsi
	
	pass = 1
	for pass <= jumlahData-1{
		idx = pass-1
		i = pass
		for i < jumlahData{
			if dataSkripsi[idx].Tahun > dataSkripsi[i].Tahun{
				idx = i
			}
			i++
		}
		temp = dataSkripsi[pass-1]
		dataSkripsi[pass-1]= dataSkripsi[idx]
		dataSkripsi[idx]= temp
		pass++
	}
}

func selectionTahunDescending(jumlahData int){
	var pass, i, idx int
	var temp Skripsi
	
	pass = 1
	for pass <= jumlahData-1{
		idx = pass-1
		i = pass
		for i < jumlahData{
			if dataSkripsi[idx].Tahun < dataSkripsi[i].Tahun{
				idx = i
			}
			i++
		}
		temp = dataSkripsi[pass-1]
		dataSkripsi[pass-1]= dataSkripsi[idx]
		dataSkripsi[idx]= temp
		pass++
	}
}

func insertionNamaAscending(jumlahData int){
	var pass, i int
	var temp Skripsi
	
	for pass = 1; pass < jumlahData; pass++ {
		temp = dataSkripsi[pass]
		i = pass - 1
		for i >= 0 && dataSkripsi[i].Nama > temp.Nama {
			dataSkripsi[i+1] = dataSkripsi[i]	
			i--
		}
		dataSkripsi[i+1] = temp
	}
}

func insertionNamaDescending(jumlahData int){
	var pass, i int
	var temp Skripsi
	
	for pass = 1; pass < jumlahData; pass++ {
		temp = dataSkripsi[pass]
		i = pass - 1
		for i >= 0 && dataSkripsi[i].Nama < temp.Nama {
			dataSkripsi[i+1] = dataSkripsi[i]	
			i--
		}
		dataSkripsi[i+1] = temp
	}
}

func insertionTahunAscending(jumlahData int){
	var pass, i int
	var temp Skripsi
	
	for pass = 1; pass < jumlahData; pass++ {
		temp = dataSkripsi[pass]
		i = pass - 1
		for i >= 0 && dataSkripsi[i].Tahun > temp.Tahun {
			dataSkripsi[i+1] = dataSkripsi[i]	
			i--
		}
		dataSkripsi[i+1] = temp
	}
}

func insertionTahunDescending(jumlahData int){
	var pass, i int
	var temp Skripsi
	
	for pass = 1; pass < jumlahData; pass++ {
		temp = dataSkripsi[pass]
		i = pass - 1
		for i >= 0 && dataSkripsi[i].Tahun < temp.Tahun {
			dataSkripsi[i+1] = dataSkripsi[i]	
			i--
		}
		dataSkripsi[i+1] = temp
	}
}


func statistik(jumlahData int) {
	var tahunlist [NMAX] int
	var jumlahpertahun [NMAX] int
	var tahun, tipetahun, i, j int
	var temptahun, tempjumlah, pass, idx int
	var ditemukan bool
	
	tipetahun = 0
	fmt.Println("\n+==========================================+")
    fmt.Println("|               STATISTIK SKRIPSI          |")
    fmt.Println("+==========================================+")
	
	if jumlahData == 0{
		fmt.Println("|            Belum ada data!               |")
        fmt.Println("+==========================================+")
	}else{
		
		for i = 0; i < jumlahData; i++ {
        tahun = dataSkripsi[i].Tahun
        ditemukan = false
			for j = 0; j < tipetahun; j++ {
				if tahunlist[j] == tahun {
					jumlahpertahun[j]++
					ditemukan = true
				}
			}
			if !ditemukan {
				tahunlist[tipetahun] = tahun
				jumlahpertahun[tipetahun] = 1
				tipetahun++
			}
		}
		for pass = 1; pass < tipetahun; pass++ {
			idx = pass - 1
			for i = pass; i < tipetahun; i++ {
				if tahunlist[idx] > tahunlist[i] {
					idx = i
				}
			}
			temptahun = tahunlist[pass-1]
			tempjumlah = jumlahpertahun[pass-1]
			tahunlist[pass-1] = tahunlist[idx]
			jumlahpertahun[pass-1] = jumlahpertahun[idx]
			tahunlist[idx] = temptahun
			jumlahpertahun[idx] = tempjumlah
		}
		
		
		fmt.Println("\n| Statistik Per Tahun :                    |")
        fmt.Println("+==========================================+")
		
		for i = 0; i < tipetahun; i++{
			fmt.Printf("| Tahun %d : %d Judul                     |\n", tahunlist[i], jumlahpertahun[i])
		}
		
		fmt.Printf("| Total Dokumen Tersimpan : %d Data         |\n", jumlahData)
		fmt.Println("+==========================================+")
	}
}

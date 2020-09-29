# IKA SMAN Situraja

Website Ikatan Alumni SMA Negeri Situraja

## Plan Fitur

1. Kelola Acara, Pendaftaran Acara, dan Download Tiket Acara
2. Testimoni Kisah Sukses Alumni
3. Pendaftaran dan Verifikasi Data Alumni
4. Artikel / Berita
5. Chat Forum (Websocket?)
6. Notifikasi Pemberitahuan Pengumuman & Surat
7. Histori Data Penting

## BackEnd

Back-End API dibuat dengan framework Golang Fiber (https://github.com/gofiber/fiber) dan GORM (https://gorm.io/). 

Selain itu, mendukung Hot Reload (kompilasi otomatis) saat development dengan Air (https://github.com/cosmtrek/air).

Golang itu struktural programming, bukan OOP, jadi jangan terlalu buat modul generik (nanti 2021 Golang rilis resmi generic).

### Cara Jalankan BackEnd

1. Masuk direktori proyek BackEnd

2. Buat Database Kosongan pakai MySQL

3. Jalankan BackEnd
  Mode Biasa : ```go run backend```
  Hot Reload : ```air```
  
4. Kompilasi Backend 
  Ubah OS Target : ```set GOOS=windows``` atau ```set GOOS=linux```
  Kompilasi      : ```go run build``` 
  Binary hasil kompilasi bisa dideploy


## FrontEnd

FrontEnd dibuat dengan ReactJS HTML+CSS biasa (tampilan awal, tanpa framework bootstrap/material-ui), dan React-Admin untuk bagian Admin (https://marmelab.com/react-admin/)

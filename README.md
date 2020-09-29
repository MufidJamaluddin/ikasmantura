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

Untuk download PDF pakai WKHTMLTOPDF (wrapper nya: https://github.com/SebastiaanKlippert/go-wkhtmltopdf, file instalan-nya: https://wkhtmltopdf.org/), bisa nggak include pas development.

Selain itu, mendukung Hot Reload (kompilasi otomatis saat file diubah) saat development dengan Air (https://github.com/cosmtrek/air).


### Cara Jalankan BackEnd

1. Masuk direktori proyek BackEnd

2. Buat Database Kosongan pakai MySQL

3. Jalankan BackEnd
  Mode Biasa : ```go run backend```
  Hot Reload : ```air```
  
4. Kompilasi

  ```go build```
  
### Deployment BackEnd

1. Tanpa Docker
  
  Deploy manual saja binary hasil kompilasi
  
2. Dengan Docker

  Buat Docker Image.
  Contoh : ```docker build --tag alumni_backend:1.0 .```


## FrontEnd

FrontEnd dibuat dengan ReactJS HTML+CSS biasa (tampilan awal, tanpa framework bootstrap/material-ui), dan React-Admin untuk bagian Admin (https://marmelab.com/react-admin/)

### Cara Jalankan FrontEnd

Frontend dibuat dengan instalan NodeJS (v12.18.x) dan sebelumnya udah tersedia NPM (v6.14.x) atau Yarn (2.x).

1. Install Library yang Dibutuhkan

  ```yarn install```

2. Jalankan Aplikasi dan Edit Kode (Hot Reload)
 
  ```yarn start```
  
3. Build Aplikasi, nanti jadi asset statik di folder ```/build``` dan siap deploy tanpa NodeJS.

  ```yarn build```
  
### Deployment FrontEnd

1. Tanpa Docker

  Deploy semua isi folder /build ke File Server / Apache / Nginx sebagai static files.

2. Dengan Docker

  Buat Docker Image
  Contoh : ```docker build --tag alumni_frontend:1.0 .```

<!-- INTRO -->
<br />
<div align="center">
  <h1 align="center">Tugas Besar 2 IF2211 Strategi Algoritma Tahun Ajaran 2023/2024</h1>

  <p align="center">
    <h3> Pemanfaatan Algoritma IDS dan BFS dalam Permainan WikiRace </h3>
    <p>Program made using Go Language</p>
    <br />
    <a href="https://github.com/zultopia/Tubes2_Stima.git">Report Bug</a>
    ·
    <a href="https://github.com/zultopia/Tubes2_Stima.git">Request Feature</a>
<br>
<br>

[![MIT License][license-shield]][license-url]

  </p>
</div>

<!-- CONTRIBUTOR -->
<div align="center" id="contributor">
  <strong>
    <h3>Dibuat oleh Kelompok "GO Blocks" :</h3>
    <table align="center">
      <tr>
        <td>NIM</td>
        <td>Nama</td>
      </tr>
      <tr>
        <td>Aurelius Justin Philo Fanjaya</td>
        <td>13522020</td>
     </tr>
     <tr>
        <td>Marzuli Suhada M</td>
        <td>13522070</td>
    </tr>
    <tr>
        <td>Fedrianz Dharma</td>
        <td>13522090</td>
    </tr>
    </table>
  </strong>
</div>

## Deskripsi Program

WikiRace atau Wiki Game adalah permainan yang melibatkan Wikipedia, sebuah ensiklopedia daring gratis yang dikelola oleh berbagai relawan di dunia, dimana pemain mulai pada suatu artikel Wikipedia dan harus menelusuri artikel-artikel lain pada Wikipedia (dengan mengeklik tautan di dalam setiap artikel) untuk menuju suatu artikel lain yang telah ditentukan sebelumnya dalam waktu paling singkat atau klik (artikel) paling sedikit.

Dokumentasi lengkap tentang program dapat dilihat pada [link berikut](https://docs.google.com/document/d/1ngvbhIkR53FhFmPfbCIbx-_Uy2z9ultUqlJMn2pOEmw/edit)
   
## Penjelasan Algoritma

## 1. Algoritma BFS

Ketika menggunakan algoritma BFS, misalkan mulai dari artikel Joko Widodo dan artikel tujuannya adalah Elon Musk. Kita mulai dengan mengunjungi artikel Joko Widodo dan melakukan scraping pada artikel Joko Widodo untuk mendapat semua link artikel yang ada pada artikel Joko Widodo tersebut. Setiap link yang didapatkan akan dicatat pada sebuah map dan dimasukkan ke dalam rute saat ini, kemudian dimasukkan (enqueue) ke dalam sebuah queue. Lalu, kita akan mengunjungi semua link terakhir pada rute yang ada pada queue dan melakukan scraping secara concurrent. Hasil scraping tersebut akan dimasukkan pada rute saat ini dan dimasukkan pada sebuah queue baru jika link yang didapatkan belum ada pada map.

Setiap queue merepresentasikan semua rute dari link awal hingga semua link/simpul yang ada pada satu level. Proses akan dilakukan terus menerus hingga mendapatkan solusi rute dari artikel awal hingga tujuan. Jika dilakukan pencarian single solution, maka proses pencarian akan dihentikan dan langsung mengembalikan solusi yang didapatkan. Pada contoh kasus artikel awal Joko Widodo dan artikel tujuan Elon Musk, solusi yang didapatkan adalah Joko Widodo, List of international presidential trips made by Joko Widodo, Elon Musk.  Namun, jika dilakukan pencarian dengan multiple solution, maka proses pencarian akan dilanjutkan pada level tersebut hingga semua rute pada queue tersebut selesai diperiksa untuk mendapatkan alternatif solusi lainnya. Pada contoh kasus akan ditemukan 8 solusi lainnya, salah satunya adalah Joko Widodo, The New York Times, Elon Musk.

## 2. Algoritma IDS

Algoritma IDS kami implementasikan dengan menggunakan algoritma DLS dengan maximum depth yang bertambah pada tiap iterasi IDS, dimulai dari depth 1. Kami mengimplementasikan algoritma DLS menggunakan pendekatan rekursif sebagai berikut. Misalkan kasus pencarian dengan artikel awal Joko Widodo dan artikel tujuan Elon Musk. Pertama, pada fungsi DLS artikel Joko Widodo akan dicek apakah sudah ada di Map atau belum, jika sudah maka akan mengambil link-link yang menjadi value di map tersebut. Jika belum, maka artikel Joko Widodo akan dikunjungi dan dilakukan scraping untuk mendapat semua link artikel yang ada pada artikel tersebut dan dimasukkan ke Map dengan key nama artikel (dalam kasus ini key nya adalah Joko Widodo). 

Setiap link yang didapatkan akan dicek apakah sudah sesuai dengan artikel tujuan, jika belum maka akan dengan memanggil fungsi DLS secara rekursif dengan start artikel tersebut dan tujuan tetap sama hingga menyentuh limit depth atau artikel ditemukan. Jika pada tiap path DLS tidak ditemukan solusinya sampai depth maximum, maka akan dilakukan pencarian ulang DLS dengan depth maximum bertambah 1 pada iterasi IDS. Jika dilakukan pencarian single solution, maka proses pencarian akan dihentikan ketika pertama kali ditemukan solution. Jika dilakukan pencarian multiple solution, maka proses pencarian akan dilanjutkan iterasi IDS tersebut hingga selesai untuk mendapatkan solusi lainnya. 

## Setup and Installation

1. Clone repo

```
git clone git@github.com:zultopia/Tubes2_GO-Blocks.git
```

2. Compile the program

```
cd src/frontend
yarn
yarn build
```

3. Run the program

```
cd ../backend
go run .
```

4. Open in Browser

```
http://localhost:3000/
```

<!-- LICENSE -->
## Licensing

The code in this project is licensed under MIT license.  
Code dalam projek ini berada di bawah lisensi MIT.

<br>
<h3 align="center"> TERIMA KASIH! </h3>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/zultopia/Tubes2_Stima/blob/main/LICENSE
# Technical Test GITS Software Engineer Intern - BE (Go)

# Muhammad Rieza Fachrezi

## Penjelasan Kompleksitas Soal No. 3
Kompleksitas pada solusi soal No.3 adalah sebagai berikut:
### Kompleksitas Waktu: O(n)
Kompleksitas waktu adalah **O(n)** dikarenakan fungsi melakukan satu iterasi dengan panjang string n pada loop `for i := 0; i < len(s); i++`. Pengecekan masing-masing karakter tepat sekali dan juga operasi pada stack, yaitu push & pop, yang dilakukan dalam waktu konstan O(1). Pada fungsi juga tidak terdapat nested loop sehingga menjadikan keseluruhan fungsi tersebut memiliki kompleksitas waktu **O(n)**
### Kompleksitas Ruang: O(n)
Kompleksitas ruang untuk fungsi tersebut adalah **O(n)** dikarenakan fungsi harus menyimpan semua bracket pembuka yang ada pada stack karakter sampai nantinya ditemukan pasangan bracket penutup yang valid. Skenario terburuk adalah ketika input string semua merupakan bracket pembuka sehingga semua n karakter pada string harus disimpan dalam stack.

**Oleh karena itu, kompleksitas waktu dan ruang fungsi tersebut masing-masing adalah O(n).**
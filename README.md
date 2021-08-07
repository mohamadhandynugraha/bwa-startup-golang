# bwa-startup-golang
this is bwa startup golang API for backend only
belajar konsep "layering" pada golang

// contoh layering yang biasa di golang
// handler panggil -> service (bisnis proses) panggil -> repository panggil -> db

kemungkinan nanti akan ada folder user, campaign, campaign_image, transaction. Setiap folder itu terdapat layering yang di comment. Masing-masing folder itu "melambangkan" entity
dan setiap entity mempunyai handler, service, repository. 

Nanti juga terdapat folder handler.

Misal contoh API register account. Langkah pertama user akan input terlebih dahulu seperti name, email dll. Langkah kedua handler akan menangkap mapping input tersebut supaya bisa dikenali oleh golang disini mapping input ke struct. Langkah ketiga, struct yang di langkah kedua di passing ke service ini berarti sudah masuk bisnis prosesnya yaitu mendaftarkan user (mengambil object mapping tadi), struct pada langkah ketiga ini akan di passing menjadi struct User. Langkah keempat save struct User ke db. langkah ke 5 disimpan di database.

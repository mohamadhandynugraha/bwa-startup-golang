# bwa-startup-golang
this is bwa startup golang API for backend only
belajar konsep "layering" pada golang

// contoh layering yang biasa di golang
// handler panggil -> service (bisnis proses) panggil -> repository panggil -> db

kemungkinan nanti akan ada folder user, campaign, campaign_image, transaction. Setiap folder itu terdapat layering yang di comment. Masing-masing folder itu "melambangkan" entity
dan setiap entity mempunyai handler, service, repository. 

Nanti juga terdapat folder handler untuk. 

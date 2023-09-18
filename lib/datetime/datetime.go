package datetime

import "time"

/*
	DateTimeNow()
	Dipakai saat create atau update data,
	digunakan untuk mengisi kolom DateIn & DateUP
*/
func DateTimeNow() string {
	return time.Now().UTC().Format("2006-01-02 15:04:05")
}

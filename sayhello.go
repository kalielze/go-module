package gomodule

import "strconv"

func SayHello() string {
	return "Belajar module golang"
}

func Materi(version int) string {
	return "Module golang" + strconv.Itoa(version)
}

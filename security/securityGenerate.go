package main

import (
	auth "github.com/abbot/go-http-auth"
	"log"
)

func main() {
	salt := []byte("YeNsbWdH")
	magic := []byte("$1$")
	log.Printf("%s %s %s", auth.MD5Crypt([]byte("password"), salt, magic), salt, magic)
}

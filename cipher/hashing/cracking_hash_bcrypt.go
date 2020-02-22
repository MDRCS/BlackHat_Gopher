package main

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
)

var storedHash = "$2a$10$Zs3ZwsjV/nF.KuvSUE.5WuwtDrK6UVXcBpQrH84V8q3Opg1yNdWLu"

func main() {
	var password string
	if len(os.Args) != 2 {
		log.Fatalln("Usage: bcrypt password")
	}
	password = os.Args[1]
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("hash = %s\n", hash)
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		log.Println("[!] Authentication failed")
		return
	}
	log.Println("[+] Authentication successful")
}


//go run cracking_hash_bcrypt.go someC0mpl3xP@ssw0rd
//hash = $2a$10$W.ebSN57XEz2fY/iZRJHLOJSBgTocXy1IUkcFNllzwGp2.KxFxf7a // you can observe that hash of the password is diffrent from storedhash, answer is that the same hash + salt just to make sure that ther is no collision.
//[+] Authentication successful

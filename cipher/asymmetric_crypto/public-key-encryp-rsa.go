package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"hash"
	"log"
)

type Block struct {
	err												 error
	privateKey										 *rsa.PrivateKey
	publicKey										 *rsa.PublicKey
	message, plaintext, ciphertext, signature, label []byte
}


func (blk *Block) GenerateKeys(keylength int) {
	if blk.privateKey,blk.err = rsa.GenerateKey(rand.Reader,keylength) ; blk.err != nil {
		log.Fatalln(blk.err)
	}
	blk.publicKey = &blk.privateKey.PublicKey
}

func (blk *Block) EncryptRSA() {
	if blk.ciphertext,blk.err = rsa.EncryptOAEP(sha256.New(),rand.Reader,blk.publicKey,blk.message,blk.label); blk.err != nil {
		log.Fatalln(blk.err)
	}

}

func (blk *Block) DecryptRSA() {
	if blk.plaintext,blk.err = rsa.DecryptOAEP(sha256.New(),rand.Reader,blk.privateKey,blk.ciphertext,blk.label); blk.err != nil {
		log.Fatalln(blk.err)
	}
}

func (blk *Block) GetSignature(hash hash.Hash) {
	if blk.signature,blk.err = rsa.SignPSS(rand.Reader,blk.privateKey,crypto.SHA256,hash.Sum(nil),nil) ; blk.err != nil {
		log.Fatalln(blk.err)
	}
}


func (blk *Block) VerifySign(hash hash.Hash) {
	if err := rsa.VerifyPSS(blk.publicKey,crypto.SHA256,hash.Sum(nil),blk.signature,nil); err != nil {
		log.Fatalln(err)
	}
}
func main() {

	blk := Block{}

	//Generate Public & Private keys with a length of 2048 bits
	blk.GenerateKeys(2048)

	//Message that we want to communicate
	blk.label = []byte("")
	blk.message = []byte("Some super secret message, maybe a session key even")

	//Encrypt the message using public key so the only way to decrypt it is using private key
	blk.EncryptRSA()
	fmt.Printf("ciphertext : %x\n",blk.ciphertext)

	//Decrypt the message using private key to get a plaintext
	blk.DecryptRSA()
	fmt.Printf("plaintext : %x\n",blk.plaintext)

	h := sha256.New()
	h.Write(blk.message)

	//Get a signature using private key
	blk.GetSignature(h)
	fmt.Printf("signature : %x\n",blk.signature)

	//Verify the signature using public key for the integrity of the message ..
	blk.VerifySign(h)
	fmt.Println("Signature verified\n")

}

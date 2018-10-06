package crypt

import (
    "fmt"
    "io"
    "crypto/rand"

    "golang.org/x/crypto/scrypt"
)

type HashStruct struct {
  Hash         string
  Salt         string
	}

const (
    PW_SALT_BYTES = 32
    PW_HASH_BYTES = 64
)

func Hash(u HashStruct) (HASH, SALT string) {
    var salt []byte
    if u.Salt == "" {
      salt = make([]byte, PW_SALT_BYTES)
      _, err := io.ReadFull(rand.Reader, salt)
      if err != nil {
        fmt.Println("Error generating salt")
        fmt.Println(err)
      }
    } else {
      salt = []byte(u.Salt)
    }


    hash, err := scrypt.Key([]byte(u.Hash), salt, 1<<14, 8, 1, PW_HASH_BYTES)
    if err != nil {
        fmt.Println("Error while hashing")
        fmt.Println(err)
    }

      //  converts to string so the values can be returned and stored
    has := string(hash[:64])
    sal := string(salt[:32])

    /*
        fmt.Println("\n\n")
        fmt.Printf("%x\n", hash)    //  hash in readable format
        fmt.Printf("%x\n", u.Salt)  //  salt in readable format
        fmt.Println(has)            //  hash in string format, not readble
        fmt.Println(salt)           //  salt in string format, not readble
        fmt.Println("\n\n")

    */

    return  has, sal
}

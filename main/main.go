package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func RandPass(passwordLength int)string {
	var (
		lowerCaseChar  = "abcdefghijklmnopqrstuvwxyz"
		upperCaseChar  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		specialCharSet = "!@#$%&*"
		numberSet 	   = "0123456789"
		completeCharSet = lowerCaseChar + upperCaseChar + specialCharSet + numberSet
		password strings.Builder
	)
	rand.Seed(time.Now().Unix())
	specialCharMin := 1
	upperCaseMin := 1
	numericMin := 3

	//Set Special Char
	for i := 0; i<specialCharMin; i++ {
		randomSpecialChar := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[randomSpecialChar]))
	}

	//Set Uppercase Char
	for i := 0; i<upperCaseMin; i++ {
		randomUpperCase := rand.Intn(len(upperCaseChar))
		password.WriteString(string(upperCaseChar[randomUpperCase]))
	}

	//Set Numeric
	for i :=0; i<numericMin; i++ {
		randomNumeric := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[randomNumeric]))
	}

	//Check remaining length left
	remainingSize := passwordLength - specialCharMin - upperCaseMin - numericMin

	//Set remaining char for password
	for i := 0; i<remainingSize; i++ {
		random := rand.Intn(len(completeCharSet))
		password.WriteString(string(completeCharSet[random]))
	}
	psw := []rune(password.String())
	rand.Shuffle(len(psw), func(i, j int) {
		psw[i], psw[j] = psw[j], psw[i]
	})
	return string(psw)
}

func main() {
	var (
		input int
		size  int
	)
	fmt.Println("Nama  : Pande Putu Devo Punda Maheswara")
	fmt.Println("NIM   : 2008561107                     ")
	fmt.Println("Kelas : B                              ")
	for true {
		fmt.Println("Menu Program")
		fmt.Println("1. Buat random password")
		fmt.Println("2. Cari password asli")
		fmt.Println("3. Tampilkan riwayat password")
		fmt.Print("Input : ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			return 
		}

		switch input {
		case 1:
			fmt.Print("Masukkan panjang password :")
			_, err := fmt.Scanln(&size)
			if err != nil {
				return 
			}
			password := RandPass(size)
			fmt.Println("Password :", password)
			break
		}
	}
}

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// HashFunc -> to convert password value into unique code
// that stored in Hash Table later on.
func HashFunc(passValue int)string {
	charSet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%&*"
	var tmp int
	var hashValue string
	hashValue = ""
	for passValue != 0 {
		tmp = passValue % len(charSet)
		passValue = passValue / 3
		hashValue += string(charSet[tmp])
	}
	return hashValue
}

// InsertToHashTable -> using map to store unique code
// and value of the code that is the password
func InsertToHashTable(hashTable* map[string]string, hashedPassword* string, realPassword string) {
	(*hashTable)[*hashedPassword] = realPassword
}

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
		uniqueCode string
	)
	uniqueCode = ""
	hashTablePassword := make(map[string]string)
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
			valuePassword := 0
			for _, v := range password {
				valuePassword += int(v)
			}
			hashedPass := HashFunc(valuePassword)
			fmt.Println("Berikut kode unik password anda :[",hashedPass,"]")
			InsertToHashTable(&hashTablePassword, &hashedPass, password)
			break
		case 2:
			fmt.Print("Ketik kode unik yang telah diberikan :")
			_, err := fmt.Scanln(&uniqueCode)
			if err != nil {
				return
			}
			if uniqueCode == "" {
				fmt.Println("Tidak boleh kosong!")
				break
			}
			if value, found := hashTablePassword[uniqueCode] ; found ==true {
				fmt.Println("Password asli :", value)
			}
			break
		case 3:
			fmt.Println("Daftar password yang telah dibuat")
			for key, value := range hashTablePassword {
				fmt.Println("Unique Code :",key,"Value :", value)
			}
		}
	}
}

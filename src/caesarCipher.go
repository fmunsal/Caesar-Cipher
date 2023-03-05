package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const lowerAlphabet string = "abcdefghijklmnopqrstuvwxyz"
const upperAlphabet string = "ABCEEFGHIJKLMNOPQRSTUVWXYZ"

func CaesarCipher() {
	menu()
}
func menu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("e:Encrypt\nd:Decypt\nChoose(q for quit): ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		choice = strings.TrimSpace(choice)
		switch choice {
		case "e", "E":
			plaintext := getMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			plaintext = strings.TrimSpace(plaintext)
			fmt.Println("Encypt:", encrypt(plaintext))

		case "d", "D":
			ciphertext := getMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			ciphertext = strings.TrimSpace(ciphertext)
			fmt.Println("Decypt:", decrypt(ciphertext))

		case "q", "Q":
			fmt.Println("Exiting the program")
			return

		default:
			fmt.Println("Wrong choise..")
		}
	}
}

// get message from input
func getMessage() (x string) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Get Message: ")
	input, _ := inputReader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	return input
}

// Encrypt function
func encrypt(text string) (newt string) {
	newText := ""
	for i := 0; i < len(text); i++ {
		if strings.Contains(lowerAlphabet, string(text[i])) {
			newText += string(lowerAlphabet[((strings.Index(lowerAlphabet, string(text[i])))+4)%26])
		} else if strings.Contains(upperAlphabet, string(text[i])) {
			newText += string(upperAlphabet[((strings.Index(upperAlphabet, string(text[i])))+4)%26])
		} else {
			newText += string(text[i])
		}
	}
	return newText
}

// Decrypt function
func decrypt(text string) (newt string) {
	newText := ""
	for i := 0; i < len(text); i++ {
		if strings.Contains(lowerAlphabet, string(text[i])) {
			newText += string(lowerAlphabet[((strings.Index(lowerAlphabet, string(text[i])))-4)%26])
		} else if strings.Contains(upperAlphabet, string(text[i])) {
			newText += string(upperAlphabet[((strings.Index(upperAlphabet, string(text[i])))-4)%26])
		} else {
			newText += string(text[i])
		}
	}
	return newText
}

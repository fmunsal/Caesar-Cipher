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
	choose := ""

	for (choose != "q") && (choose != "Q") {
		fmt.Print("1-Encrypt\n2-Decrpyt\nChoose(e,d or q for quit):")
		fmt.Scanln(&choose)
		if (choose == "E") || (choose == "e") {
			input := getMessage()
			fmt.Println("Encrpyt Text:", encrypt(input))
		} else if (choose == "D") || (choose == "d") {
			input := getMessage()
			fmt.Println("Decrpyt Text:", decrypt(input))
		}
	}
}

// get message from input
func getMessage() (x string) {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Get Message:")
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

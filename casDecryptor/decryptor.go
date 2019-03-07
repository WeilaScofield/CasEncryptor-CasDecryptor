package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Decryptor struct {
	EncryptAlg string
	Ciphertext string
	Key        int
}

func (en *Decryptor) getCiphertext(name string) (e *Decryptor, error int) {

	fileName := name
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("OPENING FAILED:", err)
		return en, 1
	}

	defer file.Close()

	bs := make([]byte, 1024*8, 1024*8)
	n := -1

	for {
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("READING FINISHED")
			fmt.Println()
			break
		}
		en.Ciphertext = fmt.Sprintln(string(bs[:n]))
	}

	return en, 0
}

func Decrypt(para int, en *Decryptor) (res string) {

	fmt.Println("CIPHERTEXT:", en.Ciphertext)

	sli := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	tab := map[string]string{
		sli[0+para]:  "a",
		sli[1+para]:  "b",
		sli[2+para]:  "c",
		sli[3+para]:  "d",
		sli[4+para]:  "e",
		sli[5+para]:  "f",
		sli[6+para]:  "g",
		sli[7+para]:  "h",
		sli[8+para]:  "i",
		sli[9+para]:  "j",
		sli[10+para]: "k",
		sli[11+para]: "l",
		sli[12+para]: "m",
		sli[13+para]: "n",
		sli[14+para]: "o",
		sli[15+para]: "p",
		sli[16+para]: "q",
		sli[17+para]: "r",
		sli[18+para]: "s",
		sli[19+para]: "t",
		sli[20+para]: "u",
		sli[21+para]: "v",
		sli[22+para]: "w",
		sli[23+para]: "x",
		sli[24+para]: "y",
		sli[25+para]: "z",
		" ":          " ",
		",":          ",",
		".":          ".",
		"?":          "?",
		"!":          "!",
		"'":          "'",
	}

	en.EncryptAlg = fmt.Sprintln(tab)

	s := []string{}

	for _, v := range en.Ciphertext {
		sl := fmt.Sprintf("%s", tab[string(v)])
		s = append(s, sl)
	}

	return strings.Join(s, "")
}

func main() {
	fmt.Println("WELCOME TO USE CAESAR DECRYPTOR")
	fmt.Println()
	en := Decryptor{}

	fmt.Print("SET KEY:")
	inputKey := bufio.NewScanner(os.Stdin)

	for inputKey.Scan() {
		para := inputKey.Text()
		fmt.Println("KEY:", para)
		if para != "" {
			n, _ := strconv.Atoi(para)
		RES1:
			fmt.Print("SET FILEPATH:")
			inputPath := bufio.NewScanner(os.Stdin)
			for inputPath.Scan() {

				name := inputPath.Text()
				fmt.Println("FILEPATH:", name)
				e, i := en.getCiphertext(name)
				if i != 0 {
					goto RES1
				}
				plaintext := Decrypt(n, e)

				fmt.Println("PLAINTEXT:", plaintext)
				fmt.Println()
				fmt.Println("ENCRYPTALG:", en.EncryptAlg)
				fmt.Println("DECRYPT COMPLETED")
				fmt.Println()
			RES2:
				fmt.Print("SET SAVEPATH:")
				inputSavePath := bufio.NewScanner(os.Stdin)

				for inputSavePath.Scan() {
					name := inputSavePath.Text()
					file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModePerm)
					if err != nil {
						fmt.Println("GENERATE FILE FAILED")
						goto RES2
					}
					_, err = file.WriteString(plaintext)
					if err != nil {
						fmt.Println("WRITING FAILED")
						return
					}
					fmt.Println()
					fmt.Println("SAVING COMPLETED")
					return
				}
			}
		}
	}

}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Encryptor struct {
	EncryptAlg string
	Plaintext  string
	Key        int
}

func (en *Encryptor) getPlaintext(name string) (e *Encryptor, error int) {

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
		en.Plaintext = fmt.Sprintln(string(bs[:n]))
	}

	return en, 0
}

func encrypt(para int, en *Encryptor) (res string) {

	fmt.Println("PLAINTEXT:", en.Plaintext)

	sli := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
		"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	tab := map[string]string{
		"a": sli[0+para],
		"b": sli[1+para],
		"c": sli[2+para],
		"d": sli[3+para],
		"e": sli[4+para],
		"f": sli[5+para],
		"g": sli[6+para],
		"h": sli[7+para],
		"i": sli[8+para],
		"j": sli[9+para],
		"k": sli[10+para],
		"l": sli[11+para],
		"m": sli[12+para],
		"n": sli[13+para],
		"o": sli[14+para],
		"p": sli[15+para],
		"q": sli[16+para],
		"r": sli[17+para],
		"s": sli[18+para],
		"t": sli[19+para],
		"u": sli[20+para],
		"v": sli[21+para],
		"w": sli[22+para],
		"x": sli[23+para],
		"y": sli[24+para],
		"z": sli[25+para],
		" ": " ",
		",": ",",
		".": ".",
		"?": "?",
		"!": "!",
		"'": "'",
	}

	en.EncryptAlg = fmt.Sprintln(tab)

	s := []string{}

	for _, v := range en.Plaintext {
		sl := fmt.Sprintf("%s", tab[string(v)])
		s = append(s, sl)
	}

	return strings.Join(s, "")
}

func main() {
	fmt.Println("WELCOME TO USE CAESAR ENCRYPTOR")
	fmt.Println()
	en := Encryptor{}

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
				e, i := en.getPlaintext(name)
				if i != 0 {
					goto RES1
				}
				ciphertext := encrypt(n, e)

				fmt.Println("CIPHERTEXT:", ciphertext)
				fmt.Println()
				fmt.Println("ENCRYPTALG:", en.EncryptAlg)
				fmt.Println("ENCRYPT COMPLETED")
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
					_, err = file.WriteString(ciphertext)
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

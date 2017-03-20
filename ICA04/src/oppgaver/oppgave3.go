package oppgaver

//////////////////////////////////
// 	ICA04, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////
import (
	"bytes"
	"fmt"
	"sort"
	"os"
	"log"
	"bufio"
)

func Oppgave3a() {
	fmt.Println("Ved bruk av os og ioutil pakkene kan man behandle filer. " +
		"I disse pakken finner vi metoder som kan gjøre det meste med filer" +
		"os.Create" +
		"os.Remove" +
		"os.Rename" +
		"os.OpenFile" +
		"os.Open" +
		"os.Chmod")
}
func Oppgave3b() {
	fmt.Println("Skriv inn filen du ønsker å scanne:")
	var filelenght = bufio.NewScanner(os.Stdin)
	var fileinput string = "Ingen fil valgt"
	for filelenght.Scan() {
		fileinput = string(filelenght.Text())
		break
	}

	fmt.Print("Antall linjeskift i filen: ")
	fmt.Println(findInText(fileinput, "\n"))

	// Henter map av filen
	m := countFile(fileinput)
	// Sorterer og printer resultat:
	sortAndPrint(m)

}
func Oppgave3c() {
	//
	//
	// HUSK Å ENDRE DENNE
	//
	//
	fmt.Println("Kjør benchmark på egne filer")
}

// Oppgave 3b (1/3)
// Finner tekst i fil.
func findInText(filename string, find string) int {
	// Last inn filen
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	// Stor buffer for å håndtere store filer
	buf := make([]byte, 32*1024)
	count := 0
	// Gjør om string til byte for søk
	search := []byte(find)
	// Leter etter tekst:
	c, _ := file.Read(buf)
	count += bytes.Count(buf[:c], search)
	return count
}

// Oppgave 3b (2/3)
// Lager en map med resultatet for hver rune, og returnerer denne
func countFile(filename string) map[int]string {
	m := make(map[int]string)
	count := 0
	// Runer for Ascii code valgt:
	for i := 0x20; i <= 0x7F; i++ {
		count = findInText(filename, string(i))
		rune := string(i)
		m[count] = rune
	}
	fmt.Println(count)
	return m
}

// Oppgave 3b (3/3)
// Sorterer map listen med append, og printer ut de fem største verdier.
func sortAndPrint(m map[int]string) {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	//// Print Alle:
	//for _, k := range keys {
	//	fmt.Println("Key:", k, "Value:", m[k])
	//}

	// Print 5 største:
	fmt.Println("5 mest brukte runes:")
	value1 := keys[len(keys)-1]
	value2 := keys[len(keys)-2]
	value3 := keys[len(keys)-3]
	value4 := keys[len(keys)-4]
	value5 := keys[len(keys)-5]
	fmt.Println("Antall:", value1, "Rune:", m[value1])
	fmt.Println("Antall:", value2, "Rune:", m[value2])
	fmt.Println("Antall:", value3, "Rune:", m[value3])
	fmt.Println("Antall:", value4, "Rune:", m[value4])
	fmt.Println("Antall:", value5, "Rune:", m[value5])
}

package main

//importa pacchetti
import (
	// "bufio"
	"encoding/json"
	"fmt"
	// "github.com/360EntSecGroup-Skylar/excelize"
	"io/ioutil"
	"log"
	// "math/rand"
	"os"
	// "strings"
	// "time"
)

//funzione di errore
func checkErrors(err error) {
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

//funzione per copiare il file excel dnd
func copyFile(src string, dst string) error {
	input, err := ioutil.ReadFile(src)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(dst, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

type razza struct {
	Elfo []struct {
		Maschio []string `json:"maschio,omitempty"`
		Femmina []string `json:"femmina,omitempty"`
	} `json:"elfo,omitempty"`
	Nano []struct {
		Maschio []string `json:"maschio,omitempty"`
		Femmina []string `json:"femmina,omitempty"`
	} `json:"nano,omitempty"`
}

type configurations struct {
	Classe       []string `json:"classe"`
	Razza        []string `json:"razza"`
	Genere       []string `json:"genere"`
	Allineamento []string `json:"allineamento"`
	Taglia       []string `json:"taglia"`
	Dio          []string `json:"dio"`
	Nomi         struct {
		Nano struct {
			Maschio []string `json:"maschio"`
			Femmina []string `json:"femmina"`
			Dubbio  []string `json:"dubbio"`
		} `json:"nano"`
		Elfo struct {
			Maschio []string `json:"maschio"`
			Femmina []string `json:"femmina"`
			Dubbio  []string `json:"dubbio"`
		} `json:"elfo"`
	} `json:"nomi"`
}

var Conf configurations

//ReadFromJSON function load a json file into a struct or return error
func ReadFromJSON(t interface{}, filename string) error {

	jsonFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonFile), t)
	if err != nil {
		return err
	}

	return nil
}

type Personaggio struct {
	Nome         string
	Utente       string
	Razza        string
	Genere       string
	Allineamento string
	Taglia       string
	Dio          string
	Classe       string
}

func (p *Personaggio) Genera(utente, genere string) error {
	p.Utente = utente
	p.Genere = genere
	p.Nome = "pippo"

	return nil
}

// func (*Personaggio) GeneraNome(Nome, Razza, Genere string) error {

// 	return nil
// }

func proponiDomanda( domanda string, valoreObiettivo *string, selezioniPossibili []string ) error {
	for *valoreObiettivo == "" {
		//chiede nel terminale la preferenza dell'utente
		fmt.Println(domanda)
		for i, valore := range selezioniPossibili {
			fmt.Println(i, ": ", valore)
		}
		fmt.Print("(0-X): ")

		//Legge quello che l'utente ha scelto
		var inputUtente int
		_, err := fmt.Scanf("%d", &inputUtente)
		checkErrors(err)

		for i, valore := range selezioniPossibili {
			if i == inputUtente {
				*valoreObiettivo = valore
				break
			}
		}

		if *valoreObiettivo == "" {
			fmt.Println("Valore inserito non riconosciuto")
		}
	}

	return nil
}

func init() {
	checkErrors(ReadFromJSON(&Conf, "conf.json"))
}

func main() {
	personaggio := new(Personaggio)

	checkErrors(proponiDomanda("inserisci qui la tua razza.", &personaggio.Razza, Conf.Razza))

	

	checkErrors(personaggio.Genera("io", "maschio"))
	log.Println(personaggio)

	// //copia il file dnd, producendo la schedaPersonaggio
	// checkErrors(copyFile("dnd.xlsx", "schedaPersonaggio.xlsx"))

	// //apre la schedaPersonaggio
	// schedaPersonaggio, err := excelize.OpenFile("./schedaPersonaggio.xlsx")
	// checkErrors(err)
	// //chiede nel terminale la preferenza dell'utente
	// razza := bufio.NewReader(os.Stdin)
	// fmt.Print("Inserisci qui la razza del tuo personaggio\n\r1: Elfo\r\n2: Nano\r\n3: Umano\r\n4: Halfling\r\n(1-4): ")
	// //Legge quello che l'utente ha scelto
	// razzaScelta, err := (razza.ReadString('\n'))
	// checkErrors(err)
	// //toglie lo spazio
	// razzaPresa := strings.TrimSpace(razzaScelta)

	// genere := bufio.NewReader(os.Stdin)
	// fmt.Print("Inserisci qui il sesso del tuo personaggio (Maschio, Femmina): ")
	// genereScelto, err := genere.ReadString('\n')
	// checkErrors(err)
	// generePreso := strings.TrimSpace(genereScelto)

	// nomeGiocatore := bufio.NewReader(os.Stdin)
	// fmt.Print("Inserisci qui il nome del giocatore: ")
	// nomeGiocatoreScelto, err := nomeGiocatore.ReadString('\n')
	// checkErrors(err)
	// nomeGiocatorePreso := strings.TrimSpace(nomeGiocatoreScelto)
	// //funzione random che utilizza l'orario del compueter per ottenere un numero intero
	// rand.Seed(time.Now().UnixNano())
	// z := Classe[rand.Intn(len(Classe))]
	// schedaPersonaggio.SetCellValue("Foglio1", "A3", z)
	// //ciclo switch che opera in base alla scelta dell'utente
	// switch razzaPresa {
	// case "Umano":
	// 	switch generePreso {
	// 	case "Maschio":
	// 		//funzione random sulla lunghezza dell'array
	// 		um := UmanoM[rand.Intn(len(UmanoM))]

	// 		//Importa valore sul file excel
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", um)
	// 		//schedaPersonaggio.AddPicture("Sheet1", "A10", "./drizzit_Guerr.jpg")
	// 		//checkErrors(err)
	// 	case "Femmina":
	// 		uf := UmanoF[rand.Intn(len(UmanoF))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", uf)
	// 	default:
	// 		log.Println("Il genere non è definito")
	// 	}
	// case "Elfo":
	// 	switch generePreso {
	// 	case "Maschio":
	// 		em := ElfoM[rand.Intn(len(ElfoM))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", em)

	// 	case "Femmina":
	// 		ef := ElfoF[rand.Intn(len(ElfoF))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", ef)
	// 	default:
	// 		log.Println("Il genere non è definito")
	// 	}

	// case "Nano":
	// 	switch generePreso {
	// 	case "Maschio":
	// 		nm := NanoM[rand.Intn(len(NanoM))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", nm)
	// 	case "Femmina":
	// 		nf := NanoF[rand.Intn(len(NanoF))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", nf)
	// 	default:
	// 		log.Println("Il genere non è definito")
	// 	}

	// case "Halfling":
	// 	switch generePreso {
	// 	case "Maschio":
	// 		hm := HalfM[rand.Intn(len(HalfM))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", hm)
	// 	case "Femmina":
	// 		hf := HalfF[rand.Intn(len(HalfF))]
	// 		schedaPersonaggio.SetCellValue("Foglio1", "A1", hf)
	// 	default:
	// 		log.Println("Il genere non è definito")
	// 	}

	// default:
	// 	log.Println("Niente Panico! Riprova!!")
	// }
	// //setting di altre scelte dell'utente e altre funzioni random
	// schedaPersonaggio.SetCellValue("Foglio1", "B1", nomeGiocatorePreso)
	// schedaPersonaggio.SetCellValue("Foglio1", "B3", 1)
	// schedaPersonaggio.SetCellValue("Foglio1", "A5", razzaPresa)
	// schedaPersonaggio.SetCellValue("Foglio1", "C7", generePreso)

	// p := Classe[rand.Intn(len(Taglia))]
	// schedaPersonaggio.SetCellValue("Foglio1", "D7", p)
	// a := Classe[rand.Intn(len(Allineamento))]
	// schedaPersonaggio.SetCellValue("Foglio1", "B5", a)
	// b := Classe[rand.Intn(len(Divinità))]
	// schedaPersonaggio.SetCellValue("Foglio1", "A7", b)
	// //salva scheda
	// filename := "./scheda_" + nomeGiocatorePreso
	// schedaPersonaggio.SaveAs(filename + ".xlsx")
	// checkErrors(err)
	// fmt.Println("La scheda è pronta! Ora puoi stamparla!")
}

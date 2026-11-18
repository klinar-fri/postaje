package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/klinar-fri/postaje/vreme"
	"github.com/urfave/cli/v3"
)

func narediInDodajPostaje(slovar_postaj map[string]vreme.Postaja, minTemp float64, maxTemp float64) {
	// vreme.DodajPostajo("LJ01", "Ljubljana FRI", "Ljubljana", []float64{22.4, 21.4, 18.5, 12.9, 10.7}, slovar_postaj)
	// vreme.DodajPostajo("LJ02",  "Ljubljana BTC", "Ljubljana",[]float64{21.4, 24.4, 19.5, 11.9, 11.7}, slovar_postaj)
	// vreme.DodajPostajo("KR01",  "Kranj center", "Kranj",[]float64{20.1, 22.8, 20.1, 12.9, 13.8}, slovar_postaj)
	// vreme.DodajPostajo("BL01",  "Bled jezero", "Bled", []float64{-10.8, 1.3, 2.2, 6.7, 5.7}, slovar_postaj)
	// vreme.DodajPostajo("BH01",  "Bohinj jezero", "Bohinjska Bistrica", []float64{10.8, 8.3, 13.2, 12.7, 15.7}, slovar_postaj)
	// vreme.DodajPostajo("PR01",  "Piran zaliv", "Piran", []float64{33.8, 28.3, 24.2, 21.7, 32.7}, slovar_postaj)
	// vreme.DodajPostajo("MBP01", "Mariborsko pohorje", "Maribor", []float64{4.3, 2.7}, slovar_postaj)


	oznake := []string{"LJ01", "LJ02", "KR01", "BL01", "BH01", "PR01", "MBP01"}
	imena := []string{"Ljubljana FRI", "Ljubljana BTC", "Kranj center", "Bled jezero", "Bohinj jezero", "Piran zaliv", "Mariborsko pohorje"}
	kraji := []string{"Ljubljana", "Ljubljana", "Kranj", "Bled", "Bohinjska Bistrica", "Piran", "Maribor"}

	meritve := [][]float64{
		{22.4, 21.4, 18.5, 12.9, 10.7},
		{21.4, 24.4, 19.5, 11.9, 11.7},
		{20.1, 22.8, 20.1, 12.9, 13.8},
		{-10.8, 1.3, 2.2, 6.7, 5.7},
		{10.8, 8.3, 13.2, 12.7, 15.7},
		{33.8, 28.3, 24.2, 21.7, 32.7},
		{4.3, 2.7},
	}


	for i := range oznake{
		vreme.DodajPostajo(oznake[i], imena[i], kraji[i], slovar_postaj)
		for j := range meritve[i] {
			if meritve[i][j] >= minTemp && meritve[i][j] <= maxTemp{
				vreme.DodajMeritev(slovar_postaj, oznake[i], meritve[i][j])
			}else{
				fmt.Printf("NAPAKA: meritev %.1f v postaji %s z oznako %s ni v intervalu (%.1f, %.1f)\n", meritve[i][j], imena[i], oznake[i], minTemp, maxTemp)
			}
		}
	}
}

func izpisiMeritveInPorocilo(slovar_postaj map[string]vreme.Postaja){
	vreme.IzpisMeritev(slovar_postaj)
	vreme.IzpisPorocila(slovar_postaj)
}
//
// Dodal sem izločanje meritev (min, max) glede na stikala že tukaj v mainu,
// da ostaneta funkciji DodajPostajo in DodajMeritev neodvisni od programa, ki uporablja paket vreme.
// Za funckijo IzpisPorocila, pa to ne moremo storiti tukaj, zato nastavimo javno spremenljivko iz vreme.go MinStMeritev
// Take spremenljivke bi lahko naredil tudi za minTemp in maxTemp ... sploh nisem na to pomislil.
// Vseeno sem jih dodal, da se lahko uporabljajo iz paketa, main sem pustil kot je.

func main(){
	cmd := &cli.Command{
		Name: "vremenske postaje",
		Usage: "Meritve na različnih vremenskih postajah.",
		Flags: []cli.Flag{
			&cli.FloatFlag{
				Name: "maxTemp",
				Usage: "Maksimalna temperatura meritve.",
				Value: 50,
			},
			&cli.FloatFlag{
				Name: "minTemp",
				Usage: "Minimalna temperatura meritve",
				Value: -50,
			},
			&cli.IntFlag{
				Name: "stMeritev",
				Usage: "Najmanjše število izmerjenih meritev.",
				Value: 3,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			minTemp :=  cmd.Float("minTemp")
			maxTemp :=  cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			vreme.MinStMeritev = stMeritev
			vreme.MaxTemp = maxTemp
			vreme.MinTemp = minTemp

			var slovar_postaj = make(map[string]vreme.Postaja)
			narediInDodajPostaje(slovar_postaj, minTemp, maxTemp)
			izpisiMeritveInPorocilo(slovar_postaj)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

/* IZPIS
.\main.exe --minTemp 0 --maxTemp 35 --stMeritev 5

NAPAKA: meritev -10.8 v postaji Bled jezero z oznako BL01 ni v intervalu (0.0, 35.0)
-----MERITVE-----:
PR01 - Piran zaliv (Piran): [33.8 28.3 24.2 21.7 32.7]
MBP01 - Mariborsko pohorje (Maribor): [4.3 2.7]
LJ01 - Ljubljana FRI (Ljubljana): [22.4 21.4 18.5 12.9 10.7]
LJ02 - Ljubljana BTC (Ljubljana): [21.4 24.4 19.5 11.9 11.7]
KR01 - Kranj center (Kranj): [20.1 22.8 20.1 12.9 13.8]
BL01 - Bled jezero (Bled): [1.3 2.2 6.7 5.7]
BH01 - Bohinj jezero (Bohinjska Bistrica): [10.8 8.3 13.2 12.7 15.7]
-----POROCILO-----:
Piran zaliv: min 21.7 °C, max 33.8 °C, povprečje 28.14 °C -> Vroče
Mariborsko pohorje: Imamo premalo meritev!
Ljubljana FRI: min 10.7 °C, max 22.4 °C, povprečje 17.18 °C -> Toplo
Ljubljana BTC: min 11.7 °C, max 24.4 °C, povprečje 17.78 °C -> Toplo
Kranj center: min 12.9 °C, max 22.8 °C, povprečje 17.94 °C -> Toplo
Bled jezero: Imamo premalo meritev!
Bohinj jezero: min 8.3 °C, max 15.7 °C, povprečje 12.14 °C -> Hladno

*/
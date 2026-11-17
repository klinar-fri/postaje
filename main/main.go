package main

import (
	"context"
	"log"
	"os"

	"github.com/klinar-fri/postaje/vreme"
	"github.com/urfave/cli/v3"
)

func narediInDodajPostaje(slovar_postaj map[string]vreme.Postaja, s vreme.Stikala) {
	// vreme.DodajPostajo("LJ01", "Ljubljana FRI", "Ljubljana", []float64{22.4, 21.4, 18.5, 12.9, 10.7}, slovar_postaj)
	// vreme.DodajPostajo("LJ02",  "Ljubljana BTC", "Ljubljana",[]float64{21.4, 24.4, 19.5, 11.9, 11.7}, slovar_postaj)
	// vreme.DodajPostajo("KR01",  "Kranj center", "Kranj",[]float64{20.1, 22.8, 20.1, 12.9, 13.8}, slovar_postaj)
	// vreme.DodajPostajo("BL01",  "Bled jezero", "Bled", []float64{-10.8, 1.3, 2.2, 6.7, 5.7}, slovar_postaj)
	// vreme.DodajPostajo("BH01",  "Bohinj jezero", "Bohinjska Bistrica", []float64{10.8, 8.3, 13.2, 12.7, 15.7}, slovar_postaj)
	// vreme.DodajPostajo("PR01",  "Piran zaliv", "Piran", []float64{33.8, 28.3, 24.2, 21.7, 32.7}, slovar_postaj)
	// vreme.DodajPostajo("MBP01", "Mariborsko pohorje", "Maribor", []float64{4.3, 2.7}, slovar_postaj)

	podatki := []struct {
		oznaka, ime, kraj string
		meritve []float64
	}{
		{"LJ01", "Ljubljana FRI", "Ljubljana", []float64{22.4, 21.4, 18.5, 12.9, 10.7}},
		{"LJ02",  "Ljubljana BTC", "Ljubljana",[]float64{21.4, 24.4, 19.5, 11.9, 11.7}},
		{"KR01",  "Kranj center", "Kranj",[]float64{20.1, 22.8, 20.1, 12.9, 13.8}},
		{"BL01",  "Bled jezero", "Bled", []float64{-10.8, 1.3, 2.2, 6.7, 5.7}},
		{"BH01",  "Bohinj jezero", "Bohinjska Bistrica", []float64{10.8, 8.3, 13.2, 12.7, 15.7}},
		{"PR01",  "Piran zaliv", "Piran", []float64{33.8, 28.3, 24.2, 21.7, 32.7}},
		{"MBP01", "Mariborsko pohorje", "Maribor", []float64{4.3, 2.7}},

	}

	for _, p := range podatki {
		vreme.DodajPostajo(p.oznaka, p.ime, p.kraj, slovar_postaj)
		for _, t := range p.meritve {
			vreme.DodajMeritev(slovar_postaj, p.oznaka, t, s)
		}
	}
}
func izpisiMeritveInPorocilo(slovar_postaj map[string]vreme.Postaja, s vreme.Stikala){
	vreme.IzpisMeritev(slovar_postaj)
	vreme.IzpisPorocila(slovar_postaj, s)
}

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
			s := vreme.Stikala{
				MinTemp :  cmd.Float("minTemp"),
				MaxTemp :  cmd.Float("maxTemp"),
				StMeritev : cmd.Int("stMeritev"),
			}
			var slovar_postaj = make(map[string]vreme.Postaja)
			narediInDodajPostaje(slovar_postaj, s)
			izpisiMeritveInPorocilo(slovar_postaj, s)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
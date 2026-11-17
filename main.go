package main

import (
	"github.com/klinar-fri/postaje/vreme"
)

func main(){
	

	var slovar_postaj = make(map[string]vreme.Postaja)
	vreme.DodajPostajo("LJ01", "Ljubljana FRI", "Ljubljana", []float64{22.4, 21.4, 18.5, 12.9, 10.7}, slovar_postaj)
	vreme.DodajPostajo("LJ02",  "Ljubljana BTC", "Ljubljana",[]float64{21.4, 24.4, 19.5, 11.9, 11.7}, slovar_postaj)
	vreme.DodajPostajo("KR01",  "Kranj center", "Kranj",[]float64{20.1, 22.8, 20.1, 12.9, 13.8}, slovar_postaj)
	vreme.DodajPostajo("BL01",  "Bled jezero", "Bled", []float64{8.8, 1.3, 2.2, 6.7, 5.7}, slovar_postaj)
	vreme.DodajPostajo("BH01",  "Bohinj jezero", "Bohinjska Bistrica", []float64{10.8, 8.3, 13.2, 12.7, 15.7}, slovar_postaj)
	vreme.DodajPostajo("PR01",  "Piran zaliv", "Piran", []float64{33.8, 28.3, 24.2, 21.7, 32.7}, slovar_postaj)
	vreme.DodajPostajo("MBP01", "Mariborsko pohorje", "Maribor", []float64{4.3, 2.7}, slovar_postaj)

	vreme.IzpisMeritev(slovar_postaj)
	vreme.IzpisPorocila(slovar_postaj)

}
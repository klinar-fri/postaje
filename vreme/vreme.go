// Paket vsebuje funckije za delo z merilnimi postajami
// 
// Omogoča dodajanje novih postaj in meritev. Glej funkciji DodajPostajo in DodajMeritev
// Omogoča izpis vseh meritev v dani postaji. Glej funkcijo IzpisMeritev.
// Omogoča izpis poročila za celoten seznam. Poročilo vsebuje minTemp, maxTemp, avgTemp in real-feel. Glej funkcijo IzpisPorocila
// 
package vreme 

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

type Stikala struct {
	MinTemp float64
	MaxTemp float64
	StMeritev int
}

// DodajPostajo doda novo postajo v seznam.
func DodajPostajo(oznaka string, ime string, kraj string, postaje map[string]Postaja) {
	postaje[oznaka] = Postaja{ime: ime, kraj: kraj} 
}

// DodajMeritev doda novo meritev v tabelo mertiev dane postaje.
func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64, s Stikala){
	e, ok := postaje[oznaka] 
	if ok {
		if temperatura < s.MaxTemp && temperatura > s.MinTemp{
			// fmt.Printf("Dolzina :%d, Kapaciteta :%d | %v\n", len(e.meritve), cap(e.meritve), e.meritve)
			e.meritve = append(e.meritve, temperatura)
			// fmt.Printf("Dolzina :%d, Kapaciteta :%d | %v\n", len(e.meritve), cap(e.meritve), e.meritve)
			postaje[oznaka] = e
		}else{
			fmt.Printf("Meritev ni veljavna! -> Temperatura je %.1f, vendar mora biti med %.1f °C in %.1f °C\n", temperatura, s.MinTemp, s.MaxTemp)
		}
	} else {
		fmt.Println("Ta postaja (ključ) ne obstaja!")
	}
}

// statistika izračuna statistiko meritev na posamezni postaji. Vrne minTemp, maxTemp, avgTemp in ok (error value)
func statistika(postaje map[string]Postaja, oznaka string, s Stikala) (najnizja, najvisja, povprecna float64, ok bool){
	var avg float64 = 0
	var max float64 = -1000 
	var min float64 = 1000
	e, ok := postaje[oznaka]

	if ok {
		if len(e.meritve) < s.StMeritev {
			ok = false
			return 
		}else{
			for i := 0; i < len(e.meritve); i++ {
				if e.meritve[i] > max {
					max = e.meritve[i]
				}
				if e.meritve[i] < min{
					min = e.meritve[i]
				}
				avg += e.meritve[i]
			}
			return min, max , avg/float64(len(e.meritve)), true
		}
	}else{
		fmt.Println("Ta postaja (ključ) ne obstaja!")
		return 
	}
}

// IzpisMeritev izpiše vse meritve za vse postaje v seznamu.
func IzpisMeritev(postaje map[string]Postaja){
	fmt.Println("-----MERITVE-----:")
	for i, v := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", i, v.ime, v.kraj, v.meritve)
	}
}

// IzpisPorocila izpiše poročilo o vseh posatajah v seznamu. Porocilo vsebuje imePostaje, minTemp, maxTemp, avgTemp in real-feel: (Vroče, Toplo, Hladno, Mrzlo)
func IzpisPorocila(postaje map[string]Postaja, s Stikala){
	fmt.Println("-----POROCILO-----:")
	for i, v := range postaje {
		min, max, avg, ok := statistika(postaje, i, s)
		var tempOpis string = ""
		if ok {
			switch {
			case avg >= 25.0:
				tempOpis = "Vroče"
			case avg >= 15.0 && avg < 25.0:
				tempOpis = "Toplo"
			case avg >= 5.0 && avg < 15.0:
				tempOpis = "Hladno"
			case avg < 5.0:
				tempOpis = "Mrzlo"
			}
			fmt.Printf("%s: min %.1f °C, max %.1f °C, povprečje %.2f °C -> %s\n", v.ime, min, max, avg, tempOpis)
		} else {
			fmt.Printf("%s: Imamo premalo meritev, potrebnih je vsaj %d!\n", v.ime, s.StMeritev)
		}
	}
}
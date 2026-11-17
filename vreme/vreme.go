package vreme 

import "fmt"

type Postaja struct {
	ime     string
	kraj    string
	meritve []float64
}

func DodajPostajo(novaPostaja Postaja, oznaka string, postaje map[string]Postaja) map[string]Postaja{
	postaje[oznaka] = novaPostaja
	return postaje
}

func DodajMeritev(postaje map[string]Postaja, oznaka string, temperatura float64){
	e, ok := postaje[oznaka] 
	if ok {
		if temperatura < 50 && temperatura > -50{
			// fmt.Printf("Dolzina :%d, Kapaciteta :%d | %v\n", len(e.meritve), cap(e.meritve), e.meritve)
			e.meritve = append(e.meritve, temperatura)
			// fmt.Printf("Dolzina :%d, Kapaciteta :%d | %v\n", len(e.meritve), cap(e.meritve), e.meritve)
			postaje[oznaka] = e
		}else{
			fmt.Println("Meritev ni veljavna!")
		}
	} else {
		fmt.Println("Ta postaja (ključ) ne obstaja!")
	}
}

// skrita -> z malo
func statistika(postaje map[string]Postaja, oznaka string) (najnizja, najvisja, povprecna float64, ok bool){
	var avg float64 = 0
	var max float64 = 0
	var min float64 = 1000000000
	e, ok := postaje[oznaka]

	if ok {
		if len(e.meritve) < 3 {
			ok = false
			return 
		}else{
			for i := 0; i < len(e.meritve); i++ {
				if e.meritve[i] > max {
					max = e.meritve[i]
				}
				if e.meritve[i] < min {
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


func IzpisMeritev(postaje map[string]Postaja){
	fmt.Println("-----MERITVE-----:")
	for i, v := range postaje {
		fmt.Printf("%s - %s (%s): %v\n", i, v.ime, v.kraj, v.meritve)
	}
}

func IzpisPorocila(postaje map[string]Postaja){
	fmt.Println("-----POROCILO-----:")
	for i, v := range postaje {
		min, max, avg, ok := statistika(postaje, i)
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
			fmt.Printf("%s: Premalo meritev!\n", v.ime)
		}
	}
}
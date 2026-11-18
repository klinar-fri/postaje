## VREMENSKE POSTAJE
Paket vreme vsebuje funckije za delo z merilnimi postajami

- Omogoča dodajanje novih postaj in meritev. Glej funkciji `DodajPostajo` in `DodajMeritev`
- Omogoča izpis vseh meritev v dani postaji. Glej funkcijo `IzpisMeritev`.
- Omogoča izpis poročila za celoten seznam. Poročilo vsebuje minTemp, maxTemp, avgTemp in real-feel. <br> Glej funkcijo `IzpisPorocila`.
- Primer uporabe:
    ```go
    type Postaja struct {
        ime     string
        kraj    string
        meritve []float64
    }
    // V slovarju postaj oznake postameznih merilnih postaj ("LJ01") predstavljajo kluče za strukture Postaja
    var slovar_postaj = make(map[string]vreme.Postaja)
    vreme.DodajPostajo("LJ01", "Ljubljana FRI", "Ljubljana", []float64{22.4, 21.4, 18.5, 12.9, 10.7}, slovar_postaj)

    // Dodamo lahko tudi posamezne meritve
    vreme.DodajMeritev(slovar_postaj, oznakaPostaje, 10)
    ```
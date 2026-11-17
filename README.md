## VREMENSKE POSTAJE
Paket vreme vsebuje funckije za delo z merilnimi postajami

- Omogoča dodajanje novih postaj in meritev. Glej funkciji `DodajPostajo` in `DodajMeritev`
- Omogoča izpis vseh meritev v dani postaji. Glej funkcijo `IzpisMeritev`.
- Omogoča izpis poročila za celoten seznam. Poročilo vsebuje minTemp, maxTemp, avgTemp in real-feel. Glej funkcijo `IzpisPorocila`.
- Z zastavicami `minTemp`, `maxTemp` in `stMeritev` izločimo nepravilne meritve in omejimo število meritev za računje statistike.
<br>primer: `./main --minTemp=0 --maxTemp-25 --stMeritev=4`
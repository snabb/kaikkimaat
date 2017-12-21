Kaikki maat CSV-tiedostossa (suomeksi)
======================================

Lähteenä käytetään EU:n julkaisutoimiston suomenkielistä valtioiden,
alueiden ja rahayksiköiden luetteloa:

http://publications.europa.eu/code/fi/fi-5000500.htm

Luettelo sisältää seuraavat kentät:
1. Valtio
2. Koko nimi
3. Maakoodi
4. Pääkaupunki / hallinnollinen keskus
5. Asukkaannimi
6. Adjektiivi
7. Rahayksikkö
8. Valuuttakoodi
9. Rahan alayksikkö

Maakoodit ovat muuten ISO 3166-1 alpha-2 mukaisia, paitsi Kreikan ja
Yhdistyneen kuningaskunnan osalta käytössä on koodit `EL` ja `UK`. Tästä
ja muista EU-erikoisuuksista on tarkempia tietoja lähdedokumentissa.

CSV-muotoisen tiedoston saa päivitettyä ajamalla:

`go run genkaikkimaat.go > kaikkimaat.csv`


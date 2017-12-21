Kaikki maat CSV-tiedostossa (suomeksi)
======================================

Lähteenä käytetään Tilastokeskuksen kansallisen luokitussuosituksen alueluokitusten
valtiot ja maat -luetteloa:

http://www.stat.fi/meta/luokitukset/valtio/001-2012/index.html

Luettelo sisältää seuraavat kentät:
1. numeerinen maakoodi
2. 2-kirjaiminen maakoodi
3. 3-kirjaiminen maakoodi
4. maan nimi suomeksi

Maakoodit ovat ISO 3166-1 mukaisia.

CSV-muotoisen tiedoston saa päivitettyä ajamalla:

`go run genkaikkimaat.go > kaikkimaat.csv`


Kaikki maat CSV-tiedostossa (suomeksi)
======================================

Lähteenä Tilastokeskuksen kansallinen luokitussuositus, alueluokitukset:
valtiot ja maat:

http://www.stat.fi/meta/luokitukset/valtio/001-2012/index.html

Luettelo sisältää seuraavat kentät:
1. numeerinen maakoodi (ISO 3166-1 numeric)
2. kaksikirjaiminen maakoodi (ISO 3166-1 alpha-2)
3. kolmikirjaiminen maakoodi (ISO 3166-1 alpha-3)
4. maan nimi suomeksi

CSV-muotoisen tiedoston saa päivitettyä ajamalla:

`go run genkaikkimaat.go > kaikkimaat.csv`


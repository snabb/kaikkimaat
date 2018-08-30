Kaikki maat CSV-tiedostossa (suomeksi)
======================================

Kaikki maailman valtiot ja maat lueteltuna suomeksi ISO-3166
maakoodien kera CSV-muotoisessa tiedostossa UTF-8 -merkistöllä
koodattuna.

Lähteenä Tilastokeskuksen kansallinen luokitussuositus, alueluokitukset,
valtiot ja maat: http://www.stat.fi/meta/luokitukset/valtio/001-2012/index.html

Tästä voi olla hyötyä, jos esimerkiksi tarvitset suomenkieliseen
ohjelmaasi valikon, josta voi valita minkä tahansa maan.

Luettelo sisältää seuraavat kentät:
1. numeerinen maakoodi (ISO 3166-1 numeric)
2. kaksikirjaiminen maakoodi (ISO 3166-1 alpha-2)
3. kolmikirjaiminen maakoodi (ISO 3166-1 alpha-3)
4. maan nimi suomeksi

CSV-tiedostossa ei ole otsikkoriviä, eli maaluettelo alkaa ensimmäiseltä
riviltä.

CSV-muotoisen tiedoston saa tarvittaessa luotua tai päivitettyä seuraavasti:
1. Tarvitset Go-kääntäjän, jonka asennusohje löytyy seuraavasta
   osoitteesta: https://golang.org/doc/install)
2. Aja sen jälkeen: `go run genkaikkimaat.go > kaikkimaat.csv`


package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

//Meeste code komt overeen dus tegen de helft aan geen toelichting meer.
//Er zijn nog een aantal println waar ik een nummer laat uitprinten in command line.
//Deze heb ik expres erin gelaten want het helpt me met troubleshooten
//Ik kan zien in welke func de fout bevind.
//Product data types voor exporteren
type Product struct {
    ID           int
    Productnaam  string
    Beschrijving string
    Geleenddoor  string
    Geleendtot   string
}

//Dit is nodig voor bepaalde functies -> zoals sql.Open
var tpl *template.Template
var db *sql.DB

func main() {
    tpl, _ = template.ParseGlob("templates/*html")
    //Met ParseGlob creeer ik een nieuwe template en ontleed de template definities die door bestanden word geindentificeerd.
    var err error

    //Hieronder open ik SQL en log ik in mijn account. poortnummer DBnaam zijn van belang
    //Variabelen db sql.DB en err zijn aangemaakt
    //Standaard error code eronder
    db, err = sql.Open("mysql", "root:Welkom01!@tcp(db:3306)/sabahattin")
    if err != nil {
        panic(err.Error())
    }
    //Wanneer we eenmaal een connectie hebben met de DB, sluiten we met defer db.Close de DB
    //Deze Handlefuncs verwachten een functie, ze refereren met html website en zorgen dat de aangegeven code word weergeven
    //In principe wanneer een verzoek komt naar de webservice met de gematchde url, zal golang http pakket de handler roepen
    //met r(read) en w(write)
    //Dus zoekHandler is wat alles laat zien in de producttabel
    defer db.Close()
    http.HandleFunc("/zoek", zoekHandler)
    http.HandleFunc("/producttoevoegen", ProducttoevoegenHandler)
    http.HandleFunc("/aanpassen", updateHandler)
    http.HandleFunc("/updateresultaat", updateresultaatHandler)
    http.HandleFunc("/verwijder", verwijderHandler)
    http.HandleFunc("/", homepaginaHandler)
    http.ListenAndServe(":8095", nil)
    //Nodig om website te bezoeken door bovenstaande in te vullen
    fmt.Println("1")
}

func zoekHandler(w http.ResponseWriter, r *http.Request) {
    stmt := "Select * FROM producten"
    //Deze statement gebruik ik voor mysql. Select = SQL code. Je wilt hier alle attributen gebruiken
    rows, err := db.Query(stmt)
    if err != nil {
        panic(err)
        //De error is bedoeld als de statement niet werkt.
    }
    //Vervolgens sluit ik de variabel rows
    defer rows.Close()
    var products []Product
    //Product is mijn structnaam
    //var products is een slice om rows scan in te voeren
    for rows.Next() {
        var p Product
        //Hij pakt weer de statement (defer rows.Next)
        //Hij pakt de volgende kolommen in de DB
        //Data word gestopt in onderstaande beschrijvingen
        //& = locatie
        //p = products
        err = rows.Scan(&p.ID, &p.Productnaam, &p.Beschrijving, &p.Geleenddoor, &p.Geleendtot)
        if err != nil {
            panic(err)
        }
        //Via append voeg ik alles in mijn products slice
        products = append(products, p)
    }
    tpl.ExecuteTemplate(w, "selecteer.html", products)
    //Hier geven we products aan zodat de template word uitgevoerd/execute
    fmt.Println("2")
}

func ProducttoevoegenHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        //ZOEKEN
        tpl.ExecuteTemplate(w, "Producttoevoegen.html", nil)
        return
        //Hij voert specifiek deze template uit voer website
    }
    r.ParseForm()
    //In html heb ik een form en met Pars  (ontleed) ik de form
    //Deze is dus geparsed met producttoevegen
    //In principe word dus gevraagd om de FormValue die hij ophaalt.
    productnaam := r.FormValue("productnaam")
    beschrijving := r.FormValue("beschrijving")
    geleenddoor := r.FormValue("geleenddoor")
    geleendtot := r.FormValue("geleendtot")
    var err error
    // Hier krijg ik foutmelding als er niks word ingevuld -> Dus controleer alle velden
    if productnaam == "" || beschrijving == "" || geleenddoor == "" || geleendtot == "" {
        fmt.Println("Fout bij toevoegen rij:", err)
        tpl.ExecuteTemplate(w, "Producttoevoegen.html", "Fout bij toevoegen data, controleer alle velden.")
        return
        //Return is als mijn if niet werkt, gaat die het nog X opnieuw doen
    }
    var ins *sql.Stmt
    //Hier word het geinporteerd naar DB
    ins, err = db.Prepare("INSERT INTO `sabahattin`.`producten` (`Productnaam`, `Beschrijving`, `Geleenddoor`, `Geleendtot`) VALUES (?, ?, ?, ?);")
    if err != nil {
        panic(err)
    }
    //Vervolgens sluit ik de lijn
    defer ins.Close()
    //Defer is er om een functie te sluiten
    //res = result geeft informatie,
    res, err := ins.Exec(productnaam, beschrijving, geleenddoor, geleendtot)
    //hier breng ik result en error terug om Exec te uitvoeren.
    //RowsAffected slaat op in rowsAffec
    //Controleren of error nill, als het niet is 1 row affecten.
    //Hier wil ik weten of ik wel 1 rij affect.
    rowsAffec, _ := res.RowsAffected()
    if err != nil || rowsAffec != 1 {
        fmt.Println("Fout bij toevoegen rij:", err)
        tpl.ExecuteTemplate(w, "Producttoevoegen.html", "Fout bij toevoegen data, controleer alle velden.")
        return
    }
    lastInserted, _ := res.LastInsertId()
    rowsAffected, _ := res.RowsAffected()
    //Prinln kijken of ze affected zijn.
    fmt.Println("ID van laatste rij toegevoegd:", lastInserted)
    fmt.Println("nummer van de rijen zijn getroffen :", rowsAffected)
    tpl.ExecuteTemplate(w, "Producttoevoegen.html", "Product is succesvol toegevoegd")
    fmt.Println("3")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {

    r.ParseForm()
    id := r.FormValue("idproducten")
    row := db.QueryRow("Select * FROM sabahattin.producten WHERE ID = ?;", id)
    var p Product
    fmt.Print(id)

    err := row.Scan(&p.ID, &p.Productnaam, &p.Beschrijving, &p.Geleenddoor, &p.Geleendtot)
    if err != nil {
        fmt.Println(err)
        http.Redirect(w, r, "/zoek", http.StatusTemporaryRedirect)
        return

    }
    tpl.ExecuteTemplate(w, "update.html", p)
    fmt.Println("4")
}

func updateresultaatHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    id := r.FormValue("idproducten")
    productnaam := r.FormValue("productnaam")
    beschrijving := r.FormValue("beschrijving")
    geleenddoor := r.FormValue("geleenddoor")
    geleendtot := r.FormValue("geleendtot")
    fmt.Println(id, productnaam, beschrijving, geleenddoor, geleendtot)
    upStmt := "UPDATE `sabahattin`.`producten` SET `Productnaam` = ?, `Beschrijving` = ?, `Geleenddoor` = ?, `Geleendtot` = ?  WHERE (`ID` = ?);"
    stmt, err := db.Prepare(upStmt)
    if err != nil {
        fmt.Println("Fout bij voorbereiden stmt")
        panic(err)

    }

    fmt.Println("db.Prepare err:", err)
    fmt.Println("db.Prepare stmt:", stmt)
    defer stmt.Close()

    var res sql.Result

    res, err = stmt.Exec(productnaam, beschrijving, geleenddoor, geleendtot, id)
    rowsAff, _ := res.RowsAffected()

    if err != nil || rowsAff != 1 {
        fmt.Println(err)
        tpl.ExecuteTemplate(w, "resultaat.html", "Er was een probleem bij het updaten van het product.")
        return

    }
    tpl.ExecuteTemplate(w, "resultaat.html", "Product is succesvol geupdated")
    fmt.Println("3")
}

func verwijderHandler(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    id := r.FormValue("idproducten")
    fmt.Println(id)

    del, err := db.Prepare("DELETE FROM `sabahattin`.`producten` WHERE (`ID` = ?);")
    if err != nil {
        panic(err)
    }

    defer del.Close()
    var res sql.Result
    res, err = del.Exec(id)
    rowsAff, _ := res.RowsAffected()
    fmt.Println("rowsAff:", rowsAff)

    if err != nil || rowsAff != 1 {
        fmt.Fprint(w, "Probleem bij verwijderen product, neem contact op met de servicedesk")
        return
    }

    fmt.Println("err:", err)
    tpl.ExecuteTemplate(w, "resultaat.html", "Product is succesvol verwijderd")
    fmt.Println("4")
}

func homepaginaHandler(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/zoek", http.StatusTemporaryRedirect)
    fmt.Println(10)
    //http.StatusTemporaryRedirect is 307.
}


package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "dylanbroad"
	password = "7rybRRPo8UeIIq7"
	dbname = "cryptoExchange"
)

func main () {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)	
	CheckError(err)

	defer db.Close()

    // insert
    // hardcoded
    insertStmt := `INSERT into "users"("name", "username", "email") values('Taylor', 'taylorcarnegie', 'taylorcarnegie@gmail.com')`
    _, e := db.Exec(insertStmt)
    CheckError(e)
 
    /* dynamic
    insertDynStmt := `insert into "Students"("Name", "Roll_Number") values($1, $2)`
    _, e = db.Exec(insertDynStmt, "Jack", 21)
    CheckError(e)
	*/
}
 
func CheckError(err error) {	
    if err != nil {
        panic(err)
    }
}


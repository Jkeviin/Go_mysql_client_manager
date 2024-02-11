package connect

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Guardar coneccion para conectarnos a la DB
var Db *sql.DB

func Connect() {
	ErrorVar := godotenv.Load()

	// Si hay un error al cargar el archivo .env
	if ErrorVar != nil {
		panic(ErrorVar)
	}

	// Hacer coneccion a la DB
	connection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		panic(err)
	}

	Db = connection

}

// Cerrar coneccion a la DB
func CloseConnection() {
	Db.Close()
}

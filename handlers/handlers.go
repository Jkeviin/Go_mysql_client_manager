package handlers

import (
	"bufio"
	"fmt"
	"golang_mysql_driver/connect"
	"golang_mysql_driver/models"
	"os"
	"strconv"
	"strings"
)

func ListClients() {
	connect.Connect()

	sql := "SELECT id, name, email, phone FROM clients WHERE active = 1 ORDER BY id ASC"

	dataList, err := connect.Db.Query(sql)

	if err != nil {
		fmt.Println(err)
	}

	defer connect.CloseConnection()

	//Clients := models.Clients{}

	// Listado de clientes
	fmt.Println("LISTADO DE CLIENTES")
	fmt.Println("ID\tNOMBRE\t\tEMAIL\t\t\tTELEFONO")
	for dataList.Next() {
		var data models.Client
		err := dataList.Scan(&data.Id, &data.Name, &data.Email, &data.Phone)
		if err != nil {
			fmt.Println(err)
		}

		//Clients = append(Clients, data)
		fmt.Printf("%d\t%s\t%s\t%s\n", data.Id, data.Name, data.Email, data.Phone)
	}
}

// clientes por id
func ClientById(id int) {
	connect.Connect()

	sql := "SELECT id, name, email, phone FROM clients WHERE id = ? AND active = 1 ORDER BY id ASC"

	dataList, err := connect.Db.Query(sql, id)

	if err != nil {
		fmt.Println(err)
	}

	if !dataList.Next() {
		fmt.Println("No se encontró el cliente")
		return
	}

	defer connect.CloseConnection()

	// Listado de clientes
	fmt.Println("CLIENTE POR ID")
	fmt.Println("ID\tNOMBRE\t\tEMAIL\t\t\tTELEFONO")

	var data models.Client
	err = dataList.Scan(&data.Id, &data.Name, &data.Email, &data.Phone)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d\t%s\t%s\t%s\n", data.Id, data.Name, data.Email, data.Phone)
}

// Insertar cliente

func InsertClient(client models.Client) {
	connect.Connect()

	sql := "INSERT INTO clients(name, email, phone) VALUES(?, ?, ?)"

	insert, err := connect.Db.Query(sql, client.Name, client.Email, client.Phone)

	if err != nil {
		fmt.Println(err)
	}

	defer insert.Close()
	defer connect.CloseConnection()

	fmt.Printf("Cliente %s insertado correctamente", client.Name)
}

// Actualizar cliente
func UpdateClient(client models.Client) {
	connect.Connect()
	defer connect.CloseConnection()

	if client.Id == 0 {
		fmt.Println("No se proporcionó el ID del cliente")
		return
	}

	var columns []string
	var values []interface{}

	// Construye dinámicamente la parte SET de la consulta
	if client.Name != "" {
		columns = append(columns, "name = ?")
		values = append(values, client.Name)
	}
	if client.Email != "" {
		columns = append(columns, "email = ?")
		values = append(values, client.Email)
	}
	if client.Phone != "" {
		columns = append(columns, "phone = ?")
		values = append(values, client.Phone)
	}

	if len(columns) == 0 {
		fmt.Println("No se proporcionaron campos para actualizar")
		return
	}

	// Construye la consulta SQL
	sql := fmt.Sprintf("UPDATE clients SET %s WHERE id = ?", strings.Join(columns, ", ")) // sprintf para concatenar strings
	values = append(values, client.Id)

	// Ejecuta la consulta
	_, err := connect.Db.Exec(sql, values...)
	if err != nil {
		fmt.Println("Error al actualizar el cliente:", err)
		return
	}

	fmt.Printf("Cliente %s actualizado correctamente", client.Name)
}

// Eliminar cliente
func DeleteClient(id int) {
	connect.Connect()
	defer connect.CloseConnection()

	// Consultar si el cliente está activo
	var active int
	err := connect.Db.QueryRow("SELECT active FROM clients WHERE id = ?", id).Scan(&active)

	if err != nil {
		fmt.Println("Error al consultar el cliente:", err)
		return
	}

	// Verificar si el cliente está activo
	if active != 1 {
		fmt.Println("El cliente no está activo o no existe")
		return
	}

	// Eliminar cliente
	_, err = connect.Db.Exec("UPDATE clients SET active = 0 WHERE id = ?", id)
	if err != nil {
		fmt.Println("Error al eliminar el cliente:", err)
		return
	}

	fmt.Printf("Cliente %d eliminado correctamente\n", id)
}

// Menu de opciones
func Menu() {
	for {
		Scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("\nMENU DE OPCIONES")
		fmt.Println("1. Listado de clientes")
		fmt.Println("2. Cliente por id")
		fmt.Println("3. Insertar cliente")
		fmt.Println("4. Actualizar cliente")
		fmt.Println("5. Eliminar cliente")
		fmt.Println("6. Salir")

		fmt.Print("\nIngrese una opción: ")

		if Scanner.Scan() {
			fmt.Println("\n******************")
			switch Scanner.Text() {
			case "1":

				fmt.Println("")
				ListClients()
				break
			case "2":
				fmt.Print("Ingrese el id del cliente que deseas consultar: ")
				var ID int
				if Scanner.Scan() {
					ID, _ = strconv.Atoi(Scanner.Text())
				}

				fmt.Println("")
				ClientById(ID)
				break
			case "3":
				fmt.Print("Ingrese el nombre del cliente: ")
				var Name string
				if Scanner.Scan() {
					Name = Scanner.Text()
				}
				fmt.Print("Ingrese el email del cliente: ")
				var Email string
				if Scanner.Scan() {
					Email = Scanner.Text()
				}
				fmt.Print("Ingrese el telefono del cliente: ")
				var Phone string
				if Scanner.Scan() {
					Phone = Scanner.Text()
				}
				var newCliente = models.Client{
					Name:  Name,
					Email: Email,
					Phone: Phone,
				}

				fmt.Println("")
				InsertClient(newCliente)
				break
			case "4":
				fmt.Print("Ingrese el id del cliente que deseas actualizar: ")
				var ID int
				if Scanner.Scan() {
					ID, _ = strconv.Atoi(Scanner.Text())
				}
				fmt.Print("Ingrese el nombre del cliente: ")
				var Name string
				if Scanner.Scan() {
					Name = Scanner.Text()
				}
				fmt.Print("Ingrese el email del cliente: ")
				var Email string
				if Scanner.Scan() {
					Email = Scanner.Text()
				}
				fmt.Print("Ingrese el telefono del cliente: ")
				var Phone string
				if Scanner.Scan() {
					Phone = Scanner.Text()
				}
				var updateCliente = models.Client{
					Id:    ID,
					Name:  Name,
					Email: Email,
					Phone: Phone,
				}

				fmt.Println("")
				UpdateClient(updateCliente)
				break
			case "5":
				fmt.Print("Ingrese el id del cliente que deseas eliminar: ")
				var ID int
				if Scanner.Scan() {
					ID, _ = strconv.Atoi(Scanner.Text())
				}

				fmt.Println("")
				DeleteClient(ID)
				break
			case "6":
				fmt.Println("")
				fmt.Println("Gracias por usar el sistema")
				return
			default:
				fmt.Println("")
				fmt.Println("Opción no válida")
				break
			}
		}
	}
}

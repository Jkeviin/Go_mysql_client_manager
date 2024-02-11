# Gestor de Clientes MySQL en Go

Este proyecto contiene un conjunto de funciones para interactuar con una base de datos MySQL y administrar clientes. A continuación se detalla la funcionalidad proporcionada por cada función:

## Funciones

### ListClients
- Descripción: Esta función lista todos los clientes activos almacenados en la base de datos.
- Uso: `ListClients()`

### ClientById
- Descripción: Esta función busca y muestra la información de un cliente específico según su ID.
- Parámetros:
  - `id`: ID del cliente a buscar.
- Uso: `ClientById(id int)`

### InsertClient
- Descripción: Esta función inserta un nuevo cliente en la base de datos.
- Parámetros:
  - `client`: Objeto `models.Client` que contiene la información del cliente a insertar.
- Uso: `InsertClient(client models.Client)`

### UpdateClient
- Descripción: Esta función actualiza la información de un cliente existente en la base de datos.
- Parámetros:
  - `client`: Objeto `models.Client` que contiene la información actualizada del cliente.
- Uso: `UpdateClient(client models.Client)`

### DeleteClient
- Descripción: Esta función elimina un cliente de la base de datos.
- Parámetros:
  - `id`: ID del cliente a eliminar.
- Uso: `DeleteClient(id int)`

### Menu
- Descripción: Esta función muestra un menú interactivo para utilizar las funciones mencionadas anteriormente.
- Uso: `Menu()`

## Uso

Para utilizar estas funciones, puedes importar el paquete `handlers` en tu código y llamar a las funciones según tus necesidades. Además, asegúrate de tener las credenciales de la base de datos configuradas correctamente en el archivo `connect.go`.

**Ejemplo:**
```go
package main

import (
    "golang_mysql_driver/handlers"
    "golang_mysql_driver/models"
)

func main() {
    // Ejemplo de uso de la función ListClients
    handlers.ListClients()

    // Ejemplo de uso de la función InsertClient
    newClient := models.Client{Name: "Nuevo Cliente", Email: "cliente@example.com", Phone: "123456789"}
    handlers.InsertClient(newClient)

    // Otros usos...
}
```

## Requisitos

- Tener acceso a una base de datos MySQL.
- Configurar correctamente las credenciales de la base de datos en `connect.go`.

## Notas

- Asegúrate de manejar los errores adecuadamente al utilizar estas funciones, especialmente al interactuar con la base de datos.
- Se proporciona un menú interactivo para facilitar la utilización de las funciones, pero también puedes llamar a las funciones directamente desde tu código según sea necesario.






#
#
#
#
#
#





# MySQL Client Manager in Go

This project contains a set of functions to interact with a MySQL database and manage clients. Below is a detailed description of the functionality provided by each function:

## Functions

### ListClients
- Description: This function lists all active clients stored in the database.
- Usage: `ListClients()`

### ClientById
- Description: This function searches and displays information for a specific client based on its ID.
- Parameters:
  - `id`: ID of the client to search.
- Usage: `ClientById(id int)`

### InsertClient
- Description: This function inserts a new client into the database.
- Parameters:
  - `client`: `models.Client` object containing the information of the client to insert.
- Usage: `InsertClient(client models.Client)`

### UpdateClient
- Description: This function updates the information of an existing client in the database.
- Parameters:
  - `client`: `models.Client` object containing the updated information of the client.
- Usage: `UpdateClient(client models.Client)`

### DeleteClient
- Description: This function deletes a client from the database.
- Parameters:
  - `id`: ID of the client to delete.
- Usage: `DeleteClient(id int)`

### Menu
- Description: This function displays an interactive menu to use the aforementioned functions.
- Usage: `Menu()`

## Usage

To use these functions, you can import the `handlers` package into your code and call the functions according to your needs. Additionally, make sure to have the database credentials correctly configured in the `connect.go` file.

**Example:**
```go
package main

import (
    "golang_mysql_driver/handlers"
    "golang_mysql_driver/models"
)

func main() {
    // Example usage of ListClients function
    handlers.ListClients()

    // Example usage of InsertClient function
    newClient := models.Client{Name: "New Client", Email: "client@example.com", Phone: "123456789"}
    handlers.InsertClient(newClient)

    // Other usages...
}
```

## Requirements

- Access to a MySQL database.
- Correctly configure database credentials in `connect.go`.

## Notes

- Make sure to handle errors appropriately when using these functions, especially when interacting with the database.
- An interactive menu is provided to facilitate the use of the functions, but you can also call the functions directly from your code as needed.

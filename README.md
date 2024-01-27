# API con Go y MySQL

- TODO: Terminar el proyecto, este readme es el proyecto finalizado

Este proyecto es una API REST desarrollada en Go, diseñada para interactuar con la base de datos `ticket_sys_db`, la cual gestiona un sistema de tickets. La base de datos incluye tablas para departamentos, roles, usuarios, clientes, rutas, viajes de clientes y tickets.

## Características

- **CRUD para Departamentos**: Crear, leer, actualizar y eliminar departamentos.
- **Gestión de Usuarios**: Registrar y autenticar usuarios, asignar roles.
- **Manejo de Clientes**: Administrar información de clientes.
- **Rutas y Viajes**: Registrar rutas y asociar viajes a clientes.
- **Tickets**: Crear y gestionar tickets.

## Tecnologías Utilizadas

- **Lenguaje de Programación**: Go
- **Base de Datos**: MySQL
- **Librerías Principales**: `net/http` para el servidor HTTP, `database/sql` con `github.com/go-sql-driver/mysql` para la conexión a MySQL.

## Estructura del Proyecto

```
api_go/
├── controllers/
│   ├── cliente.go
│   ├── departamento.go
│   ├── error.go
│   └── usuario.go
├── models/
│   ├── cliente.go
│   ├── conexion.go
│   ├── departamento.go
│   └── usuario.go
├── routes/
│   └── routes.go
├── test/
│   ├── routes_test.go
│   └── tutorial.go
├── views/
│   └── error.json
├── db.sql
├── go.mod
├── go.sum
└── main.go
```

## Configuración

Para configurar y ejecutar este proyecto en tu entorno local, sigue estos pasos:

1. **Clonar el Repositorio**:
   ```
   git clone https://github.com/DavidDevGt/api_go.git
   cd api_go
   ```

2. **Configurar la Base de Datos**:
   - Asegúrate de tener MySQL instalado y en ejecución.
   - Ejecuta el script `db.sql` para crear y configurar la base de datos y las tablas.

3. **Instalar Dependencias**:
   ```
   go mod tidy
   ```

4. **Ejecutar la Aplicación**:
   ```
   go run main.go
   ```

## Pruebas

Para ejecutar las pruebas automatizadas, usa el siguiente comando:

```
go test ./test/...
```

## Contribuir

Si deseas contribuir a este proyecto, por favor haz un fork del repositorio, realiza tus cambios y envía un Pull Request.
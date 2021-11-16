# Bienvenidos
Este es un API REST simple que usa el framework Echo de Golang y el ODM Kamva. Está conectado a una base de datos en la nube de Mongo Atlas.

Para poder ejecutar el REST server, colocar el siguiente comando:

```
go run main/server.go
```
La estructura de cada curso es simple:

```
{
    "id": "6193f7e608530287f48bcdf0",
    "Name": "Course 3",
    "Description": "Description 3"
}
```


Los puntos a probar son:

- Consultar todos los cursos:

> METHOD: GET

```
localhost:8080/courses
```
- Crear curso

> METHOD: POST (by form-data)

```
localhost:8080/course
```
- Consultar curso por ID

> METHOD: GET

```
localhost:8080/course/:id
```
- Eliminar curso por ID

> METHOD: DELETE

```
localhost:8080/course/:id
```
- Actualizar curso por ID

> METHOD: PUT (by form-data)

```
localhost:8080/course/:id
```
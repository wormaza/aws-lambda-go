# aws-lambda-go

Ejemplos de pruebas para uso de AWS

## Tecnologías

Dentro de las tecnologías utilizadas se destacan las siguientes

* [Golang](https://golang.org/) - Lenguaje utilizado para las funciones lambda
* [AWS Lambda](https://aws.amazon.com/es/lambda/) - Funciones lambda
* [DynamoDb](https://aws.amazon.com/es/dynamodb/) - Base de datos no relacional

## Funciones

A continuación se indican las funciones lambda y su uso:

* [autorizador-code](autorizador-code): implementa una validación simple de un token de acceso (por código)
* [lambda-code-desactivar](lambda-code-desactivar): modifica el valor del campo activo de un registro a falso
* [lambda-code-getlicitacion](lambda-code-getlicitacion): obtiene licitaciones por rut de proveedor.
* [lambda-code-getproveedor](lambda-code-getproveedor): obtiene proveedores por código de licitación.
* [lambda-code-postlicitacion](lambda-code-postlicitacion): ingresa los datos de licitación para generar las recomendaciones (se delegan a sistema externo).
* [lambda-code-postproveedor](lambda-code-postproveedor): ingresa los datos de una recomendación para que sean alamcenados en la base de datos.

## Estructura eventos

### lambda-code-desactivar

* Entrada
 
```
{
	"Code":<string>
}

```

* Salida
```
{
	"Message":"Ok"
}
```

### lambda-code-getlicitacion

* Entrada
 
```
<URL>/licitacion?RutProveedor=rut

```

* Salida
```
[
    {
        "code": "8041959",
        "externalcode": "500977-138-E216",
        "relevancia": 60
    }
]

```

### lambda-code-getproveedor

* Entrada
 
```
<URL>/proveedor?Codigo=codigo

```

* Salida
```
[
    {
        "rut": "76.753.070-6",
        "relevancia": 60
    }    
]

```

### lambda-code-postlicitacion

* Entrada
 
```
{
  "Codigo": <string>,
  "CodigoExterno": <string>,
  "Nombre": <string>,
  "Descripcion": <string>,
  "CodigoOrganismo": <numeric>,
  "NombreOrganismo": <string>,
  "RutUnidad": <string>,
  "CodigoUnidad": <numeric>,
  "NombreUnidad": <string>,
  "CodigoTipo": <numeric>,
  "Tipo": <string>,
  "CodigoProducto": <numeric>,
  "NombreProducto": <string>,
  "CodigoCategoria": <numeric>,
  "NombreCategoria": <string>,
  "DescripcionItem": <string>
}


```

* Salida
```
{
	"Message":"Ok"
}
```

### lambda-code-postproveedor

* Entrada
 
```
{
	"Code":<string>,
	"ExternalCode":<string>,
	"Rut":<string>,
	"Rubros":<string>,
	"Relevancia":<numeric>,
	"Activo":<boolean>
}

```

* Salida
```
{
	"Message":"Ok"
}
```
## DynamoDb

Corresponde a una base de datos no relacional, en la nube. La estructura es la siguiente:

```
{
	"Code":<string>,
	"ExternalCode":<string>,
	"Rut":<string>,
	"Rubros":<string>,
	"Relevancia":<numeric>,
	"Activo":<boolean>
}
```
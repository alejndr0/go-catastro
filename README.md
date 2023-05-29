# go-catastro

Libreria en Go para el acceso al Catastro.

### Documentacion 

Basado en la version 2.1 de los [servicios web libres](https://www.catastro.meh.es/ws/webservices_libres.pdf) de la Sede Electrónica del Catastro. Los struct mantienen en el mismo esquema que los esquema XML que se pueden encontrar en el PDF.

### Ejemplo de uso  

```go
package main

import (
	"log"

	"github.com/alejndr0/go-catastro/catastro"
	"github.com/alejndr0/go-catastro/tipos"
)

func main() {

	resp, err := catastro.ObtenerCallejero(catastro.ObtenerCallejeroRequest{
		Provincia: "VALENCIA",
		Municipio: "VALENCIA",
		TipoVia:   tipos.Calle,
		NombreVia: "MAYOR",
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)

}

```
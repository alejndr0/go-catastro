package catastro

import "errors"

// Listado de PROVINCIAS
func ObtenerProvincias() (ObtenerProvinciasResponse, error) {
	var response ObtenerProvinciasResponse
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/ObtenerProvincias", nil, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (r ObtenerProvinciasResponse) hasFailed() (bool, error) {
	if r.ConsultaProvincieroResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaProvincieroResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Listado de MUNICIPIOS de una provincia
func ObtenerMunicipios(req ObtenerMunicipiosRequest) (ObtenerMunicipiosResponse, error) {
	var response ObtenerMunicipiosResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/ObtenerMunicipios", queryparams, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
func (r ObtenerMunicipiosResponse) hasFailed() (bool, error) {
	if r.ConsultaMunicipieroResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaMunicipieroResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Listado de los NÚMEROS de una vía
func ObtenerNumerero(req ObtenerNumereroRequest) (ObtenerNumereroResponse, error) {
	var response ObtenerNumereroResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
		"TipoVia":   req.TipoVia,
		"NomVia":    req.NombreVia,
		"Numero":    req.Numero,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/ObtenerNumerero", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ObtenerNumereroResponse) hasFailed() (bool, error) {
	if r.ConsultaNumereroResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaNumereroResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Listado de VÍAS de un municipio
func ObtenerCallejero(req ObtenerCallejeroRequest) (ObtenerCallejeroResponse, error) {
	var response ObtenerCallejeroResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
		"TipoVia":   req.TipoVia,
		"NomVia":    req.NombreVia,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/ObtenerCallejero", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ObtenerCallejeroResponse) hasFailed() (bool, error) {
	if r.ConsultaCallejeroResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaCallejeroResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Consulta de DATOS CATASTRALES NO PROTEGIDOS de un inmueble identificado por su localización
func ConsultaDNPLOC(req ConsultaDNPLOCRequest) (ConsultaDNPLOCResponse, error) {
	var response ConsultaDNPLOCResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
		"Sigla":     req.Sigla,
		"Calle":     req.Calle,
		"Numero":    req.Numero,
		"Bloque":    req.Bloque,
		"Escalera":  req.Escalera,
		"Planta":    req.Planta,
		"Puerta":    req.Puerta,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/Consulta_DNPLOC", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ConsultaDNPLOCResponse) hasFailed() (bool, error) {
	if r.ConsultaDnplocResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaDnplocResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Consulta de DATOS CATASTRALES NO PROTEGIDOS de un inmueble identificado por su Referencia Catastral
func ConsultaDNPRC(req ConsultaDNPRCRequest) (ConsultaDNPRCResponse, error) {
	var response ConsultaDNPRCResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
		"RefCat":    req.ReferenciaCatastral,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/Consulta_DNPRC", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ConsultaDNPRCResponse) hasFailed() (bool, error) {
	if r.ConsultaDnprcResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaDnprcResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Consulta de DATOS CATASTRALES NO PROTEGIDOS de un inmueble identificado por polígono-parcela
func ConsultaDNPPP(req ConsultaDNPPPRequest) (ConsultaDNPPPResponse, error) {
	var response ConsultaDNPPPResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
		"Poligono":  req.Poligono,
		"Parcela":   req.Parcela,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/Consulta_DNPPP", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ConsultaDNPPPResponse) hasFailed() (bool, error) {
	if r.ConsultaDnpppResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaDnpppResult.ErrorList[0].Desc)
	}
	return false, nil
}

package catastro

// Listado de PROVINCIAS
func ObtenerProvinciasCodigos() (ObtenerProvinciasResponse, error) {
	var response ObtenerProvinciasResponse
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejeroCodigos.svc/json/ObtenerProvincias", nil, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// Listado de MUNICIPIOS de una provincia
func ObtenerMunicipiosCodigos(req ObtenerMunicipiosCodigosRequest) (ObtenerMunicipiosResponse, error) {
	var response ObtenerMunicipiosResponse
	queryparams := map[string]string{
		"CodigoProvincia":    req.CodigoProvincia,
		"CodigoMunicipio":    req.CodigoMunicipio,
		"CodigoMunicipioINE": req.CodigoMunicipioINE,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejero.svc/json/ObtenerMunicipios", queryparams, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}

// Listado de los NÚMEROS de una vía
func ObtenerNumereroCodigos(req ObtenerNumereroCodigosRequest) (ObtenerNumereroResponse, error) {
	var response ObtenerNumereroResponse
	queryparams := map[string]string{
		"CodigoProvincia":    req.CodigoMunicipio,
		"CodigoMunicipio":    req.CodigoMunicipio,
		"CodigoMunicipioINE": req.CodigoMunicipioINE,
		"CodigoVia":          req.CodigoVia,
		"Numero":             req.Numero,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejeroCodigos.svc/json/ObtenerNumereroCodigos", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Listado de VÍAS de un municipio
func ObtenerCallejeroCodigos(req ObtenerCallejeroCodigosRequest) (ObtenerCallejeroResponse, error) {
	var response ObtenerCallejeroResponse
	queryparams := map[string]string{
		"CodigoProvincia":  req.CodigoProvincia,
		"CodigoMunicipio":  req.CodigoMunicipio,
		"CodigoMunipioNIE": req.CodigoMunicipioNIE,
		"CodigoVia":        req.CodigoVia,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejeroCodigos.svc/json/ObtenerCallejeroCodigos", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Consulta de DATOS CATASTRALES NO PROTEGIDOS de un inmueble identificado por su localización
func ConsultaDNPLOCCodigos(req ConsultaDNPLOCCodigosRequest) (ConsultaDNPLOCResponse, error) {
	var response ConsultaDNPLOCResponse
	queryparams := map[string]string{
		"CodigoProvincia":    req.CodigoProvincia,
		"CodigoMunicipio":    req.CodigoMunicipio,
		"CodigoMunicipioNIE": req.CodigoMunicipioNIE,
		"CodigoVia":          req.CodigoVia,
		"Numero":             req.Numero,
		"Bloque":             req.Bloque,
		"Escalera":           req.Escalera,
		"Planta":             req.Planta,
		"Puerta":             req.Puerta,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejeroCodigos.svc/json/Consulta_DNPLOC_Codigos", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Consulta de DATOS CATASTRALES NO PROTEGIDOS de un inmueble identificado por su Referencia Catastral
func ConsultaDNPRCCodigos(req ConsultaDNPRCCodigosRequest) (ConsultaDNPRCResponse, error) {
	var response ConsultaDNPRCResponse
	queryparams := map[string]string{
		"CodigoProvincia":    req.CodigoProvincia,
		"CodigoMunicipio":    req.CodigoMunicipio,
		"CodigoMunicipioNIE": req.CodigoMunicipioNIE,
		"RefCat":             req.ReferenciaCatastral,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejeroCodigos.svc/json/Consulta_DNPRC_Codigos", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

// Consulta de DATOS CATASTRALES NO PROTEGIDOS de un inmueble identificado por polígono-parcela
func ConsultaDNPPPCodigos(req ConsultaDNPPPCodigosRequest) (ConsultaDNPPPResponse, error) {
	var response ConsultaDNPPPResponse
	queryparams := map[string]string{
		"CodigoProvincia":    req.CodigoProvincia,
		"CodigoMunicipio":    req.CodigoMunicipio,
		"CodigoMunicipioINE": req.CodigoMunicipioINE,
		"Poligono":           req.Poligono,
		"Parcela":            req.Parcela,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCallejeroCodigos.svc/json/Consulta_DNPPP_Codigos", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}

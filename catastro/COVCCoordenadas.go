package catastro

import "errors"

// Servicio de consulta de Referencia Catastral por Coordenadas
func ConsultaRCCOOR(req ConsultaRCCOORRequest) (ConsultaRCCOORResponse, error) {
	var response ConsultaRCCOORResponse
	queryparams := map[string]string{
		"CoorX": req.CoorX,
		"CoorY": req.CoorY,
		"SRS":   req.SRS,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCoordenadas.svc/json/Consulta_RCCOOR", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ConsultaRCCOORResponse) hasFailed() (bool, error) {
	if r.ConsultaRCCOORResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaRCCOORResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Servicio de consulta lista de Referencias Catastrales por distancia a unas Coordenadas.
func ConsultaRCCOORDistancia(req ConsultaRCCOORDistanciaRequest) (ConsultaRCCOORDistanciaResponse, error) {
	var response ConsultaRCCOORDistanciaResponse
	queryparams := map[string]string{
		"CoorX": req.CoorX,
		"CoorY": req.CoorY,
		"SRS":   req.SRS,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCoordenadas.svc/json/Consulta_RCCOOR_Distancia", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ConsultaRCCOORDistanciaResponse) hasFailed() (bool, error) {
	if r.ConsultaRCCOORDistanciaResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaRCCOORDistanciaResult.ErrorList[0].Desc)
	}
	return false, nil
}

// Servicio de consulta de coordenadas por Provincia, Municipio y Referencia Catastral.
func ConsultaCPMRC(req ConsultaCPMRCRequest) (ConsultaCPMRCResponse, error) {
	var response ConsultaCPMRCResponse
	queryparams := map[string]string{
		"Provincia": req.Provincia,
		"Municipio": req.Municipio,
		"SRS":       req.SRS,
		"Parcela":   req.Parcela,
		"RefCat":    req.ReferenciaCatastral,
	}
	err := getJson("OVCServWeb/OVCWcfCallejero/COVCCoordenadas.svc/json/Consulta_CPMRC", queryparams, &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
func (r ConsultaCPMRCResponse) hasFailed() (bool, error) {
	if r.ConsultaCPMRCResult.Control.Cuerr != 0 {
		return true, errors.New(r.ConsultaCPMRCResult.ErrorList[0].Desc)
	}
	return false, nil
}

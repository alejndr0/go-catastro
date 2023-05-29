package catastro

type ErrorList struct {
	Err []struct {
		Cod  string `json:"cod"`
		Desc string `json:"des"`
	} `json:"err"`
}

type ObtenerProvinciasResponse struct {
	ConsultaProvincieroResult struct {
		Control struct {
			Cuprov int `json:"cuprov"`
			Cuerr  int `json:"cuerr"`
		} `json:"control"`
		Provinciero struct {
			Prov []struct {
				Cpine string `json:"cpine"`
				Np    string `json:"np"`
			} `json:"prov"`
		} `json:"provinciero"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_provincieroResult"`
}

type ObtenerMunicipiosResponse struct {
	ConsultaMunicipieroResult struct {
		Control struct {
			Cumun int `json:"cumun"`
			Cuerr int `json:"cuerr"`
		} `json:"control"`
		Municipiero struct {
			Muni []struct {
				Locat struct {
					Cd  string `json:"cd"`
					Cmc string `json:"cmc"`
				} `json:"locat"`
				Loine struct {
					Cp string `json:"cp"`
					Cm string `json:"cm"`
				} `json:"loine"`
				Nm string `json:"nm"`
			} `json:"muni"`
		} `json:"municipiero"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_municipieroResult"`
}

type ObtenerMunicipiosRequest struct {
	Provincia string
	Municipio string
}

type ObtenerMunicipiosCodigosRequest struct {
	CodigoMunicipio    string
	CodigoMunicipioINE string
	CodigoProvincia    string
}

type ObtenerNumereroResponse struct {
	ConsultaNumereroResult struct {
		Control struct {
			Cunum int `json:"cunum"`
			Cuerr int `json:"cuerr"`
		} `json:"control"`
		Nump []struct {
			Pc struct {
				Pc1 string `json:"pc1"`
				Pc2 string `json:"pc2"`
			} `json:"pc"`
			Num struct {
				Pnp string `json:"pnp"`
			} `json:"num"`
		} `json:"nump"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_numereroResult"`
}
type ObtenerNumereroRequest struct {
	Provincia string
	Municipio string
	TipoVia   string
	NombreVia string
	Numero    string
}
type ObtenerNumereroCodigosRequest struct {
	CodigoProvincia    string
	CodigoMunicipioINE string
	CodigoMunicipio    string
	CodigoVia          string
	Numero             string
}

type ObtenerCallejeroResponse struct {
	ConsultaCallejeroResult struct {
		Control struct {
			Cuca  int `json:"cuca"`
			Cuerr int `json:"cuerr"`
		} `json:"control"`
		Callejero struct {
			Calle []struct {
				Loine struct {
					Cp string `json:"cp"`
					Cm string `json:"cm"`
				} `json:"loine"`
				Dir struct {
					Cv string `json:"cv"`
					Tv string `json:"tv"`
					Nv string `json:"nv"`
				} `json:"dir"`
			} `json:"calle"`
		} `json:"callejero"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_callejeroResult"`
}

type ObtenerCallejeroRequest struct {
	Provincia string
	Municipio string
	TipoVia   string
	NombreVia string
}

type ObtenerCallejeroCodigosRequest struct {
	CodigoProvincia    string
	CodigoMunicipio    string
	CodigoMunicipioNIE string
	CodigoVia          string
}

type ConsultaDNPLOCResponse struct {
	ConsultaDnplocResult struct {
		Control struct {
			Cunum int `json:"cunum"`
			Cuerr int `json:"cuerr"`
		} `json:"control"`
		Numerero struct {
			Nump []struct {
				Pc struct {
					Pc1 string `json:"pc1"`
					Pc2 string `json:"pc2"`
				} `json:"pc"`
				Num struct {
					Pnp string `json:"pnp"`
				} `json:"num"`
			} `json:"nump"`
		} `json:"numerero"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_dnplocResult"`
}

type ConsultaDNPLOCRequest struct {
	Provincia string
	Municipio string
	Sigla     string
	Calle     string
	Numero    string
	Bloque    string
	Escalera  string
	Planta    string
	Puerta    string
}

type ConsultaDNPLOCCodigosRequest struct {
	CodigoProvincia    string
	CodigoMunicipio    string
	CodigoMunicipioNIE string
	CodigoVia          string
	Numero             string
	Bloque             string
	Escalera           string
	Planta             string
	Puerta             string
}

type ConsultaDNPRCResponse struct {
	ConsultaDnprcResult struct {
		Control struct {
			Cudnp  int `json:"cudnp"`
			Cucons int `json:"cucons"`
			Cuerr  int `json:"cuerr"`
		} `json:"control"`
		Bico struct {
			Bi struct {
				Idbi struct {
					Cn string `json:"cn"`
					Rc struct {
						Pc1 string `json:"pc1"`
						Pc2 string `json:"pc2"`
						Car string `json:"car"`
						Cc1 string `json:"cc1"`
						Cc2 string `json:"cc2"`
					} `json:"rc"`
				} `json:"idbi"`
				Dt struct {
					Loine struct {
						Cp string `json:"cp"`
						Cm string `json:"cm"`
					} `json:"loine"`
					Cmc  string `json:"cmc"`
					Np   string `json:"np"`
					Nm   string `json:"nm"`
					Locs struct {
						Lous struct {
							Lourb struct {
								Dir struct {
									Cv  string `json:"cv"`
									Tv  string `json:"tv"`
									Nv  string `json:"nv"`
									Pnp string `json:"pnp"`
									Snp string `json:"snp"`
								} `json:"dir"`
								Loint struct {
									Es string `json:"es"`
									Pt string `json:"pt"`
									Pu string `json:"pu"`
								} `json:"loint"`
								Dp string `json:"dp"`
								Dm string `json:"dm"`
							} `json:"lourb"`
						} `json:"lous"`
					} `json:"locs"`
				} `json:"dt"`
				Ldt  string `json:"ldt"`
				Debi struct {
					Luso string `json:"luso"`
					Sfc  string `json:"sfc"`
					Cpt  string `json:"cpt"`
					Ant  string `json:"ant"`
				} `json:"debi"`
			} `json:"bi"`
			Lcons []struct {
				Lcd string `json:"lcd"`
				Dt  struct {
					Lourb struct {
						Loint struct {
							Es string `json:"es"`
							Pt string `json:"pt"`
							Pu string `json:"pu"`
						} `json:"loint"`
					} `json:"lourb"`
				} `json:"dt,omitempty"`
				Dfcons struct {
					Stl string `json:"stl"`
				} `json:"dfcons"`
			} `json:"lcons"`
		} `json:"bico"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_dnprcResult"`
}

type ConsultaDNPRCRequest struct {
	Provincia           string
	Municipio           string
	ReferenciaCatastral string
}

type ConsultaDNPRCCodigosRequest struct {
	CodigoProvincia     string
	CodigoMunicipio     string
	CodigoMunicipioNIE  string
	ReferenciaCatastral string
}

type ConsultaDNPPPResponse struct {
	ConsultaDnpppResult struct {
		Control struct {
			Cudnp int `json:"cudnp"`
			Cucul int `json:"cucul"`
			Cuerr int `json:"cuerr"`
		} `json:"control"`
		Bico struct {
			Bi struct {
				Idbi struct {
					Cn string `json:"cn"`
					Rc struct {
						Pc1 string `json:"pc1"`
						Pc2 string `json:"pc2"`
						Car string `json:"car"`
						Cc1 string `json:"cc1"`
						Cc2 string `json:"cc2"`
					} `json:"rc"`
				} `json:"idbi"`
				Dt struct {
					Loine struct {
						Cp string `json:"cp"`
						Cm string `json:"cm"`
					} `json:"loine"`
					Cmc  string `json:"cmc"`
					Np   string `json:"np"`
					Nm   string `json:"nm"`
					Locs struct {
						Lors struct {
							Lorus struct {
								Czc string `json:"czc"`
								Cpp struct {
									Cpo string `json:"cpo"`
									Cpa string `json:"cpa"`
								} `json:"cpp"`
								Npa  string `json:"npa"`
								Cpaj string `json:"cpaj"`
							} `json:"lorus"`
						} `json:"lors"`
					} `json:"locs"`
				} `json:"dt"`
				Ldt  string `json:"ldt"`
				Debi struct {
					Luso string `json:"luso"`
					Sfc  string `json:"sfc"`
				} `json:"debi"`
			} `json:"bi"`
			Lspr []struct {
				Cspr string `json:"cspr"`
				Dspr struct {
					Ccc string `json:"ccc"`
					Dcc string `json:"dcc"`
					IP  string `json:"ip"`
					Ssp string `json:"ssp"`
				} `json:"dspr"`
			} `json:"lspr"`
		} `json:"bico"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"consulta_dnpppResult"`
}

type ConsultaDNPPPRequest struct {
	Provincia string
	Municipio string
	Poligono  string
	Parcela   string
}

type ConsultaDNPPPCodigosRequest struct {
	CodigoProvincia    string
	CodigoMunicipio    string
	CodigoMunicipioINE string
	Poligono           string
	Parcela            string
}

type ConsultaRCCOORResponse struct {
	ConsultaRCCOORResult struct {
		Control struct {
			Cucoor int `json:"cucoor"`
			Cuerr  int `json:"cuerr"`
		} `json:"control"`
		Coordenadas struct {
			Coord []struct {
				Pc struct {
					Pc1 string `json:"pc1"`
					Pc2 string `json:"pc2"`
				} `json:"pc"`
				Geo struct {
					Xcen string `json:"xcen"`
					Ycen string `json:"ycen"`
					Srs  string `json:"srs"`
				} `json:"geo"`
				Ldt string `json:"ldt"`
			} `json:"coord"`
		} `json:"coordenadas"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"Consulta_RCCOORResult"`
}

type ConsultaRCCOORRequest struct {
	CoorX string
	CoorY string
	SRS   string
}

type ConsultaRCCOORDistanciaResponse struct {
	ConsultaRCCOORDistanciaResult struct {
		Control struct {
			Cucoor int `json:"cucoor"`
			Cuerr  int `json:"cuerr"`
		} `json:"control"`
		CoordenadasDistancias struct {
			Coordd []struct {
				Geo struct {
					Xcen string `json:"xcen"`
					Ycen string `json:"ycen"`
					Srs  string `json:"srs"`
				} `json:"geo"`
				Lpcd []struct {
					Pc struct {
						Pc1 string `json:"pc1"`
						Pc2 string `json:"pc2"`
					} `json:"pc"`
					Dt struct {
						Loine struct {
							Cp string `json:"cp"`
							Cm string `json:"cm"`
						} `json:"loine"`
						Lourb struct {
							Dir struct {
								Cv  string `json:"cv"`
								Pnp string `json:"pnp"`
							} `json:"dir"`
						} `json:"lourb"`
					} `json:"dt"`
					Ldt string `json:"ldt"`
					Dis string `json:"dis"`
				} `json:"lpcd"`
			} `json:"coordd"`
		} `json:"coordenadas_distancias"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"Consulta_RCCOOR_DistanciaResult"`
}
type ConsultaRCCOORDistanciaRequest struct {
	CoorX string
	CoorY string
	SRS   string
}

type ConsultaCPMRCResponse struct {
	ConsultaCPMRCResult struct {
		Control struct {
			Cucoor int `json:"cucoor"`
			Cuerr  int `json:"cuerr"`
		} `json:"control"`
		Coordenadas struct {
			Coord []struct {
				Pc struct {
					Pc1 string `json:"pc1"`
					Pc2 string `json:"pc2"`
				} `json:"pc"`
				Geo struct {
					Xcen string `json:"xcen"`
					Ycen string `json:"ycen"`
					Srs  string `json:"srs"`
				} `json:"geo"`
				Ldt string `json:"ldt"`
			} `json:"coord"`
		} `json:"coordenadas"`
		ErrorList ErrorList `json:"lerr"`
	} `json:"Consulta_CPMRCResult"`
}

type ConsultaCPMRCRequest struct {
	Provincia           string
	Municipio           string
	SRS                 string
	Parcela             string
	ReferenciaCatastral string
}

package domain

import "context"

// type BuscaCEPAPIResponse struct {
// 	APIName   string             `json:"api_name"`
// 	BrasilAPI *BrasilAPIResponse `json:"brasilapi_cep_info,omitempty"`
// 	ViaCEPAPI *ViaCEPAPIResponse `json:"viaapi_cep_info,omitempty"`
// 	Error     string             `json:"error,omitempty"`
// }

type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

type ViaCEPAPIResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type BuscaCEPAPIClient interface {
	GetBrasilAPICEP(ctx context.Context, cep string) (*BrasilAPIResponse, error)
	GetViaAPICEP(ctx context.Context, cep string) (*ViaCEPAPIResponse, error)
}

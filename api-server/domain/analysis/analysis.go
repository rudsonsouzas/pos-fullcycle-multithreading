package analysis

import (
	"api-server/domain"
	"context"
	"fmt"
	"log"
	"time"
)

type analysisService struct {
	buscaCEPAPIClient domain.BuscaCEPAPIClient
	log               *log.Logger
}

func NewAnalysisService(buscaCEPAPIClient domain.BuscaCEPAPIClient, log *log.Logger) *analysisService {

	return &analysisService{
		buscaCEPAPIClient: buscaCEPAPIClient,
		log:               log,
	}
}

func (s *analysisService) RunAnalysis(c context.Context, cep string) (string, error) {
	type result struct {
		Source string
		Data   interface{}
		Err    error
	}

	// Timeout de 1 segundo para chamada das APIs
	apiCtx, apiCancel := context.WithTimeout(c, 1*time.Second)
	defer apiCancel()

	resultCh := make(chan result, 2)

	// Chamada BrasilAPI
	go func() {
		data, err := s.buscaCEPAPIClient.GetBrasilAPICEP(apiCtx, cep)
		resultCh <- result{Source: "BrasilAPI", Data: data, Err: err}
	}()

	// Chamada ViaCEP
	go func() {
		data, err := s.buscaCEPAPIClient.GetViaAPICEP(apiCtx, cep)
		resultCh <- result{Source: "ViaCEP", Data: data, Err: err}
	}()

	select {
	case res := <-resultCh:
		if res.Err != nil {
			s.log.Printf("Erro ao buscar CEP na API %s: %v", res.Source, res.Err)
			return "", res.Err
		}
		s.log.Printf("Resposta recebida da API %s: %+v", res.Source, res.Data)
		return fmt.Sprintf("Fonte: %s, Dados: %+v", res.Source, res.Data), nil
	case <-apiCtx.Done():
		s.log.Printf("Timeout ao buscar CEP nas APIs")
		return "", apiCtx.Err()
	}

}

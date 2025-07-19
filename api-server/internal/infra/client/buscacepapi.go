package client

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"api-server/domain"
	httpclient "api-server/pkg/http_client"

	"github.com/cenkalti/backoff"
)

type BuscaCEPAPIClient struct {
	httpClient httpclient.HTTPClient
	log        *log.Logger
}

// type BuscaCEPAPI interface {
// 	GetDolarQuote(c context.Context) (*domain.BuscaCEPAPIResponse, error)
// }

func NewBuscaCEPAPIClient(httpClient httpclient.HTTPClient, log *log.Logger) *BuscaCEPAPIClient {
	return &BuscaCEPAPIClient{
		httpClient: httpClient,
		log:        log,
	}
}

func (awc *BuscaCEPAPIClient) getCEP(ctx context.Context, url string) ([]byte, error) {
	var resBody []byte

	ebo := backoff.NewExponentialBackOff()
	ebo.MaxInterval = 1 * time.Second

	if err := backoff.Retry(func() error {

		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			awc.log.Printf("erro ao criar a requisição para buscar o CEP através da URL: %s. [Erro]: %s", url, err.Error())
			return err
		}

		req.Header.Add("Content-Type", "application/json")

		res, err := awc.httpClient.Do(req)
		if err != nil {
			awc.log.Printf("erro ao realizar a requisição para buscar o CEP através da URL: %s. [Erro]: %s", url, err.Error())
			return err
		}
		defer func() {
			err = res.Body.Close()
			if err != nil {
				awc.log.Printf("erro ao encerrar o response body da consulta através da URL: %s. [Erro]: %s", url, err.Error())
				return
			}
		}()

		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			awc.log.Printf("erro ao ler o response body através da URL: %s. [Erro]: %s", url, err.Error())
			return err
		}

		if res.StatusCode != 200 {
			awc.log.Printf("Busca CEP através da URL [%s]- API status code %d: %s", url, res.StatusCode, bodyBytes)
			return err
		}

		resBody = bodyBytes
		return nil

	}, backoff.WithContext(backoff.WithMaxRetries(ebo, uint64(5)), context.Background())); err != nil {
		awc.log.Printf("erro ao requisitar a busca pela CEP através da URL: %s. [Erro]: %s", url, err.Error())
		return []byte{}, err
	}

	return resBody, nil
}

func (awc *BuscaCEPAPIClient) GetBrasilAPICEP(ctx context.Context, cep string) (*domain.BrasilAPIResponse, error) {
	var brasilAPIResponse *domain.BrasilAPIResponse
	brasilAPIUrl := "https://brasilapi.com.br/api/cep/v1/" + cep

	resBody, err := awc.getCEP(ctx, brasilAPIUrl)
	if err != nil {
		awc.log.Printf("erro ao buscar o CEP na BrasilAPI: %s. [Erro]: %s", cep, err.Error())
		return nil, err
	}

	// bodyBytes, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	awc.log.Printf("erro ao ler o response body oriundo da BrasilAPI. [Erro]: %s", err.Error())
	// 	return nil, err
	// }

	if err := json.Unmarshal(resBody, &brasilAPIResponse); err != nil {
		awc.log.Printf("erro ao traduzir o response Body para o padrão esperado na BrasilAPI: %s", err.Error())
		return nil, err
	}

	return brasilAPIResponse, nil
}

func (awc *BuscaCEPAPIClient) GetViaAPICEP(ctx context.Context, cep string) (*domain.ViaCEPAPIResponse, error) {
	var vaiAPIResponse *domain.ViaCEPAPIResponse
	viaAPIUrl := "https://viacep.com.br/ws/" + cep + "/json/"

	resBody, err := awc.getCEP(ctx, viaAPIUrl)
	if err != nil {
		awc.log.Printf("erro ao buscar o CEP na ViaAPI: %s. [Erro]: %s", cep, err.Error())
		return nil, err
	}

	// bodyBytes, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	awc.log.Printf("erro ao ler o response body oriundo da ViaAPI. [Erro]: %s", err.Error())
	// 	return nil, err
	// }

	if err := json.Unmarshal(resBody, &vaiAPIResponse); err != nil {
		awc.log.Printf("erro ao traduzir o response Body para o padrão esperado na ViaAPI: %s", err.Error())
		return nil, err
	}

	return vaiAPIResponse, nil
}

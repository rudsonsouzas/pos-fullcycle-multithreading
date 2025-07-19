package domain

import "context"

type AnalysisService interface {
	RunAnalysis(c context.Context, cep string) (string, error)
}

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) RunAnalysis(c *gin.Context) {
	// Call the analysis service to run
	cepInfo, err := h.analisysService.RunAnalysis(c.Request.Context(), c.Param("cep"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "falha ao buscar o CEP: " + c.Param("cep") + ".", "detalhes": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "CEP encontrado com sucesso!", "data": cepInfo})
	c.Next()
}

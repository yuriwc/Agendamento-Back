package controllers

import (
	"back_agendamento/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmpresaController struct {
	NomeEmpresa      string `json:"nome_empresa" binding:"required"`
	DescricaoEmpresa string `json:"descricao_empresa" binding:"required"`
	HoraInicio       string `json:"hora_inicio" binding:"required"`
	HoraFim          string `json:"hora_fim" binding:"required"`
	Foto             string `json:"foto"`
	IntervaloServico string `json:"intervalo_servico" binding:"required"`
}

func saveEmpresaToDB(empresa EmpresaController, c *gin.Context) (models.Empresa, error) {
	var empresaDB models.Empresa

	empresaDB.NomeEmpresa = empresa.NomeEmpresa
	empresaDB.DescricaoEmpresa = empresa.DescricaoEmpresa
	empresaDB.HoraInicio = empresa.HoraInicio
	empresaDB.HoraFim = empresa.HoraFim
	empresaDB.Foto = empresa.Foto
	empresaDB.IntervaloServico = empresa.IntervaloServico

	result, erro := models.CriarEmpresa(empresaDB)
	if erro != nil {
		c.JSON(400, gin.H{"error": erro.Error()})
		return empresaDB, erro
	}

	return result, nil
}

func CriarEmpresa(c *gin.Context) {
	var empresa EmpresaController
	err := c.BindJSON(&empresa)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	_, erro := saveEmpresaToDB(empresa, c)
	if erro != nil {
		return
	}

	c.JSON(200, gin.H{"message": "Empresa criada com sucesso!", "data": empresa})
}

func BuscarEmpresa(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 64)
	empresa, err := models.BuscarEmpresa(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": empresa})
}

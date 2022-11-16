package controllers

import (
	"back_agendamento/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClienteController struct {
	Cpf           string `json:"cpf" binding:"required"`
	PrimeiroNome  string `json:"primeiroNome" binding:"required"`
	UltimoNome    string `json:"ultimoNome" binding:"required"`
	Celular       string `json:"celular" binding:"required"`
	Latitude      string `json:"latitude"`
	Longitude     string `json:"longitude"`
	Foto          string `json:"foto"`
	Email         string `json:"email" binding:"required"`
	IdTipoCliente uint   `json:"idTipoCliente" binding:"required"`
}

func saveToDB(cliente ClienteController, c *gin.Context) (models.Cliente, error) {
	var clienteDB models.Cliente

	clienteDB.PrimeiroNome = cliente.PrimeiroNome
	clienteDB.UltimoNome = cliente.UltimoNome
	clienteDB.Cpf = cliente.Cpf
	clienteDB.Celular = cliente.Celular
	clienteDB.Latitude = cliente.Latitude
	clienteDB.Longitude = cliente.Longitude
	clienteDB.Foto = cliente.Foto
	clienteDB.Email = cliente.Email
	clienteDB.IdTipoCliente = cliente.IdTipoCliente

	result, erro := models.CriarCliente(clienteDB)
	if erro != nil {
		c.JSON(400, gin.H{"error": erro.Error()})
		return clienteDB, erro
	}

	return result, nil
}

func CriarCliente(c *gin.Context) {
	var cliente ClienteController
	err := c.BindJSON(&cliente)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Cliente criado com sucesso!", "data": cliente})
}

func BuscarCliente(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 0, 64)

	cliente, err := models.BuscarCliente(uint(id))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": cliente})
}

func BuscarTodosClientes(c *gin.Context) {
	clientes, err := models.BuscarTodosClientes()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": clientes})
}

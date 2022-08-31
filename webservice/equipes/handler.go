package equipes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Brun0Nasc/sys-projetos/domain/equipes"
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"

	"github.com/gin-gonic/gin"
)

func novaEquipe(c *gin.Context) {
	fmt.Println("Tentando adicionar nova equipe")
	req := modelApresentacao.ReqEquipe{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"message": "Could not create. Parameters were not passed correctly " + err.Error(),
		})
		return
	}

	equipe.NovaEquipe(&req, c)
	c.JSON(http.StatusCreated, gin.H{"OK":"Registro adicionado com Sucesso!"})
}

func listarEquipes(c *gin.Context) {
	fmt.Println("Tentando listar equipes") 
	if equipes, err := equipe.ListarEquipes(); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, equipes)
	}
}

func buscarEquipe(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("Tentando encontrar equipe")
	if equipe, err := equipe.BuscarEquipe(id); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"message":"Nenhum registro encontrado", "err":err.Error()})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error":"" + err.Error()})
		}
	} else {
		c.JSON(http.StatusOK, equipe)
	}
}
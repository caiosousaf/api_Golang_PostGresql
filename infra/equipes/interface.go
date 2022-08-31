package equipes

import (
	modelApresentacao "github.com/Brun0Nasc/sys-projetos/domain/equipes/model"
	"github.com/gin-gonic/gin"
)

type IEquipe interface {
	NovaEquipe(req *modelApresentacao.ReqEquipe, c *gin.Context)
	ListarEquipes() ([]modelApresentacao.ReqEquipe, error)
	BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error)
}
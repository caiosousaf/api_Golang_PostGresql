package projetos

import (
	"gerenciadorDeProjetos/config/database"
	modelApresentacao "gerenciadorDeProjetos/domain/projetos/model"
	"gerenciadorDeProjetos/infra/projetos"
	"github.com/gin-gonic/gin"
)

func NovoProjeto(req *modelApresentacao.ReqProjeto, c *gin.Context) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	projetosRepo.NovoProjeto(req, c)
}

func ListarProjetos() ([]modelApresentacao.ReqProjetos, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.ListarProjetos()
}

func ListarProjeto(id string) (*modelApresentacao.ReqProjetos, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)

	return projetosRepo.ListarProjeto(id)
}

func ListarProjetosComStatus(status string) ([]modelApresentacao.ReqStatusProjeto, error) {
	db := database.Conectar()
	defer db.Close()
	projetosRepo := projetos.NovoRepo(db)
	
	return projetosRepo.ListarProjetosComStatus(status)
}
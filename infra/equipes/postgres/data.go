package postgres

import (
	"database/sql"
	"fmt"

	modelApresentacao "gerenciadorDeProjetos/domain/equipes/model"
	modelPessoa "gerenciadorDeProjetos/domain/pessoas/model"
	modelData "gerenciadorDeProjetos/infra/equipes/model"
)

type DBEquipes struct {
	DB *sql.DB
}

func (postgres *DBEquipes) NovaEquipe(req *modelData.Equipe) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `INSERT INTO equipes
	(nome_equipe)
	VALUES($1::TEXT) RETURNING *`

	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		return nil, err
	}
	return equipe, nil
}

func (postgres *DBEquipes) ListarEquipes() ([]modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes ORDER BY id_equipe`
	var res = []modelApresentacao.ReqEquipe{}
	var equipe = modelApresentacao.ReqEquipe{}

	rows, err := postgres.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Listagem de todas as equipes deu certo!")
	return res, nil
}

func (postgres *DBEquipes) BuscarEquipe(id string) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `SELECT * FROM equipes WHERE id_equipe = $1`
	var equipe = &modelApresentacao.ReqEquipe{}

	pessoas, err := postgres.BuscarMembrosDeEquipe(id)
	if err != nil {
		if err == sql.ErrNoRows {
			pessoas = nil
		} else {
			return nil, err
		}
	}
	projetos, err := postgres.BuscarProjetosDeEquipe(id)

	if err != nil {
		if err == sql.ErrNoRows {
			projetos = nil
		} else {
			return nil, err
		}
	}

	tasks, err := postgres.BuscarTasksDeEquipe(id)

	if err != nil {
		if err == sql.ErrNoRows {
			projetos = nil
		} else {
			return nil, err
		}
	}

	equipe.Pessoas = &pessoas
	equipe.Projetos = &projetos
	equipe.Tarefas = &tasks

	row := postgres.DB.QueryRow(sqlStatement, id)
	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}
	fmt.Println("Busca deu certo!")
	return equipe, nil
}

func (postgres *DBEquipes) BuscarMembrosDeEquipe(id string) ([]modelPessoa.ReqMembros, error) {
	sqlStatement := `select id_pessoa, nome_pessoa, funcao_pessoa, equipe_id, data_contratacao from pessoas WHERE equipe_id = $1`
	var res = []modelPessoa.ReqMembros{}
	var equipe = modelPessoa.ReqMembros{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		if err := rows.Scan(&equipe.ID_Pessoa, &equipe.Nome_Pessoa, &equipe.Funcao_Pessoa, &equipe.Equipe_ID, &equipe.Data_Contratacao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Busca de membros de uma equipe deu certo!")
	return res, nil
}

func (postgres *DBEquipes) BuscarProjetosDeEquipe(id string) ([]modelApresentacao.ReqEquipeProjetos, error) {
	sqlStatement := `select eq.nome_equipe, pr.id_projeto, pr.nome_projeto, pr.status, pr.descricao_projeto, pr.data_criacao, pr.data_conclusao, pr.prazo_entrega 
	from equipes as eq 
	inner join projetos as pr on eq.id_equipe = pr.equipe_id where eq.id_equipe = $1`
	var res = []modelApresentacao.ReqEquipeProjetos{}
	var equipe = modelApresentacao.ReqEquipeProjetos{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {

		if err := rows.Scan(&equipe.Nome_Equipe, &equipe.ID_Projeto, &equipe.Nome_Projeto, &equipe.Status, &equipe.Descricao_Projeto,
			&equipe.Data_Criacao, &equipe.Data_Conclusao, &equipe.Prazo_Entrega); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Busca dos projetos de uma equipe deu certo!")
	return res, nil
}

func (postgres *DBEquipes) BuscarTasksDeEquipe(id string) ([]modelApresentacao.ReqTasksbyTeam, error) {
	sqlStatement := `select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, tk.status,
	tk.data_criacao, tk.prazo_entrega, tk.data_conclusao
					 from tasks tk 
					 inner join pessoas pe on pe.id_pessoa = tk.pessoa_id 
					 inner join equipes eq on eq.id_equipe = pe.equipe_id	
					 where eq.id_equipe = $1`
	var res = []modelApresentacao.ReqTasksbyTeam{}
	var equipe = modelApresentacao.ReqTasksbyTeam{}

	rows, err := postgres.DB.Query(sqlStatement, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(&equipe.ID_Task, &equipe.Descricao_Task, &equipe.Pessoa_ID, &equipe.Nome_Pessoa, &equipe.Projeto_ID,
			&equipe.Status, &equipe.Data_Criacao, &equipe.Prazo_Entrega, &equipe.Data_Conclusao); err != nil {
			if err == sql.ErrNoRows {
				return nil, err
			} else {
				return nil, err
			}
		}
		res = append(res, equipe)
	}
	fmt.Println("Busca das tarefas de uma equipe deu certo!")
	return res, nil
}

func (postgres *DBEquipes) DeletarEquipe(id string) error {
	sqlStatement := `DELETE FROM equipes WHERE id_equipe = $1`

	_, err := postgres.DB.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	fmt.Println("Tudo certo em deletar uma equipe!!")
	return nil
}

func (postgres *DBEquipes) AtualizarEquipe(id string, req *modelData.UpdateEquipe) (*modelApresentacao.ReqEquipe, error) {
	sqlStatement := `UPDATE equipes SET nome_equipe = $1 
	WHERE id_equipe = $2 RETURNING *`
	var equipe = &modelApresentacao.ReqEquipe{}

	row := postgres.DB.QueryRow(sqlStatement, req.Nome_Equipe, id)

	if err := row.Scan(&equipe.ID_Equipe, &equipe.Nome_Equipe, &equipe.Data_Criacao); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		} else {
			return nil, err
		}
	}

	return equipe, nil
}

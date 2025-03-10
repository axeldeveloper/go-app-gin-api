package repository

import (
	"database/sql"
	"errors"
	"example-go-gin/domain"
)

type CompanyRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *CompanyRepository {
	return &CompanyRepository{db}
}

func (cr *CompanyRepository) GetCompanies() ([]domain.Company, error) {
	var companies []domain.Company

	rows, err := cr.db.Query("SELECT id, cnpj_basico, razao_social, natureza_juridica, qualificacao_responsavel, capital, porte, ente_federativo FROM company")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var company domain.Company
		if err := rows.Scan(
			&company.ID,
			&company.CNPJBasico,
			&company.RazaoSocial,
			&company.NaturezaJuridica,
			&company.QualificacaoResponsavel,
			&company.Capital, &company.Porte,
			&company.EnteFederativo); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return companies, nil
}

func (ur *CompanyRepository) Save(user domain.Company) error {
	// Implementação para salvar um usuário no banco de dados
	// Exemplo: _, err := ur.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return nil
}

func (ur *CompanyRepository) FindByID(id int) (*domain.Company, error) {
	// Implementação para buscar um usuário pelo ID no banco de dados
	// Exemplo: row := ur.db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", id)
	// var user domain.User
	// err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	// return &user, err
	return nil, errors.New("not implemented")
}

func (ur *CompanyRepository) Update(user domain.Company) error {
	// Implementação para atualizar um usuário no banco de dados
	// Exemplo: _, err := ur.db.Exec("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4", user.Name, user.Email, user.Password, user.ID)
	return nil
}

func (ur *CompanyRepository) Delete(id int) error {
	// Implementação para deletar um usuário do banco de dados
	// Exemplo: _, err := ur.db.Exec("DELETE FROM users WHERE id = $1", id)
	return nil
}

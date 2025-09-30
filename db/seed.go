package db

import (
	"database/sql"
	"example-go-gin/domain"
	"log"
)

func SeedCompanies(db *sql.DB) {
	companies := []domain.Company{
		{RazaoSocial: "Empresa Alpha"},
		{RazaoSocial: "Empresa Beta"},
		{RazaoSocial: "Empresa Gamma"},
	}

	for _, company := range companies {
		_, err := db.Exec("INSERT INTO companies (RazaoSocial) VALUES ($1) ON CONFLICT (name) DO NOTHING", company.RazaoSocial)
		if err != nil {
			log.Printf("Erro ao inserir empresa %s: %v", company.RazaoSocial, err)
		}
	}
	log.Println("Seed de empresas concluído!")
}

package domain

// todo represents data about a task in the Company list
type Company struct {
	ID                      string `json:"id"`
	CNPJBasico              string `json:"cnpj_basico"`
	RazaoSocial             string `json:"razao_social"`
	NaturezaJuridica        string `json:"natureza_juridica"`
	QualificacaoResponsavel string `json:"qualificacao_responsavel"`
	Capital                 string `json:"capital"`
	Porte                   string `json:"porte"`
	EnteFederativo          string `json:"ente_federativo"`
}

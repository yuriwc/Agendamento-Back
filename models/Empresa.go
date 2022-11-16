package models

import "back_agendamento/services"

type Empresa struct {
	IdEmpresa        uint   `db:"idEmpresa"`
	NomeEmpresa      string `db:"nomeEmpresa"`
	DescricaoEmpresa string `db:"descricaoEmpresa"`
	HoraInicio       string `db:"horaInicio"`
	HoraFim          string `db:"horaFim"`
	Foto             string `db:"foto"`
	IntervaloServico string `db:"intervaloServico"`
}

func CriarEmpresa(empresa Empresa) (Empresa, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("INSERT INTO empresa (nomeEmpresa, descricaoEmpresa, horaInicio, horaFim, foto, intervaloServico) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Empresa{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(empresa.NomeEmpresa, empresa.DescricaoEmpresa, empresa.HoraInicio, empresa.HoraFim, empresa.Foto, empresa.IntervaloServico)
	if err != nil {
		return Empresa{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Empresa{}, err
	}

	empresa.IdEmpresa = uint(id)

	return empresa, nil
}

func BuscarEmpresa(id uint) (Empresa, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("SELECT * FROM empresa WHERE idEmpresa = ?")
	if err != nil {
		return Empresa{}, err
	}
	defer stmt.Close()

	empresa := Empresa{}
	err = stmt.QueryRow(id).Scan(&empresa.IdEmpresa, &empresa.NomeEmpresa, &empresa.DescricaoEmpresa, &empresa.HoraInicio, &empresa.HoraFim, &empresa.Foto, &empresa.IntervaloServico)
	if err != nil {
		return Empresa{}, err
	}

	return empresa, nil
}

func BuscarTodasEmpresas() ([]Empresa, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("SELECT * FROM empresa")
	if err != nil {
		return []Empresa{}, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return []Empresa{}, err
	}
	defer rows.Close()

	var empresas []Empresa
	for rows.Next() {
		empresa := Empresa{}
		err = rows.Scan(&empresa.IdEmpresa, &empresa.NomeEmpresa, &empresa.DescricaoEmpresa, &empresa.HoraInicio, &empresa.HoraFim, &empresa.Foto, &empresa.IntervaloServico)
		if err != nil {
			return []Empresa{}, err
		}
		empresas = append(empresas, empresa)
	}

	return empresas, nil
}

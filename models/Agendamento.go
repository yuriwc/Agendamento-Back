package models

import "back_agendamento/services"

type Agendamento struct {
	IdAgendamento uint   `db:"idAgendamento"`
	IdCliente     uint   `db:"idCliente"`
	IdEmpresa     uint   `db:"idEmpresa"`
	HoraInicio    string `db:"horaInicio"`
	Data          string `db:"data"`
	HoraFim       string `db:"horaFim"`
}

func CriarAgendamento(agendamento Agendamento) (Agendamento, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("INSERT INTO agendamento (idCliente, idEmpresa, horaInicio, data, horaFim) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return Agendamento{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(agendamento.IdCliente, agendamento.IdEmpresa, agendamento.HoraInicio, agendamento.Data, agendamento.HoraFim)
	if err != nil {
		return Agendamento{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Agendamento{}, err
	}

	agendamento.IdAgendamento = uint(id)

	return agendamento, nil
}

func ApagarAgendamento(id uint) error {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("DELETE FROM agendamento WHERE idAgendamento = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

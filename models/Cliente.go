package models

import "back_agendamento/services"

type Cliente struct {
	IdCliente     uint   `db:"idCliente"`
	PrimeiroNome  string `db:"primeiroNome"`
	UltimoNome    string `db:"ultimoNome"`
	Cpf           string `db:"cpf"`
	Celular       string `db:"celular"`
	Latitude      string `db:"latitude"`
	Longitude     string `db:"longitude"`
	Foto          string `db:"foto"`
	Email         string `db:"email"`
	IdTipoCliente uint   `db:"idTipoCliente"`
}

func CriarCliente(cliente Cliente) (Cliente, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("INSERT INTO cliente (primeiroNome, ultimoNome, cpf, celular, latitude, longitude, foto, email, idTipoCliente) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return Cliente{}, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(cliente.PrimeiroNome, cliente.UltimoNome, cliente.Cpf, cliente.Celular, cliente.Latitude, cliente.Longitude, cliente.Foto, cliente.Email, cliente.IdTipoCliente)
	if err != nil {
		return Cliente{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return Cliente{}, err
	}

	cliente.IdCliente = uint(id)

	return cliente, nil
}

func BuscarCliente(id uint) (Cliente, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("SELECT * FROM cliente WHERE idCliente = ?")
	if err != nil {
		return Cliente{}, err
	}
	defer stmt.Close()

	cliente := Cliente{}
	err = stmt.QueryRow(id).Scan(&cliente.IdCliente, &cliente.PrimeiroNome, &cliente.UltimoNome, &cliente.Cpf, &cliente.Celular, &cliente.Latitude, &cliente.Longitude, &cliente.Foto, &cliente.Email, &cliente.IdTipoCliente)
	if err != nil {
		return Cliente{}, err
	}

	return cliente, nil
}

func BuscarTodosClientes() ([]Cliente, error) {
	db := services.ConnectToDB()

	stmt, err := db.Prepare("SELECT * FROM cliente")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clientes := []Cliente{}
	for rows.Next() {
		cliente := Cliente{}
		err = rows.Scan(&cliente.IdCliente, &cliente.PrimeiroNome, &cliente.UltimoNome, &cliente.Cpf, &cliente.Celular, &cliente.Latitude, &cliente.Longitude, &cliente.Foto, &cliente.Email, &cliente.IdTipoCliente)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}

	return clientes, nil
}

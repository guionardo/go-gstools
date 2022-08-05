package cep

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/guionardo/go-gstools/tools"
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

type CEPRepositorySQLite struct {
	db   *sql.DB
	lock sync.Mutex
}

func (r *CEPRepositorySQLite) GetCEP(cep string) (*CEP, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	cep = tools.JustNumbers(cep)
	row, err := r.db.Query("SELECT logradouro,tipo_logradouro,bairro,municipio,uf,data_requisicao FROM cep WHERE cep = ?", cep)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var log, tipo, bairro, municipio, uf string
		var dataRequisicao time.Time
		row.Scan(&log, &tipo, &bairro, &municipio, &uf, &dataRequisicao)
		return &CEP{
			CEP:            cep,
			Logradouro:     log,
			TipoLogradouro: tipo,
			Bairro:         bairro,
			Municipio:      municipio,
			UF:             uf,
			DataRequisicao: dataRequisicao,
		}, nil
	}
	return nil, fmt.Errorf("CEP %s não encontrado", cep)
}

func (r *CEPRepositorySQLite) SaveCEP(cep *CEP) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	insertStudentSQL := `INSERT INTO cep(cep, logradouro, tipo_logradouro, bairro, municipio, uf, data_requisicao) 
	VALUES (?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(cep) DO UPDATE SET logradouro = ?, tipo_logradouro = ?, bairro = ?, municipio = ?, uf = ?, data_requisicao = ?`
	statement, err := r.db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		return err
	}
	_, err = statement.Exec(tools.JustNumbers(cep.CEP), cep.Logradouro, cep.TipoLogradouro, cep.Bairro, cep.Municipio, cep.UF, cep.DataRequisicao,
		cep.Logradouro, cep.TipoLogradouro, cep.Bairro, cep.Municipio, cep.UF, cep.DataRequisicao) // Execute statement
	if err != nil {
		return err
	}
	return nil
}

func NewCEPRepositorySQLite(connectionString string) (repo *CEPRepositorySQLite, err error) {
	repo = &CEPRepositorySQLite{}
	repo.db, err = sql.Open("sqlite3", connectionString)
	if err != nil {
		return
	}
	repo.lock.Lock()
	defer repo.lock.Unlock()
	createTable(repo.db)

	return repo, nil
}

func createTable(db *sql.DB) error {

	createCepTableSQL := `CREATE TABLE IF NOT EXISTS cep (
		cep TEXT NOT NULL PRIMARY KEY,		
		logradouro TEXT,
		tipo_logradouro TEXT,
		bairro TEXT,
		municipio TEXT,
		uf TEXT,
		data_requisicao DATETIME
	  );` // SQL Statement for Create Table

	statement, err := db.Prepare(createCepTableSQL) // Prepare SQL Statement
	if err != nil {
		return err
	}
	_, err = statement.Exec() // Execute SQL Statements
	return err
}

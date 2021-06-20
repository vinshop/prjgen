package prjgen

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinshop/prjgen/pkg/logger"
)

const (
	DriverName = "mysql"

	TableNameField = "Tables_in_db"
)

type Generator struct {
	Protocol string
	DB       string
	Host     string
	Port     int
	User     string
	Password string
	Tables   []string

	ModelDir      string
	RepoDir       string
	ServiceDir    string
	ControllerDir string
	DtoDir        string
	MapperDir     string
	OutputDir     string

	db *sql.DB

	tables []string
}

func (g *Generator) exec() error {
	if err := g.connect(); err != nil {
		return err
	}
	defer func() {
		if err := g.close(); err != nil {
			logger.Errorw("error when close connection", "error", err)
		}
	}()

	if err := g.getAllDBTable(); err != nil {
		return err
	}

	return nil
}

func (g *Generator) connect() error {
	connString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", g.User, g.Password, g.Protocol, g.Host, g.Port, g.DB)
	logger.Infow("connecting to database", "connection string", connString)
	db, err := sql.Open(DriverName, connString)
	if err != nil {
		return err
	}
	g.db = db
	return nil
}

func (g *Generator) close() error {
	return g.db.Close()
}

func (g *Generator) getAllDBTable() error {
	rows, err := g.db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	g.tables = make([]string, 0)
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return err
		}
		g.tables = append(g.tables, table)
	}
	logger.Infow("list all table", "tables", g.tables)

	return nil
}

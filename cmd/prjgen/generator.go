package prjgen

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinshop/prjgen/cmd/prjgen/models"
	"github.com/vinshop/prjgen/pkg/logger"
	"github.com/vinshop/prjgen/pkg/util"
	"os"
	"os/exec"
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

	module string

	db *sql.DB

	tables []string
}

func (g *Generator) exec() error {

	if err := g.getModule(); err != nil {
		logger.Errorw("error when get current module", "error", err)
		return err
	}

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

	for _, table := range g.tables {
		if err := g.getTable(table); err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) getModule() error {
	cmd := exec.Command("bash", "-c", "go mod edit -json > gomod.json")
	if err := cmd.Run(); err != nil {
		return err
	}
	defer func() {
		if err := exec.Command("bash", "-c", "rm -f gomod.json").Run(); err != nil {
			logger.Errorw("error when remove gomod.json", "error", err)
		}
	}()

	f, err := os.Open("gomod.json")
	if err != nil {
		return err
	}
	defer f.Close()

	var mod struct {
		Module struct {
			Path string
		}
	}

	if err := json.NewDecoder(f).Decode(&mod); err != nil {
		return err
	}

	g.module = mod.Module.Path

	if g.OutputDir != "." {
		g.module += "/" + g.OutputDir
	}

	logger.Info("current module: ", g.module)

	return nil
}

func (g *Generator) connect() error {
	connString := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", g.User, g.Password, g.Protocol, g.Host, g.Port, "information_schema")
	logger.Info("connecting to database with connection string: ", connString)
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
	q, err := g.db.Prepare("SELECT `TABLE_NAME` FROM `TABLES` WHERE `TABLE_SCHEMA` = ?")
	if err != nil {
		logger.Errorw("Error when create prepare statement", "error", err)
		return err
	}
	rows, err := q.Query(g.DB)
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

func (g *Generator) getTable(name string) error {
	q, err := g.db.Prepare("SELECT `COLUMN_NAME`, `DATA_TYPE`, `IS_NULLABLE`, `COLUMN_KEY` FROM `COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` =?")
	if err != nil {
		logger.Errorw("Error when create prepare statement", "error", err)
		return err
	}

	rows, err := q.Query(g.DB, name)
	if err != nil {
		return err
	}

	//m := models.Model{
	//	Base: models.Base{
	//		Pkg: g.module,
	//		Dir: g.module + "/" + g.ModelDir,
	//	},
	//	Name:  util.ToSnakeCase(name),
	//	Table: name,
	//}

	for rows.Next() {

		f := models.ModelField{

		}

		var rName, rType, rNullable, rColKey string

		if err := rows.Scan(&rName, &rType, &rNullable, &rColKey); err != nil {
			return err
		}

		f.Name = util.ToSnakeCase()

		logger.Info(rName, rType, rNullable, rColKey)

	}
	return nil

}

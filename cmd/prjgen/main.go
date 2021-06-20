package prjgen

import (
	"github.com/spf13/cobra"
	"github.com/vinshop/prjgen/pkg/logger"
	"os"
)

var Cmd = &cobra.Command{
	Use:   "prjgen",
	Short: "Generate project",
	Run: func(cmd *cobra.Command, args []string) {
		if err := g.exec(); err != nil {
			logger.Errorw("error when exec command", "error", err)
			os.Exit(1)
		}
	},
}

var g = new(Generator)

func init() {
	flag := Cmd.Flags()
	flag.StringSliceVar(&g.Tables, "table", nil, "tables to generate, empty to generate all")
	flag.StringVar(&g.ModelDir, "model", "models", "models directory")
	flag.StringVar(&g.ControllerDir, "controller", "controllers", "controller directory")
	flag.StringVar(&g.RepoDir, "repo", "repositories", "repositories directory")
	flag.StringVar(&g.ServiceDir, "service", "services", "services directory")
	flag.StringVar(&g.MapperDir, "mapper", "mappers", "mappers directory")
	flag.StringVar(&g.DtoDir, "dto", "dtos", "DTOs directory")
	flag.StringVarP(&g.OutputDir, "output", "o", ".", "output directory")

	flag.StringVar(&g.Protocol, "protocol", "tcp", "database protocol")
	flag.StringVar(&g.DB, "db", "db", "database name")
	flag.StringVar(&g.Host, "host", "localhost", "database host")
	flag.IntVar(&g.Port, "port", 3306, "database port")
	flag.StringVar(&g.User, "user", "root", "database user")
	flag.StringVar(&g.Password, "pass", "admin", "database password")

}

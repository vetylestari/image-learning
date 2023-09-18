package console

import (
	"fmt"
	"log"

	"github.com/Renos-id/go-starter-template/database"
	"github.com/spf13/cobra"
)

var initSchema *cobra.Command

func init() {
	initSchema = &cobra.Command{
		Use:   "init-schema",
		Short: "Init Schema PostgreSQL",
		Long:  "Init Schema PostgreSQL",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running Init Schema command")
			_, err := database.CreateSchema()
			if err != nil {
				log.Fatal("init schema failed: ", err)
			}

			fmt.Println("Init Schema with success")
		},
	}
	RootCmd.AddCommand(initSchema)
}

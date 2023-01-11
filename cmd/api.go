package cmd

import (
	"database/sql"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func API() *cobra.Command {
	return &cobra.Command{
		Use:   "api",
		Short: "start get ground party service in api mode",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info("Starting get ground service api")
			runAPI()
		},
	}
}

func runAPI() {
	// init mysql.
	db, err := sql.Open("mysql", "user:password@/getground")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// ping
	http.HandleFunc("/ping", handlerPing)
	http.ListenAndServe(":3000", nil)
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

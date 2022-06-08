package main

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg"
	"matchmove.com/demo/common/config"
	"matchmove.com/demo/migration/model"
)

// list of init queries
var initQueries = []string{
	// Sequences
	model.CreateSequenceRole,
	model.CreateSequenceStatus,
	model.CreateSequenceUser,
	model.CreateSequenceTokens,

	// Tables
	model.CreateRoleTable,
	model.CreateStatusTable,
	model.CreateUsersTable,
	model.CreateTokensTable,
	model.CreateReqCountTable,
}

func main() {
	c := config.GetConfigs()
	psqlInfo := c.Database
	dbClient := pg.Connect(&pg.Options{
		Addr:     fmt.Sprint(psqlInfo.Host, ":", psqlInfo.Port),
		User:     fmt.Sprint(psqlInfo.User),
		Password: fmt.Sprint(psqlInfo.Password),
		Database: fmt.Sprint(psqlInfo.Database),
	})
	defer dbClient.Close()

	for _, query := range initQueries {
		_, err := dbClient.Exec(query, nil)
		if err != nil {
			if !(strings.Contains(err.Error(), "schema") || strings.Contains(err.Error(), "already exists")) {
				panic(err)
			}
		}
	}

}

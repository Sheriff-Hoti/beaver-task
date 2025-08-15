package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/Sheriff-Hoti/beaver-task/config"
	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func main() {
	fmt.Println("init")
	help := flag.Bool("help", false, "show help")
	configPath := flag.String("config", config.GetDefaultConfigPath(), "path to config file")
	//read config file

	flag.Parse()
	cfg, err := config.ReadConfigFile(*configPath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}

	ctx := context.Background()

	db, err := sql.Open("sqlite", cfg.DataDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("config path:", *configPath)
	fmt.Println("help:", *help)
	fmt.Println("config vals:", cfg)
	//do db operations
	//initialize bubbletea app
}

package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"

	"github.com/Sheriff-Hoti/beaver-task/config"
	"github.com/Sheriff-Hoti/beaver-task/database"
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

	if *help {
		flag.Usage()
		return
	}

	cfg, err := config.ReadConfigFile(*configPath)
	if err != nil {
		log.Fatal("error reading config file", err)
		return
	}

	ctx := context.Background()

	db, err := sql.Open("sqlite", cfg.DataDir)
	if err != nil {
		fmt.Println("error opening database", err)
		return
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
		return
	}

	defer db.Close()

	queries := database.New(db)

	res, err := queries.CreateTask(ctx, database.CreateTaskParams{
		Title:       "Sample Task",
		Description: sql.NullString{String: "This is a sample task", Valid: true},
	})

	fmt.Println("config path:", *configPath)
	fmt.Println("help:", *help)
	fmt.Println("config vals:", cfg)
	fmt.Println("res:", res)
	fmt.Println("err:", err)
	//do db operations
	//initialize bubbletea app
}

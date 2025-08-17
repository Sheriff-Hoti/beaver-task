package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Sheriff-Hoti/beaver-task/config"
	"github.com/Sheriff-Hoti/beaver-task/database"
	tui "github.com/Sheriff-Hoti/beaver-task/tui"
	tea "github.com/charmbracelet/bubbletea"
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
		log.Fatal("error opening database", err)
		return
	}

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal("error creating tables:", err)
		return
	}

	defer db.Close()

	queries := database.New(db)

	initialTasks, err := queries.ListTasks(ctx, 1)
	tui.InitialModel(initialTasks)
	if err != nil {
		log.Fatal("error listing tasks:", err)
		return
	}
	//initialte the background and foreground here
	manager := &tui.Manager{}

	p := tea.NewProgram(
		// tui.InitialModel(initialTasks)
		manager, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	fmt.Println("config path:", *configPath)
	fmt.Println("help:", *help)
	fmt.Println("config vals:", cfg)
	// fmt.Println("res:", res)
	// fmt.Println("err:", err)
	//do db operations
	//initialize bubbletea app
}

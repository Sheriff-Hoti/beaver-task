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
	if err != nil {
		log.Fatal("error listing tasks:", err)
		return
	}
	//initialte the background and foreground here
	// background := tui.NewBackground()
	manager := tui.NewManager(queries, tui.FromDatabaseTasks(initialTasks))

	p := tea.NewProgram(manager, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

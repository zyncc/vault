package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"github.com/pressly/goose/v3"
	"github.com/zyncc/vault/cmd"
	"github.com/zyncc/vault/db"
	"github.com/zyncc/vault/password"
	_ "modernc.org/sqlite"
)

func main() {
	if err := bootstrapVault(); err != nil {
		fmt.Println("Failed to initialize Vault, ", err.Error())
		return
	}
	cmd.Execute()
}

func bootstrapVault() error {
	dbPath, err := initStorage()
	if err != nil {
		return err
	}

	if !isFirstRun(dbPath) {
		return nil
	}

	fmt.Println("🔐 First time setuppppp")
	fmt.Println("Set a master password to protect your vault")

	master, err := promptMasterPassword()
	if err != nil {
		return err
	}

	if err := runMigrations(dbPath); err != nil {
		return err
	}

	hash, salt, err := password.HashPassword(master)
	if err != nil {
		log.Fatal("failed to hash master password")
	}

	q := db.Init()
	if err := q.CreateMasterPassword(context.Background(), db.CreateMasterPasswordParams{
		Password: hash,
		Salt:     salt,
	}); err != nil {
		log.Fatal("Failed to create Master Password")
	}

	fmt.Println("✅ Vault initialized successfully")
	return nil
}

func initStorage() (string, error) {
	// Create Vault Folder in User's Config Folder
	base, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	dir := filepath.Join(base, "vault")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return "", err
	}

	dbPath := filepath.Join(dir, "vault.db")
	return dbPath, nil
}

func isFirstRun(dbPath string) bool {
	_, err := os.Stat(dbPath)
	return os.IsNotExist(err)
}

//go:embed migrations/*.sql
var MigrationsFolder embed.FS

func runMigrations(dbPath string) error {
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return err
	}
	defer db.Close()

	goose.SetDialect("sqlite3")
	goose.SetBaseFS(MigrationsFolder)
	return goose.Up(db, "migrations")
}

func promptMasterPassword() (string, error) {
	p1 := promptui.Prompt{
		Label: "Master Password",
		Mask:  '*',
	}

	pass1, err := p1.Run()
	if err != nil {
		return "", err
	}

	p2 := promptui.Prompt{
		Label: "Confirm Master Password",
		Mask:  '*',
	}

	pass2, err := p2.Run()
	if err != nil {
		return "", err
	}

	if pass1 != pass2 {
		return "", errors.New("passwords do not match")
	}

	if len(pass1) < 8 {
		return "", errors.New("password too short")
	}

	return pass1, nil
}

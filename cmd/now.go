package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
    "os/exec"

	"github.com/spf13/cobra"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Today's journal",
	Long: "Either opens today's entry, or creates a new one to edit",
	Run: func(cmd *cobra.Command, args []string) {
	    getNow()
	},
}

func getNow() {
    now := time.Now()
    date := now.Format(time.DateOnly)

    home, err := os.UserHomeDir()
    if err != nil {
        log.Fatal(err)
    }
    journalDir := home + "/Dropbox/journal"
    fileName := date + ".md"

    if _, err := os.Stat(journalDir); errors.Is(err, os.ErrNotExist) {
        fmt.Println("path does not exist")
    }
    if _, err := os.Stat(journalDir + "/" + fileName); errors.Is(err, os.ErrNotExist) {
        fmt.Println("here")
        createEntry(journalDir, fileName)
    }

    editEntry(journalDir, fileName)
}

func editEntry(dir string, name string) {
    path := dir + "/" + name 
    cmd := exec.Command("nvim", path)
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    err := cmd.Run()
    fmt.Println(err)
}

func createEntry(dir string, name string) {
    path := dir + "/" + name

    f, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()
}

func init() {
	rootCmd.AddCommand(nowCmd)
}

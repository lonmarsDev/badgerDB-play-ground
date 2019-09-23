package main

import (
	"bufio"
	"log"
	"os"

	badger "github.com/dgraph-io/badger"
	y "github.com/dgraph-io/badger/y"
)

func main() {
	// Open DB
	var backupFile string
	db, err := badger.Open(badger.DefaultOptions("badger31"). 
		WithValueDir("badger31").
		WithTruncate(false) )
	
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	backupFile = "backup.bak"
	// Create File
	f, err := os.Create(backupFile)
	if err != nil {
		log.Fatal(err)
	}

	bw := bufio.NewWriterSize(f, 64<<20)
	if _, err = db.Backup(bw, 0); err != nil {
		log.Fatal(err)
	}

	if err = bw.Flush(); err != nil {
		log.Fatal(err)
	}

	if err = y.FileSync(f); err != nil {
		log.Fatal(err)
	}

	f.Close()

}

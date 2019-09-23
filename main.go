package main

import (
	//	"bufio"
	"log"
	//	"os"

	badger "github.com/dgraph-io/badger"
	//	y "github.com/dgraph-io/badger/y"
)

func main() {
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	db, err := badger.Open(badger.DefaultOptions("badger31"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Start a writable transaction.
	txn := db.NewTransaction(true)
	defer txn.Discard()

	// Use the transaction...
	err = txn.Set([]byte("answer"), []byte("42"))
	if err != nil {
		panic(err)
	}

	// Commit the transaction and check for error.
	if err := txn.Commit(); err != nil {
		panic(err)
	}

	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("43"))
		return err
	})
	// Your code here…

	//   err := db.View(func(txn *badger.Txn) error {
	// 	// Your code here…
	// 	return nil
	//   })

}

// func doBackup() error {
// 	// Open DB
// 	var backupFile string
// 	db, err := badger.Open(badger.DefaultOptions("badger31"))
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()
// 	backupFile = "backup.bak"
// 	// Create File
// 	f, err := os.Create(backupFile)
// 	if err != nil {
// 		return err
// 	}

// 	bw := bufio.NewWriterSize(f, 64<<20)
// 	if _, err = db.Backup(bw, 0); err != nil {
// 		return err
// 	}

// 	if err = bw.Flush(); err != nil {
// 		return err
// 	}

// 	if err = y.FileSync(f); err != nil {
// 		return err
// 	}

// 	return f.Close()
// }

package main

import (
	"fmt"
	"os"

	badger "github.com/dgraph-io/badger"
)

func main() {

	bak, err := os.Open("backup.bak")
	if err != nil {
		panic(err)
	}
	defer bak.Close()

	opts := badger.DefaultOptions("./data-restored")
	opts.Dir = "./data-restored"
	opts.ValueDir = "./data-restored"
	//opts.Logger = nil
	//opts.Truncate = true
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Load(bak, 16)
	if err != nil {
		panic(err)
	}

	err = db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte("marlon1"))
		if err != nil {
			panic(err)
		}

		var valNot, valCopy []byte
		err = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.

			// Accessing val here is valid.
			fmt.Printf("The answer is: %s\n", val)

			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			valNot = val // Do not do this.
			return nil
		})
		if err != nil {
			panic(err)
		}

		// DO NOT access val here. It is the most common cause of bugs.
		fmt.Printf("NEVER do this. %s\n", valNot)

		// You must copy it to use it outside item.Value(...).
		fmt.Printf("The answer is: %s\n", valCopy)

		// Alternatively, you could also use item.ValueCopy().
		valCopy, err = item.ValueCopy(nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("The answer is: %s\n", valCopy)

		return nil

	})

	if err != nil {
		panic(err)
	}

}

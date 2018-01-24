package main

import (
	"context"
	"flag"
	"log"

	"github.com/maksharma/bigtable/lib"
	"github.com/maksharma/bigtable/model"
)

func main() {
	project := flag.String("project", "", "The Google Cloud Platform project ID. Required.")
	instance := flag.String("instance", "", "The Google Cloud Bigtable instance ID. Required.")
	flag.Parse()

	for _, f := range []string{"project", "instance"} {
		if flag.Lookup(f).Value.String() == "" {
			log.Fatalf("The %s flag is required.", f)
		}
	}

	ctx := context.Background()

	model.Init(ctx, *project, *instance)

	err := model.CreateIfNotExists(ctx)
	if err != nil {
		log.Fatalf("Could not create table %s: %v", lib.TABLE_NAME, err)
	}

	err = model.CreateColumnFamily(ctx)
	if err != nil {
		log.Fatalf("Could not create column family %s: %v", lib.COLUMN_FAMILY_NAME, err)
	}

	err = model.InsertAndDisplay(ctx, *project, *instance)
	if err != nil {
		log.Fatalf("Could not insert and display data %s: %v", lib.TABLE_NAME, err)
	}

	log.Printf("Deleting a row")
	if err = model.DeleteRow(ctx); err != nil {
		log.Fatalf("Could not delete table %s: %v", lib.TABLE_NAME, err)
	}

	log.Printf("Deleting the table")
	if err = model.DeleteTable(ctx); err != nil {
		log.Fatalf("Could not delete table %s: %v", lib.TABLE_NAME, err)
	}

	if err = model.CloseConnections(ctx); err != nil {
		log.Fatalf("Could not close connections table %s: %v", lib.TABLE_NAME, err)
	}

}

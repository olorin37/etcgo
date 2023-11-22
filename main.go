package main

import (
	"fmt"
	"log"

	"github.com/itchyny/gojq"
	"github.com/go-git/go-git/v5"
)

func main() {
	query, err := gojq.Parse(".foo | ..")
	if err != nil {
		log.Fatalln(err)
	}
	input := map[string]any{"foo": []any{1, 2, 3}}
	iter := query.Run(input) // or query.RunWithContext
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", v)
	}

	r, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL: "https://github.com/git-fixtures/basic.git",
        })
        if err != nil {
        	log.Fatal(err)
        }
	fmt.Printf("%s\n", r)
}

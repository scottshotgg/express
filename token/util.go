package token

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintTokens(tokens []Token, jsonIndent string) {
	for _, t := range tokens {
		if t.Type == Block || t.Type == Array || t.Type == Group || t.Type == Function || t.Type == Attribute {
			jsonIndent += "\t"

			tks, ok := t.Value.True.([]Token)
			if !ok {
				fmt.Println("could not assert interface to []token")
				os.Exit(9)
			}

			fmt.Println()
			fmt.Println(jsonIndent[0:len(jsonIndent)-1] + t.Type)
			PrintTokens(tks, jsonIndent)

			jsonIndent = jsonIndent[0 : len(jsonIndent)-1]
			continue
		}

		tokenJSON, err := json.Marshal(t)
		if err != nil {
			fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", t)
			os.Exit(9)
		}
		fmt.Println(jsonIndent + string(tokenJSON))
	}
}

func PrintValues(tokens []Value, jsonIndent string) {
	for _, t := range tokens {
		if t.Type == Block || t.Type == Array || t.Type == Group || t.Type == Function || t.Type == Attribute {
			jsonIndent += "\t"

			tks, ok := t.True.([]Value)
			if !ok {
				fmt.Println("could not assert value")
				os.Exit(9)
			}

			fmt.Println()
			fmt.Println(jsonIndent[0:len(jsonIndent)-1] + t.Type)
			PrintValues(tks, jsonIndent)

			jsonIndent = jsonIndent[0 : len(jsonIndent)-1]
			continue
		}

		tokenJSON, err := json.Marshal(t)
		if err != nil {
			fmt.Printf("\nERROR: Could not marshal JSON from token: %#v\n", t)
			os.Exit(9)
		}
		fmt.Println(jsonIndent + string(tokenJSON))
	}
}

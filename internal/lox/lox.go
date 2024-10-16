package lox

import (
	"fmt"

	"github.com/zweix123/golox/internal/scanner"
)

func Run(source string) error {
	scanner := scanner.NewScanner(source)
	tokens, err := scanner.ScanTokens()
	if err != nil {
		return err
	}
	for _, token := range tokens {
		fmt.Printf("%s\n", token.String())
	}
	return nil
}

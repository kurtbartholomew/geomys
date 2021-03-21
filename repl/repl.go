package repl

import (
	"bufio"
	"fmt"
	"io"
	"geomys/lexer"
	"geomys/token"
)

const PROMPT = ">>"

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexPtr := lexer.New(line)

		for currentToken := lexPtr.NextToken(); currentToken.Type != token.EOF; currentToken = lexPtr.NextToken() {
			fmt.Printf("%+v\n", currentToken)
		}
	}
}
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/tophatsteve/monkey/token"

	"github.com/tophatsteve/monkey/lexer"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		if line == "quit" {
			fmt.Println("Exiting")
			return
		}

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

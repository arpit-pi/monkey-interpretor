package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/arpit-pi/monkey-interpretor/token"
	"github.com/arpit-pi/monkey-interpretor/lexer"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// to print struct fields
			fmt.Printf("%+v\n", tok)
		}
	}
}
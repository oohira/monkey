package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/oohira/monkey/lexer"
	"github.com/oohira/monkey/token"
)

// PROMPT is characters to prompt users input.
const PROMPT = ">> "

// Start starts REPL: reads user's input from in and writes the result to out.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

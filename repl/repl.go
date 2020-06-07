package repl

import (
	"bufio"
	"fmt"
	"github.com/xxnmxx/gotax/evaluator"
	"github.com/xxnmxx/gotax/lexer"
	"github.com/xxnmxx/gotax/parser"
	"github.com/xxnmxx/gotax/object"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnviroment()

	for {
		fmt.Print(PROMPT)
		sccaned := scanner.Scan()
		if !sccaned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

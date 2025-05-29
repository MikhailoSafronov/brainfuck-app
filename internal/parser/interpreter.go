package parser

import (
	"fmt"

	"github.com/MikhailoSafronov/Project-lab4/internal/ast"
	"github.com/MikhailoSafronov/Project-lab4/internal/lexer"
)

// Interpret виконує обробку AST
func Interpret(nodes []ast.Node) {
	const tapeSize = 30000
	tape := make([]byte, tapeSize)
	ptr := 0

	var exec func(nodes []ast.Node)
	exec = func(nodes []ast.Node) {
		for _, n := range nodes {
			switch node := n.(type) {
			case *ast.Command:
				switch node.Kind {
				case lexer.GT:
					ptr++
					if ptr >= tapeSize {
						ptr = 0 // wrap around
					}
				case lexer.LT:
					if ptr == 0 {
						ptr = tapeSize - 1
					} else {
						ptr--
					}
				case lexer.PLUS:
					tape[ptr]++
				case lexer.MINUS:
					tape[ptr]--
				case lexer.DOT:
					fmt.Printf("%c", tape[ptr])
				case lexer.COMMA:
					var input byte
					fmt.Scanf("%c", &input)
					tape[ptr] = input
				}
			case *ast.Loop:
				for tape[ptr] != 0 {
					exec(node.Body)
				}
			}
		}
	}

	exec(nodes)
}

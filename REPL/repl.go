package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime/pprof"
	"time"

	"github.com/AmjedChakhis/GoLang-interpreter/core/debug"
	"github.com/AmjedChakhis/GoLang-interpreter/core/lexer"
	"github.com/AmjedChakhis/GoLang-interpreter/core/parser/parserImpl"
	"github.com/AmjedChakhis/GoLang-interpreter/core/runtime"
	"github.com/AmjedChakhis/GoLang-interpreter/core/types"
)

// REPL
/*
* Function to start the repl
 */

const PROMPT = ">_"

var ctx = types.NewContext()

const (
	enableCpuProfiling = true
	enableMemProfiling = true
	isDebugging        = false
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	fmt.Println(`   ____ ____  _____ 
  / ___| __ )| |
 | |  _|  _ \| |  
 | |_| | |_) | | 
  \____|____/|_|

  GBI - Go Based Interpreter
`)
	fmt.Println("------------- Welcome to GBI cmd : Tap your commands now... ------------")
	fmt.Println("                                 👋                              ")

	if enableCpuProfiling {

		f, err := os.Create("cpu.pprof")
		if err != nil {
			panic(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

	}
	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		start := time.Now() // start time
		line := scanner.Text()

		if line == "exit" || line == "exit;" {
			break
		}

		replLexer := lexer.InitLexer(line)
		repParser := parserImpl.InitParser(replLexer)

		pr := repParser.Parse()
		errs := repParser.Errors()

		if len(errs) != 0 {
			io.WriteString(out, fmt.Sprintf("%d errors ❌ occurred while parsing your input \n", len(errs)))
			for idx, e := range errs {
				io.WriteString(out, fmt.Sprintf("error number:%d with message: %s \n", idx, e.Message))
			}
			continue
		}

		afterParsing := time.Since(start)

		if isDebugging {
			fmt.Printf("parsing step for %s took %s \n", line, afterParsing)
		}

		evaluated, err := runtime.Eval(pr, ctx)
		if err != debug.NOERROR {
			io.WriteString(out, fmt.Sprintf("error while evaluating your input: %s \n", err.Error()))
			continue
		}

		if isDebugging {
			fmt.Printf("expression evaluations  step for %s took %s \n", line, afterParsing)
		}

		if evaluated != nil {
			io.WriteString(out, evaluated.ToString())
			io.WriteString(out, "\n")
		}

		// take memory snapshot
		if enableMemProfiling {
			f, err := os.Create("mem.pprod")

			if err != nil {
				panic(err)
			}
			pprof.WriteHeapProfile(f)
			f.Close()
		}
	}
}

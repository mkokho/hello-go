package main

import (
	"fmt"
	"os"
	"strings"
	"time"
	"strconv"
)

var copmletionScript = `# Copmletion function for bash shell
_godate()
{
    local cur prev opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    opts="` + strings.Join(timezones, " ") + `"

		COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
}
complete -F _godate godate
`

func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Print("Examples: \n")
		fmt.Fprint(os.Stdout, strings.Join(examples(), "\n"))
		os.Exit(1)
	}

	if len(os.Args) == 2 && os.Args[1] == "completion" {
		fmt.Print(copmletionScript)
		os.Exit(0)
	}

	res, err := run(os.Args[1:])
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, res)
	os.Exit(0)
}

func run(args []string) (string, error) {
	if len(args) == 0 {
		return time.Now().String(), nil
	} else if len(args) == 1 {
		nanos, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			loc, err := time.LoadLocation(args[0])
			if err != nil {
				return "", err
			}
			return time.Now().In(loc).String(), nil
		}
		return time.Unix(0, nanos).String(), nil
	} else if len(args) == 2 {
		nanos, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			nanos, err = strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return "could not parse nanoseconds from any of the arguments", nil
			}
		}

		loc, err := time.LoadLocation(args[0])
		if err != nil {
			loc, err = time.LoadLocation(args[1])
			if err != nil {
				return "could not parse location from any of the arguments", nil
			}
		}

		return time.Unix(0, nanos).In(loc).String(), nil
	}
	return "not yet implemetnted", nil
}



func examples() []string{
	supportedArgs := [][]string{
		[]string{},
		[]string{"US/Central"},
		[]string{"UTC"},
		[]string{"1481687107832160800"},
		[]string{"1481687107832160800", "UTC"},
		[]string{"UTC", "1481687107832160800"},
	}

	res := []string{}
	for _, args := range supportedArgs {
		x, err := run(args)
		if err != nil {
			x = err.Error()
		}
		res = append(res, fmt.Sprintf("> godate %s\n%s", strings.Join(args, " "), x))
	}
	return res
}

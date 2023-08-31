package flags

import (
	calculator "calculator/calc"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
)

type commandLine struct {
	*flag.FlagSet
	formal map[string]*CustomFlag
}

type CustomFlag struct {
	flag.Flag
	Example   string
	GroupName string
}

func newCmdLine() *commandLine {
	var cmd *commandLine
	flagset := flag.NewFlagSet("Shell Calculator", flag.ExitOnError)
	flagset.SetOutput(os.Stdout)
	cmd = &commandLine{FlagSet: flagset}
	flagset.Usage = cmd.PrintDefault
	return cmd
}

func sortFlags(flags map[string]*CustomFlag) []*CustomFlag {

	result := make([]*CustomFlag, len(flags))
	i := 0
	for _, f := range flags {
		result[i] = f
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}

func (c *commandLine) VisitAll(fn func(f *CustomFlag)) {
	for _, flag := range sortFlags(c.formal) {
		fn(flag)
	}
}

func (f *commandLine) Var(value flag.Value, name, usage string, example ...string) {
	var customflag *CustomFlag

	if strings.HasPrefix(name, "-") {
		panic(fmt.Sprintf("flag %q begins with -", name))
	} else if strings.Contains(name, "=") {
		panic(fmt.Sprintf("flag %q contains =", name))
	}

	if strings.Contains(name, ",") {
		names := splitIntoMultiname(name)

		for _, n := range names {
			customflag = &CustomFlag{Flag: flag.Flag{Name: n, Usage: usage, Value: value, DefValue: value.String()}}
			customflag.GroupName = names[0]

			if len(example) != 0 {
				customflag.Example = example[0]
			}

			if f.formal == nil {
				f.formal = make(map[string]*CustomFlag)
			}
			f.formal[n] = customflag
			f.FlagSet.Var(customflag.Value, customflag.Name, customflag.Usage)
		}

	} else {
		customflag = &CustomFlag{Flag: flag.Flag{Name: name, Usage: usage, Value: value, DefValue: value.String()}}

		if len(example) != 0 {
			customflag.Example = example[0]
		}

		if f.formal == nil {
			f.formal = make(map[string]*CustomFlag)
		}
		f.formal[name] = customflag
		f.FlagSet.Var(customflag.Value, customflag.Name, customflag.Usage)
	}
}

func (c *commandLine) Usage() {
	fmt.Fprintf(c.Output(), "\n%s Usage:\n\n", c.Name())
}

func (c *commandLine) PrintDefault() {
	c.Usage()
	hasSubnames := false

	var names = make([]string, 0)
	var groupName = make(map[string]string)
	if len(c.formal) > 0 {
		for fln, flag := range c.formal {
			if flag.GroupName != "" && groupName[fln] != flag.GroupName {
				groupName[fln] = flag.GroupName
				hasSubnames = true
			}

			if groupName[fln] == flag.GroupName {
				names = append(names, fln)
			}
		}
	}

	c.VisitAll(func(f *CustomFlag) {
		var b strings.Builder

		var fname string
		if hasSubnames {
			for i := range names {
				names[i] = fmt.Sprintf("-%s", names[i])
			}

			fname = strings.Join(names, ", ")
			fmt.Fprintf(&b, "   %s:", fname)
		} else {
			fname = f.Name
			fmt.Fprintf(&b, "   -%s:", fname)
		}

		name, usage := flag.UnquoteUsage(&f.Flag)
		if len(name) > 0 {
			b.WriteString(" ")
			b.WriteString(name)
		}

		if f.DefValue == "" {
			fmt.Fprintf(&b, " (default %q)\n", f)
		} else {
			fmt.Fprintf(&b, " (default %v)\n", f.DefValue)
		}

		b.WriteString("\n    \t")
		b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

		if f.Example != "" {
			b.WriteString("\n    \tExample:\n")
			b.WriteString(strings.ReplaceAll(f.Example, "\n", "\n   \t"))
		}

		fmt.Fprint(c.Output(), b.String(), "\n")
		if hasSubnames {
			os.Exit(0)
		}
	})
}

func splitIntoMultiname(fname string) []string {
	fname = strings.ReplaceAll(fname, " ", "")
	return strings.Split(fname, ",")
}

func (f *PrecisionFlag) Set(p string) error {
	var verb string

	fmt.Sscanf(p, "%s", &verb)

	if !strings.Contains(verb, "%") {
		verb = strings.Replace(verb, verb, "%."+verb+"f", 1)
	}
	f.CalculatorPrecision = calculator.CalculatorPrecision(verb)
	return nil
}

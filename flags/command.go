package flags

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"golang.org/x/exp/slices"
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
	flagset := flag.NewFlagSet("Go Calculator", flag.ExitOnError)
	flagset.SetOutput(os.Stderr)
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

	var groupName = make(map[string]string)
	if len(c.formal) > 0 {
		for fln, flag := range c.formal {
			if flag.GroupName != "" && groupName[fln] != flag.GroupName {
				groupName[fln] = flag.GroupName
			}
		}
	}

	i := 0
	seen := make([]string, 0)
	c.VisitAll(func(f *CustomFlag) {
		if hasSubnames {
			// os.Exit(0)
			if i == len(groupName)-1 {
				return
			}
		}

		var b strings.Builder
		var fname string
		if groupName[f.Name] == f.GroupName {
			names := make([]string, 0)
			for flagName, gn := range groupName {
				if gn == f.GroupName {
					names = append(names, fmt.Sprintf("-%s", flagName))
				}
			}

			fname = strings.Join(names, ", ")
			fmt.Fprintf(&b, "   \033[01;05m%s:\033[00m", fname)

			hasSubnames = true
		} else {
			fname = f.Name
			fmt.Fprintf(&b, "   \033[01;05m-%s:\033[00m", fname)
			hasSubnames = false
		}

		if !slices.Contains[[]string](seen, f.GroupName) {

			name, usage := flag.UnquoteUsage(&f.Flag)
			if len(name) > 0 {
				b.WriteString(" \033[01;05;33m")
				b.WriteString(" ")
				b.WriteString(name)
				b.WriteString(" \033[00m")
			}

			b.WriteString(" \033[01;04;35m")
			if f.DefValue == "" {
				fmt.Fprintf(&b, "(default %q)\n", f)
			} else {
				fmt.Fprintf(&b, "(default %v)\n", f.DefValue)
			}
			b.WriteString("\033[00m")

			b.WriteString("\n    \t\033[01;32m")
			b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))
			b.WriteString("\033[00m")

			if f.Example != "" {
				b.WriteString("\033[05;36m")
				b.WriteString("\n    \tExample:\n")
				b.WriteString(strings.ReplaceAll(f.Example, "\n", "\n   \t"))
				b.WriteString("\033[00m\n")
			}

			fmt.Fprintf(c.Output(), "%s\n\033[00m", b.String())
		}
		i++
		seen = append(seen, f.GroupName)
	})
}

func splitIntoMultiname(fname string) []string {
	fname = strings.ReplaceAll(fname, " ", "")
	return strings.Split(fname, ",")
}

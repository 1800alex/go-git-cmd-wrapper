package types

import (
	"context"
	"log"
	"os"
	"os/exec"
	"strings"
)

type logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
}

// Executor The Git command call function.
type Executor func(ctx context.Context, dir string, name string, debug bool, args ...string) (string, error)

// NewCmd Creates a new Cmd.
func NewCmd(name string) *Cmd {
	g := &Cmd{
		Debug:   false,
		Base:    "git",
		Dir:     "",
		Options: []string{name},
		Logger:  log.New(os.Stdout, "", 0),
	}
	g.Executor = defaultExecutor(g)

	return g
}

// Cmd Command.
type Cmd struct {
	Debug    bool
	Base     string
	Dir      string
	Options  []string
	Logger   logger
	Executor Executor
}

// Option Command option.
type Option func(g *Cmd)

// AddOptions Add one command option.
func (g *Cmd) AddOptions(option string) {
	if strings.HasPrefix(option, "--working-dir=") {
		g.Dir = strings.TrimPrefix(option, "--working-dir=")
	}

	g.Options = append(g.Options, option)
}

// ApplyOptions Apply command options.
func (g *Cmd) ApplyOptions(options ...Option) {
	for _, opt := range options {
		opt(g)
	}
}

// Exec Execute the Git command call.
func (g *Cmd) Exec(ctx context.Context, dir string, name string, debug bool, args ...string) (string, error) {
	return g.Executor(ctx, dir, name, debug, args...)
}

func defaultExecutor(g *Cmd) Executor {
	return func(ctx context.Context, dir string, name string, debug bool, args ...string) (string, error) {
		if debug {
			g.Logger.Println(name, strings.Join(args, " "))
		}

		cmd := exec.CommandContext(ctx, name, args...)
		if dir != "" {
			cmd.Dir = dir
		}

		output, err := cmd.CombinedOutput()

		return string(output), err
	}
}

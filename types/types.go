package types

import (
	"context"
	"io"
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

type StreamHandler struct {
	Stdout io.Writer
	Stderr io.Writer
	Stdin  io.Reader
}

// Executor The Git command call function.
type Executor func(ctx context.Context, name string, debug bool, streamHandler StreamHandler, args ...string) (string, error)

// NewCmd Creates a new Cmd.
func NewCmd(name string) *Cmd {
	g := &Cmd{
		Debug:   false,
		Base:    "git",
		Options: []string{name},
		Logger:  log.New(os.Stdout, "", 0),
	}
	g.Executor = defaultExecutor(g)

	return g
}

// Cmd Command.
type Cmd struct {
	Debug         bool
	Base          string
	Options       []string
	StreamHandler StreamHandler
	Logger        logger
	Executor      Executor
}

// Option Command option.
type Option func(g *Cmd)

// AddOptions Add one command option.
func (g *Cmd) AddOptions(option string) {
	g.Options = append(g.Options, option)
}

// ApplyOptions Apply command options.
func (g *Cmd) ApplyOptions(options ...Option) {
	for _, opt := range options {
		opt(g)
	}
}

// Exec Execute the Git command call.
func (g *Cmd) Exec(ctx context.Context, name string, debug bool, streamHandler StreamHandler, args ...string) (string, error) {
	return g.Executor(ctx, name, debug, streamHandler, args...)
}

func defaultExecutor(g *Cmd) Executor {
	return func(ctx context.Context, name string, debug bool, streamHandler StreamHandler, args ...string) (string, error) {
		if debug {
			g.Logger.Println(name, strings.Join(args, " "))
		}
		//output, err := exec.CommandContext(ctx, name, args...).CombinedOutput()

		var combineOutputBuilder strings.Builder
		cmd := exec.CommandContext(ctx, name, args...)
		if streamHandler.Stdin != nil {
			cmd.Stdin = streamHandler.Stdin
		}

		cmd.Stdout = &combineOutputBuilder
		if streamHandler.Stdout != nil {
			cmd.Stdout = io.MultiWriter(&combineOutputBuilder, streamHandler.Stdout)
		}

		cmd.Stderr = &combineOutputBuilder
		if streamHandler.Stderr != nil {
			cmd.Stderr = io.MultiWriter(&combineOutputBuilder, streamHandler.Stderr)
		}

		err := cmd.Run()

		return combineOutputBuilder.String(), err
	}
}

// NewFanOutPipe is a helper function to create a single writer with multiple readers.
func NewFanOutPipe(readerCount int) (io.Writer, []io.Reader) {
	sourceReader, sourceWriter := io.Pipe()
	destinationReaders := make([]io.Reader, 0)
	destinationWriters := make([]*io.PipeWriter, 0)
	for i := 0; i < readerCount; i++ {
		destinationReader, destinationWriter := io.Pipe()
		destinationReaders = append(destinationReaders, destinationReader)
		destinationWriters = append(destinationWriters, destinationWriter)
	}

	go func() {
		defer func() {
			_ = sourceWriter.Close()
			for _, destinationWriter := range destinationWriters {
				_ = destinationWriter.Close()
			}
		}()
		// Copy the sourceReader to all the destinationWriters
		// type assertion
		_writers := make([]io.Writer, len(destinationWriters))
		for i, w := range destinationWriters {
			_writers[i] = w
		}
		_, _ = io.Copy(io.MultiWriter(_writers...), sourceReader)
	}()

	return sourceWriter, destinationReaders
}

package main

import (
	"fmt"
	"os"
	"strings"

	mowcli "github.com/jawher/mow.cli"

	"github.com/novln/soba"
)

func main() {
	app := mowcli.App("soba", "CLI for pretty-printing and filtering soba logs")
	app.Version("v version", "soba v1.0.0")

	colorOpt := getCLIColorOption(app)
	includeOpts := getCLIIncludeOption(app)
	excludeOpts := getCLIExcludeOption(app)

	app.Action = func() {
		color := NewPrettyHandler(parseCLIColorOption(*colorOpt))
		filter := NewFilterHandler(parseCLIFilterOptions(*includeOpts, *excludeOpts))

		pipeline := NewPipeline(os.Stdin, os.Stdout, filter, color)
		err := pipeline.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "soba: cannot read standard input:", err)
			mowcli.Exit(1)
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, "soba: cannot parse command line flags:", err)
		os.Exit(2)
	}
}

func getCLIColorOption(app *mowcli.Cli) *string {
	return app.StringOpt("c color", "always", fmt.Sprint(
		`Colorize the output: "always" or "never"`,
	))
}

func getCLIIncludeOption(app *mowcli.Cli) *[]string {
	return app.StringsOpt("i include", []string{}, fmt.Sprint(
		`Add include filters:`, "\n",
		`- "`, soba.LoggerKey, `:repositories.users"`, "\n",
		`- "`, soba.LevelKey, `:debug"`, "\n",
		`- "`, soba.MessageKey, `:Foobar and stuff"`, "\n",
	))
}

func getCLIExcludeOption(app *mowcli.Cli) *[]string {
	return app.StringsOpt("e exclude", []string{}, fmt.Sprint(
		`Add exclude filters:`, "\n",
		`- "`, soba.LoggerKey, `:repositories.users"`, "\n",
		`- "`, soba.LevelKey, `:debug"`, "\n",
		`- "`, soba.MessageKey, `:Foobar and stuff"`, "\n",
	))
}

func parseCLIColorOption(opt string) *ColorOptions {
	switch opt {
	case "never", "off", "no", "disable":
		return &ColorOptions{
			EnableColor: false,
		}

	default:
		return &ColorOptions{
			EnableColor: true,
		}
	}
}

func parseCLIFilterOptions(include []string, exclude []string) *FilterOptions {
	opts := &FilterOptions{}

	loggerKey := fmt.Sprintf("%s:", soba.LoggerKey)
	levelKey := fmt.Sprintf("%s:", soba.LevelKey)
	messageKey := fmt.Sprintf("%s:", soba.MessageKey)

	for _, val := range include {
		switch {
		case strings.HasPrefix(val, loggerKey):
			opts.Loggers.Include = append(opts.Loggers.Include, strings.TrimPrefix(val, loggerKey))

		case strings.HasPrefix(val, levelKey):
			opts.Levels.Include = append(opts.Levels.Include, strings.TrimPrefix(val, levelKey))

		case strings.HasPrefix(val, messageKey):
			opts.Messages.Include = append(opts.Messages.Include, strings.TrimPrefix(val, messageKey))
		}
	}

	for _, val := range exclude {
		switch {
		case strings.HasPrefix(val, loggerKey):
			opts.Loggers.Exclude = append(opts.Loggers.Exclude, strings.TrimPrefix(val, loggerKey))

		case strings.HasPrefix(val, levelKey):
			opts.Levels.Exclude = append(opts.Levels.Exclude, strings.TrimPrefix(val, levelKey))

		case strings.HasPrefix(val, messageKey):
			opts.Messages.Exclude = append(opts.Messages.Exclude, strings.TrimPrefix(val, messageKey))
		}
	}

	return opts
}

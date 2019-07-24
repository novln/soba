package main

import (
	"bufio"
	"fmt"
	"io"
)

// Pipeline is components that read from stdin (line by line) and dispatch the workload on
// three goroutines, streaming the lines: "filter", "color" and "print".
//
// Filter is the first step/task: it preserve valid json on the stream that satisfies a list of custom rules.
// For example, these custom rules could be to discard a log level, eliminate a log message by its content
// or retain specific loggers.
//
// Color is the second step/task: it colorizes and formats the input json into a more human readable format
// where each element is on it's own line with clear indentation.
//
// Print is the final step/task: it receives the stream of lines and write them on stdout.
type Pipeline struct {
	input         io.Reader
	writer        io.Writer
	filterHandler *FilterHandler
	colorHandler  *PrettyHandler
	filterChan    chan []byte
	filterDone    chan struct{}
	colorChan     chan []byte
	colorDone     chan struct{}
	printChan     chan []byte
	printDone     chan struct{}
}

// NewPipeline creates a new Pipeline instance.
func NewPipeline(input io.Reader, writer io.Writer, filter *FilterHandler, color *PrettyHandler) *Pipeline {
	return &Pipeline{
		input:         input,
		writer:        writer,
		filterHandler: filter,
		colorHandler:  color,
	}
}

// Run executes the pipeline.
func (pipeline *Pipeline) Run() error {

	pipeline.startPrint()
	pipeline.startColor()
	pipeline.startFilter()

	scanner := bufio.NewScanner(pipeline.input)
	for scanner.Scan() {
		pipeline.filterChan <- scanner.Bytes()
	}

	err := scanner.Err()

	pipeline.stopFilter()
	pipeline.stopColor()
	pipeline.stopPrint()

	return err
}

// startPrint launches "filter" goroutine.
func (pipeline *Pipeline) startFilter() {
	pipeline.filterChan = make(chan []byte, 1024)
	pipeline.filterDone = make(chan struct{})
	go pipeline.doFilter()
}

// stopPrint shutdowns "filter" goroutine.
func (pipeline *Pipeline) stopFilter() {
	close(pipeline.filterChan)
	<-pipeline.filterDone
}

// doFilter receives a line from the main process, filters the input using custom rules,
// and send the new line to the "color" goroutine.
func (pipeline *Pipeline) doFilter() {
	handler := pipeline.filterHandler
	for json := range pipeline.filterChan {
		output := handler.Filter(json)
		if len(output) > 0 {
			pipeline.colorChan <- output
		}
	}
	close(pipeline.filterDone)
}

// startPrint launches "color" goroutine.
func (pipeline *Pipeline) startColor() {
	pipeline.colorChan = make(chan []byte, 1024)
	pipeline.colorDone = make(chan struct{})
	go pipeline.doColor()
}

// stopPrint shutdowns "color" goroutine.
func (pipeline *Pipeline) stopColor() {
	close(pipeline.colorChan)
	<-pipeline.colorDone
}

// doColor receives a line from the "filter" goroutine, converts the input into a more human readable format,
// and send the new line to the "print" goroutine.
func (pipeline *Pipeline) doColor() {
	handler := pipeline.colorHandler
	for json := range pipeline.colorChan {
		output := handler.Pretty(json)
		pipeline.printChan <- output
	}
	pipeline.colorDone <- struct{}{}
}

// startPrint launches "print" goroutine.
func (pipeline *Pipeline) startPrint() {
	pipeline.printChan = make(chan []byte, 1024)
	pipeline.printDone = make(chan struct{})
	go pipeline.doPrint()
}

// stopPrint shutdowns "print" goroutine.
func (pipeline *Pipeline) stopPrint() {
	close(pipeline.printChan)
	<-pipeline.printDone
}

// doPrint receives a line to print on stdout.
func (pipeline *Pipeline) doPrint() {
	for json := range pipeline.printChan {
		fmt.Fprint(pipeline.writer, string(json))
	}
	pipeline.printDone <- struct{}{}
}

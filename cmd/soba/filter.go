package main

import (
	"bytes"

	fjson "github.com/valyala/fastjson"

	"github.com/novln/soba"
)

// DefaultFilterOptions is the default options for the filter handler.
var DefaultFilterOptions = &FilterOptions{
	Loggers: FilterAction{
		Include: []string{},
		Exclude: []string{},
	},
	Messages: FilterAction{
		Include: []string{},
		Exclude: []string{},
	},
	Levels: FilterAction{
		Include: []string{},
		Exclude: []string{},
	},
}

// FilterOptions is the configuration for the filter handler.
type FilterOptions struct {
	// Loggers are the filters used for the logger name.
	// This filter type will checks that the name start (or don't) with the given prefix.
	Loggers FilterAction
	// Levels are the filters used for the entry priority.
	// This filter type will checks that the level is equals (or not) to the given value.
	Levels FilterAction
	// Messages are the filters used for the entry message.
	// This filter type will checks that the message contains (or not) the given value.
	Messages FilterAction
}

// HasFilters returns true is this options has filters.
func (opts FilterOptions) HasFilters() bool {
	return opts.Loggers.HasFilters() || opts.Levels.HasFilters() || opts.Messages.HasFilters()
}

// FilterAction is a list of inclusion or exclusion values.
type FilterAction struct {
	// Include is a list of values used to retain a json line.
	Include []string
	// Exclude is a list of values used to discard a json line.
	Exclude []string
}

// HasFilters returns true is this action has a inclusion or exclusion list.
func (action FilterAction) HasFilters() bool {
	return len(action.Include) > 0 || len(action.Exclude) > 0
}

// Filter inspect the input json to establish if it's required using a list of custom rules.
// This function return the json if it's desired, nil otherwise.
func Filter(json []byte, opts *FilterOptions) []byte {
	return NewFilterHandler(opts).Filter(json)
}

// FilterHandler will inspects the input json to establish if it's required using a list of custom rules.
// It will discard the other ones.
type FilterHandler struct {
	parser  *fjson.Parser
	filters FilterOptions
}

// NewFilterHandler creates a new filter instance.
func NewFilterHandler(opts *FilterOptions) *FilterHandler {
	if opts == nil {
		opts = DefaultFilterOptions
	}

	handler := &FilterHandler{
		parser:  &fjson.Parser{},
		filters: *opts,
	}

	return handler
}

// Filter inspect the input json to establish if it's required using a list of custom rules.
// This function return the json if it's desired, nil otherwise.
func (handler *FilterHandler) Filter(json []byte) []byte {
	err := fjson.ValidateBytes(json)
	if err != nil {
		return nil
	}

	if !handler.filters.HasFilters() {
		return json
	}

	v, err := handler.parser.ParseBytes(json)
	if err != nil {
		return json
	}

	level := v.GetStringBytes(soba.LevelKey)
	if filter(bytes.Equal, handler.filters.Levels, level) {
		return nil
	}

	logger := v.GetStringBytes(soba.LoggerKey)
	if filter(bytes.HasPrefix, handler.filters.Loggers, logger) {
		return nil
	}

	message := v.GetStringBytes(soba.MessageKey)
	if filter(bytes.Contains, handler.filters.Messages, message) {
		return nil
	}

	return json
}

func filter(operation func([]byte, []byte) bool, filters FilterAction, value []byte) bool {

	// If there is no filter, no need to excute the exclusion/inclusion.
	if !filters.HasFilters() {
		return false
	}

	// By default, accept everything.
	filter := false

	// If there is no exclusion filter, only do inclusion.
	if len(filters.Exclude) == 0 {
		filter = true
	}

	// If there is no inclusion filter, only do exclusion.
	if len(filters.Include) == 0 {
		filter = false
	}

	// Finally, start with the exclusion pattern to finish by the inclusion pattern.
	// This allows to have global exclusion and specific inclusion.

	for _, exclude := range filters.Exclude {
		v := operation(value, []byte(exclude))
		if v {
			filter = true
			break
		}
	}

	for _, include := range filters.Include {
		v := operation(value, []byte(include))
		if v {
			filter = false
			break
		}
	}

	return filter
}

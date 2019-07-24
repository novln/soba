package main

import (
	"strings"

	"github.com/zchee/color"
)

// Forked from https://github.com/tidwall/pretty

// DefaultColorOptions is the default options for the pretty handler.
var DefaultColorOptions = &ColorOptions{
	EnableColor:  true,
	KeyColor:     color.FgGreen,
	StringColor:  color.FgHiYellow,
	NumberColor:  color.FgHiCyan,
	BooleanColor: color.FgHiMagenta,
	NullColor:    color.FgHiMagenta,
}

// ColorOptions is the configuration for the pretty handler.
type ColorOptions struct {
	// EnableColor will colorize the JSON.
	EnableColor bool
	// KeyColor defines the color used for JSON key type.
	KeyColor color.Attribute
	// StringColor defines the color used for JSON string type.
	StringColor color.Attribute
	// StringColor defines the color used for JSON number type.
	NumberColor color.Attribute
	// BooleanColor defines the color used for JSON boolean type.
	BooleanColor color.Attribute
	// NullColor defines the color used for JSON null value.
	NullColor color.Attribute
}

// Pretty converts the input json into a more human readable format where each
// element is on it's own line with clear indentation.
func Pretty(json []byte, opts *ColorOptions) []byte {
	return NewPrettyHandler(opts).Pretty(json)
}

// PrettyHandler will colorizes and formats the input json into a more human readable format
// where each element is on it's own line with clear indentation.
type PrettyHandler struct {
	// Content
	Buffer *strings.Builder
	// Colors
	KeyColor     *color.Color
	StringColor  *color.Color
	NumberColor  *color.Color
	BooleanColor *color.Color
	NullColor    *color.Color
}

// NewPrettyHandler creates a new PrettyHandler instance with given ColorOptions.
func NewPrettyHandler(opts *ColorOptions) *PrettyHandler {
	opts = getColorOptions(opts)

	handler := &PrettyHandler{
		Buffer:       &strings.Builder{},
		KeyColor:     color.New(opts.KeyColor, color.Bold),
		StringColor:  color.New(opts.StringColor),
		NumberColor:  color.New(opts.NumberColor),
		BooleanColor: color.New(opts.BooleanColor),
		NullColor:    color.New(opts.NullColor),
	}

	// Disable colors if it was explicit.
	if !opts.EnableColor {
		handler.KeyColor.DisableColor()
		handler.StringColor.DisableColor()
		handler.NumberColor.DisableColor()
		handler.BooleanColor.DisableColor()
		handler.NullColor.DisableColor()
	}

	return handler
}

// Pretty converts the input json into a more human readable.
func (handler *PrettyHandler) Pretty(json []byte) []byte {
	handler.prettyAny(json, 0, "  ", 0)
	if handler.Buffer.Len() > 0 {
		handler.writeByte('\n')
	}

	output := []byte(handler.Buffer.String())
	handler.Buffer.Reset()
	return output
}

// nolint: gocyclo
func (handler *PrettyHandler) prettyAny(json []byte, i int, indent string, tabs int) int {
	for i < len(json) {
		if json[i] <= ' ' {
			i++
			continue
		}
		if json[i] == '"' {
			return handler.prettyString(json, i)
		}
		if (json[i] >= '0' && json[i] <= '9') || json[i] == '-' {
			return handler.prettyNumber(json, i)
		}
		if json[i] == '{' {
			return handler.prettyObject(json, i, '{', '}', indent, tabs)
		}
		if json[i] == '[' {
			return handler.prettyObject(json, i, '[', ']', indent, tabs)
		}
		if json[i] == 't' {
			handler.writeColor(handler.BooleanColor, "true")
			return i + 4
		}
		if json[i] == 'f' {
			handler.writeColor(handler.BooleanColor, "false")
			return i + 5
		}
		if json[i] == 'n' {
			handler.writeColor(handler.NullColor, "null")
			return i + 4
		}
		i++
	}
	return i
}

func (handler *PrettyHandler) prettyNumber(json []byte, i int) int {
	x := i
	y := i + 1

	for y < len(json) {
		if json[y] <= ' ' || json[y] == ',' || json[y] == ':' || json[y] == ']' || json[y] == '}' {
			break
		}
		y++
	}

	handler.writeColor(handler.NumberColor, string(json[x:y]))
	return y
}

func (handler *PrettyHandler) prettyString(json []byte, i int) int {
	x, y := handler.getString(json, i)
	handler.writeColor(handler.StringColor, string(json[x:y]))
	return y
}

func (handler *PrettyHandler) prettyKey(json []byte, i int) int {
	x, y := handler.getString(json, i)
	handler.writeColor(handler.KeyColor, string(json[x:y]))
	return y
}

func (handler *PrettyHandler) prettyObject(json []byte, i int, open byte, close byte, indent string, tabs int) int {

	handler.writeByte(open)
	i++

	n := 0
	for i < len(json) {
		if json[i] <= ' ' {
			i++
			continue
		}
		if json[i] == close {
			if n > 0 {
				handler.writeByte('\n')
			}
			handler.tabs(indent, tabs)
			handler.writeByte(close)
			return i + 1
		}
		if open == '[' || json[i] == '"' {
			if n > 0 {
				handler.writeString(",")
			}

			handler.writeByte('\n')
			handler.tabs(indent, tabs+1)

			if open == '{' {
				i = handler.prettyKey(json, i)
				handler.writeString(": ")
			}

			i = handler.prettyAny(json, i, indent, tabs+1)
			i--
			n++
		}
		i++
	}
	return i
}

func (handler *PrettyHandler) tabs(indent string, tabs int) {
	for i := 0; i < tabs; i++ {
		handler.writeString(indent)
	}
}

func (handler *PrettyHandler) writeString(msg string) {
	_, err := handler.Buffer.WriteString(msg)
	if err != nil {
		panic(err)
	}
}

func (handler *PrettyHandler) writeByte(msg byte) {
	err := handler.Buffer.WriteByte(msg)
	if err != nil {
		panic(err)
	}
}

func (handler *PrettyHandler) writeColor(color *color.Color, msg string) {
	_, err := color.Fprint(handler.Buffer, msg)
	if err != nil {
		panic(err)
	}
}

func (PrettyHandler) getString(json []byte, i int) (int, int) {
	x := i
	y := i + 1

	for y < len(json) {
		if json[y] == '"' {
			sc := 0
			for j := y - 1; j > x; j-- {
				if json[j] == '\\' {
					sc++
				} else {
					break
				}
			}
			y++
			if sc%2 == 1 {
				continue
			}
			break
		}
		y++
	}

	return x, y
}

func getColorOptions(opts *ColorOptions) *ColorOptions {
	if opts == nil {
		opts = DefaultColorOptions
	}
	if opts.KeyColor == 0 {
		opts.KeyColor = DefaultColorOptions.KeyColor
	}
	if opts.StringColor == 0 {
		opts.StringColor = DefaultColorOptions.StringColor
	}
	if opts.NumberColor == 0 {
		opts.NumberColor = DefaultColorOptions.NumberColor
	}
	if opts.BooleanColor == 0 {
		opts.BooleanColor = DefaultColorOptions.BooleanColor
	}
	if opts.NullColor == 0 {
		opts.NullColor = DefaultColorOptions.NullColor
	}
	return opts
}

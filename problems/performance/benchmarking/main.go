package main

import (
	"bytes"
	"strings"
	"unicode"
)

// StringProcessor defines the interface for string processing implementations
type StringProcessor interface {
	Process(s string) string
}

// NaiveProcessor uses basic string operations
type NaiveProcessor struct{}

func (p *NaiveProcessor) Process(s string) string {
	result := ""
	for _, r := range s {
		if unicode.IsUpper(r) {
			result += string(unicode.ToLower(r))
		} else {
			result += string(unicode.ToUpper(r))
		}
	}
	return result
}

// BuilderProcessor uses strings.Builder
type BuilderProcessor struct{}

func (p *BuilderProcessor) Process(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, r := range s {
		if unicode.IsUpper(r) {
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(unicode.ToUpper(r))
		}
	}
	return builder.String()
}

// BytesProcessor uses bytes.Buffer
type BytesProcessor struct{}

func (p *BytesProcessor) Process(s string) string {
	var buffer bytes.Buffer
	buffer.Grow(len(s))
	for _, r := range s {
		if unicode.IsUpper(r) {
			buffer.WriteRune(unicode.ToLower(r))
		} else {
			buffer.WriteRune(unicode.ToUpper(r))
		}
	}
	return buffer.String()
}

// PreallocProcessor uses pre-allocated slice
type PreallocProcessor struct{}

func (p *PreallocProcessor) Process(s string) string {
	result := make([]rune, len(s))
	for i, r := range s {
		if unicode.IsUpper(r) {
			result[i] = unicode.ToLower(r)
		} else {
			result[i] = unicode.ToUpper(r)
		}
	}
	return string(result)
}

// MapProcessor uses a map for character lookup
type MapProcessor struct {
	upperToLower map[rune]rune
	lowerToUpper map[rune]rune
}

func NewMapProcessor() *MapProcessor {
	upperToLower := make(map[rune]rune)
	lowerToUpper := make(map[rune]rune)
	for r := 'A'; r <= 'Z'; r++ {
		upperToLower[r] = unicode.ToLower(r)
		lowerToUpper[unicode.ToLower(r)] = r
	}
	return &MapProcessor{
		upperToLower: upperToLower,
		lowerToUpper: lowerToUpper,
	}
}

func (p *MapProcessor) Process(s string) string {
	var builder strings.Builder
	builder.Grow(len(s))
	for _, r := range s {
		if unicode.IsUpper(r) {
			builder.WriteRune(p.upperToLower[r])
		} else {
			builder.WriteRune(p.lowerToUpper[r])
		}
	}
	return builder.String()
}

// GenerateTestString creates a test string of specified length
func GenerateTestString(length int) string {
	var builder strings.Builder
	builder.Grow(length)
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			builder.WriteRune('A' + rune(i%26))
		} else {
			builder.WriteRune('a' + rune(i%26))
		}
	}
	return builder.String()
}

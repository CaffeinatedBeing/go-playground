package main

import (
	"testing"
)

func TestFileLoggerSatisfiesPlugin(t *testing.T) {
	var _ Plugin = &FileLogger{}
}

func TestMemoryStorageSatisfiesInterfaces(t *testing.T) {
	var _ Plugin = &MemoryStorage{store: make(map[string]string)}
	var _ Storage = &MemoryStorage{store: make(map[string]string)}
}

func TestNilPluginBehavior(t *testing.T) {
	var p Plugin = (*NilPlugin)(nil)
	if p != nil {
		// This is a subtle Go gotcha: typed nil != untyped nil
		// But calling methods on p should not panic
		if p.Name() != "nil-plugin" {
			t.Errorf("Expected 'nil-plugin', got '%s'", p.Name())
		}
	}
}

func TestTypeAssertions(t *testing.T) {
	plugins := []Plugin{
		&FileLogger{filename: "test.log"},
		&MemoryStorage{store: make(map[string]string)},
	}

	for _, p := range plugins {
		// Safe type assertion
		if logger, ok := p.(Logger); ok {
			// Should not panic if Log is implemented
			// FileLogger does not implement Log, so this should not be true
			if _, isFileLogger := p.(*FileLogger); isFileLogger {
				t.Error("FileLogger should not satisfy Logger interface")
			}
			_ = logger // avoid unused variable
		}

		if storage, ok := p.(Storage); ok {
			err := storage.Save("foo", "bar")
			if err != nil {
				t.Errorf("Save failed: %v", err)
			}
			val, err := storage.Load("foo")
			if err != nil || val != "bar" {
				t.Errorf("Load failed: %v, val: %s", err, val)
			}
		}
	}
}

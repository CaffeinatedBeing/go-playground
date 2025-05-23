package main

import (
	"fmt"
)

// Plugin is the base interface for all plugins
type Plugin interface {
	Name() string
	Run() error
}

// Logger is an interface for logging plugins
type Logger interface {
	Plugin
	Log(message string) error
}

// Storage is an interface for storage plugins
type Storage interface {
	Plugin
	Save(key, value string) error
	Load(key string) (string, error)
}

// FileLogger is a logger plugin
// Implements Logger interface
// Implements Plugin interface
// Implements Log method
// Implements Run method
// Implements Name method
type FileLogger struct {
	filename string
}

func (f *FileLogger) Name() string { return "file-logger" }
func (f *FileLogger) Run() error   { fmt.Println("FileLogger running"); return nil }
func (f *FileLogger) Log(message string) error {
	fmt.Printf("[FileLogger] %s: %s\n", f.filename, message)
	return nil
}

// MemoryStorage is a storage plugin
type MemoryStorage struct {
	store map[string]string
}

func (m *MemoryStorage) Name() string                 { return "memory-storage" }
func (m *MemoryStorage) Run() error                   { fmt.Println("MemoryStorage running"); return nil }
func (m *MemoryStorage) Save(key, value string) error { m.store[key] = value; return nil }
func (m *MemoryStorage) Load(key string) (string, error) {
	v, ok := m.store[key]
	if !ok {
		return "", fmt.Errorf("not found")
	}
	return v, nil
}

// NilPlugin is a plugin that is nil (to demonstrate nil interface bugs)
type NilPlugin struct{}

func (n *NilPlugin) Name() string { return "nil-plugin" }
func (n *NilPlugin) Run() error   { return nil }

func main() {
	var plugins []Plugin
	plugins = append(plugins, &FileLogger{filename: "app.log"})
	plugins = append(plugins, &MemoryStorage{store: make(map[string]string)})
	plugins = append(plugins, nil)               // Intentionally add a nil plugin
	plugins = append(plugins, (*NilPlugin)(nil)) // Typed nil

	for _, p := range plugins {
		if p == nil {
			fmt.Println("Found nil plugin (untyped nil)")
			continue
		}

		fmt.Printf("Running plugin: %s\n", p.Name())
		p.Run()

		// Try type assertion to Logger
		if logger, ok := p.(Logger); ok {
			fmt.Println("This is a Logger plugin. Logging...")
			logger.Log("Hello from logger!")
		}

		// Try type switch
		switch v := p.(type) {
		case Storage:
			fmt.Println("This is a Storage plugin. Saving and loading...")
			v.Save("foo", "bar")
			val, err := v.Load("foo")
			if err != nil {
				fmt.Println("Load error:", err)
			} else {
				fmt.Println("Loaded value:", val)
			}
		case *NilPlugin:
			fmt.Println("This is a NilPlugin (typed nil)")
		default:
			fmt.Println("Unknown plugin type")
		}
	}
}

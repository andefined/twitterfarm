package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// ExitOnError : Terminate Program with Error
func ExitOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetAllConfigs ...
func GetAllConfigs(paths chan string) {
	home := GetHomeDir()
	go (func() error {
		defer close(paths)
		return filepath.Walk(home, func(p string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			select {
			case paths <- p:
			}
			return nil
		})
	})()
}

// SetHomeDir : Create the .twitterfarm directory under $HOME
func SetHomeDir() {
	home, err := homedir.Dir()
	ExitOnError(err)

	path := home + "/.twitterfarm"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}

	fmt.Printf("twitterfarm configuration folder: %s\n", path)
}

// GetHomeDir :
func GetHomeDir() string {
	home, _ := homedir.Dir()
	return home + "/.twitterfarm"
}

// TruncateString ...
func TruncateString(n int, s string) string {
	if n < 0 {
		return s
	}
	r := []rune(s)
	l := len(r)
	if n >= l {
		return s
	}
	if n > 3 && l > 3 {
		return string(r[:n-3]) + "..."
	}
	return string(r[:n])
}

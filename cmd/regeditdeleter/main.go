package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/pflag"

	"golang.org/x/sys/windows/registry"
)

var path string
var rootKey string

func init() {
	pflag.StringVar(&path, "path", "", "Path after the root key to the key to be deleted.")
	pflag.StringVar(&rootKey, "root", "", "Root Key of the registry entry.")
	pflag.Parse()
}

func main() {
	if len(strings.Trim(path, " ")) == 0 || len(strings.Trim(rootKey, " ")) == 0 {
		log.Fatal("Either path or root not set")
	}
	root, err := getRoot(rootKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleting HKEY_%s\\%s", rootKey, path)

	if err := registry.DeleteKey(*root, path); err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully deleted HKEY_%s\\%s", rootKey, path)

}

func getRoot(rootKey string) (*registry.Key, error) {
	switch rootKey {
	case "CLASSES_ROOT":
		return genKeyPtr(registry.CLASSES_ROOT), nil
	case "CURRENT_USER":
		return genKeyPtr(registry.CURRENT_USER), nil
	case "LOCAL_MACHINE":
		return genKeyPtr(registry.LOCAL_MACHINE), nil
	case "USERS":
		return genKeyPtr(registry.USERS), nil
	case "CURRENT_CONFIG":
		return genKeyPtr(registry.CURRENT_CONFIG), nil
	case "PERFORMANCE_DATA":
		return genKeyPtr(registry.PERFORMANCE_DATA), nil
	default:
		return nil, fmt.Errorf("Root %s not found", rootKey)
	}
}

func genKeyPtr(key registry.Key) *registry.Key {
	return &key
}

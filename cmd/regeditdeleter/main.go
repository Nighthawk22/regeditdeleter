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
var subKey bool

func init() {
	pflag.StringVar(&path, "path", "", "Path after the root key to the key to be deleted.")
	pflag.StringVar(&rootKey, "root", "", "Root Key of the registry entry.")
	pflag.BoolVar(&subKey, "sub", false, "Set flag if key is subkey.")
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

	if subKey {
		deleteSubKey(*root, path)

	} else {
		deleteKey(*root, path)

	}

	log.Printf("Successfully deleted HKEY_%s\\%s", rootKey, path)

}

func deleteSubKey(root registry.Key, path string) {
	k, err := registry.OpenKey(root, removeLastKey(path), registry.WRITE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	if err = k.DeleteValue(getLastKey(path)); err != nil {
		log.Fatal(err)
	}
}

func deleteKey(root registry.Key, path string) {
	if err := registry.DeleteKey(root, path); err != nil {
		log.Fatal(err)
	}

}

func getLastKey(path string) string {
	splitted := strings.Split(path, "\\")
	return strings.Trim(splitted[len(splitted)-1], " ")
}

func removeLastKey(path string) string {
	splittedPath := strings.Split(path, "\\")
	splittedPath = splittedPath[:len(splittedPath)-1]
	newPath := ""
	for i, p := range splittedPath {
		if i == 0 {
			newPath += p
		} else {
			newPath += "\\" + p
		}
	}
	return newPath
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

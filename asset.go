package main

import (
	"fmt"
	"os"
	"path"
)

func loadAsset(name string) string {
	dat, err := os.ReadFile(path.Join(AssetDir, fmt.Sprintf("%s.txt", name)))
	if err != nil {
		return name
	}
	return string(dat)
}

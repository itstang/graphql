//go:generate go-bindata -ignore=\.go -pkg=schema -o=bindata.go ./...
package schema

import (
	"log"
	"sort"
)

func String() string {
	b := []byte{}
	names := AssetNames()

	sort.Strings(names)
	for _, name := range names {
		data, err := Asset(name)
		if err != nil {
			log.Fatal(err)
		}
		b = append(b, data...)
	}
	return string(b)
}

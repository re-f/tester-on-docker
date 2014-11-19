package docker

import (
	"fmt"
	"testing"
)

func TestSearchConfig(t *testing.T) {
	RunTestCase(t, func(t *testing.T) {
		sections := getSections()
		for _, sec := range sections {
			fmt.Println(sec)
		}
	})

}

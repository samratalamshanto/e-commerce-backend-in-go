package common

import (
	"fmt"

	"github.com/jinzhu/copier"
)

// S = source type, R = result/target/Destincation type
func Mapper[S, R any](source S) (*R, error) {
	var destination R
	if err := copier.Copy(&destination, &source); err != nil {
		return nil, fmt.Errorf("failed to map %T to %T: %w", source, destination, err)
	}

	fmt.Println(destination)
	return &destination, nil
}

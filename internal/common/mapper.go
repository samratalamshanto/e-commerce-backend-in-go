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

	return &destination, nil
}

// S = source type, R = result/target/Destincation type
func MapperForList[S, R any](sourceList []S) ([]R, error) {
	var destinationList []R = make([]R, 0, len(sourceList)) //pre-allocation

	/*
		for index, value := range slice { ... }
		If only use one variable, it gives only the index, not the element.
	*/
	for _, singleSource := range sourceList {
		var singleDestination R
		if err := copier.Copy(&singleDestination, &singleSource); err != nil {
			return nil, fmt.Errorf("failed to map %T to %T: %w", sourceList, destinationList, err)
		}
		destinationList = append(destinationList, singleDestination)
	}

	return destinationList, nil
}

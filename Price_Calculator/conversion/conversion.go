//Outsourcing Sharable Logic Into A Package
package conversion

import (
	"strconv"
	"errors"
)
func StringsToFloats (strings []string) ([]float64, error) {

	var floats []float64
	for _, stringVal := range strings{

		floatVal, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("failed to convert stringf to float")	
		}
		floats= append(floats, floatVal)
	}
	return floats, nil

}

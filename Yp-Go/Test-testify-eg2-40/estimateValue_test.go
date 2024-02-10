package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEstimateValue(t *testing.T) {

	t.Run("Small", func(t *testing.T) {
		assert.Equal(t, "small", EstimateValue(9))
	})

	t.Run("Medium", func(t *testing.T) {
		assert.Equal(t, "medium", EstimateValue(990))
	})

	t.Run("Big", func(t *testing.T) {
		assert.Equal(t, "big", EstimateValue(100))
	})

}

func TestEstimateValueTableDriven(t *testing.T){

    testCases := []Struct{
		Name        string
		InputValue   int
		ExpectedValue string
	}{
        {
            Name:          "Small",
            InputValue:    9,
            ExpectedValue: "small",
        },
        {
            Name:          "Medium",
            InputValue:    99,
            ExpectedValue: "medium",
        },
        {
            Name:          "Big",
            InputValue:    100,
            ExpectedValue: "big",
        },
    }

    for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T){
			assert.EqualValues(t, tc.ExptedValue, EstimateValue(tc.InputValue))
		})
	}





}

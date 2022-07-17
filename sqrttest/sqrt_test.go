package sqrt

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type testCase struct {
	input          float64
	expectedOutput float64
}

func TestMany(t *testing.T) {
	for _, testcase := range populateTestCases(t) {
		t.Run(fmt.Sprintf("%f", testcase.input), func(t *testing.T) {
			actualOp, err := Sqrt(testcase.input)
			require.NoError(t, err)
			require.InDelta(t, actualOp, testcase.expectedOutput, 0.001)
		})
	}
}

func populateTestCases(t *testing.T) []testCase {
	f, err := os.Open("./sqrt_cases.csv")
	require.NoError(t, err)
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	require.NoError(t, err)

	var testcases []testCase
	for _, v := range data {
		ip, _ := strconv.ParseFloat(v[0], 64)
		op, _ := strconv.ParseFloat(v[1], 64)
		newcase := testCase{ip, op}
		testcases = append(testcases, newcase)
	}
	return testcases
}

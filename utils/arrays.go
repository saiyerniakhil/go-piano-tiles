package utils

import "fmt"

func ReplaceRow[T any](matrix [][]T, rowIndex int, newRow []T) ([][]T, error) {
	// Check if the rowIndex is within the bounds of the matrix.
	if rowIndex < 0 || rowIndex >= len(matrix) {
		return nil, fmt.Errorf("rowIndex %d is out of bounds for matrix with %d rows", rowIndex, len(matrix))
	}

	// Optional: Check if the new row has the same number of columns.
	// This is good practice for maintaining a rectangular matrix structure.
	if len(matrix) > 0 && len(matrix[rowIndex]) != 0 && len(newRow) != len(matrix[rowIndex]) {
		return nil, fmt.Errorf("new row length %d does not match existing row length %d", len(newRow), len(matrix[rowIndex]))
	}

	matrix[rowIndex] = newRow
	return matrix, nil
}

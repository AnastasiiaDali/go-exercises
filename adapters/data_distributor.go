package adapters

func DataDistributor(ArrayOfFileNamesFromCLI []string, ArrayOfNumbersFromCLI []string) []int {
	var numbers []int

	if len(ArrayOfNumbersFromCLI) != 0 && len(ArrayOfFileNamesFromCLI) != 0 {
		return nil
	} else if len(ArrayOfNumbersFromCLI) != 0 {
		numbers = GetNumbersFromCLI()
	} else if len(ArrayOfFileNamesFromCLI) != 0 {
		numbers = GetNumbersFromFile()
	}
	return numbers
}

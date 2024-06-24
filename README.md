# Statistics Calculator Project

## Overview
This project is designed to calculate several statistical measures from a given dataset read from a file. The statistical measures include:

- Average
- Median
- Variance
- Standard Deviation

The program reads a dataset from a file provided as a command-line argument, computes the specified statistics, and prints the results.

## Instructions

### Input File Format
The input file should contain one integer per line, representing a statistical population. Here is an example of the expected format:

```cmd
189 113 121 114 145 110 ...
```


### Running the Program
To run the program, use the following command format (assuming the program is implemented in Go):

```cmd
$ go run main.go data.txt
```

### Output Format
After reading the data from the file, the program should calculate and print the statistics in the following format (example values):

```cmd
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```
These values should be printed as rounded integers.

## Requirements
- The program must read data from the file specified as the command-line argument.
- It must calculate the average, median, variance, and standard deviation of the dataset.
- The results should be printed as rounded integers.

## Testing
The program will be tested using a set of randomly generated numbers. An auditor will compare the output of your program with expected results to verify correctness.

## Learning Outcomes
By completing this project, you will gain an understanding of:
- Basic statistical concepts: Average, Median, Variance, Standard Deviation
- Reading and processing data from a file using programming languages
- Implementing statistical calculations programmatically

---

Feel free to customize this README.md file further based on specific implementation details or additional instructions relevant to your project setup.


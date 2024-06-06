Statistics Calculator
Project Overview
This project is designed to calculate the following statistical measures from a given dataset:

Average
Median
Variance
Standard Deviation

The program reads data from a file provided as a command-line argument, calculates the required statistics, and prints the results.

Instructions
Input File Format
The input file should contain one integer per line, representing a statistical population. Here is an example of the expected format:

189
113
121
114
145
110
...


Running the Program
To run the program, use the following command format (if the program is implemented in Go):

$ go run your-program.go data.txt


Output Format
After reading the data from the file, the program should calculate and print the statistics in the following format (example values):

Average: 35
Median: 4
Variance: 5
Standard Deviation: 65


Requirements
The program must read data from the file specified as the command-line argument.
It must calculate the average, median, variance, and standard deviation.
The results should be printed as rounded integers.

Testing
The program will be tested using a set of randomly generated numbers. An auditor will compare the output of your program with expected results.

Learning Outcomes
By completing this project, you will gain an understanding of:

Basic statistical concepts: Average, Median, Variance, Standard Deviation
Reading and processing data from a file
Implementing statistical calculations programmatically
# Hapax Analysis

This project performs analysis on text files in a specified directory, extracting hapax legomena (words that appear only once) and saving the results in an Excel file.

## Description

The project reads all `.txt` files from a specified directory, processes the text by tokenizing it, optionally lemmatizing words, and calculates the frequency of each word. Only words that appear once (hapax legomena) are written to an Excel file.

## Features

- **Read text files** from a specified directory.
- **Tokenize text** and **lemmatize** (optional).
- **Count word frequency**, filtering out non-hapax words.
- **Save results** in an Excel file, listing hapax legomena.

## Installation

1. Clone this repository.
2. Make sure Go is installed. If not, install Go from the official website: https://golang.org/dl/
3. Install the required Go packages:
    ```bash
    go get github.com/aaaton/golem/v4
    go get github.com/xuri/excelize/v2
    ```
4. Install the required dependencies for your project:
    ```bash
    go mod tidy
    ```

## Usage

1. Change the `absDirPath` in `main.go` to the path of your text files directory.
2. Run the main program:
    ```bash
    go run main.go
    ```

3. After running the program, check the output file `hapax_list.xlsx` for the list of hapax legomena.


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- This project uses the [Golem library](https://github.com/aaaton/golem) for lemmatization.
- The output is saved using the [excelize library](https://github.com/xuri/excelize).

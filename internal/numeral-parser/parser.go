package numeral_parser

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"io"
	"regexp"
	"strings"
)

type NumeralParser struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func (p NumeralParser) Close() error {
	return p.writer.Flush()
}

func NewParser(r io.Reader, w io.Writer) *NumeralParser {
	return &NumeralParser{
		reader: bufio.NewReader(r),
		writer: bufio.NewWriter(w),
	}
}

func (p *NumeralParser) ParseAllLines() error {
	var stop bool = false
	for !stop {
		line, err := p.reader.ReadString('\n')
		if errors.Is(err, io.EOF) {
			stop = true
		} else if err != nil {
			return err
		}

		if line == "" {
			continue
		}

		newLine, err := parseLineAndChaneNumeralRussian(line)
		if err != nil {
			return err
		}

		_, err = p.writer.WriteString(newLine)
		if err != nil {
			return err
		}
	}

	return nil
}

var (
	//numeralRussianRegex = regexp.MustCompile(
	//	fmt.Sprintf(
	//		`(?i)(?:%s)`,
	//		strings.Join(
	//			lo.Map(
	//				RussianNumbers, func(item string, _ int) string {
	//					return fmt.Sprintf(`([^а-яА-Я]|^)%s(\s|$)`, item)
	//				},
	//			),
	//			"|",
	//		),
	//	),
	//)
	re                    = regexp.MustCompile(`\s`)
	russianNumeralPrinter = message.NewPrinter(language.Russian)
)

func parseLineAndChaneNumeralRussian(line string) (newLine string, err error) {
	indexes := re.FindAllStringIndex(line, -1)
	wordsGroup := readLineWordsByIndexes(line, indexes)
	newLine = line
	for _, words := range wordsGroup {
		if len(words) == 0 {
			continue
		}

		numeral, err := numeralWordsToNumeral(words)
		if err != nil {
			return "", err
		}

		wordsLine := strings.Join(words, " ")
		newLine = strings.ReplaceAll(newLine, wordsLine, russianNumeralPrinter.Sprint(numeral))
		newLine = strings.ReplaceAll(newLine, "\u00a0", " ")
	}
	return newLine, nil
}

func readLineWordsByIndexes(line string, bounds [][]int) [][]string {
	groups := make([][]string, len(bounds))

	left := 0
	i := 0
	for _, bound := range bounds {
		// read word
		word := line[left:bound[0]]
		left = bound[1]
		// check that numeral
		normalizeWord := strings.ToLower(word)
		_, isNumeral := NumeralByRussianNumber[normalizeWord]
		if !isNumeral && groups[i] != nil {
			i++
		}

		if !isNumeral {
			continue
		}

		groups[i] = append(groups[i], word)
	}
	return groups[:i+1]
}

func numeralWordsToNumeral(words []string) (uint64, error) {
	numerals := make([]Numeral, 0, len(words))
	for _, word := range words {
		numeral, find := NumeralByRussianNumber[strings.ToLower(word)]
		if !find {
			return 0, fmt.Errorf("not found numeral by %s", word)
		}
		numerals = append(numerals, numeral)
	}

	var number uint64
	index := 0
	for index < len(numerals) {
		numeral := numerals[index]

		if index+1 < len(numerals) {
			nextNumeral := numerals[index+1]
			if nextNumeral.IsMultiplier {
				index += 2
				number += numeral.Value * nextNumeral.Value
				continue
			}
		}

		number += numeral.Value
		index++
	}

	return number, nil
}

package cli

import(
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"

	"wruffic/api"
)

func ScannerHelper(prompt string, scanner *bufio.Scanner) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", err
		}
		return "", io.EOF
	}
	return strings.TrimSpace(scanner.Text()), nil
}

func GetSearchQuery() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	SearchQuery, err := ScannerHelper("Enter search query: ", scanner)
	if err != nil {
		return "", err
	}
	return SearchQuery, nil
}

func buildChoices(r api.SearchResult) []Choice {
	var choices []Choice

	for _, a := range r.Artists {
		if a.AC == 0 {
			continue
		}
		choices = append(choices, Choice{
			Label: fmt.Sprintf("%s (%d albums)", a.Name, a.AC),
			Kind:  "artist",
			ID:    a.ID,
		})
	}

	for _, al := range r.Albums {
		choices = append(choices, Choice{
			Label: fmt.Sprintf("%s - %s (%d)", al.Name, al.Artist, al.Year),
			Kind:  "album",
			ID:    al.ID,
		})
	}

	for _, s := range r.Songs {
		choices = append(choices, Choice{
			Label: fmt.Sprintf("%s - %s", s.Title, s.Artist),
			Kind:  "song",
			ID:    s.ID,
			Song:  s,
		})
	}

	return choices
}

func printChoices(choices []Choice) {
	last := ""
	for i, c := range choices {
		if c.Kind != last {
			fmt.Printf("\n== %ss ==\n", c.Kind)
			last = c.Kind
		}
		fmt.Printf("%2d) %s\n", i+1, c.Label)
	}
	fmt.Println("\n 0) back")
}

func readPick(choices []Choice) (Choice, bool, error) {
	sc := bufio.NewScanner(os.Stdin)
	line, err := ScannerHelper("> ", sc)
	if err != nil {
		return Choice{}, false, err
	}

	n, convErr := strconv.Atoi(line)
	if convErr != nil || n < 1 || n > len(choices) {
		return Choice{}, false, nil
	}
	return choices[n-1], true, nil
}

func PickSongs(client *api.Client, results api.SearchResult,) (string, string, error) {
	choices := buildChoices(results)
	printChoices(choices)

	pick, ok, err := readPick(choices)
	if err != nil {
		return "", "", fmt.Errorf("i ewwowed sawwyyyy 3:\nError: %v", err)
	}
	if !ok {
		return "", "", nil
	}

	switch pick.Kind {
	case "album":
		fmt.Printf("Song kind ID is: %v\n", pick.ID)
		return pick.ID, pick.Kind, nil
	case "song":
		fmt.Printf("Song kind ID is: %v\n", pick.ID)
		return pick.ID, pick.Kind, nil
	case "artist":
		fmt.Println("artist browsing no no exist")
	}
	return "", "", nil
}

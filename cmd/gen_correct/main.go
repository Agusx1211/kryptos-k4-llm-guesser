package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"

	llm "github.com/agusx1211/kryptos-k4-llm-guesser"
)

// Transform any two or more spaces into a single space
// if the string starts with a space, leave it there
// if it is only a space, (or only spaces), return am empty string
func collapseSpaces(str string) string {
	trimmed := strings.TrimSpace(str)

	// If the string contains only spaces, return an empty string
	if trimmed == "" {
		return ""
	}

	// Replace multiple spaces with a single space, excluding leading space
	re := regexp.MustCompile(`\s{2,}`)
	collapsed := re.ReplaceAllString(trimmed, " ")

	// Preserve leading space if it exists
	if strings.HasPrefix(str, " ") {
		return " " + collapsed
	}

	return collapsed
}

func cleanup(str string) string {
	str = strings.ReplaceAll(str, "\n", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\\n", "")

	str = strings.Map(func(r rune) rune {
		if r == '?' {
			return r
		}
		if r >= 'a' && r <= 'z' {
			return r
		}
		if r >= 'A' && r <= 'Z' {
			return r
		}
		if r >= '0' && r <= '9' {
			return r
		}
		return ' '
	}, str)

	// Convert all to lowercase
	str = strings.ToLower(str)

	return collapseSpaces(str)
}

// Component defines a structure that can hold either text or an integer.
type Component struct {
	Text  string
	Num   int
	IsNum bool
}

// extractComponents converts the template to a slice of Component structs.
func extractComponents(template string) []Component {
	// Regular expression to match the numbers inside curly braces `{}`.
	re := regexp.MustCompile(`\{(\d+)\}`)
	// Split the template based on the pattern `{number}`.
	parts := re.Split(template, -1)

	// Find all numbers inside curly braces `{}`.
	matches := re.FindAllStringSubmatch(template, -1)

	// Result slice containing both text and numbers.
	result := make([]Component, 0, len(parts)+len(matches))

	// Iterate over the parts and matches concurrently.
	for i, part := range parts {
		if part != "" {
			result = append(result, Component{Text: part, IsNum: false})
		}
		if i < len(matches) {
			num, err := strconv.Atoi(matches[i][1])
			if err != nil {
				panic(err)
			}
			result = append(result, Component{Num: num, IsNum: true})
		}
	}

	return result
}

func main() {
	// Parse command line arguments
	var url string
	var template string
	var fill bool
	var min float64
	var jobs int64

	flag.StringVar(&url, "url", "http://192.168.100.202:8080/", "URL of the llama.cpp server")
	flag.StringVar(&template, "template", "", "Template to use")
	flag.BoolVar(&fill, "fill", false, "Fill in the middle (only works if the model supports it)")
	flag.Float64Var(&min, "min", 0.01, "Minimum probability to consider")
	flag.Int64Var(&jobs, "jobs", 1, "Number of concurrent LLM jobs")
	flag.Parse()

	completer := llm.NewCompleter(url, int(jobs))

	components := extractComponents(template)

	// First component determines the prefix
	prefixLen := len(components[0].Text)

	completeTemplate(completer, components, prefixLen, fill, min)
}

func completeTemplate(completer *llm.Completer, components []Component, ignoreIndex int, fill bool, min float64) {
	// The first string element of components is the prefix
	// the first number element of components is the target
	prefix := ""
	target := 0

	if components[0].IsNum {
		target = components[0].Num
		components = components[1:]
	} else {
		prefix = components[0].Text
		if len(components) == 1 {
			fmt.Println(prefix[ignoreIndex:])
			return
		}
		if !components[1].IsNum {
			panic("Invalid template")
		}

		target = components[1].Num
		components = components[2:]
	}

	suffix := ""
	if fill && len(components) > 0 {
		suffix = components[0].Text
		suffix = strings.TrimLeft(suffix, " ")
		suffix = strings.TrimRight(suffix, " ")
	}

	// Trim trailing space of prefix
	prefix = strings.TrimRight(prefix, " ")

	results := make(chan string)
	go generateCombinations(completer, prefix, suffix, "", target, min, results)

	wg := sync.WaitGroup{}
	for res := range results {
		nextComponents := make([]Component, len(components))
		copy(nextComponents, components)

		// Append the result to the template
		// either as suffix of a text component
		// or as a text component itself
		if len(nextComponents) == 0 || nextComponents[0].IsNum {
			nextComponents = append([]Component{{Text: prefix + res, IsNum: false}}, nextComponents...)
		} else {
			nextComponents[0] = Component{Text: prefix + res + nextComponents[0].Text, IsNum: false}
		}

		wg.Add(1)
		go func(nextComponents []Component) {
			defer wg.Done()
			completeTemplate(completer, nextComponents, ignoreIndex, fill, min)
		}(nextComponents)
	}

	wg.Wait()
}

func generateCombinations(completer *llm.Completer, prefix, suffix string, ongoing string, target int, min float64, results chan<- string) {
	noSpaces := strings.ReplaceAll(ongoing, " ", "")
	if len(noSpaces) == target {
		results <- ongoing
		return
	}

	if len(noSpaces) > target {
		return
	}

	resp, err := completer.MakeCompletionRequest(prefix+ongoing, suffix, 1, 1024)
	if err != nil {
		panic(fmt.Sprintf("Error: %v\n", err))
	}

	mask := make(map[string]struct{}, len(resp.CompletionProbs[0].Probs))
	for _, prob := range resp.CompletionProbs[0].Probs {
		if prob.Prob < min {
			continue
		}

		cleaned := cleanup(prob.TokenStr)
		if cleaned == "" {
			continue
		}

		// Filter out duplicates after cleaning
		if _, ok := mask[cleaned]; ok {
			continue
		}
		mask[cleaned] = struct{}{}

		next := ongoing + prob.TokenStr
		generateCombinations(completer, prefix, suffix, next, target, min, results)
	}

	if ongoing == "" {
		close(results)
	}
}

package render

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data map[string]interface{}) {

	listTemplate, _ := CreateTemplateCache()

	var buf bytes.Buffer
	err := listTemplate.ExecuteTemplate(&buf, tmpl, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error ExecuteTemplate  cache: %v", err)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Cache-Control", "max-age=3600") // Cache for 1 hour
	w.Header().Set("Expires", time.Now().Add(time.Hour).Format(http.TimeFormat))

	_, err = io.Copy(w, &buf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("Error Writing to browsercache: %v", err)
		return
	}
}

func CreateTemplateCache() (*template.Template, error) {
	var listPath []string

	listPath = append(listPath, "templates/*.tmpl")
	TemplateCache, err := ParseEachFolders(listPath)
	if err != nil {
		return nil, err
	}

	return TemplateCache, nil
}

func ParseEachFolders(listPath []string) (*template.Template, error) {

	funcs := map[string]any{
		"contains":                strings.Contains,
		"hasPrefix":               strings.HasPrefix,
		"hasSuffix":               strings.HasSuffix,
		"daysAgo":                 daysAgo,
		"formatFloatWithCommas":   formatFloatWithCommas,
		"formatIntegerWithCommas": formatIntegerWithCommas,
		"addFloat":                addFloat,
		"add":                     add,
		"asteriksString":          asteriksString,
		"formatCurrency":          formatCurrency,
		"firstWord":               firstWord,
		"formatDate":              formatDate,
		"getMonthRange":           getMonthRange,
		"displayString":           displayString,
	}

	templateBuilder := template.New("").Funcs(funcs)

	for _, value := range listPath {
		_, err := templateBuilder.ParseGlob(value)

		if err != nil {
			log.Printf("Error parser glob %s", err)
		}
	}

	return templateBuilder, nil
}

func displayString(value string) string {
	if strings.TrimSpace(value) == "" {
		return "-"
	}
	return value
}

func firstWord(words string) string {
	wordsList := strings.Split(words, " ")
	if len(wordsList) > 0 {
		return wordsList[0]
	}

	return ""
}

func formatCurrency(amount int, currency string) string {
	formattedAmount := formatIntegerWithCommas(amount)
	fmt.Println("formatted amount: " + formattedAmount)

	return fmt.Sprintf("%s %s", currency, formattedAmount)
}

func formatFloatWithCommas(amountFloat float64) string {
	amount := fmt.Sprintf("%.0f", amountFloat)

	// Format the number with commas
	parts := strings.Split(amount, ".")
	integerPart := parts[0]
	formatted := ""
	for i, c := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			formatted += ","
		}
		formatted += string(c)
	}
	if len(parts) > 1 {
		formatted += "." + parts[1]
	}
	return formatted
}

func asteriksString(s string) string {
	const maxLength = 8
	firstChar := s[:1]

	asterisks := strings.Repeat("*", maxLength)

	return firstChar + asterisks
}

func getMonthRange(start time.Time, end time.Time) int {
	start = time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), 1, 0, 0, 0, 0, end.Location())

	monthCount := 0

	for current := start; current.Before(end) || current.Equal(end); current = current.AddDate(0, 1, 0) {
		monthCount++
	}

	return monthCount
}

func formatIntegerWithCommas(amount int) string {
	amountStr := fmt.Sprintf("%d", amount)

	// Format the number with commas
	parts := strings.Split(amountStr, ".")
	integerPart := parts[0]
	formatted := ""
	for i, c := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			formatted += ","
		}
		formatted += string(c)
	}
	return formatted
}

func daysAgo(date time.Time) string {
	now := time.Now().UTC()

	now = now.Truncate(24 * time.Hour)
	date = date.UTC().Truncate(24 * time.Hour)

	days := int(now.Sub(date).Hours() / 24)

	// Return the formatted string
	return fmt.Sprintf("%d hari yang lalu", days)
}

func formatDate(date time.Time) string {

	return date.Format("02 Jan 2006")
}

func addFloat(x, y float64) float64 {
	return x + y
}

func add(x int, y int) int {
	return x + y
}

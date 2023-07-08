package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func GenerateSlug(title string) string {

	slug := strings.ToLower(title)

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, " ")

	slug = strings.TrimSpace(slug)

	slug = strings.ReplaceAll(slug, " ", "-")

	return slug
}

func GenerateSlugCamelCase(title string) string {

	// Create a regular expression pattern to match uppercase letters followed by lowercase letters
	re := regexp.MustCompile("([A-Z])")

	// Use the pattern to replace the matches with lowercase letter and underscore
	output := re.ReplaceAllStringFunc(title, func(s string) string {
		return "_" + strings.ToLower(s)
	})
	fmt.Println(output)

	// Remove leading underscore if present
	output = strings.TrimPrefix(output, "_")

	return output
}

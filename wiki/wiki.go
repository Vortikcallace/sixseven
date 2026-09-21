package wiki

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type WikipediaSummary struct {
	Title   string `json:"title"`
	Extract string `json:"extract"`
}

var httpClient = &http.Client{Timeout: 10 * time.Second}

func Meme() (string, string, error) {
	topics := []string{
		"Alzheimer's disease",
		"Autism spectrum",
		"Tourette syndrome",
		"Parkinson's disease",
		"Huntington's disease",
		"Multiple sclerosis",
		"Schizophrenia",
		"Bipolar disorder",
		"Obsessive–compulsive disorder",
		"Attention deficit hyperactivity disorder",
		"Epilepsy",
		"Amyotrophic lateral sclerosis",
		"Asperger syndrome",
		"Dyslexia",
	}

	randomTopic := topics[rand.Intn(len(topics))]

	title := strings.ReplaceAll(randomTopic, " ", "_")
	encodedTopic := url.PathEscape(title)

	apiURL := fmt.Sprintf("https://en.wikipedia.org/api/rest_v1/page/summary/%s", encodedTopic)

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("User-Agent", "MyGoWikiApp/1.0 (contact@example.com)")
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return "", "", fmt.Errorf("failed to fetch from wikipedia: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("wikipedia API returned status: %s", resp.Status)
	}

	var summary WikipediaSummary
	if err := json.NewDecoder(resp.Body).Decode(&summary); err != nil {
		return "", "", fmt.Errorf("failed to decode response: %w", err)
	}

	if summary.Extract == "" {
		return "", "", fmt.Errorf("empty extract for topic %q", summary.Title)
	}

	return summary.Title, summary.Extract, nil
}

func Main() {
	title, text, err := Meme()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("TOPIC: %s\n\n", title)
	fmt.Println(text)
}

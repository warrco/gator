package commands

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Set up HTTP client
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Build request with context
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	req.Header.Set("User-Agent", "gator")
	if err != nil {
		return nil, fmt.Errorf("unable to construct request: %w", err)
	}

	// Retrieve response with HTTP client
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("unable to receive a response: %w", err)
	}
	defer res.Body.Close()

	// Retrieve body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response body: %w", err)
	}

	// Unmarshal XML content
	d := RSSFeed{}
	err = xml.Unmarshal(data, &d)
	if err != nil {
		return nil, fmt.Errorf("failed to decode XML body: %w", err)
	}

	unescaped_data := unescapeHTML(&d)

	return unescaped_data, nil
}

func unescapeHTML(r *RSSFeed) *RSSFeed {
	newFeed := RSSFeed{}

	newFeed.Channel.Title = html.UnescapeString(r.Channel.Title)
	newFeed.Channel.Description = html.UnescapeString(r.Channel.Description)
	newFeed.Channel.Link = r.Channel.Link

	for _, item := range r.Channel.Item {

		newItem := RSSItem{
			Title:       html.UnescapeString(item.Title),
			Link:        item.Link,
			Description: html.UnescapeString(item.Description),
			PubDate:     item.PubDate,
		}
		newFeed.Channel.Item = append(newFeed.Channel.Item, newItem)
	}
	return &newFeed
}

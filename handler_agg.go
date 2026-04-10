package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"html"
	"io"
	"net/http"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("time_between_reqs required")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return err
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	// rssFeed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return err
	// }

	// rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	// rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	// for i, item := range rssFeed.Channel.Item {
	// 	rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
	// 	rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	// }

	// fmt.Printf("%+v\n", rssFeed)
	// return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	rssFeed := RSSFeed{}
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return nil, err
	}

	return &rssFeed, nil

}

func scrapeFeeds(s *state) error {
	nextFeedToFetch, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		ID:            nextFeedToFetch.ID,
		LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	rssFeed, err := fetchFeed(context.Background(), nextFeedToFetch.Url)
	if err != nil {
		return err
	}

	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	// rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)

	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		// rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	fmt.Printf("%+v\n", rssFeed)
	return nil
}

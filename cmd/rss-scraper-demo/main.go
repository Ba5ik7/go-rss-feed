package main

import (
	"fmt"

	"github.com/Ba5ik7/go-rss-feed/config"
	"github.com/Ba5ik7/go-rss-feed/internal/scraper"
	"github.com/Ba5ik7/go-rss-feed/internal/utils"
)

func main() {
	logger := utils.InitLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config", "error", err)
		return
	}

	fmt.Printf("Scraping RSS scraping...")
	scraper.ScrapeFeeds(cfg.Feeds)
	fmt.Printf("Scraping complete.")
}
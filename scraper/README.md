# Cisco IOS Command Scraper

This directory contains a Go application that scrapes the official Cisco IOS command reference website to generate the data used by the LSP.

## How it Works

The scraper uses the `colly` library to crawl the Cisco documentation, starting from a main command reference page. It follows links to individual command pages and extracts information such as the command name, description, syntax, and usage guidelines.

Once the data is collected, it outputs a JSON file containing the scraped data. This JSON file is then used by the LSP to provide features like autocompletion and hover documentation.

## Running the Scraper

To run the scraper and regenerate the `commands.json` file, follow these steps:

1.  Navigate to the `scraper` directory:
    ```sh
    cd scraper
    ```

2.  Run the application, redirecting the output to the `commands.json` file:
    ```sh
    go run . > ../lsp/ios/commands.json
    ```

**Note:** The scraper uses a local cache (`./cache`) to avoid re-downloading pages unnecessarily. If you want to force a fresh scrape, delete the `cache` directory.

## Risks and Dependencies

This scraper is highly dependent on the structure and layout of the Cisco IOS command reference website. Any changes to the HTML structure of the website could break the scraper, requiring updates to the scraping logic.

#!/bin/bash

# Base directory for the HTML standard specification
HTML_SPEC_PARENT_DIR="standards/specs/html"
# Directory to save downloaded HTML files
FETCHED_HTML_DIR="${HTML_SPEC_PARENT_DIR}/fetched_pages"
# Directory where the main Markdown files converted from HTML will be stored
MD_PAGES_DIR="${HTML_SPEC_PARENT_DIR}/md_pages"

# Create the necessary directories if they don't exist
mkdir -p "${FETCHED_HTML_DIR}"
mkdir -p "${MD_PAGES_DIR}"

echo "Starting processing of HTML standard files..."
BASE_URL="https://html.spec.whatwg.org/multipage"

# Get the list of HTML files to process
# Assuming extract_html_links.py is in the current working directory
# and outputs one filename per line.
HTML_FILES=$(python3 extract_html_links.py)

if [ -z "$HTML_FILES" ]; then
    echo "No HTML files found by extract_html_links.py. Exiting."
    exit 1
fi

# Loop through each HTML file
while IFS= read -r html_filename; do
    if [ -z "$html_filename" ]; then
        continue # Skip empty lines if any
    fi

    # Construct URL and path for the fetched HTML file
    fetch_url="${BASE_URL}/${html_filename}"
    full_html_path="${FETCHED_HTML_DIR}/${html_filename}"

    echo "Processing: ${html_filename}"
    echo "  Downloading from: ${fetch_url}"
    curl -s -L -o "${full_html_path}" "${fetch_url}"
    if [ $? -ne 0 ]; then
        echo "  Download failed for ${fetch_url}. Skipping this file."
        echo "----------------------------------------"
        continue
    fi
    
    # Construct Markdown filename (e.g., introduction.html -> introduction.md)
    markdown_filename="${html_filename%.html}.md"
    full_markdown_path="${MD_PAGES_DIR}/${markdown_filename}"

    echo "Processing: ${full_html_path}"

    # Step 1: Convert HTML to Markdown using pandoc
    echo "  Converting to Markdown: ${full_markdown_path}"
    pandoc "${full_html_path}" -f html -t markdown -o "${full_markdown_path}"
    if [ $? -ne 0 ]; then
        echo "  Pandoc failed for ${full_html_path}. Skipping this file."
        echo "----------------------------------------"
        continue
    fi

    # Step 2: Split the Markdown file into sections
    # The split_markdown.py script will create a 'sections' subdirectory
    # relative to the location of full_markdown_path
    echo "  Splitting Markdown into sections..."
    python3 split_markdown.py "${full_markdown_path}"
    if [ $? -ne 0 ]; then
        echo "  split_markdown.py failed for ${full_markdown_path}."
    else
        echo "  Successfully processed ${html_filename}"
    fi
    echo "----------------------------------------"

done <<< "$HTML_FILES"

echo "HTML standard processing complete."
echo "Downloaded HTML files are in: ${FETCHED_HTML_DIR}"
echo "Main Markdown files are in: ${MD_PAGES_DIR}"
echo "Sectioned Markdown files are in subdirectories like: ${MD_PAGES_DIR}/filename/sections/"

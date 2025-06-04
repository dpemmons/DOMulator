import os
import re
from html.parser import HTMLParser

class LinkExtractor(HTMLParser):
    def __init__(self):
        super().__init__()
        self.unique_links = set()

    def handle_starttag(self, tag, attrs):
        if tag == 'a':
            attrs_dict = dict(attrs)
            href = attrs_dict.get('href')
            if href:
                # Ignore fragment-only links, javascript links, and external URLs
                if href.startswith('#') or href.startswith('javascript:') or \
                   href.startswith('http://') or href.startswith('https://') or \
                   href.startswith('mailto:') or href.startswith('tel:'):
                    return
                
                # Remove fragment identifiers from the URL
                href_without_fragment = href.split('#')[0]

                # Ensure it's an HTML file and not an empty string after removing fragment
                if href_without_fragment and href_without_fragment.endswith('.html'):
                    self.unique_links.add(href_without_fragment)

def extract_local_html_links_stdlib(html_file_path):
    """
    Parses an HTML file using html.parser and extracts unique local HTML links.
    Ignores fragment-only links and external URLs.
    """
    parser = LinkExtractor()
    with open(html_file_path, 'r', encoding='utf-8') as f:
        html_content = f.read()
        parser.feed(html_content)
            
    return sorted(list(parser.unique_links))

if __name__ == "__main__":
    index_file = "standards/specs/html/html-living-standard.html"
    links = extract_local_html_links_stdlib(index_file)
    
    # Output the links, one per line
    for link in links:
        print(link)

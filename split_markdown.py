import re
import os

def sanitize_filename(name):
    """Sanitizes a string to be a valid filename."""
    name = name.lower()
    name = re.sub(r'\[\.?secno\]', '', name)  # Remove section numbers like [1. ]
    name = re.sub(r'\[.*?\]\(.*?\)', '', name) # Remove markdown links
    name = re.sub(r'\{.*?\}', '', name)      # Remove attributes like {#id .class}
    name = re.sub(r'[^\w\s-]', '', name)     # Remove non-alphanumeric characters except spaces and hyphens
    name = re.sub(r'\s+', '-', name)         # Replace spaces with hyphens
    name = name.strip('-')
    return name

def split_markdown_into_chunks(input_filepath, output_dir):
    """
    Splits a large markdown file into smaller chunks based on H2 headings.
    Each chunk is saved as a separate .md file in the output_dir.
    """
    if not os.path.exists(output_dir):
        os.makedirs(output_dir)

    with open(input_filepath, 'r', encoding='utf-8') as f:
        lines = f.readlines()

    current_chunk_lines = []
    current_filename = None
    h2_pattern = re.compile(r"^(##\s+.*)")
    section_counter = 0

    for line in lines:
        match = h2_pattern.match(line)
        if match:
            # If we have a current chunk, write it to a file
            if current_filename and current_chunk_lines:
                with open(os.path.join(output_dir, current_filename), 'w', encoding='utf-8') as chunk_file:
                    chunk_file.writelines(current_chunk_lines)
                current_chunk_lines = []

            # Start a new chunk
            heading_text = match.group(1).strip()
            # Try to extract a number for ordering, otherwise use a generic name
            num_match = re.match(r"##\s*\[?(\d+(\.\d+)*)\.?\s*\]?", heading_text) # Matches ## [1. ], ## 1.2.
            raw_name_part = re.sub(r"##\s*\[?(\d+(\.\d+)*)\.?\s*\]?", "", heading_text).strip() # Get text after number

            if num_match:
                section_number_str = num_match.group(1).replace('.', '_')
                # Pad with zero if it's a single digit main section for sorting
                if '.' not in section_number_str and len(section_number_str) == 1:
                     section_number_str = f"0{section_number_str}"
                base_name = sanitize_filename(raw_name_part)
                current_filename = f"{section_number_str}-{base_name}.md"
                if not base_name: # Handle cases like "## [2. ]{.secno}"
                    current_filename = f"{section_number_str}-section.md"

            else: # For unnumbered sections like "Acknowledgments"
                section_counter += 1 # Use a counter for potentially duplicate unnumbered sections
                base_name = sanitize_filename(raw_name_part)
                if not base_name: # If sanitizing results in empty, use a generic name
                    current_filename = f"unnumbered-section-{section_counter}.md"
                else:
                    current_filename = f"{base_name}.md"
            
            print(f"Starting new chunk: {current_filename}")


        if current_filename: # Only add lines if we've found at least one H2
            current_chunk_lines.append(line)

    # Write the last chunk
    if current_filename and current_chunk_lines:
        with open(os.path.join(output_dir, current_filename), 'w', encoding='utf-8') as chunk_file:
            chunk_file.writelines(current_chunk_lines)
    
    print(f"Markdown splitting complete. Chunks saved in {output_dir}")

if __name__ == "__main__":
    import argparse

    parser = argparse.ArgumentParser(description="Splits a Markdown file into H2-based sections.")
    parser.add_argument("input_file", help="Path to the input Markdown file.")
    args = parser.parse_args()

    input_filepath = args.input_file
    base_input_dir = os.path.dirname(input_filepath)
    output_directory_name = "sections" # Changed from chunks to sections
    output_dir_path = os.path.join(base_input_dir, output_directory_name)
    
    print(f"Input file: {input_filepath}")
    print(f"Output directory: {output_dir_path}")

    split_markdown_into_chunks(input_filepath, output_dir_path)

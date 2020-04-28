# Read from the file file.txt and output all valid phone numbers to stdout.
cat file.txt | awk '$0~/^[0-9]{3}-[0-9]{3}-[0-9]{4}$/ || $0~/^\([0-9]{3}\) [0-9]{3}-[0-9]{4}$/'
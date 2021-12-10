# DATA MANIPULATOR

Split data into smaller chunks.

### Usage:

- `--help`: Show this help.
- `-f` or `--filepath`: The file to split (prefixed with `@`). e.g. `./data_manipulator -f "@/path/to/file.csv"`.
- `-s` or `--split`: The number of chunks to split the data into. e.g. `./data_manipulator -f "@/path/to/file.csv" -s 10`. Default: `100`.

Files are splitted and saved into `./chunks/` directory.
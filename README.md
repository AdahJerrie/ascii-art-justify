# ascii-art-justify

A command-line tool that renders text as ASCII art with terminal alignment support.

---

## Usage

```bash
go run . [--align=ALIGNMENT] "TEXT"
```

### Examples

```bash
# Default left alignment
go run . "Hello"

# Right align
go run . --align=right "Hello"

# Center align
go run . --align=center "Hello"

# Justify
go run . --align=justify "Hello World"
```

---

## Supported Alignments

| Flag | Description |
|---|---|
| `left` | Default, no padding |
| `right` | Pads content to the right edge of the terminal |
| `center` | Centers content in the terminal |
| `justify` | Distributes spaces between characters to fill the terminal width |

---

## How Alignment Works

The terminal width is detected automatically via a syscall. Content width is measured from the rendered ASCII art output. Padding is then calculated and applied per row — since each ASCII art character spans 8 rows, every row receives the same padding to maintain correct structure.

For **justify**, spaces are distributed evenly across the gaps between characters. If spaces don't divide evenly, the leftmost gaps receive the extra space.

---

## Requirements

- Go 1.18+
- Linux/Unix terminal with ANSI support
- `standard.txt` banner file in the project root

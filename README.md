![Example](res/example.png)

# Printy

Printy is an improvement on 'echo' that parses color codes and converts them to ANSI codes. Color codes are (currently) in the format of `#[0-9, b or r]#`. The numeric color codes set the foreground color, `b` enables bold text, and `r` resets.

## Installation

For **\*nix** users, there's a nice little script in `scripts/build.sh` which not only compiles `main.go` for you, but also links it to `/usr/local/bin`. I've only tested this on macOS so far, but it should work on all platforms.

## Codes

| Printy | ANSI       | Description                              |
| ------ | ---------- | ---------------------------------------- |
| `#r#`  | `\033[0m`  | Resets all formatting                    |
| `#b#`  | `\033[1m`  | Enables bold text                        |
| `#0#`  | `\033[30m` | Sets the foreground color to black       |
| `#1#`  | `\033[31m` | Sets the foreground color to red         |
| `#2#`  | `\033[32m` | Sets the foreground color to green       |
| `#3#`  | `\033[33m` | Sets the foreground color to yellow      |
| `#4#`  | `\033[34m` | Sets the foreground color to blue        |
| `#5#`  | `\033[35m` | Sets the foreground color to magenta     |
| `#6#`  | `\033[36m` | Sets the foreground color to cyan        |
| `#7#`  | `\033[37m` | Sets the foreground color to white       |
| `#9#`  | `\033[39m` | Sets the foreground color to the default |

## Notes

This is a heavy WIP project, and almost everything about this is subject to change. As of writing this README, this is only my second day using the Go programming language.

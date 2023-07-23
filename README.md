# Lines

A CLI tool that takes text and breaks into multiple lines of a desired length.

## Dependencies

* Go v1.20

## Getting Started

1. Clone this repository.

```zsh
git clone https://github.com/mnys176/lines.git --depth 1
```

2. Check the value of the `GOBIN` environment variable.

```zsh
go env GOBIN
```

This is the path that will need to appear in your system `PATH` in order to invoke the tool. Update this value if necessary.

```zsh
go env -w GOBIN=<install-path>
```

3. From the root of the repository, install the tool to the `GOBIN` directory.

```zsh
cd <repository-root>
go install
```

4. Get some text.

```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.
```

5. Pipe it through the tool, e.g., from a file.

```zsh
cat lorem.txt | lines
```

```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod
tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim
veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea
commodo consequat. Duis aute irure dolor in reprehenderit in voluptate
velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint
occaecat cupidatat non proident, sunt in culpa qui officia deserunt
mollit anim id est laborum.
```

## "Polishing" Text

The `lines` tool has to make some opinionated adjustments to the input text in order to process it effectively.

1. The tool defines a "word" to match the regular expression `/\S+/`. Extra whitespace must be stripped to identify words.
2. Multiple words are concatenated together with space characters to form a "paragraph". Again, paragraphs are stripped of leading or trailing whitespace.
3. Multiple paragraphs are joined together by two newline characters (`\n\n`) to form an "essay".

The following input text...

```


		    There should be no whitespace before this paragraph.



This paragraph		  	    should be      joined together.

This paragraph
should remain as-is
because single
newlines do not
delimit paragraphs.

At lease two newlines do.

There should not be extra whitespace after this paragraph.





```

...will be cleaned up to this.

```
There should be no whitespace above this paragraph.

This paragraph should be joined together.

This paragraph
should remain as-is
because single
newlines do not
delimit paragraphs.

At lease two newlines do.

There should not be extra whitespace after this paragraph.
```

## Setting Length

The default line length is 72 characters. This can be changed using the `--length` flag.

```zsh
cat lorem.txt | lines --length 24
```

```
Lorem ipsum dolor sit
amet, consectetur
adipiscing elit, sed do
eiusmod tempor
incididunt ut labore et
dolore magna aliqua. Ut
enim ad minim veniam,
quis nostrud
exercitation ullamco
laboris nisi ut aliquip
ex ea commodo consequat.
Duis aute irure dolor in
reprehenderit in
voluptate velit esse
cillum dolore eu fugiat
nulla pariatur.
Excepteur sint occaecat
cupidatat non proident,
sunt in culpa qui
officia deserunt mollit
anim id est laborum.
```

**NOTE:** Words longer than this length are omitted from the output.

# Prefixing and Suffixing

Each line can be framed with a prefix and suffix using the `--prefix` and `--suffix` flags. These optional additions will be taken into account as to not exceed the configured length.

```zsh
# It also supports UTF-8 encoding.
lines -p '🔥' -s '🌙' "
Hello
Goodbye
I am back again!
"
```

```
🔥Hello🌙
🔥Goodbye🌙
🔥I'm back again!🌙
```

// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["grep"] = model.Subcommand{
		Name:        []string{"grep"},
		Description: `Matches patterns in input text. Supports simple patterns and regular expressions`,
		Args: []model.Arg{{
			Name: "search pattern",
		}, {
			Templates: []model.Template{model.TemplateFilepaths},
			Name:      "file",
		}},
		Options: []model.Option{{
			Name:        []string{"--help"},
			Description: `Print a usage message briefly summarizing these command-line options and the bug-reporting address, then exit`,
		}, {
			Name:        []string{"-E", "--extended-regexp"},
			Description: `Interpret PATTERN as an extended regular expression (-E is specified by POSIX.)`,
		}, {
			Name:        []string{"-F", "--fixed-string"},
			Description: `Interpret PATTERN as a list of fixed strings, separated by newlines, any of which is to be matched. (-F is specified by POSIX.)`,
		}, {
			Name:        []string{"-G", "--basic-regexp"},
			Description: `Interpret PATTERN as a basic regular expression (BRE, see below). This is the default`,
		}, {
			Name:        []string{"-e", "--regexp"},
			Description: `Use PATTERN as the pattern. This can be used to specify multiple search patterns, or to protect a pattern beginning with a hyphen (-). (-e is specified by POSIX.)`,
			Args: []model.Arg{{
				Name: "pattern",
			}},
		}, {
			Name:        []string{"-i", "--ignore-case", "-y"},
			Description: `Ignore case distinctions in both the PATTERN and the input files. (-i is specified by POSIX.)`,
		}, {
			Name:        []string{"-v", "--invert-match"},
			Description: `Invert the sense of matching, to select non-matching lines. (-v is specified by POSIX.)`,
		}, {
			Name:        []string{"-w", "--word-regexp"},
			Description: `Select only those lines containing matches that form whole words. The test is that the matching substring must either be at the beginning of the line, or preceded by a non-word constituent character. Similarly, it must be either at the end of the line or followed by a non-word constituent character. Word-constituent characters are letters, digits, and the underscore`,
		}, {
			Name:        []string{"-x", "--line-regexp"},
			Description: `Select only those matches that exactly match the whole line. (-x is specified by POSIX.)`,
		}, {
			Name:        []string{"-c", "--count"},
			Description: `Suppress normal output; instead print a count of matching lines for each input file. With the -v, --invert-match option, count non-matching lines. (-c is specified by POSIX.)`,
		}, {
			Name:        []string{"--color"},
			Description: `Surround the matched (non-empty) strings, matching lines, context lines, file names, line numbers, byte offsets, and separators (for fields and groups of context lines) with escape sequences to display them in color on the terminal. The colors are defined by the environment variable GREP_COLORS. The deprecated environment variable GREP_COLOR is still supported, but its setting does not have priority`,
			Args: []model.Arg{{
				Name:        "WHEN",
				Suggestions: []model.Suggestion{{Name: []string{`never`}}, {Name: []string{`always`}}, {Name: []string{`auto`}}},
			}},
		}, {
			Name:        []string{"-L", "--files-without-match"},
			Description: `Suppress normal output; instead print the name of each input file from which no output would normally have been printed. The scanning will stop on the first match`,
			ExclusiveOn: []string{"-l", "--files-with-matches"},
		}, {
			Name:        []string{"-l", "--files-with-matches"},
			Description: `Suppress normal output; instead print the name of each input file from which output would normally have been printed. The scanning will stop on the first match. (-l is specified by POSIX.)`,
			ExclusiveOn: []string{"-L", "--files-without-match"},
		}, {
			Name:        []string{"-m", "--max-count"},
			Description: `Stop reading a file after NUM matching lines. If the input is standard input from a regular file, and NUM matching lines are output, grep ensures that the standard input is positioned to just after the last matching line before exiting, regardless of the presence of trailing context lines. This enables a calling process to resume a search. When grep stops after NUM matching lines, it outputs any trailing context lines. When the -c or --count option is also used, grep does not output a count greater than NUM. When the -v or --invert-match option is also used, grep stops after outputting NUM non-matching lines`,
			Args: []model.Arg{{
				Name: "NUM",
			}},
		}, {
			Name:        []string{"-o", "--only-matching"},
			Description: `Print only the matched (non-empty) parts of a matching line, with each such part on a separate output line`,
		}, {
			Name:        []string{"-q", "--quiet", "--silent"},
			Description: `Quiet; do not write anything to standard output. Exit immediately with zero status if any match is found, even if an error was detected. Also see the -s or --no-messages option. (-q is specified by POSIX.)`,
		}, {
			Name:        []string{"-s", "--no-messages"},
			Description: `Suppress error messages about nonexistent or unreadable files. Portability note: unlike GNU grep, 7th Edition Unix grep did not conform to POSIX, because it lacked -q and its -s option behaved like GNU grep's -q option. USG -style grep also lacked -q but its -s option behaved like GNU grep. Portable shell scripts should avoid both -q and -s and should redirect standard and error output to /dev/null instead. (-s is specified by POSIX.)`,
		}, {
			Name:        []string{"-b", "--byte-offset"},
			Description: `Print the 0-based byte offset within the input file before each line of output. If -o (--only-matching) is specified, print the offset of the matching part itself`,
		}, {
			Name:        []string{"-H", "--with-filename"},
			Description: `Print the file name for each match. This is the default when there is more than one file to search`,
		}, {
			Name:        []string{"-h", "--no-filename"},
			Description: `Suppress the prefixing of file names on output. This is the default when there is only one file (or only standard input) to search`,
		}, {
			Name:        []string{"--label"},
			Description: `Display input actually coming from standard input as input coming from file LABEL. This is especially useful when implementing tools like zgrep, e.g., gzip -cd foo.gz | grep --label=foo -H something`,
			Args: []model.Arg{{
				Name: "LABEL",
			}},
		}, {
			Name:        []string{"-n", "--line-number"},
			Description: `Prefix each line of output with the 1-based line number within its input file. (-n is specified by POSIX.)`,
		}, {
			Name:        []string{"-T", "--initial-tab"},
			Description: `Make sure that the first character of actual line content lies on a tab stop, so that the alignment of tabs looks normal. This is useful with options that prefix their output to the actual content: -H,-n, and -b. In order to improve the probability that lines from a single file will all start at the same column, this also causes the line number and byte offset (if present) to be printed in a minimum size field width`,
		}, {
			Name:        []string{"-u", "--unix-byte-offsets"},
			Description: `Report Unix-style byte offsets. This switch causes grep to report byte offsets as if the file were a Unix-style text file, i.e., with CR characters stripped off. This will produce results identical to running grep on a Unix machine. This option has no effect unless -b option is also used; it has no effect on platforms other than MS-DOS and MS -Windows`,
		}, {
			Name:        []string{"--null"},
			Description: `Output a zero byte (the ASCII NUL character) instead of the character that normally follows a file name. For example, grep -lZ outputs a zero byte after each file name instead of the usual newline. This option makes the output unambiguous, even in the presence of file names containing unusual characters like newlines. This option can be used with commands like find -print0, perl -0, sort -z, and xargs -0 to process arbitrary file names, even those that contain newline characters`,
		}, {
			Name:        []string{"-A", "--after-context"},
			Description: `Print num lines of trailing context after each match`,
			Args: []model.Arg{{
				Name: "NUM",
			}},
		}, {
			Name:        []string{"-B", "--before-context"},
			Description: `Print num lines of leading context before each match. See also the -A and -C options`,
			Args: []model.Arg{{
				Name: "NUM",
			}},
		}, {
			Name:        []string{"-C", "--context"},
			Description: `Print NUM lines of output context. Places a line containing a group separator (--) between contiguous groups of matches. With the -o or --only-matching option, this has no effect and a warning is given`,
			Args: []model.Arg{{
				Name: "NUM",
			}},
		}, {
			Name:        []string{"-a", "--text"},
			Description: `Treat all files as ASCII text. Normally grep will simply print ""Binary file ... matches'' if files contain binary characters. Use of this option forces grep to output lines matching the specified pattern`,
		}, {
			Name:        []string{"--binary-files"},
			Description: `Controls searching and printing of binary files`,
			Args: []model.Arg{{
				Name: "value",
				Suggestions: []model.Suggestion{{
					Name:        []string{`binary`},
					Description: `Search binary files but do not print them`,
				}, {
					Name:        []string{`without-match`},
					Description: `Do not search binary files`,
				}, {
					Name:        []string{`text`},
					Description: `Treat all files as text`,
				}},
			}},
		}, {
			Name:        []string{"-D", "--devices"},
			Description: `Specify the demanded action for devices, FIFOs and sockets`,
			Args: []model.Arg{{
				Name: "action",
				Suggestions: []model.Suggestion{{
					Name:        []string{`read`},
					Description: `Read as if they were normal files`,
				}, {
					Name:        []string{`skip`},
					Description: `Devices will be silently skipped`,
				}},
			}},
		}, {
			Name:        []string{"-d", "--directories"},
			Description: `Specify the demanded action for directories`,
			Args: []model.Arg{{
				Name: "action",
				Suggestions: []model.Suggestion{{
					Name:        []string{`read`},
					Description: `Directories are read in the same manner as normal files`,
				}, {
					Name:        []string{`skip`},
					Description: `Silently ignore the directories`,
				}, {
					Name:        []string{`recurse`},
					Description: `Read directories recursively`,
				}},
			}},
		}, {
			Name:        []string{"--exclude"},
			Description: `Note that --exclude patterns take priority over --include patterns, and if no --include pattern is specified, all files are searched that are not excluded. Patterns are matched to the full path specified, not only to the filename component`,
			Args: []model.Arg{{
				Name:       "GLOB",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--exclude-dir"},
			Description: `If -R is specified, only directories matching the given filename pattern are searched.  Note that --exclude-dir patterns take priority over --include-dir patterns`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFolders},
				Name:       "dir",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-I"},
			Description: `Ignore binary files. This option is equivalent to --binary-file=without-match option`,
		}, {
			Name:        []string{"--include"},
			Description: `If specified, only files matching the given filename pattern are searched. Note that --exclude patterns take priority over --include patterns. Patterns are matched to the full path specified, not only to the filename component`,
			Args: []model.Arg{{
				Name:       "GLOB",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"--include-dir"},
			Description: `If -R is specified, only directories matching the given filename pattern are searched. Note that --exclude-dir patterns take priority over --include-dir patterns`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFolders},
				Name:       "dir",
				IsOptional: true,
			}},
		}, {
			Name:        []string{"-R", "-r", "--recursive"},
			Description: `Recursively search subdirectories listed`,
		}, {
			Name:        []string{"--line-buffered"},
			Description: `Force output to be line buffered. By default, output is line buffered when standard output is a terminal and block buffered otherwise`,
		}, {
			Name:        []string{"-U", "--binary"},
			Description: `Search binary files, but do not attempt to print them`,
		}, {
			Name:        []string{"-J", "-bz2decompress"},
			Description: `Decompress the bzip2(1) compressed file before looking for the text`,
		}, {
			Name:        []string{"-V", "--version"},
			Description: `Print version number of grep to the standard output stream`,
		}, {
			Name:        []string{"-P", "--perl-regexp"},
			Description: `Interpret pattern as a Perl regular expression`,
		}, {
			Name:        []string{"-f", "--file"},
			Description: `Obtain patterns from FILE, one per line. The empty file contains zero patterns, and therefore matches nothing. (-f is specified by POSIX.)`,
			Args: []model.Arg{{
				Templates: []model.Template{model.TemplateFilepaths},
				Name:      "FILE",
			}},
		}},
	}
}
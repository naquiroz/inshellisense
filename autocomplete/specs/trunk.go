// Code generated by autocomplete/extract/extract.ts. DO NOT EDIT.

package specs

import (
	"github.com/cpendery/clac/autocomplete/model"
)

func init() {
	Specs["trunk"] = model.Subcommand{
		Name:        []string{"trunk"},
		Description: `An all-in-one tool for scalably checking, formatting, and monitoring code`,
		Options: []model.Option{{
			Name:         []string{"-h", "--help"},
			Description:  `Usage information`,
			IsPersistent: true,
		}, {
			Name:         []string{"--version"},
			Description:  `The version`,
			IsPersistent: true,
		}, {
			Name:        []string{"-m", "--monitor"},
			Description: `Enable the trunk daemon to monitor file changes in your repo`,
			Args: []model.Arg{{
				Name:        "value",
				Suggestions: []model.Suggestion{{Name: []string{`true`}}, {Name: []string{`false`}}},
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--ci"},
			Description:  `Run in continuous integration mode`,
			IsPersistent: true,
		}, {
			Name:        []string{"-o", "--output"},
			Description: `Output format`,
			Args: []model.Arg{{
				Name:        "format",
				Suggestions: []model.Suggestion{{Name: []string{`text`}}, {Name: []string{`json`}}},
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"--no-progress"},
			Description:  `Don't show progress updates`,
			IsPersistent: true,
			ExclusiveOn:  []string{"--ci-progress"},
		}, {
			Name:         []string{"--ci-progress"},
			Description:  `Show updates every 30 seconds without clearing terminal screen (implicit with --ci)`,
			IsPersistent: true,
			ExclusiveOn:  []string{"--no-progress"},
		}, {
			Name:        []string{"--action_timeout"},
			Description: `How long actions (downloads, lint runs, etc.) have to time out, with units (s, ms, etc)`,
			Args: []model.Arg{{
				Name: "timeout",
			}},
			IsPersistent: true,
		}, {
			Name:         []string{"-v", "--verbose"},
			Description:  `Output details about what's happening under the hood`,
			IsPersistent: true,
		}},
		Subcommands: []model.Subcommand{{
			Name:        []string{"init"},
			Description: `Setup trunk in this repo`,
			Options: []model.Option{{
				Name:        []string{"--lock"},
				Description: `Add sha256s to trunk.yaml for additional verification`,
			}, {
				Name:        []string{"--check-sample"},
				Description: `Run "trunk check sample" without prompting`,
				ExclusiveOn: []string{"--nocheck-sample"},
			}, {
				Name:        []string{"--nocheck-sample"},
				Description: `Do not run "trunk check sample" post-init`,
				ExclusiveOn: []string{"--check-sample"},
			}, {
				Name:        []string{"--force"},
				Description: `Overwrite existing trunk.yaml`,
			}},
		}, {
			Name:        []string{"check"},
			Description: `Universal code checker`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "paths",
				IsOptional: true,
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-y", "--fix"},
				Description: `Automatically apply all fixes without prompting`,
			}, {
				Name:        []string{"-a", "--all"},
				Description: `Run on all files instead of only changed files`,
			}, {
				Name:        []string{"-n", "--no-fix"},
				Description: `Don't automatically apply fixes`,
			}, {
				Name:        []string{"--include-existing-autofixes"},
				Description: `Show autofixes for existing issues`,
			}, {
				Name:        []string{"--force"},
				Description: `Run on all files, even if ignored`,
			}, {
				Name:        []string{"--diff"},
				Description: `Diff printing mode`,
				Args: []model.Arg{{
					Name:        "mode",
					Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`compact`}}, {Name: []string{`full`}}},
				}},
			}, {
				Name:        []string{"--filter"},
				Description: `Filter the set of executed linters and/or the returned codes; use a leading '-' to exclude a linter or code`,
				Args: []model.Arg{{
					Name:       "linter or code",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--exclude"},
				Description: `Shorthand for an inverse --filter`,
				Args: []model.Arg{{
					Name:       "linter or code",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"-j", "--jobs"},
				Description: `Number of concurrent jobs (does not affect background linting)`,
				Args: []model.Arg{{
					Name: "number",
				}},
			}, {
				Name:        []string{"--sample"},
				Description: `Run each linter on N files (implies --no-fix and --all if no paths are given)`,
				Args: []model.Arg{{
					Name: "N",
				}, {
					Templates:  []model.Template{model.TemplateFilepaths},
					Name:       "paths",
					IsOptional: true,
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--upstream"},
				Description: `Upstream branch used to compute changed files (autodetected by default)`,
				Args: []model.Arg{{
					Name: "branch",
				}},
			}},
			Subcommands: []model.Subcommand{{
				Name:        []string{"upgrade"},
				Description: `Upgrade all linters to latest versions`,
				Args: []model.Arg{{
					Name:        "linters",
					Description: `Linter(s) to upgrade (upgrades all if none specified)`,
					IsVariadic:  true,
				}},
			}, {
				Name:        []string{"download"},
				Description: `Download all files needed for trunk to work offline`,
				Args: []model.Arg{{
					Name:        "tools",
					Description: `Tool(s) to download (if omitted, downloads all configured tools)`,
					IsVariadic:  true,
				}},
			}, {
				Name:        []string{"enable"},
				Description: `Enable linters`,
				Args: []model.Arg{{
					Name:        "linters",
					Description: `Linter(s) to enable`,
					IsVariadic:  true,
				}},
			}, {
				Name:        []string{"disable"},
				Description: `Disable linters`,
				Args: []model.Arg{{
					Name:        "linters",
					Description: `Linter(s) to disable`,
					IsVariadic:  true,
				}},
			}},
		}, {
			Name:        []string{"fmt"},
			Description: `Universal code formatter`,
			Args: []model.Arg{{
				Templates:  []model.Template{model.TemplateFilepaths},
				Name:       "paths",
				IsOptional: true,
				IsVariadic: true,
			}},
			Options: []model.Option{{
				Name:        []string{"-a", "--all"},
				Description: `Run on all files instead of only changed files`,
			}, {
				Name:        []string{"-n", "--no-fix"},
				Description: `Don't automatically apply fixes`,
			}, {
				Name:        []string{"--include-existing-autofixes"},
				Description: `Show autofixes for existing issues`,
			}, {
				Name:        []string{"--force"},
				Description: `Run on all files, even if ignored`,
			}, {
				Name:        []string{"--diff"},
				Description: `Diff printing mode`,
				Args: []model.Arg{{
					Name:        "mode",
					Suggestions: []model.Suggestion{{Name: []string{`none`}}, {Name: []string{`compact`}}, {Name: []string{`full`}}},
				}},
			}, {
				Name:        []string{"--filter"},
				Description: `Filter the set of executed linters and/or the returned codes; use a leading '-' to exclude a linter or code`,
				Args: []model.Arg{{
					Name:       "linter or code",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--exclude"},
				Description: `Shorthand for an inverse --filter`,
				Args: []model.Arg{{
					Name:       "linter or code",
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"-j", "--jobs"},
				Description: `Number of concurrent jobs (does not affect background linting)`,
				Args: []model.Arg{{
					Name: "number",
				}},
			}, {
				Name:        []string{"--sample"},
				Description: `Run each linter on N files (implies --no-fix and --all if no paths are given)`,
				Args: []model.Arg{{
					Name: "N",
				}, {
					Templates:  []model.Template{model.TemplateFilepaths},
					Name:       "paths",
					IsOptional: true,
					IsVariadic: true,
				}},
			}, {
				Name:        []string{"--upstream"},
				Description: `Upstream branch used to compute changed files (autodetected by default)`,
				Args: []model.Arg{{
					Name: "branch",
				}},
			}},
		}, {
			Name:        []string{"upgrade"},
			Description: `Upgrade trunk to the latest release`,
		}, {
			Name:        []string{"git-hooks"},
			Description: `Git hooks`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"install"},
				Description: `Install trunk git hooks`,
			}},
		}, {
			Name:        []string{"cache"},
			Description: `Cache management`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"clean"},
				Description: `Clean the cache`,
				Options: []model.Option{{
					Name:        []string{"--all"},
					Description: `Delete all files (including results cache)`,
				}, {
					Name:        []string{"-n", "--dry-run"},
					Description: `Print all directories that would be cleaned out without running deletion`,
				}},
			}},
		}, {
			Name:        []string{"print-config"},
			Description: `Print the resolved trunk config`,
		}, {
			Name:        []string{"daemon"},
			Description: `Daemon management`,
			Subcommands: []model.Subcommand{{
				Name:        []string{"launch"},
				Description: `Start the trunk daemon if its not already running`,
			}, {
				Name:        []string{"shutdown"},
				Description: `Shutdown the trunk daemon if it is running`,
			}, {
				Name:        []string{"status"},
				Description: `Report daemon status`,
			}},
		}},
	}
}
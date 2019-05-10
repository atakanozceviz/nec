# Nec

Nec helps you to speed up your CI for Visual Studio projects by looking up changes (using git diff) and finds out which solutions needs to build and tests needs to run. After that executes user-defined commands for the test projects and solutions.

Nec parses all the solutions (.sln) and projects (.csproj) in a folder and creates dependency graph, then uses that graph for finding dependencies.

## Install

```
go get github.com/atakanozceviz/nec
```

## Getting Started

You will need to define your commands in a configuration file. Nec supports JSON, TOML, YAML, HCL, and Java Properties files.

##### Example JSON:

```json
{
  "commands": {
    "build": {
      "description":"Builds a project and all of its dependencies.",
      "workers": 1,
      "onerror": "exit",
      "name": "dotnet",
      "args": [
        "build"
      ]
    },
    "test": {
      "description":".NET test driver used to execute unit tests.",
      "workers": 5,
      "onerror": "continue",
      "name": "dotnet",
      "args": [
        "test",
        "--logger:trx"
      ]
    }
  }
}
```

In the above example there are 2 commands, `build` and `test`. Do not change these names!

`"description"` is short description for that command.
`"workers"`: defines how many commands can be run in parallel. (Setting to 1 would run one command at a time)
`"onerror"`: defines what to do when error occurs. Can be `"exit"` or `"continue"`.
- `"exit"`: stops all the existing works and exits with error when an error occurs.
- `"continue"`: continues until all the works are done, if error occurs just prints out.
`"name"` is the command name. Must be executable.
`"args"` is the list of arguments for the command.

For the above sample:
`"test"` command will be interpreted like so:

```command
$ dotnet test --logger:trx
```

and the `"build"` command is:

```command
$ dotnet build
```

## How to use

Nec tries to load "nec.json" file in the current directory. You can specify json file path using the `-s` flag.

```console
$ nec
Nec helps you to speed up your CI for
Visual Studio projects by looking up
changes (using git diff) and finds out
which solutions needs to build and tests
needs to run. After that executes user-defined
commands for the test projects and solutions.

Nec parses all the solutions (.sln) and
projects (.csproj) in a folder and creates
dependency graph, then uses that graph for
finding dependencies.

Usage:
  nec [command]

Available Commands:
  build       Run build command.
  help        Help about any command
  test        Run test command.

Flags:
  -c, --commit string      git commit id to find affected projects (default "HEAD^")
  -h, --help               help for nec
  -i, --ignore string      ignore list file
  -s, --settings string    settings file (default is nec.json)
  -w, --walk-path string   the path to start the search for .sln files (default ".")

Use "nec [command] --help" for more information about a command.
```

While using `-c` flag, commit ID can be path to project folder with leading ID:

```console
$ nec -s path/to/*.json build -c ../GitRepositoryRoot/HEAD^^^
```

This would use `"GitRepositoryRoot"` folder to get project changes. (Runs git diff HEAD^^^ in that folder)

For the above sample json:

```console
$ nec -s path/to/*.json build -c HEAD^^
```

Command would check which projects are affected since two commits (HEAD^^) then run `dotnet build` command in the necessary solutions folders.

```console
$ nec -s path/to/*.json test -c HEAD^^
```

Command would check which projects are affected since two commits (HEAD^^) then run `dotnet test --logger:trx` command in the necessary test projects folders.

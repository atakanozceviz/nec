# Nec

Nec helps you to speed up your CI for Visual Studio projects by looking up changes (using git diff) and finds out which solutions needs to build and tests needs to run. After that executes user-defined commands for the test projects and solutions.

Nec parses all the solutions (.sln) and projects (.csproj) in a folder and creates dependency graph, then uses that graph for finding dependencies.

## Install

```
go get github.com/atakanozceviz/nec/nec
```

## Getting Started

You will need to define your commands in a json file.

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

`"description"` is short description for that command and will be shown in cli like this: 

```console
...
Available Commands:
  build       Builds a project and all of its dependencies.
  test        .NET test driver used to execute unit tests.
...
```

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
  build       Builds a project and all of its dependencies.
  help        Help about any command
  test        .NET test driver used to execute unit tests.

Flags:
  -c, --c string        Git commit id for getting changes. (default "HEAD^")
      --config string   config file (default is $HOME/.nec.yaml)
  -h, --help            help for nec
  -s, --s string        Path to settings file. (default "nec.json")
  -w, --w string        The path to start the search for .sln files. (default ".")

Use "nec [command] --help" for more information about a command.
```

`"Available Commands"` are generated from json file.

For the above sample json:

```console
$ nec -s path/to/*.json build -c HEAD^^^
```

Command would check which projects are affected since two commits (HEAD^^) then run `dotnet build` command in the necessary solutions folders.

```console
$ nec -s path/to/*.json test -c HEAD^^^
```

Command would check which projects are affected since two commits (HEAD^^) then run `dotnet test --logger:trx` command in the necessary test projects folders.

command would run `dotnet test --logger:trx` command in the necessary test projects folders.
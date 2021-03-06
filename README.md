# rockstar-cli

Command to make you a rockstar in less than 10 seconds,
based on [rockstar](https://github.com/avinassh/rockstar)

### Before

![before](https://raw.githubusercontent.com/3zcurdia/rockstar-cli/master/images/before.png)

### After

![after](https://raw.githubusercontent.com/3zcurdia/rockstar-cli/master/images/after.png)

### Punch Card

![punchcard](https://raw.githubusercontent.com/3zcurdia/rockstar-cli/master/images/punchcard.png)

## Features

* Random real comments
* Normal distributed commits
* Less commits on weekends
* Automatic push into github
* Cross compiled binaries
  * Darwin 386/amd64
  * Linux 386/amd64/arm
  * Windows 386/amd64

## Requirements

You must have ```git``` installed and for automatic github push you must have installed [hub](https://github.com/github/hub)

## Install

Download the binaries [here](https://github.com/3zcurdia/rockstar-cli/releases) or install through go

    go get github.com/3zcurdia/rockstar-cli

## Usage

    rockstar-cli [global options] command [command options] [arguments...]

    COMMANDS:
       help, h  Shows a list of commands or help for one command

    GLOBAL OPTIONS:
       --days, -d "500"       days of activity
       --code, -c "writeln('Go is Awesome!!!')" code base
       --filename, -f "main.go"     output file
       --help, -h         show help
       --version, -v        print the version

## Contribue to the project

To contribuer Just follow the next stepts:

* Check out the latest master to make sure the feature hasn't been implemented or the bug hasn't been fixed yet
* Check out the issue tracker to make sure someone already hasn't requested it and/or contributed it
* Fork the project
* Start a feature/bugfix branch
* Commit and push until you are happy with your contribution
* It is desired to add some tests for it.

## License

The MIT License (MIT)

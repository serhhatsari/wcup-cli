# wcup-cli (under development)

 > **wcup** is a CLI for football fans who wants to follow World Cup 2022 ⚽  


## Installation  

**Run**

```shell
$ npm install -g wcup-22
```
## Usage

### Commands available

```shell
wcup <command>

Commands:
  groups     Get groups  
  scores     Get scores of past and live fixtures
  fixtures   Get upcoming and past fixtures of a team


Options:
  -h, --help  Show help                                          [boolean]
  

```
#### Command `groups`
Get groups of World Cup 2022

```shell
Usage: wcup groups

Options:
  -h, --help  Show help                                          [boolean]
  -g  --group 

Examples:
  wcup groups -g "A"  
  
```


```
#### Command `scores`
Get scores of past and live fixtures

```shell
Usage: wcup scores [options]

Options:
  -h, --help  Show help                                          [boolean]
  -l, --live  Live scores                                        [boolean]
  -c, --country     Name of the country                                       [string]

Examples:
  wcup scores -t "por" -l
  
```

#### Command `fixtures`
Get upcoming and past fixtures of a league and team

```shell
Usage: wcup fixtures [options]

Options:
  -h, --help    Show help                                         [boolean]
  -c, --country     Name of the country                   [string]
  -n, --next    Next or upcoming matches                          [boolean]

Examples:
  wcup fixtures -t "arg" -n

```


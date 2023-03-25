# easy-jira-task-review
Easy creation of review tasks for Jira with Discord notifications.

## Installation
1) Install 
```shell
go install github.com/vasylNaumenko/ej@latest
```
2) Fill a configuration file named ```~/.ej.config.yaml``` using ```.ej.config.example.yaml``` as an example.
The command ```get``` can provide helpful information:
```
ej get [command]

Available Commands:
  projects    returns a list of projects
  reviewers   returns a reviewers list from the configuration file
  users       returns a list of users
```

## Usage
```
ej review -h
Creates tasks for review of your merge request.
	Examples:
		review [issue-id] [MR link1,MR link2...] -t=tag1,tag2 (creates tasks for the tag1 and tag2 assignees)
		review [issue-id] [MR link1,MR link2...] -t=tag1 -r=1  (creates tasks for the tag1 assignee and one random assignee)
		review [issue-id] [MR link1,MR link2...] -r=2  (creates tasks for two random reviewers)

Usage:
  ej review [flags]

Flags:
  -h, --help          help for review
  -r, --random int    random assignees count: -r=2 (takes random 2 assignee from a reviewers list)
  -t, --tags string   assignees tags: -t=[tag1,tag2,...]
```
You can combine ```tags``` and ```random``` assignee choices.

For example, the flags -t=n1,n2 -r=2 will create tasks for the users n1 and n2, plus it will create tasks for two additional random users.

## Configuration
The default configuration path is ```~/.ej.config.yaml.```

You can specify a configuration file path using the flag:
```
--config=./config.yaml
```

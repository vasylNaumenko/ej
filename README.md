# easy-jira-task-review
Easy creation of review tasks for the Jira with a Discord notifications.

## __Any ideas and help are welcome__

## Installation
1) Install 
```shell
go install github.com/vasylNaumenko/ej@latest
```
2) Fill a configuration file ```~/.ej.config.yaml``` using as example ```.ej.config.example.yaml```

Command ```get``` can provide some helpful information:
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
Creates the tasks for for review of yours merge request.
	Examples:
		review [issue-id] [MR link1,MR link2...] -t=tag1,tag2 (creates tasks for the tag1 and the tag2 assignees)
		review [issue-id] [MR link1,MR link2...] -t=tag1 -r=1  (creates tasks for the tag1 assignee and plus a random one)
		review [issue-id] [MR link1,MR link2...] -r=2  (creates tasks for the 2 random reviewers)

Usage:
  ej review [flags]

Flags:
  -h, --help          help for review
  -r, --random int    random assignees count: -r=2 (takes random 2 assignee from a reviewers list)
  -t, --tags string   assignees tags: -t=[tag1,tag2,...]
```
You can combine ```tags``` and ```random``` assignee choise.

For example flags ```-t=n1,n2 -r=2``` will create tasks for the users n1 and n2 plus it creates tasks for additional two random users.


## Configuration
Default configuration path is ```~/.ej.config.yaml```

You can specify a configuration file path with a flag:
```
--config=./config.yaml
```
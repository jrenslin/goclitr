# goclitr

## What is it about?

goclitr is a directory-specific command line-based to-do list manager with annotations, that aims to work on multi-user systems.

## Usage

To use goclitr in a directory, run `goclitr init`. This creates a hidden directory at `./goclitr` and adds the current filepath to a general, user-specific list stored at `~/.config/goclitr/dirs.json`. After this, you can add new tasks to the local task list using `goclitr add`. Further commands can be found below.

### Commands

The following is quoted from Goclitr's help message.

```
Command                                Description
init                                   Initialize
teardown                               Tear down
current                                Lists currently active (=not completed) issues
help                                   Print this message
list                                   List current tasks
completed                              List completed
add          <text>                    Add a task
new                                    Add a task (same as add)
delete       <ID>                      Delete task with the given ID
remove       <ID>                      Delete task (same as delete)
modify       <ID>         <text>       Modify the task's text
progress     <ID>         <int: 0-10>  Edit progress of the task's text
annotate     <ID>         <text>       Annotate a task
done         <ID>                      Finish task
complete     <ID>                      Finish task (same as done) 
listall                                Lists all projects you've worked on
project      <Project ID>              Return path of project X
```

## Inspiration / Thanks

- [Taskwarrior](https://taskwarrior.org)

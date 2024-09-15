# Task Tracker

This project is a sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/febryansambuari/task-tracker-cli.git
cd task-tracker-cli
```

Run the following command to build and run the project:

```bash
# Build the project
go build -o task-tracker-cli

# Add a task
./task-tracker-cli add [task description]
example: task-tracker-cli add "Create a new Golang project"

# Update a task
./task-tracker-cli update [task id] [task description]
example: task-tracker-cli update 1 "Create a new Javascript project"

# Delete a task
./task-tracker delete [task id]
example: task-tracker-cli delete 1

# Mark a task as done
./task-tracker-cli mark [task id] done
example: task-tracker-cli 1 done

# Mark a task as in-progress
./task-tracker-cli mark [task id] in-progress
example: task-tracker-cli 1 in-progress

# Mark a task as todo
./task-tracker-cli mark [task id] todo
example: task-tracker-cli 1 todo

# List all tasks
./task-tracker-cli list

# List all tasks that are done
./task-tracker-cli list done

# List all tasks that are in-progress
./task-tracker-cli list in-progress

# List all tasks that are todo
./task-tracker-cli list todo
```
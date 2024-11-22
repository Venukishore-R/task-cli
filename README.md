# CLI Task Manager

A simple command-line application for managing tasks, including adding, updating, deleting, and changing the status of tasks.

## Features
- **Add Tasks**: Add new tasks to the task list.
- **Update Tasks**: Modify existing task details.
- **Delete Tasks**: Remove tasks from the list.
- **Update Status**: Change the status of tasks (e.g., Todo, In Progress, Done).

## Technologies Used
- **Programming Language**: Go
- **CLI Framework**: Cobra
- **Data Storage**: JSON file

## Installation
Follow these steps to install the CLI Task Manager:

1. Clone this repository:
   ```bash
   git clone https://github.com/Venukishore-R/task-cli.git
   cd task-cli
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Build and install the project:
   ```bash
   go build
   go install
   ```
## Usage
Run the application using the following command:
```bash
task-cli
```
### 1. Add a Task
To add a new task, use:
```bash
task-cli add "YOUR TASK"
```
### 2. Update a Task
To update an existing task, use:
```bash
task-cli update <TASK_ID> <UPDATE>
```
### 3. Delete a Task
To delete a task from the list, use:
```bash
task-cli delete <TASK_ID>
```
### 4. Update the Status
To update the status of a task:
- To mark a task as "In Progress":
  ```bash
  task-cli mark-in-progress <TASK_ID>
  ```
- To mark a task as "Done":
  ```bash
  task-cli mark-done <TASK_ID>
  ```
### 5. List All Tasks
To list all tasks, use:
```bash
task-cli list
```
### 6. List Tasks by Status
To list tasks based on their status (e.g., done, in-progress, todo):
```bash
task-cli list <STATUS>
```
**Note**: The status must be one of the following: "done", "in-progress", or "todo".

### 7. Help Command
For help and more information about commands, use:
```bash
task-cli --help
```
## Conclusion
The CLI Task Manager provides a straightforward interface for managing tasks directly from the command line, making it a useful tool for productivity and organization.
For more details about the project, check out the [Task Tracker Project on Roadmap.sh](https://roadmap.sh/projects/task-tracker).

# Todo List CLI

## Project Overview
This is the third project in a **100 Projects Series to Super Master Golang**. It’s a command-line todo list application written in Golang that allows users to add, list, and delete tasks, with persistent storage in a text file (`todos.txt`). The project introduces file handling and data persistence.

### Purpose
- Learn file I/O operations for persistent data storage.
- Practice struct usage and basic data management.
- Enhance CLI development with multiple commands.

### Difficulty Level
- **Beginner**: Suitable for those with basic Golang knowledge.

---

## Tech Stack
- **Language**: Golang (using the standard library)
- **Packages**:
- `fmt`: For input/output formatting.
- `os`: For CLI arguments and file operations.
- `bufio`: For reading file contents line-by-line.
- `strings`: For string manipulation.
- **Tools**:
- Go compiler (version 1.21 or later recommended).
- Any text editor or IDE (e.g., VS Code with Go extension).

---

## Prerequisites
- **Go Installed**: Verify with `go version`.
- **Basic Terminal Knowledge**: Ability to run commands in a terminal.

---

## Setup Instructions
1. **Create the Project Directory**:
```bash
mkdir todo-cli
cd todo-cli
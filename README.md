[![Coverage Status](https://coveralls.io/repos/github/TristanShz/flow/badge.svg?branch=main)](https://coveralls.io/github/TristanShz/flow?branch=main)

<p align="center">
  <img src="assets/banner.png" />
</p>

---

Flow is a CLI tool that helps you manage your time spent developing your projects. It's designed to be simple, fast and easy to use.

It is written in GO and uses File system to store the data.

The project is in its early stages and is still under development.

## Installation

### Mac OS X

```bash
brew tap TristanShz/flow-tracker

brew install flow-tracker
```

### Linux

I'm working on a way to set up a package manager for Linux. For now, you can install it using the following command:

```bash
curl -sSf https://raw.githubusercontent.com/TristanShz/flow/main/install.sh | sudo sh
```
## Available commands

### `flow start [project] [tags]`

Starts a new flow session for the specified project.

| name | default | description                     |
| ---- | ------- | ------------------------------- |
| tags | \       | Tags to be used for the session |

example:

```bash
flow start my-project +tag1 +tag2
```

### `flow stop`

Stops the current flow session.

### `flow status`

See the status of the current flow session.

### `flow report`

View a user-friendly report of sessions.

| name              | default | description                                           |
| ----------------- | ------- | ----------------------------------------------------- |
| --format [format] | by-day  | Format of the report. Options: `by-day`, `by-project` |
| --day             | /       | Get a report for all sessions of the current day      |
| --week            | /       | Get a report for all sessions of the current week     |
| --project         | /       | Get a report for all sessions of the given project    |
| --since [date]    | /       | Get a report for all sessions since the given date    |
| --until [date]    | /       | Get a report for all sessions until the given date    |

### `flow edit [session-id (optional)]`

Open the session with given ID in the default editor. If no ID is provided, it will open the last session

### `flow abort`

Abort the current session.

## Roadmap

- [x] Start a flow session
- [x] Stop a flow session
- [x] View current session status
- [x] List all projects
- [x] View a report of all sessions
- [x] View a report of all sessions for a given project
- [x] View a report of sessions in a given time range
- [x] Edit a session
- [x] Abort a session
- [ ] Pause a session
- [ ] Resume a session
- [ ] Start session with attach mode
- [ ] Export report to CSV
- [ ] Export report to JSON

# Old School
Old-school style screensaver effect in Go and Ebitengine.

## Preview
See it live at <https://ewaldhorn.github.io/oldschool/>

![Screen Shot](./screenshot.png "Screenshot")

### Technologies

I used a variety of technologies to make this all happen:

| Tech         | Where                        | Why                                |
| ------------ | ---------------------------- | ---------------------------------- |
| Go           | <https://go.dev/>            | Development language of choice     |
| Ebitengine   | <https://ebitengine.org/>    | Great 2D engine, with WASM support |
| Task         | <https://taskfile.dev/>      | Build tool of choice               |
| GoLangCILint | <https://golangci-lint.run/> | Go linter for code cleanup         |

### Tasks

For convenience, I have a Task file. This helps save me from having to remember commands. Instead, Task does that for me!

Run `task` to see a list of available tasks. Some are:

| Task            | Action taken                                                                    |
| --------------- | ------------------------------------------------------------------------------- |
| buildproduction | Builds a production WASM binary, minus debug information                        |
| buildwasm       | Builds the WASM project                                                         |
| clean           | Removes the './bin/' folder                                                     |
| coverreport     | Generates the test coverage report                                              |
| default         | Lists available tasks                                                           |
| lint            | Runs the 'golangci-lint' tool on the source code                                |
| play            | Runs the desktop project, handy for quick local testing                         |
| run             | Runs the files in the ./bin/ folder as-is, no build steps envoked. On port 9000 |
| runwasm         | Runs the project in WASM on port 9000                                           |
| setupexecjs     | Copies the wasm_exec.js and html files to the './bin/' folder                   |
| test            | Runs the project tests                                                          |

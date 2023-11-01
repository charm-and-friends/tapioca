# Glitter

Add some dynamic [bubbbletea](https://github.com/charmbracelet/bubbletea)
elements to your existing application.

## Why should you use Glitter?

- You have a useful application with a pragmatic command line interface. It
doesn't matter if it is complex and uses
[cobra](https://github.com/spf13/cobra) or if it is simple and uses
[flag](https://pkg.go.dev/flag).

- You already have set in place some logging mechanism. Again, it doesn't
matter if you are using [log](https://pkg.go.dev/log) or
[zerolog](https://github.com/rs/zerolog).

- You want to add some dynamic elements. A simple progress bar or something way
  more complex.

Add some Glitter to your application! With all the flexibility of
[bubbbletea](https://github.com/charmbracelet/bubbletea) but without its
complexity. As simple as two lines of code:

```go
program := glitter.NewProgram(glitter.NewSpinner()).GoRun()
defer program.QuitAndWait()
```

## Pre-defined models
Glitter extends [bubbles](https://github.com/charmbracelet/bubbles) with some
pre-defined models for fast prototyping:

- TODO: Progress with title and ETA.
- TODO: Spinner with title.
- TODO: Prompt.

These models are usable out of the box and quit after pressing usual quit key
combinations (`ctrl+C`/`cmd+C`). They are intended to cover basic dynamic
elements that one would like to add to an existing application. 

However they are just the tip of the iceberg of what can be done with
Bubbletea and you might want a custom element. We also provide a `WrapModel`
function that will wrap your custom model to make it behave better with an
existing application (e.g. quitting with `ctrl+C`/`cmd+C`).
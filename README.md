# ossref

> *Every project stands on the shoulders of others.*

Track the intellectual lineage of your project. Not dependencies — **decisions**.

`ossref` is a simple CLI tool for documenting the open source projects that shaped your architecture, patterns, and design decisions. Like a bibliography for your codebase.

## Install

```bash
go install github.com/herfstvalt/ossref@latest
```

Or download a prebuilt binary from [Releases](https://github.com/herfstvalt/ossref/releases).

## Usage

```bash
ossref init       # Create .references.yml
ossref add        # Add a reference interactively
ossref list       # List all references
ossref graph      # Show dependency graph
ossref            # Show the zen message + references
```

## Example

```yaml
# .references.yml
references:
    - project: rustdesk/rustdesk
      learned: Remote desktop relay architecture — NAT traversal via rendezvous server
      applied: services/relay/

    - project: opencode-ai/opencode
      learned: TUI page-based navigation with tea.Model dispatch
      applied: cmd/tui.go
```

```
$ ossref graph

  myproject
  ├── rustdesk
  │   └── Remote desktop relay architecture — NAT traversal via rendezvous server
  └── opencode
      └── TUI page-based navigation with tea.Model dispatch
```

## The Zen of Open Source Referencing

```
Every project stands on the shoulders of others.
Attribution is not obligation — it is memory.
A copied pattern without a reference is a severed thread.
AI can generate code. It cannot generate gratitude.
Your dependency tree shows what you use.
Your references show what you learned.
In the age of infinite generation,
provenance is the only thing that's scarce.
```

## License

MIT

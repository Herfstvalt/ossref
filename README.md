# ossref

> *Every project stands on the shoulders of others.*

Track the intellectual lineage of your project. Not dependencies — **decisions**.

`ossref` is a simple CLI tool for documenting the open source projects that shaped your architecture, patterns, and design decisions. Like a bibliography for your codebase.

## Install

Pick whichever method matches your workflow — no Go required for any of these:

**Quick install (macOS/Linux):**
```bash
curl -sSL https://raw.githubusercontent.com/herfstvalt/ossref/main/install.sh | sh
```

**Homebrew (macOS/Linux):**
```bash
brew install herfstvalt/tap/ossref
```

**npm (JavaScript/TypeScript developers):**
```bash
npx ossref
# or install globally
npm install -g ossref
```

**pip (Python developers):**
```bash
pip install ossref
```

**Download binary:**

Grab the latest release for your platform from [Releases](https://github.com/herfstvalt/ossref/releases).

**From source (requires Go):**
```bash
go install github.com/herfstvalt/ossref@latest
```

## How it works

You maintain a `.references.yml` file in your project root that documents where your architectural decisions came from. Not your imports or dependencies — the projects you *studied*, the patterns you *adapted*, the ideas you *learned from*.

Each reference has three fields:
- **project** — the open source project you learned from
- **learned** — what you took away (architecture, pattern, approach)
- **applied** — where you applied it in your codebase

## Getting started

**1. Initialize in your project:**
```bash
cd your-project
ossref init
```
This creates an empty `.references.yml`.

**2. Add your first reference:**
```bash
ossref add
```
```
  project: rustdesk/rustdesk
  learned: Remote desktop relay architecture — NAT traversal via rendezvous server
  applied: services/relay/

  ✓ Reference added. 1 references across 1 projects.
```

**3. Keep adding as you build:**
```bash
ossref add
```
```
  project: opencode-ai/opencode
  learned: TUI page-based navigation with tea.Model dispatch
  applied: cmd/tui.go

  ✓ Reference added. 2 references across 2 projects.
```

**4. View your references:**
```bash
ossref
```
```

  The Zen of Open Source Referencing

  Every project stands on the shoulders of others.
  Attribution is not obligation — it is memory.
  A copied pattern without a reference is a severed thread.
  AI can generate code. It cannot generate gratitude.
  Your dependency tree shows what you use.
  Your references show what you learned.
  In the age of infinite generation,
  provenance is the only thing that's scarce.

  ─────────────────────────────────────────────

  This project stands on the shoulders of:

    rustdesk → Remote desktop relay architecture — NAT traversal via rendezvous server
    opencode → TUI page-based navigation with tea.Model dispatch

  2 references across 2 projects
  .references.yml

```

**5. See the dependency graph:**
```bash
ossref graph
```
```
  myproject
  ├── rustdesk
  │   └── Remote desktop relay architecture — NAT traversal via rendezvous server
  └── opencode
      └── TUI page-based navigation with tea.Model dispatch
```

**6. Generate a visual graph:**
```bash
ossref graph --svg
```
Creates a `references.svg` file you can embed in your README or docs.

## Commands

```
ossref              Show zen message and project references
ossref init         Create a new .references.yml
ossref add          Add a reference interactively
ossref list         List all references
ossref graph        Show reference dependency graph (terminal)
ossref graph --svg  Generate references.svg visual graph
ossref help         Show this help
```

## Example .references.yml

```yaml
references:
    - project: rustdesk/rustdesk
      learned: Remote desktop relay architecture — NAT traversal via rendezvous server
      applied: services/relay/

    - project: opencode-ai/opencode
      learned: TUI page-based navigation with tea.Model dispatch
      applied: cmd/tui.go

    - project: openai/codex
      learned: Git worktree isolation for parallel agent execution
      applied: internal/runtime/

    - project: nicbarker/clay
      learned: Session list UX — grouped by recency, status dots, minimal chrome
      applied: apps/mobile/lib/features/session/
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

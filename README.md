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
  project: python/cpython
  learned: import this — embedding a philosophy statement directly in the tool
  applied: zen message

  ✓ Reference added. 1 references across 1 projects.
  ✓ REFERENCES.md updated.
```

**3. Keep adding as you build:**
```bash
ossref add
```
```
  project: citation-file-format/citation-file-format
  learned: CITATION.cff schema design — YAML-based file convention for software citation metadata
  applied: .references.yml schema

  ✓ Reference added. 2 references across 2 projects.
  ✓ REFERENCES.md updated.
```

Every `ossref add` auto-generates a `REFERENCES.md` that renders on GitHub — no README needed.

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

    citation-file-format → CITATION.cff schema design — YAML-based file convention for software citation metadata
    cpython              → import this — embedding a philosophy statement directly in the tool
    go                   → Go Proverbs — concise, opinionated design principles as community culture
    spdx-spec            → SBOM standardization — structured machine-readable software metadata as a file convention
    goreleaser           → Cross-platform binary release pipeline with Homebrew tap auto-publishing
    esbuild              → npm wrapper pattern — postinstall script that downloads a platform-specific Go binary

  6 references across 6 projects
  .references.yml

```

**5. See the dependency graph:**
```bash
ossref graph
```
```
  Herfstvalt/ossref
  ├── citation-file-format
  │   └── CITATION.cff schema design — YAML-based file convention for software citation metadata
  ├── cpython
  │   └── import this — embedding a philosophy statement directly in the tool
  ├── go
  │   └── Go Proverbs — concise, opinionated design principles as community culture
  ├── spdx-spec
  │   └── SBOM standardization — structured machine-readable software metadata as a file convention
  ├── goreleaser
  │   └── Cross-platform binary release pipeline with Homebrew tap auto-publishing
  └── esbuild
      └── npm wrapper pattern — postinstall script that downloads a platform-specific Go binary
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

This is [ossref's own `.references.yml`](.references.yml) — we use the tool on itself:

```yaml
references:
    - project: citation-file-format/citation-file-format
      learned: CITATION.cff schema design — YAML-based file convention for software citation metadata
      applied: .references.yml schema

    - project: python/cpython
      learned: import this — embedding a philosophy statement directly in the tool
      applied: zen message

    - project: golang/go
      learned: Go Proverbs — concise, opinionated design principles as community culture
      applied: zen message tone and structure

    - project: spdx/spdx-spec
      learned: SBOM standardization — structured machine-readable software metadata as a file convention
      applied: .references.yml as a committed file convention

    - project: goreleaser/goreleaser
      learned: Cross-platform binary release pipeline with Homebrew tap auto-publishing
      applied: .goreleaser.yml, release workflow

    - project: esbuild/esbuild
      learned: npm wrapper pattern — postinstall script that downloads a platform-specific Go binary
      applied: npm/install.js
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

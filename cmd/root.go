package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/herfstvalt/ossref/internal/graph"
	"github.com/herfstvalt/ossref/internal/refs"
	"github.com/herfstvalt/ossref/internal/zen"
)

var scanner = bufio.NewScanner(os.Stdin)

func Execute() error {
	if len(os.Args) < 2 {
		return runDefault()
	}

	switch os.Args[1] {
	case "init":
		return runInit()
	case "add":
		return runAdd()
	case "list":
		return runList()
	case "graph":
		return runGraph()
	case "render":
		return runRender()
	case "help", "--help", "-h":
		return runHelp()
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", os.Args[1])
		return runHelp()
	}
}

func runDefault() error {
	fmt.Println(zen.Message)

	cwd, err := os.Getwd()
	if err != nil {
		return nil
	}

	path, err := refs.Find(cwd)
	if err != nil {
		fmt.Println("  No .references.yml found. Run `ossref init` to start.")
		fmt.Println()
		return nil
	}

	f, err := refs.Load(path)
	if err != nil {
		return nil
	}

	if len(f.References) == 0 {
		fmt.Println("  No references yet. Run `ossref add` to add your first.")
		fmt.Println()
		return nil
	}

	fmt.Println("  ─────────────────────────────────────────────")
	fmt.Println()
	fmt.Println("  This project stands on the shoulders of:")
	fmt.Println()

	projects := map[string]bool{}
	maxName := 0
	for _, r := range f.References {
		name := shortName(r.Project)
		projects[r.Project] = true
		if len(name) > maxName {
			maxName = len(name)
		}
	}

	for _, r := range f.References {
		name := shortName(r.Project)
		padding := strings.Repeat(" ", maxName-len(name))
		fmt.Printf("    %s%s → %s\n", name, padding, r.Learned)
	}

	fmt.Println()
	fmt.Printf("  %d references across %d projects\n", len(f.References), len(projects))
	fmt.Println("  .references.yml")
	fmt.Println()

	return nil
}

func runInit() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path, err := refs.Init(cwd)
	if err != nil {
		return err
	}

	fmt.Printf("Created %s\n", path)
	return nil
}

func runAdd() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path, err := refs.Find(cwd)
	if err != nil {
		return fmt.Errorf("no .references.yml found — run `ossref init` first")
	}

	fmt.Print("  project: ")
	project := prompt()
	if project == "" {
		return fmt.Errorf("project is required")
	}

	fmt.Print("  learned: ")
	learned := prompt()
	if learned == "" {
		return fmt.Errorf("learned is required")
	}

	fmt.Print("  applied: ")
	applied := prompt()

	ref := refs.Reference{
		Project: project,
		Learned: learned,
		Applied: applied,
	}

	projectName := detectProjectName(cwd)

	if err := refs.Add(path, ref, projectName); err != nil {
		return err
	}

	f, err := refs.Load(path)
	if err != nil {
		return err
	}

	projects := map[string]bool{}
	for _, r := range f.References {
		projects[r.Project] = true
	}

	fmt.Printf("\n  ✓ Reference added. %d references across %d projects.\n", len(f.References), len(projects))
	fmt.Println("  ✓ REFERENCES.md updated.")
	fmt.Println()
	return nil
}

func runList() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path, err := refs.Find(cwd)
	if err != nil {
		return fmt.Errorf("no .references.yml found — run `ossref init` first")
	}

	f, err := refs.Load(path)
	if err != nil {
		return err
	}

	if len(f.References) == 0 {
		fmt.Println("  No references yet. Run `ossref add` to add your first.")
		return nil
	}

	for i, r := range f.References {
		fmt.Printf("  %d. %s\n", i+1, r.Project)
		fmt.Printf("     learned: %s\n", r.Learned)
		if r.Applied != "" {
			fmt.Printf("     applied: %s\n", r.Applied)
		}
		fmt.Println()
	}

	return nil
}

func runGraph() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path, err := refs.Find(cwd)
	if err != nil {
		return fmt.Errorf("no .references.yml found — run `ossref init` first")
	}

	f, err := refs.Load(path)
	if err != nil {
		return err
	}

	if len(f.References) == 0 {
		fmt.Println("  No references to graph.")
		return nil
	}

	projectName := detectProjectName(cwd)

	// Check for --svg flag
	svgOutput := false
	outputFile := "references.svg"
	for i, arg := range os.Args {
		if arg == "--svg" {
			svgOutput = true
			if i+1 < len(os.Args) && !strings.HasPrefix(os.Args[i+1], "-") {
				outputFile = os.Args[i+1]
			}
		}
	}

	if svgOutput {
		svg := graph.RenderSVG(projectName, f)
		if err := os.WriteFile(outputFile, []byte(svg), 0644); err != nil {
			return err
		}
		fmt.Printf("  ✓ Graph written to %s\n", outputFile)
		return nil
	}

	// Terminal tree output
	fmt.Println()
	fmt.Printf("  %s\n", projectName)

	grouped := map[string][]refs.Reference{}
	order := []string{}
	for _, r := range f.References {
		if _, exists := grouped[r.Project]; !exists {
			order = append(order, r.Project)
		}
		grouped[r.Project] = append(grouped[r.Project], r)
	}

	for i, project := range order {
		entries := grouped[project]
		isLast := i == len(order)-1

		if isLast {
			fmt.Printf("  └── %s\n", shortName(project))
		} else {
			fmt.Printf("  ├── %s\n", shortName(project))
		}

		for j, e := range entries {
			isLastEntry := j == len(entries)-1
			var prefix string
			if isLast {
				prefix = "      "
			} else {
				prefix = "  │   "
			}

			if isLastEntry {
				fmt.Printf("%s└── %s\n", prefix, e.Learned)
			} else {
				fmt.Printf("%s├── %s\n", prefix, e.Learned)
			}
		}
	}

	fmt.Println()
	return nil
}

func runRender() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	path, err := refs.Find(cwd)
	if err != nil {
		return fmt.Errorf("no .references.yml found — run `ossref init` first")
	}

	f, err := refs.Load(path)
	if err != nil {
		return err
	}

	projectName := detectProjectName(cwd)
	if err := refs.WriteReferencesMarkdown(path, projectName, f); err != nil {
		return err
	}

	fmt.Println("  ✓ REFERENCES.md updated.")
	return nil
}

func runHelp() error {
	fmt.Println(`ossref — The Zen of Open Source Referencing

Usage:
  ossref              Show zen message and project references
  ossref init         Create a new .references.yml
  ossref add          Add a reference interactively
  ossref list         List all references
  ossref graph        Show reference dependency graph (terminal)
  ossref graph --svg  Generate references.svg visual graph
  ossref render       Regenerate REFERENCES.md
  ossref help         Show this help`)
	return nil
}

func shortName(project string) string {
	parts := strings.Split(project, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-1]
	}
	return project
}

func prompt() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func detectProjectName(dir string) string {
	// Try git remote first
	if f, err := os.ReadFile(dir + "/.git/config"); err == nil {
		lines := strings.Split(string(f), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "url = ") {
				url := strings.TrimPrefix(line, "url = ")
				url = strings.TrimSuffix(url, ".git")
				parts := strings.Split(url, "/")
				if len(parts) >= 2 {
					return parts[len(parts)-2] + "/" + parts[len(parts)-1]
				}
			}
		}
	}

	// Fallback to directory name
	parts := strings.Split(dir, "/")
	return parts[len(parts)-1]
}

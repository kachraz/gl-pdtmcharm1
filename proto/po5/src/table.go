/*
table.go
This file has the actual table code
*/

package src

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	C "github.com/fatih/color"
)

func TableFuncMain() {
	TableMain()
}

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("200")) // This controls the overall border color

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Color Definitions
	cr := C.New(C.FgRed).SprintFunc()
	chg := C.New(C.FgHiGreen).SprintFunc()
	cm := C.New(C.FgMagenta).SprintFunc()

	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			selectedItem := m.table.SelectedRow()[1] // This controls which item from the table is being selected starts with 0
			cmd := exec.Command("pdtm", "-install", selectedItem)

			output, err := cmd.CombinedOutput()

			// err := cmd.Run()
			if err != nil {
				cr := C.New(C.FgRed).SprintFunc() // Red color for warnings
				fmt.Println(cr("[FAIL] ", err))

				fmt.Println("Output: ", string(output))
				return m, tea.Quit
			}
			fmt.Println(cm("[SELECTED]", selectedItem))
			fmt.Println(chg(string(output)))
			fmt.Println(cr("Use pdtm -r <name> to remove"))
			return m, tea.Quit

		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

// ---------------------------- Your Edits Start --------------------------
// ---------------------------- Your Edits Start --------------------------
// ---------------------------- Your Edits Start --------------------------
/*
This is ou addition of the help functions taken from gla ,
adding the navigation help at the end of the table
*/
var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("200")).Render

func (m model) helpView() string {
	return helpStyle("\n  ↑/↓: Navigate • q: Quit • Enter: to Select Install\n")
}

// ---------------------------- Your Edits ends --------------------------
// ---------------------------- Your Edits ends --------------------------
// ---------------------------- Your Edits ends --------------------------

// This controls the views of the application
// helpView() added here
func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n" + m.helpView() + "\n"
}

// This is the main program and entry poin , this will go up in V1main()
func TableMain() {
	columns := []table.Column{
		{Title: "#", Width: 3},
		{Title: "Tool", Width: 20},
		{Title: "What?", Width: 40},
	}

	rows := []table.Row{
		{"1", "aix", "Interact LLM APIs"},
		{"2", "alterx", "Subdomain wordlist gen using DSL."},
		{"3", "asnmap", "Go Cli/Lib map network ranges ASN"},
		{"4", "cdncheck", "ID tech with dns / ip."},
		{"5", "chaos-client", "Go client comms Chaos dataset API."},
		{"6", "cloudlist", "Multi-cloud Assets from Cloud"},
		{"7", "dnsx", "DNS Toolkit"},
		{"8", "httpx", "HTTP Toolkit"},
		{"9", "interactsh-client", "Gathering Server"},
		{"10", "interactsh-server", "Gathering Client"},
		{"11", "katana", "Crawl/Spider Framework"},
		{"12", "mapcidr", "Go Cli/Lib Ops Management CIDR"},
		{"13", "naabu", "Go Por Scanner"},
		{"14", "notify", "Output Streamer"},
		{"15", "nuclei", "YAML/DSL Vuln Scanner"},
		{"16", "pdtm", "Project Discovery Tool Manager"},
		{"17", "proxify", "Swiss Army Knife Proxy"},
		{"18", "shuffledns", "Go massdns wrapper"},
		{"19", "simplehttpserver", "Go Alernative to python httpserver"},
		{"20", "subfinder", "Subdomain Discovery"},
		{"21", "tlsx", "Tls Grabber"},
		{"22", "uncover", "Exposed host disvoery via search"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10), // This controls number of rows
	)

	/*
		This section is the style it follows the ANSI 256 Color Codes
		https://www.ditig.com/256-colors-cheat-sheet
	*/
	s := table.DefaultStyles()
	s.Header = s.Header. // Affects header only
				BorderStyle(lipgloss.NormalBorder()).
				Foreground(lipgloss.Color("229")).
				BorderForeground(lipgloss.Color("82")).
				BorderBottom(true).
				Bold(false)
	s.Selected = s.Selected. // Affects Selections only
					Foreground(lipgloss.Color("229")).
					Background(lipgloss.Color("22")).
					Bold(false).
					Italic(true) // Added Italics

	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
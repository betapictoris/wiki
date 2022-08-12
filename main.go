package main

// Import pkgs
import (
	"log"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/knipferrc/teacup/statusbar"
	"github.com/charmbracelet/glamour"
	
	md "github.com/JohannesKaufmann/html-to-markdown"
)

const lang   = "en" // Lang prefix used on
const apiURL = "https://" + lang + ".wikipedia.org/api/rest_v1/" // Wikipedia API URL
const useHighPerformanceRenderer = false

// Bubble represents the properties of the UI.
type Bubble struct {
	statusbar   statusbar.Bubble
	height      int
	content     string
	title 	    string
	articleName string
}

// Init intializes the UI.
func (b Bubble) Init() tea.Cmd {
	return nil
}

// New creates a new instance of the UI.
func New() Bubble {
	sb := statusbar.New(
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Dark: "#ffffff", Light: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#F25D94", Dark: "#F25D94"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#A550DF", Dark: "#A550DF"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#6124DF", Dark: "#6124DF"},
		},
	)

	return Bubble{
		statusbar: sb,
	}
}

// Update handles all UI interactions.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.height = msg.Height
		b.statusbar.SetSize(msg.Width)
		b.statusbar.SetContent(b.title, b.articleName, "", "")

		return b, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	return b, tea.Batch(cmds...)
}

// View returns a string representation of the UI.
func (b Bubble) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		lipgloss.NewStyle().Height(b.height-statusbar.Height).Render(b.content),
		b.statusbar.View(),
	)
}

func main() {
	article := ""
	saveToFile := false

	if len(os.Args) >= 2 {
		article = os.Args[1]

		if len(os.Args) >= 3 {
			if os.Args[2] == "-s" {
				saveToFile = true
			}
		}
	} else {
		log.Fatal("Usage: wiki <Article> [-s (Save to file)]")
	}

	// Read remote URL's conts
	resp, err := http.Get(apiURL + "page/html/" + article)
	cont, err := ioutil.ReadAll(resp.Body)
	converter := md.NewConverter("", true, nil)
	content, err := converter.ConvertString(strings.Replace(string(cont), "//upload.wikimedia.org", "https://upload.wikimedia.org", -1))
	
	out, err := glamour.Render(content, "dark")
	
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if saveToFile {
		if err := ioutil.WriteFile(article + ".md", []byte(content), 0666); err != nil {
        		log.Fatal(err)
		}
	} else {
		p := tea.NewProgram(
			Bubble{content: string(out), title: "Wiki CLI", articleName: strings.Replace(article, "_", " ", -1)},
			tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
			tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
		)
	
		if err := p.Start(); err != nil {
			fmt.Println("could not run program:", err)
			os.Exit(1)
		}
	}
}

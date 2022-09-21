package main

// Import pkgs
import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/glamour"
	"github.com/knipferrc/teacup/statusbar"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

const (
	lang                       = "en"                                              // Lang prefix used on
	apiURL                     = "https://" + lang + ".wikipedia.org/api/rest_v1/" // Wikipedia API URL
	useHighPerformanceRenderer = false
)

// Bubble represents the properties of the UI.
type Bubble struct {
	statusbar   statusbar.Bubble
	viewport    viewport.Model
	height      int
	content     string
	title       string
	articleName string
	ready       bool
}

// Init intializes the UI.
func (Bubble) Init() tea.Cmd {
	return nil
}

// New creates a new instance of the UI.
func NewStatusbar() statusbar.Bubble {
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
			Background: lipgloss.AdaptiveColor{Light: "#3c3836", Dark: "#3c3836"},
		},
		statusbar.ColorConfig{
			Foreground: lipgloss.AdaptiveColor{Light: "#ffffff", Dark: "#ffffff"},
			Background: lipgloss.AdaptiveColor{Light: "#6124DF", Dark: "#6124DF"},
		},
	)

	return sb
}

// Update handles all UI interactions.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.height = msg.Height

		footerHeight := lipgloss.Height(b.footerView())
		verticalMarginHeight := footerHeight

		if !b.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			b.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			b.viewport.YPosition = 0
			b.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			b.viewport.SetContent(b.content)
			b.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			b.viewport.YPosition = 1
		} else {
			b.viewport.Width = msg.Width
			b.viewport.Height = msg.Height - verticalMarginHeight
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			cmds = append(cmds, tea.Quit)
		}
	}

	// Handle keyboard and mouse events in the viewport
	b.viewport, cmd = b.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}

// View returns a string representation of the UI.
func (b Bubble) View() string {
	if !b.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s", b.viewport.View(), b.footerView())
}

func (b Bubble) footerView() string {
	b.statusbar.SetSize(b.viewport.Width)
	b.statusbar.SetContent(b.title, b.articleName, "", fmt.Sprintf("%3.f%%", b.viewport.ScrollPercent()*100))
	return b.statusbar.View()
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
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	cont, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	converter := md.NewConverter("", true, nil)
	content, err := converter.ConvertString(strings.ReplaceAll(string(cont), "//upload.wikimedia.org", "https://upload.wikimedia.org"))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	out, err := glamour.Render(content, "dark")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if saveToFile {
		f, err := os.OpenFile(article+".md", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			f.Close()
			log.Fatal(err)
		}

		_, err = f.Write([]byte(content))
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	} else {
		p := tea.NewProgram(
			Bubble{statusbar: NewStatusbar(), content: string(out), title: "Wiki CLI", articleName: strings.ReplaceAll(article, "_", " ")},
			tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
			tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
		)

		if err := p.Start(); err != nil {
			fmt.Println("could not run program:", err)
			os.Exit(1)
		}
	}
}

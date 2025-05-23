package main

import (
	_ "embed"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
)

func htmlHandler(s *rweb.Server) {
	s.Get("/", func(ctx rweb.Context) error {
		return ctx.WriteHTML(renderer(htmlPage{}))
	})

}

func renderer(comps ...element.Component) string {
	b := element.NewBuilder()
	element.RenderComponents(b, comps...)
	return b.String()
}

type htmlPage struct{}

//go:embed assets/style.css
var styles string

//go:embed assets/script.js
var javascriptFile string

func (h htmlPage) Render(b *element.Builder) (x any) {
	b.Html().R(
		b.Head().R(
			b.Meta("charset", "UTF-8").R(),
			b.Meta("name", "viewport", "content", "width=device-width, initial-scale=1.0").R(),
			b.Title().T("Go Code Executor"),
			b.Link("rel", "stylesheet", "href", "https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.52.2/min/vs/editor/editor.main.css").R(),
			b.Style().T(styles),
			// A script tag should use the "src" attribute to load external JavaScript files or
			// include inline JavaScript code as a child element, but not both at least, that's the convention I know.
			// Putting the inline JS at the end of the body to make sure all the page elements are in place before this JS runs
			b.Script("src", "https://cdnjs.cloudflare.com/ajax/libs/monaco-editor/0.52.2/min/vs/loader.min.js").R(),
		),
		b.Body().R(

			b.Div("class", "app-container").R(
				b.Header().R(
					b.H1().T("Go Code Executor"),
					b.P().T("Welcome to the Go Code Executor! This is a web-based environment for writing and executing Go code."),
				),
				b.Main().R(
					b.Div("class", "editor-container", "style", "border:2px solid maroon").R(
						b.Div("id", "editor").R(),
						b.Div("class", "button-container").R(
							b.Button("id", "format-button").T("Format"),
							b.Button("id", "run-button").T("Run (ctrl+Enter)"),
						),
						b.Div("class", "output-container").R(
							b.Div("class", "output-header").R(
								b.H2().T("Execution Results"),
								b.Div("id", "execution-status").T("Ready"),
							),
							b.Div("class", "output-content").R(
								b.Div("class", "output-section").R(
									b.H3().T("Standard Output"),
									b.Pre("id", "stdout-output", "class", "output-area").R(),
								),
								b.Div("class", "output-section").R(
									b.H3().T("Standard Error"),
									// (Multiple classes are separated by spaces)
									b.Pre("id", "stderr-output", "class", "output-area error").R(),
								),
								b.Div("class", "execution-info").R(
									b.Div("id", "execution-time").R(),
									b.Div("id", "execution-result").R(),
								),
							),
						),
					),
				),
				b.Footer().R(
					b.P().T("Go Code Executor - A web-based Go code execution environment"),
				),
			),
			// Putting the inline JS at the end of the body to make sure all the page elements are in place before this JS runs
			b.Script().T(javascriptFile),
		),
	)

	return
}

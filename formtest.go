package main

import (
    "fmt"
    "net/http"
	"forms" // My own library!
)

func handler(w http.ResponseWriter, r *http.Request) {
    p := &Page{w, r}
    w.Header().Set("Content-Type", "text/html")
    if r.Method == "POST" {
        p.POST()
    } else {
        p.GET()
    }
}

type Page struct {
	response http.ResponseWriter
	request *http.Request
}

func (p *Page) getForm() *forms.Form {
	frm := new(forms.Form)
	email_valid := func (inp *forms.Input, r *http.Request) (bool) {
		inp.Errors = append(inp.Errors, "Invalid email")
		return false
	}
	email := &forms.Textbox{forms.Input{Name:"email", Label:"Email",
		Validators: []func (*forms.Input, *http.Request) (bool){email_valid}}}
	frm.AddInput(email)
	frm.AddInput(&forms.Password{forms.Input{Name:"password", Label:"Password"}})
	frm.AddInput(&forms.Textarea{forms.Input{Name:"message", Label:"Message"}})
	frm.AddInput(&forms.Dropdown{Input: forms.Input{Name:"gender", Label:"Gender"},
		Options:[]string{"Female", "Male"}})
	frm.AddInput(&forms.Radio{Input: forms.Input{Name:"gender2", Label:"Gender"},
		Options:[]string{"Female", "Male"}})
	frm.AddInput(&forms.Checkbox{forms.Input{Name:"optin", Label:"Send updates?"}})
	frm.AddInput(&forms.Hidden{forms.Input{Name:"price"}})
	frm.AddInput(&forms.File{forms.Input{Name:"icon", Label:"Image upload"}})
	frm.AddInput(&forms.Button{forms.Input{Name:"send", Label:"Sign Up"}})
	return frm
}

func (p *Page) GET() {
	frm := p.getForm()
	html := fmt.Sprintf(`<form action="" method="post">%s</form>`, frm.Render())
	fmt.Fprintf(p.response, html)
}

func (p *Page) POST() {
	_ = p.request.ParseForm() // XXX handle error
	frm := p.getForm()
	if !frm.Validate(p.request) {
		html := fmt.Sprintf(`<form action="" method="post">%s</form>`, frm.Render())
		fmt.Fprintf(p.response, html)
		return
	}
    fmt.Fprintf(p.response, "<h1>Bingo!</h1><dl>")
	for k, v := range p.request.Form {
		fmt.Fprintf(p.response, "<dt>%s</dt><dd>%s</dd>", k, v)
	}
    fmt.Fprintf(p.response, "</dl>")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

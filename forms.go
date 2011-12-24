package main

import (
	"fmt"
	"net/http"
)

/* Inputs */

type InputInterface interface {
	Validate(*http.Request) (bool)
	Render() (string)
	GetLabel() (string)
}

// Text Input

type Textbox struct {
	Name string
	Label string
	Errors []string
	// Validators list
}

func (inp *Textbox) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Textbox) Render() (string) {
	return fmt.Sprintf(`<input type="text" name="%s"/>`, inp.Name)
}

func (inp *Textbox) GetLabel() (string) {
	return inp.Label
}

// Password Input

type Password struct {
	Name string
	Label string
	Errors []string
	// Validators list
}

func (inp *Password) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Password) Render() (string) {
	return fmt.Sprintf(`<input type="password" name="%s"/>`, inp.Name)
}

func (inp *Password) GetLabel() (string) {
	return inp.Label
}


// Textarea Input

type Textarea struct {
	Name string
	Label string
	Errors []string
	// Validators list
}

func (inp *Textarea) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Textarea) Render() (string) {
	return fmt.Sprintf(`<textarea name="%s"></textarea>`, inp.Name)
}

func (inp *Textarea) GetLabel() (string) {
	return inp.Label
}


// Dropdown

type Dropdown struct {
	Name string
	Label string
	Errors []string
	Options []string
	// Validators list
}

func (inp *Dropdown) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Dropdown) Render() (string) {
	out := fmt.Sprintf(`<select name="%s">`, inp.Name)
	for _, opt := range inp.Options {
		out += fmt.Sprintf(`<option>%s</option>`, opt)
	}
	out += "</select>"
	return out
}

func (inp *Dropdown) GetLabel() (string) {
	return inp.Label
}


// Radios

type Radio struct {
	Name string
	Label string
	Errors []string
	Options []string
	// Validators list
}

func (inp *Radio) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Radio) Render() (string) {
	out := ""
	for _, opt := range inp.Options {
		out += fmt.Sprintf(`<label><input name="%s" type="radio" value="%s"/> %s</label>`,
			inp.Name, opt, opt)
	}
	return out
}

func (inp *Radio) GetLabel() (string) {
	return inp.Label
}


// Checkbox

type Checkbox struct {
	Name string
	Label string
	Errors []string
	// Validators list
}

func (inp *Checkbox) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Checkbox) Render() (string) {
	return fmt.Sprintf(`<input type="checkbox" name="%s"/>`, inp.Name)
}

func (inp *Checkbox) GetLabel() (string) {
	return inp.Label
}


// Button

type Button struct {
	Name string
	Label string
	// Validators list
}

func (inp *Button) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Button) Render() (string) {
	return fmt.Sprintf(`<button name="%s" type="submit">%s</button>`,
		inp.Name, inp.Label)
}

func (inp *Button) GetLabel() (string) {
	return ""
}


// Hidden

type Hidden struct {
	Name string
	Errors []string
	// Validators list
}

func (inp *Hidden) Validate(r *http.Request) (bool) {
	return true
}

func (inp *Hidden) Render() (string) {
	return fmt.Sprintf(`<input type="hidden" name="%s"/>`, inp.Name)
}

func (inp *Hidden) GetLabel() (string) {
	return ""
}


// File 

type File struct {
	Name string
	Label string
	Errors []string
	// Validators list
}

func (inp *File) Validate(r *http.Request) (bool) {
	return true
}

func (inp *File) Render() (string) {
	return fmt.Sprintf(`<input type="file" name="%s"/>`, inp.Name)
}

func (inp *File) GetLabel() (string) {
	return inp.Label
}


/* Form */

type Form struct {
	Inputs []InputInterface
	Validators []func() (bool)
}

func (frm *Form) Render() (string) {
	out := "<table>"
	for _, inp := range frm.Inputs {
		out += fmt.Sprintf(`
		<tr>
			<th>%s</th>
			<td>%s</td>
		</tr>`, inp.GetLabel(), inp.Render())
	}
	out += "\n</table>"
	return out
}

func (frm *Form) AddInput(inp InputInterface) {
	frm.Inputs = append(frm.Inputs, inp)
}

func main() {
	frm := new(Form)
	frm.AddInput(&Textbox{Name:"email", Label:"Email"})
	frm.AddInput(&Password{Name:"password", Label:"Password"})
	frm.AddInput(&Textarea{Name:"message", Label:"Message"})
	frm.AddInput(&Dropdown{Name:"gender", Label:"Gender",
		Options:[]string{"Female", "Male"}})
	frm.AddInput(&Radio{Name:"gender2", Label:"Gender",
		Options:[]string{"Female", "Male"}})
	frm.AddInput(&Checkbox{Name:"optin", Label:"Send updates?"})
	frm.AddInput(&Hidden{Name:"price"})
	frm.AddInput(&File{Name:"icon", Label:"Image upload"})
	frm.AddInput(&Button{Name:"send", Label:"Sign Up"})
	fmt.Println(frm.Render())
}

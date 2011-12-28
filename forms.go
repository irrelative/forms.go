package forms

import (
	"fmt"
	"net/http"
)

/* Inputs */

type InputInterface interface {
	Render() (string) // Render the input field
	GetLabel() (string) // Render the label
	GetValidators() ([]func (*http.Request) (bool))
}

type Input struct {
	Name string
	Label string
	Errors []string
	Validators []func (*http.Request) (bool)
}

func (inp *Input) Render() (string) {
	return "Not implemented"
}

func (inp *Input) GetLabel() (string) {
	return fmt.Sprintf(`<label for="id_%s">%s</label>`, inp.Name, inp.Label)
}

func (inp *Input) GetValidators() ([]func (*http.Request) (bool)) {
	return inp.Validators
}

// Text Input

type Textbox struct { Input }

func (inp *Textbox) Render() (string) {
	return fmt.Sprintf(`<input type="text" name="%s" id="id_%s"/>`,
		inp.Name, inp.Name)
}

// Password Input

type Password struct { Input }

func (inp *Password) Render() (string) {
	return fmt.Sprintf(`<input type="password" name="%s" id="id_%s"/>`, inp.Name, inp.Name)
}


// Textarea Input

type Textarea struct { Input }

func (inp *Textarea) Render() (string) {
	return fmt.Sprintf(`<textarea name="%s" id="id_%s"></textarea>`, inp.Name, inp.Name)
}


// Dropdown

type Dropdown struct {
	Input
	Options []string
}

func (inp *Dropdown) Render() (string) {
	out := fmt.Sprintf(`<select name="%s" id="id_%s">`, inp.Name, inp.Name)
	for _, opt := range inp.Options {
		out += fmt.Sprintf(`<option>%s</option>`, opt)
	}
	out += "</select>"
	return out
}


// Radios

type Radio struct {
	Input
	Options []string
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
	Input
}

func (inp *Checkbox) Render() (string) {
	return fmt.Sprintf(`<input type="checkbox" name="%s" id="id_%s"/>`,
		inp.Name, inp.Name)
}


// Button

type Button struct {
	Input
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
	Input
}

func (inp *Hidden) Render() (string) {
	return fmt.Sprintf(`<input type="hidden" name="%s"/>`, inp.Name)
}

func (inp *Hidden) GetLabel() (string) {
	return ""
}


// File 

type File struct {
	Input
}

func (inp *File) Render() (string) {
	return fmt.Sprintf(`<input type="file" name="%s" id="id_%s"/>`,
		inp.Name, inp.Name)
}


/* Form */

type Form struct {
	Inputs []InputInterface
	Validators []func(*http.Request) (bool)
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

func (frm *Form) Validate(r *http.Request) (bool) {
	valid := true
	for _, inp := range frm.Inputs {
		for _, validator := range inp.GetValidators() {
			valid = validator(r) && valid
		}
	}
	return valid
}

func main() {
	frm := new(Form)
	frm.AddInput(&Textbox{Input{Name:"email", Label:"Email"}})
	frm.AddInput(&Password{Input{Name:"password", Label:"Password"}})
	frm.AddInput(&Textarea{Input{Name:"message", Label:"Message"}})
	frm.AddInput(&Dropdown{Input:Input{Name:"gender", Label:"Gender"},
		Options:[]string{"Female", "Male"}})
	frm.AddInput(&Radio{Input:Input{Name:"gender2", Label:"Gender"},
		Options:[]string{"Female", "Male"}})
	frm.AddInput(&Checkbox{Input{Name:"optin", Label:"Send updates?"}})
	frm.AddInput(&Hidden{Input{Name:"price"}})
	frm.AddInput(&File{Input{Name:"icon", Label:"Image upload"}})
	frm.AddInput(&Button{Input{Name:"send", Label:"Sign Up"}})
	fmt.Println(frm.Render())
}

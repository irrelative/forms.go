package forms

import (
	"fmt"
	"regexp"
	"net/http"
)

/* Inputs */

type InputInterface interface {
	Render() (string) // Render the input field
	RenderErrors() (string) // Render the errors
	GetLabel() (string) // Render the label
	Validate(*http.Request) (bool) // Validate the input
}

type Input struct {
	Name string
	Label string
	Errors []string
	Validators []func (*Input, *http.Request) (bool)
}

func (inp *Input) Render() (string) {
	return "Not implemented"
}

func (inp *Input) RenderErrors() (string) {
	out := ""
	if len(inp.Errors) != 0 {
		out += "<ul class=\"error_list\">"
		for _, errstr := range inp.Errors {
			out += fmt.Sprintf(`<li>%s</li>`, errstr)
		}
		out += "</ul>"
	}
	return out
}

func (inp *Input) GetLabel() (string) {
	return fmt.Sprintf(`<label for="id_%s">%s</label>`, inp.Name, inp.Label)
}

func (inp *Input) Validate(r *http.Request) (bool) {
	valid := true
	for _, validator := range inp.Validators {
		valid = validator(inp, r) && valid
	}
	return valid
}

func (inp *Input) ValidateRequired() {
	inp.Validators = append(inp.Validators, func (inp *Input, r *http.Request) (bool) {
		name := inp.Name
		value, ok := r.Form[name]
		ret := false
		if ok {
			if 0 != len(value) && 0 != len(value[0]) {
				ret = true
			}
		}
		if !ret {
			inp.Errors = append(inp.Errors, fmt.Sprintf("%s field is required", inp.Label))
		}
		return ret
	})
}

func (inp *Input) ValidateMatchRegexp(re *regexp.Regexp) {
	inp.Validators = append(inp.Validators, func (inp *Input, r *http.Request) (bool) {
		name := inp.Name
		value, ok := r.Form[name]
		ret := false
		if ok {
			if 0 != len(value) && re.Match([]byte(value[0])) {
				ret = true
			}
		}
		if !ret {
			inp.Errors = append(inp.Errors, fmt.Sprintf("%s is not valid", inp.Label))
		}
		return ret
	})
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

type Checkbox struct { Input }

func (inp *Checkbox) Render() (string) {
	return fmt.Sprintf(`<input type="checkbox" name="%s" id="id_%s"/>`,
		inp.Name, inp.Name)
}


// Button

type Button struct { Input }

func (inp *Button) Render() (string) {
	return fmt.Sprintf(`<button name="%s" type="submit">%s</button>`,
		inp.Name, inp.Label)
}

func (inp *Button) GetLabel() (string) {
	return ""
}


// Hidden

type Hidden struct { Input }

func (inp *Hidden) Render() (string) {
	return fmt.Sprintf(`<input type="hidden" name="%s"/>`, inp.Name)
}

func (inp *Hidden) GetLabel() (string) {
	return ""
}


// File 

type File struct { Input }

func (inp *File) Render() (string) {
	return fmt.Sprintf(`<input type="file" name="%s" id="id_%s"/>`,
		inp.Name, inp.Name)
}


/* Validators */

type Validator func (inp *Input, r *http.Request) (bool)


/* Form */

type Form struct {
	Inputs []InputInterface
	Validators []func(*http.Request) (bool)
}

func (frm *Form) Render() (string) {
	out := "\n<table>"
	for _, inp := range frm.Inputs {
		out += fmt.Sprintf(`
        <tr>
            <th>%s</th>
            <td>%s%s</td>
        </tr>`, inp.GetLabel(), inp.RenderErrors(), inp.Render())
	}
	out += "\n</table>\n"
	return out
}

func (frm *Form) AddInput(inp InputInterface) {
	frm.Inputs = append(frm.Inputs, inp)
}

func (frm *Form) Validate(r *http.Request) (bool) {
	valid := true
	for _, inp := range frm.Inputs {
		valid = inp.Validate(r) && valid
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

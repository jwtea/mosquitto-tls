package main

import (
	"github.com/manifoldco/promptui"
	errors "golang.org/x/xerrors"
)

// PromptModifyDefaults returns wether the user has asked to change request default values
func (o *Opts) PromptModifyDefaults() (bool, error) {
	p := promptui.Prompt{
		Label:     "Modify request defaults",
		IsConfirm: true,
	}

	_, err := p.Run()

	if err != nil {
		if errors.Is(err, promptui.ErrAbort) {
			return false, nil
		}
		return false, errors.Errorf("Modify prompt failed %v", err)
	}

	if err = o.PromptLastWill(); err != nil {
		return false, errors.Errorf("Modify prompt failed %v", err)
	}
	return true, nil
}

func (o *Opts) PromptMode() error {

	p := promptui.Select{
		Label: "Mode",
		Items: []string{"ModeAuto", "ModeManual"},
	}

	m, _, err := p.Run()

	if err != nil {
		return errors.Errorf("Mode prompt failed %v", err)
	}

	o.Mode = Mode(m)
	return nil
}

func (o *Opts) PromptLastWill() error {

	p := promptui.Prompt{
		Label:     "Enable last will",
		IsConfirm: true,
	}

	_, err := p.Run()

	if err != nil {
		if errors.Is(err, promptui.ErrAbort) {
			return nil
		}
		return errors.Errorf("Last will prompt failed %v", err)
	}

	o.SetWill = true

	return nil

}

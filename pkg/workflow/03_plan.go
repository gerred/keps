package workflow

import (
	"github.com/calebamiles/keps/pkg/keps"
	"github.com/calebamiles/keps/pkg/keps/sections"
	"github.com/calebamiles/keps/pkg/settings"
)

// Plan helps an author sketch out guides for
//  - operators: who may be forced to interact with the enhancement in rage and want to understand
//               how things can go wrong, how to know its working, and what other moving parts does
//               this enhancement interact with
//  - other developers: who need to review, comprehend, and extend the enhancement over its lifetime
//  - teachers: documentation writers, technical trainers, writers, all want to understand how to talk
//              about the enhancement
// Guide templates are rendered and their locations added to the KEP metadata. Plan also adds a template
// for iterating on success criteria as the enhancement works towards general availability
func Plan(runtime settings.Runtime) error {
	kep, err := keps.Open(runtime.TargetDir())
	if err != nil {
		return err
	}

	sectionContent, err := sections.ForImplementableState(kep)
	if err != nil {
		return err
	}

	kep.AddSections(sectionContent)
	err = kep.Persist()
	if err != nil {
		return err
	}

	err = kep.Persist()
	if err != nil {
		return err
	}

	// TODO add mechanics for creating PR

	return nil
}

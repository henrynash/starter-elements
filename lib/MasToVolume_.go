// example of how to convert a density and mass to a volume
package lib

import (
	"github.com/antha-lang/antha/antha/anthalib/wunit"
	"github.com/antha-lang/antha/component"
	"github.com/antha-lang/antha/execute"
	"github.com/antha-lang/antha/inject"
	"golang.org/x/net/context"
)

// Input parameters for this protocol

// Data which is returned from this protocol

// Physical inputs to this protocol

// Physical outputs from this protocol

func _MasToVolumeRequirements() {

}

// Actions to perform before protocol itself
func _MasToVolumeSetup(_ctx context.Context, _input *MasToVolumeInput) {

}

// Core process of the protocol: steps to be performed for each input
func _MasToVolumeSteps(_ctx context.Context, _input *MasToVolumeInput, _output *MasToVolumeOutput) {

	_output.Vol = wunit.MasstoVolume(_input.MyMass, _input.MyDensity)

	_output.BacktoMass = wunit.VolumetoMass(_output.Vol, _input.MyDensity)
}

// Actions to perform after steps block to analyze data
func _MasToVolumeAnalysis(_ctx context.Context, _input *MasToVolumeInput, _output *MasToVolumeOutput) {

}

func _MasToVolumeValidation(_ctx context.Context, _input *MasToVolumeInput, _output *MasToVolumeOutput) {

}
func _MasToVolumeRun(_ctx context.Context, input *MasToVolumeInput) *MasToVolumeOutput {
	output := &MasToVolumeOutput{}
	_MasToVolumeSetup(_ctx, input)
	_MasToVolumeSteps(_ctx, input, output)
	_MasToVolumeAnalysis(_ctx, input, output)
	_MasToVolumeValidation(_ctx, input, output)
	return output
}

func MasToVolumeRunSteps(_ctx context.Context, input *MasToVolumeInput) *MasToVolumeSOutput {
	soutput := &MasToVolumeSOutput{}
	output := _MasToVolumeRun(_ctx, input)
	if err := inject.AssignSome(output, &soutput.Data); err != nil {
		panic(err)
	}
	if err := inject.AssignSome(output, &soutput.Outputs); err != nil {
		panic(err)
	}
	return soutput
}

func MasToVolumeNew() interface{} {
	return &MasToVolumeElement{
		inject.CheckedRunner{
			RunFunc: func(_ctx context.Context, value inject.Value) (inject.Value, error) {
				input := &MasToVolumeInput{}
				if err := inject.Assign(value, input); err != nil {
					return nil, err
				}
				output := _MasToVolumeRun(_ctx, input)
				return inject.MakeValue(output), nil
			},
			In:  &MasToVolumeInput{},
			Out: &MasToVolumeOutput{},
		},
	}
}

var (
	_ = execute.MixInto
	_ = wunit.Make_units
)

type MasToVolumeElement struct {
	inject.CheckedRunner
}

type MasToVolumeInput struct {
	MyDensity wunit.Density
	MyMass    wunit.Mass
}

type MasToVolumeOutput struct {
	BacktoMass wunit.Mass
	Vol        wunit.Volume
}

type MasToVolumeSOutput struct {
	Data struct {
		BacktoMass wunit.Mass
		Vol        wunit.Volume
	}
	Outputs struct {
	}
}

func init() {
	if err := addComponent(component.Component{Name: "MasToVolume",
		Constructor: MasToVolumeNew,
		Desc: component.ComponentDesc{
			Desc: "example of how to convert a density and mass to a volume\n",
			Path: "src/github.com/antha-lang/starter-elements/an/AnthaAcademy/Lesson5_Units2/B_MasstoVolume.an",
			Params: []component.ParamDesc{
				{Name: "MyDensity", Desc: "", Kind: "Parameters"},
				{Name: "MyMass", Desc: "", Kind: "Parameters"},
				{Name: "BacktoMass", Desc: "", Kind: "Data"},
				{Name: "Vol", Desc: "", Kind: "Data"},
			},
		},
	}); err != nil {
		panic(err)
	}
}

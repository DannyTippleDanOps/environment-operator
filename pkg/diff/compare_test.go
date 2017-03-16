package diff

import (
	"testing"

	"github.com/pearsontechnology/environment-operator/pkg/bitesize"
)

func TestDiffEmpty(t *testing.T) {
	a := bitesize.Environment{}
	b := bitesize.Environment{}

	if d := Compare(a, b); d != "" {
		t.Errorf("Expected diff to be empty, got: %s", d)
	}

}

func TestIgnoreTestFields(t *testing.T) {
	a := bitesize.Environment{Tests: []bitesize.Test{}}
	b := bitesize.Environment{Tests: []bitesize.Test{
		{Name: "a"},
	}}

	if d := Compare(a, b); d != "" {
		t.Errorf("Expected diff to be empty, got: %s", d)
	}
}

func TestIgnoreDeploymentFields(t *testing.T) {
	a := bitesize.Environment{Deployment: &bitesize.DeploymentSettings{}}
	b := bitesize.Environment{Deployment: &bitesize.DeploymentSettings{
		Method: "bluegreen",
	}}

	if d := Compare(a, b); d != "" {
		t.Errorf("Expected diff to be empty, got: %s", d)
	}
}

func TestDiffNames(t *testing.T) {
	a := bitesize.Environment{Name: "asd"}
	b := bitesize.Environment{Name: "asdf"}

	if Compare(a, b) == "" {
		t.Error("Expected diff, got the same")
	}
}

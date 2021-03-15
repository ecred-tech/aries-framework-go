package gabriel

import "testing"

func TestGabriel(t *testing.T) {

	got := Gabriel()
	want := "Gabriel"

	if (got != want) {
		t.Errorf("got %q want %q", got, want)
	}
}

package plugin

import "testing"

func Test(t *testing.T) {
	p := GinSwaggerPlugin{}
	if !p.Enabled() {
		t.Fatal("swagger deafult enabled")
	}
	if p.Order() != defaultConfig["order"] {
		t.Fatalf("swagger order default: %d", defaultConfig["order"])
	}
	p.Load()
}

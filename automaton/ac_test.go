package automaton

import "testing"

func TestACAutomaton(t *testing.T) {
	ac := ACAutomaton{}
	ac.Reserve(10000)
	ac.Insert("she")
	ac.Insert("he")
	ac.Insert("say")
	ac.Insert("shr")
	ac.Insert("her")
	ac.Build()
	if ac.Query("yasherhs") != 3 {
		t.Fatalf("wrong answer")
	}
}

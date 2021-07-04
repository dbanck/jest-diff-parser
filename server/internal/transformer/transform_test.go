package transformer

import "testing"

func Test_tranform_example_1(t *testing.T) {
	source := `    +   Object {
    +     "cutQty": Array [
    +       Object {
    +         "item-1": null,
    +         "item-2": 441,
    +         "timestamp": 1623933323000
    +       }
	+     ]
    +   }`
	want := `{"cutQty":[{"item-1":null,"item-2":441,"timestamp":1623933323000,}],}`

	if got := Transform(source); got != want {
		t.Errorf("tranform() = %v, want %v", got, want)
	}
}

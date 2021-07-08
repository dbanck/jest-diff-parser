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
	want := `{"cutQty":[{"item-1":null,"item-2":441,"timestamp":1623933323000,},],,},`

	if got := Transform(source); got != want {
		t.Errorf("tranform() = \n%v\nwant\n%v", got, want)
	}
}

func Test_tranform_example_2(t *testing.T) {
	source := `+   Object {
+     "dataIndex": 0,
+     "decimals": 2,
+     "i18nId": "i18n-ref-1",
+     "multiplierId": 2,
+     "ref": "ref/analog/1",
+     "type": 0,
+     "unit": undefined,
+     "unitId": "1",
+     "yMax": 33,
+     "yMin": 11
+   }`
	want := `{"dataIndex":0,"decimals":2,"i18nId":"i18n-ref-1","multiplierId":2,"ref":"ref/analog/1","type":0,"unit":undefined,"unitId":"1","yMax":33,"yMin":11,},`

	if got := Transform(source); got != want {
		t.Errorf("tranform() = \n%v\nwant\n%v", got, want)
	}
}

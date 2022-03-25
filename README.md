# GoProjektarbeit
In dieser Projektarbeit war es die Aufgabe eine bestehende implementation eines Mini-Compilers in C++ nach Go zu übersetzen (siehe [Übung SoftwareProjekt](https://sulzmann.github.io/SoftwareProjekt/schein-neu.html#(6))).

## Klassen
### [Optional](Optional/Optional.go)
Eine generische Klasse um ein potentiell leeres Objekt zu halten
### [Expression](Expression/Expression.go)
`Expression` ist ein Interface und wird für die Implementation folgender Ausdrücke verwendet:
- **`IntExp`**:<br>
Ein Ausdruck der nur genau einen ganzzahligen Wert hält.
- **`PlusExp`**:<br>
Ein Ausdruck der genau zwei weitere Ausdrücke(`Expressions`) hält und deren Addition repräsentiert.
- **`MultExp`**:<br>
Ein Ausdruck der genau zwei weitere Ausdrücke(`Expressions`) hält und deren Multiplikation repräsentiert.
### [Stack](Stack/Stack.go)
Diese Klasse kapselt den *Go*-Datentyp `Slice` und bereichert ihn mit Funktionalitäten eines Stacks für die einfachere Nutzung mit der [VM](VM/vm.go).
### [Tokenizer](Tokenizer/Tokenizer.go)
Der [Tokenizer](Tokenizer/Tokenizer.go) hat die Aufgabe einen geschriebenen Ausdruck(`Expression`) in Form eines *Strings* in eine Sequenz von *Tokens* umzuwandeln, die so einfacher vom [Parser](Parser/Parser.go) verarbeitet werden können.<br><br>
Beispiel:<br>
    *"(2 + 1) \* 2"* =>  \[`OPEN`, `TWO`, `PLUS`, `ONE`, `CLOSE`, `MULT`, `TWO`\]
### [Parser](Parser/Parser.go)
Der [Parser](Parser/Parser.go) soll nun mit Hilfe des [Tokenizers](Tokenizer/Tokenizer.go) geschriebene Ausdrücke(`Expression`) in Instanzen von [Expressions](Expression/Expression.go) umwandeln. Dies geschieht nach den folgenden Regeln:<br>
```
E ::= E + T | T

T ::= T * F | F

F ::= N | (E)
```
Sollte ein Ausdruck nicht diesen Regeln folgen, so wird eine leere [Expression](Expression/Expression.go)(Nothing\[Expression]) zurückgegeben.
### [VM](VM/vm.go)
Die [VM](VM/vm.go) soll nun eine virtuelle Maschine simulieren, die diese [Expressions](Expression/Expression.go) auf einem eigenen (simulierten)[Stack](Stack/Stack.go) verarbeitet. Dafür werden die drei *OpCodes* `PUSH`, `PLUS` und `MULT` definiert:
- **`PUSH`**:<br>
Es wird eine Zahl auf den Stack gelegt.
- **`PLUS`**:<br>
Es werden die beiden obersten Zahlen vom Stack genommen, addiert und das Ergebnis daraus wieder auf den Stack gelegt.
- **`MULT`**:<br>
Es werden die beiden obersten Zahlen vom Stack genommen, multipliziert und das Ergebnis daraus wieder auf den Stack gelegt.

## Beispiele
### Parser
Der folgende Code führt zu folgender Ausgabe:
```
display(NewParser("(1)").Parse())
display(NewParser("1").Parse())
display(NewParser("1 + 0").Parse())
display(NewParser("1 + (0) ").Parse())
display(NewParser("1 + 2 * 0 ").Parse())
display(NewParser("1 * 2 + 0 ").Parse())
display(NewParser("(1 + 2) * 0 ").Parse())
display(NewParser("(1 + 2) * 0 + 2").Parse())

// Ausgabe
// 1
// 1
// (1 + 0)
// (1 + 0)
// (1 + (2 * 0))
// ((1 * 2) + 0)
// ((1 + 2) * 0)
// (((1 + 2) * 0) + 2)
```

### VM
Der folgende Code führt zu folgender Ausgabe:
```
codes := []Code{
  NewPush(1),
  NewPush(2),
  NewPush(3),
  NewMult(),
  NewPlus(),
}
vm := NewVm(codes)
res := vm.Run()
showVMRes(res)

codes = []Code{
  NewPush(2),
  NewPush(3),
  NewPush(5),
  NewPlus(),
  NewMult(),
}
vm = NewVm(codes)
res = vm.Run()
showVMRes(res)

res = RunOnVM("(1 + 2) * 2 + 2")
showVMRes(res)

// Ausgabe
// VM stack (top): 7
// VM stack (top): 16
// VM stack (top): 8
```

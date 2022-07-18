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
```go
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
```go
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

## Vergleich der Go-Konzepte zu C++
### Klassen
In C++ werden Klassen sowie ihre zugehörigen Variablen und Funktionen direkt definiert. In Go werden nur die Klassenvariablen bei der Definition vorgegeben und dann beliebige funktionen dazu implementiert. <br><br>
**C++**
```c++
class IntExpression : public Expression
{
public:
	IntExpression(int value);
	int Eval();
	char* Pretty();
	void Compile(Stack<Code> codeStack);

protected:
	int value;
};
```
**Go**
```go
type IntExpression struct {
	Value int
}
```
### Interfaces
In C++ gibt es keine expliziten Interfaces. Man erreicht die Funktionen eines Interfaces über die definition abstrakter Oberklassen in denen pur virtuelle Funktionen definiert werden, die dann in abgeleiteten Klassen implementiert werden müssen.<br><br>
**C++**
```c++
class Expression {
public:
	virtual int Eval() = 0;
	virtual char* Pretty() = 0;
	virtual void Compile(Stack<Code> codeStack) = 0;
};
```
In Go können Interfaces direkt definiert werden. Hierbei werden die nötigen zu implementierenden Funktionen festgelegt. Die Zuordnung, welche Klassen dieses Interface implementieren geschieht dann zur Compilezeit. Jede Klasse die die Bedingungen des Interfaces erfüllt, wird zu einer Implementierung dieses Interfaces. Im Gegensatz zu C++, wo das "Interface" als Oberklasse der Implementation angegeben werden muss. <br>
**Go**
```go
type Expression interface {
	Eval() int
	Pretty() string
	Compile(codeStack Stack[Code])
}
```
### Generics
Generische Klassen heißen in C++ "templates". Sie werden über das Schlüsselwort "template" gefolgt von spitzen Klammern definiert, in denen der Generische Typ eingeschränkt werden kann.<br><br>
**C++**
```c++
template<class T> class Stack {
   ...
};
```
Seit Version 1.18 unterstützt auch Go generische Strukturen. Sie werden über eckige Klammern hinter dem Namen einer Klasse (oder Interfaces/Funktionen) markiert, in denen ebenfalls der generische Typ eingeschränkt werden kann. <br>
**Go**
```go
type Stack[T any] interface {
	...
}
```
### Anderes
#### Overloading
Go unterstützt das Überladen von Funktionen nicht. Es ist also nicht möglich zwei Funktionen mit dem selben Namen zu definieren, die sich nur in den Paramtern unterscheiden.
#### Objektorientierung
Auch wenn man in Go die Funktionalitäten von Klassen mit Hilfe von Structs und Interfaces teilweise nachbauen kann ist Go nicht Objektorientiert. Konzepte wie Vererbung sind nicht vorgesehen und nur schwer zu replizieren.
#### Concurrency
Go bietet mit seinen "Go-Routinen" eine sehr einfache Möglichkeit, Funktionen parallel laufen zu lassen. Mit dem Schlüsselwort "go" gefolgt von einem Funktionsaufruf wird dieser parallel zum Haupt Thread ausgeführt. Dazu bietet Go noch nützliche Datentypen wie [Channel](https://golangdocs.com/channels-in-golang) oder Pakete wie [sync](https://pkg.go.dev/sync), die die Parallelisierung noch weiter vereinfachen. Das Parallelisieren von Funktionen in C++ ist dagegen sehr aufwendig und komplex.

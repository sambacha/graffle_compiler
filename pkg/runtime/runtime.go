package runtime

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BuiltinType interface {
	Number() (Number, bool)
	String() String
	Bool() Bool
}

type NumericType interface {
	Number() (Number, bool)
}

type StringType interface {
	String() String
}

type LogicalType interface {
	Bool() Bool
}

type Number struct {
	float64
}

type String struct {
	string
}

type Bool struct {
	bool
}

func NewNumber(val float64) Number {
	return Number{val}
}

func NewString(val string) String {
	return String{val}
}

func NewBool(val bool) Bool {
	return Bool{val}
}

func (n Number) Val() float64 {
	return n.float64
}

func (s String) Val() string {
	return s.string
}

func (b Bool) Val() bool {
	return b.bool
}

func (n Number) Number() (Number, bool) {
	return n, true
}

func (n Number) String() String {
	return NewString(fmt.Sprintf("%f", n.Val()))
}

func (n Number) Bool() Bool {
	if n.Val() == 0 {
		return NewBool(false)
	}
	return NewBool(true)
}

func (s String) Number() (Number, bool) {
	str := strings.ReplaceAll(s.Val(), ",", ".")
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return NewNumber(0), false
	}
	return NewNumber(f), true
}

func (s String) String() String {
	return s
}

func (s String) Bool() Bool {
	str := strings.ToLower(s.Val())
	str = strings.ReplaceAll(str, " ", "")
	if str == "" || str == "false" || str == "nottrue" || str == "0" {
		return NewBool(false)
	}
	return NewBool(true)
}

func (b Bool) Number() (Number, bool) {
	if b.Val() {
		return NewNumber(1), true
	}
	return NewNumber(0), true
}

func (b Bool) String() String {
	if b.Val() {
		return NewString("True")
	}
	return NewString("False")
}

func (b Bool) Bool() Bool {
	return b
}

type Labeled interface {
	Label(l string)
}

type Vertice struct {
	label string
	val   BuiltinType
}

type Edge struct {
	label      string
	weight     Number
	isDirected bool
	v1         Vertice
	v2         Vertice
}

type Graph struct {
	label string
	v     []Vertice
	e     []Edge
}

func (g *Graph) AddVertice(nv Vertice) {
	vset := map[Vertice]struct{}{}
	for _, v := range g.v {
		vset[v] = struct{}{}
	}

	if _, b := vset[nv]; !b {
		g.v = append(g.v, nv)
	}
}

func (g *Graph) AddEdge(ne Edge) {
	eset := map[Edge]struct{}{}
	for _, e := range g.e {
		eset[e] = struct{}{}
	}

	if _, b := eset[ne]; !b {
		g.e = append(g.e, ne)
	}
}

func (g *Graph) AddGraph(ng Graph) {
	for _, v := range ng.v {
		g.AddVertice(v)
	}

	for _, e := range ng.e {
		g.AddEdge(e)
	}
}

func (g *Graph) GetSingleVertices() []Vertice {
	connv := map[Vertice]struct{}{}
	for _, e := range g.e {
		connv[e.v1] = struct{}{}
		connv[e.v2] = struct{}{}
	}

	var sv []Vertice
	for _, v := range g.v {
		if _, b := connv[v]; !b {
			sv = append(sv, v)
		}
	}
	return sv
}

func NewVertice(val interface{}) Vertice {
	bo, b := val.(bool)
	if b {
		return Vertice{val: NewBool(bo)}
	}

	bi, b := val.(BuiltinType)
	if b {
		return Vertice{val: bi}
	}

	v, b := val.(Vertice)
	if b {
		return v
	}

	log.Fatalf("Assignement error! Cannot create vertice from type: %t", val)
	return NewVertice(false)
}

func NewEdge(v1, v2 Vertice, w Number, d bool) Edge {
	return Edge{v1: v1, v2: v2, weight: w, isDirected: d}
}

func NewGraph(args ...interface{}) Graph {
	var g Graph

	for _, a := range args {
		if v, b := a.(BuiltinType); b {
			g.AddVertice(NewVertice(v))
		}
		if v, b := a.(Vertice); b {
			g.AddVertice(v)
		}
		if e, b := a.(Edge); b {
			g.AddEdge(e)
		}
		if g, b := a.(Graph); b {
			g.AddGraph(g)
		}
	}

	return g
}

func (v *Vertice) Label(l string) {
	v.label = l
}

func (v Vertice) String() String {
	str := v.val.String()
	if v.label == "" {
		return NewString(fmt.Sprintf("(%s)", str))
	}
	return NewString(fmt.Sprintf("(%s)@[%s]", str, v.label))
}

func (v Vertice) Bool() Bool {
	return v.val.Bool()
}

func (e *Edge) Label(l string) {
	e.label = l
}

func (e Edge) String() String {
	v1 := e.v1.String()
	str := fmt.Sprintf("%s -", v1)

	if e.weight.float64 != 0 {
		w := e.weight.String()
		str = fmt.Sprintf("%s[%s]-", str, w)
	}

	if e.isDirected {
		str += ">"
	}

	v2 := e.v2.String()
	str = fmt.Sprintf("%s %s", str, v2)

	if e.label != "" {
		str = fmt.Sprintf("{%s}@[%s]", str, e.label)
	}

	return NewString(str)
}

func (e Edge) Bool() Bool {
	if e.v1.Bool().bool == false && e.v2.Bool().bool == false {
		return NewBool(false)
	}
	return NewBool(true)
}

func (g *Graph) Label(l string) {
	g.label = l
}

func (g Graph) String() (String, bool) {
	var str string
	for _, v := range g.GetSingleVertices() {
		str = fmt.Sprintf("%s\n\t%s;", str, v.String())
	}

	for _, e := range g.e {
		str = fmt.Sprintf("%s\n\t%s;", str, e.String())
	}

	if g.label != "" {
		str = fmt.Sprintf("{%s\n}@[%s]", str, g.label)
	}

	return NewString(str), true
}

func (g Graph) Bool() Bool {
	if g.e == nil && g.v == nil {
		return NewBool(false)
	}
	return NewBool(true)
}

func Add(l, r interface{}) BuiltinType {
	li, lb := l.(NumericType)
	ri, rb := r.(NumericType)
	if lb && rb {
		ln, lc := li.Number()
		rn, rc := ri.Number()
		if lc && rc {
			ln.float64 += rn.float64
			return ln
		}
	}
	log.Fatalf("Error! Wrong types passed in Addition! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Subtract(l, r interface{}) BuiltinType {
	li, lb := l.(NumericType)
	ri, rb := r.(NumericType)
	if lb && rb {
		ln, lc := li.Number()
		rn, rc := ri.Number()
		if lc && rc {
			ln.float64 -= rn.float64
			return ln
		}
	}
	log.Fatalf("Error! Wrong types passed in Subtraction! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Multiply(l, r interface{}) BuiltinType {
	li, lb := l.(NumericType)
	ri, rb := r.(NumericType)
	if lb && rb {
		ln, lc := li.Number()
		rn, rc := ri.Number()
		if lc && rc {
			ln.float64 *= rn.float64
			return ln
		}
	}
	log.Fatalf("Error! Wrong types passed in Multiplication! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Divide(l, r interface{}) interface{} {
	li, lb := l.(NumericType)
	ri, rb := r.(NumericType)
	if lb && rb {
		ln, lc := li.Number()
		rn, rc := ri.Number()
		if lc && rc {
			ln.float64 /= rn.float64
			return ln
		}
	}
	log.Fatalf("Error! Wrong types passed in Dividing! left: %t, right: %t", l, r)
	return NewNumber(0)
}

func Not(a interface{}) Bool {
	i, b := a.(LogicalType)
	if b {
		v := i.Bool()
		return NewBool(!v.bool)
	}
	log.Fatalf("Error! Wrong type passed in \"Not\" operation! type: %t", a)
	return NewBool(false)
}

func And(l, r interface{}) Bool {
	li, lb := l.(LogicalType)
	ri, rb := r.(LogicalType)
	if lb && rb {
		lv := li.Bool()
		rv := ri.Bool()
		return NewBool(lv.bool && rv.bool)
	}
	log.Fatalf("Error! Wrong types passed in \"And\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func Or(l, r interface{}) Bool {
	li, lb := l.(LogicalType)
	ri, rb := r.(LogicalType)
	if lb && rb {
		lv := li.Bool()
		rv := ri.Bool()
		return NewBool(lv.bool || rv.bool)
	}
	log.Fatalf("Error! Wrong types passed in \"Or\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func Equals(l, r interface{}) Bool {
	li, lb := l.(BuiltinType)
	ri, rb := r.(BuiltinType)
	if lb && rb {
		lv, lb := li.Number()
		rv, rb := ri.Number()
		if lb && rb {
			return NewBool(lv.float64 == rv.float64)
		}

		log.Fatalf("Error! Wrong operands, cannot evaluate \"Equals\" operation! left: %t %s, right: %t %s", li, li.String(), ri, ri.String())
		return NewBool(false)
	}
	log.Fatalf("Error! Wrong types passed in \"Equals\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func Less(l, r interface{}) Bool {
	li, lb := l.(BuiltinType)
	ri, rb := r.(BuiltinType)
	if lb && rb {
		lv, lb := li.Number()
		rv, rb := ri.Number()
		if lb && rb {
			return NewBool(lv.float64 < rv.float64)
		}

		log.Fatalf("Error! Wrong operands, cannot evaluate \"Less than\" operation! left: %t %s, right: %t %s", li, li.String(), ri, ri.String())
		return NewBool(false)
	}
	log.Fatalf("Error! Wrong types passed in \"Less than\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func Greater(l, r interface{}) Bool {
	li, lb := l.(BuiltinType)
	ri, rb := r.(BuiltinType)
	if lb && rb {
		lv, lb := li.Number()
		rv, rb := ri.Number()
		if lb && rb {
			return NewBool(lv.float64 > rv.float64)
		}

		log.Fatalf("Error! Wrong operands, cannot evaluate \"Greater than\" operation! left: %t %s, right: %t %s", li, li.String(), ri, ri.String())
		return NewBool(false)
	}
	log.Fatalf("Error! Wrong types passed in \"Greater than\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func LessOrEquals(l, r interface{}) Bool {
	li, lb := l.(BuiltinType)
	ri, rb := r.(BuiltinType)
	if lb && rb {
		lv, lb := li.Number()
		rv, rb := ri.Number()
		if lb && rb {
			return NewBool(lv.float64 <= rv.float64)
		}

		log.Fatalf("Error! Wrong operands, cannot evaluate \"Less than or equals\" operation! left: %t %s, right: %t %s", li, li.String(), ri, ri.String())
		return NewBool(false)
	}
	log.Fatalf("Error! Wrong types passed in \"Less than or equals\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func GreaterOrEquals(l, r interface{}) Bool {
	li, lb := l.(BuiltinType)
	ri, rb := r.(BuiltinType)
	if lb && rb {
		lv, lb := li.Number()
		rv, rb := ri.Number()
		if lb && rb {
			return NewBool(lv.float64 >= rv.float64)
		}

		log.Fatalf("Error! Wrong operands, cannot evaluate \"Greater than or equals\" operation! left: %t %s, right: %t %s", li, li.String(), ri, ri.String())
		return NewBool(false)
	}
	log.Fatalf("Error! Wrong types passed in \"Greater than or equals\" operation! left: %t, right: %t", l, r)
	return NewBool(false)
}

func Print(i interface{}) {
	if bi, b := i.(StringType); b {
		str := bi.String()
		fmt.Println(str.string)
		return
	}

	fmt.Printf("%+v\n", i)
}

func Input() String {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">>> ")
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Panicf("Error scanning input: %v", err)
	}
	return NewString(text)
}

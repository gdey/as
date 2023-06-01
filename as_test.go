package as_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/gdey/as"
)

func TestBool(t *testing.T) {
	type tcase struct {
		it interface{}
		b  bool
		ok bool
	}
	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {
			b, ok := as.Bool(tc.it)
			if ok != tc.ok {
				t.Errorf("ok, expected %v got %v", tc.ok, ok)
				return
			}
			if b != tc.b {
				t.Errorf("b, expected %v got %v", tc.b, b)
				return
			}
		}
	}
	tests := map[string]tcase{
		"empty string": {ok: false, b: false, it: ""},
		"asdf":         {ok: false, b: false, it: "asdf"},
		"1234":         {ok: false, b: false, it: "1234"},
		"nil":          {ok: false, b: false, it: nil},
		"[]int":        {ok: false, b: false, it: make([]int, 1)},
		"False":        {ok: true, b: false, it: "False"},
		"0":            {ok: true, b: false, it: "0"},
		"#0":           {ok: true, b: false, it: int(0)},
		"uint#0":       {ok: true, b: false, it: uint(0)},
		"#-1":          {ok: true, b: false, it: int(-1)},
		"#2":           {ok: true, b: false, it: int(2)},
		"#0.0":         {ok: true, b: false, it: float64(0.0)},
		"#-1.0":        {ok: true, b: false, it: float64(-1.0)},
		"#2.0":         {ok: true, b: false, it: float64(2.0)},
		"false":        {ok: true, b: false, it: "false"},
		"F":            {ok: true, b: false, it: "F"},
		"f":            {ok: true, b: false, it: "f"},
		"bool#false":   {ok: true, b: false, it: false},
		"*bool#false":  {ok: true, b: false, it: as.Pointer(false)},
		"*bool#nil":    {ok: true, b: false, it: (*bool)(nil)},
		"true":         {ok: true, b: true, it: "true"},
		"True":         {ok: true, b: true, it: "True"},
		"T":            {ok: true, b: true, it: "T"},
		"t":            {ok: true, b: true, it: "t"},
		"1":            {ok: true, b: true, it: "1"},
		"#1":           {ok: true, b: true, it: int(1)},
		"uint#1":       {ok: true, b: true, it: uint(1)},
		"bool#true":    {ok: true, b: true, it: true},
		"*bool#true":   {ok: true, b: true, it: as.Pointer(true)},
	}
	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}

func TestInt(t *testing.T) {
	type tcase struct {
		it interface{}
		i  int64
		ok bool
	}
	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {
			t.Run("int64", func(t *testing.T) {

				i, ok := as.Int64(tc.it)
				if ok != tc.ok {
					t.Errorf("ok, expected %v got %v", tc.ok, ok)
					return
				}
				if i != tc.i {
					t.Errorf("i, expected %v got %v", tc.i, i)
					return
				}

			})
			t.Run("int", func(t *testing.T) {

				i, ok := as.Int(tc.it)
				if ok != tc.ok {
					t.Errorf("ok, expected %v got %v", tc.ok, ok)
					return
				}
				if i != int(tc.i) {
					t.Errorf("i, expected %v got %v", tc.i, i)
					return
				}

			})
		}
	}
	tests := map[string]tcase{
		"nil":        {ok: false, i: 0, it: nil},
		"abcde":      {ok: false, i: 0, it: "abcde"},
		"[]int":      {ok: false, i: 0, it: make([]int, 1)},
		"0":          {ok: true, i: 0, it: "0"},
		"#0":         {ok: true, i: 0, it: 0},
		"bool#true":  {ok: true, i: 1, it: true},
		"bool#false": {ok: true, i: 0, it: false},
		"float#0":    {ok: true, i: 0, it: 0.0},
		"float#1.0":  {ok: true, i: 1, it: 1.0},
		"uint#0.0":   {ok: true, i: 0, it: uint(0)},
		"uint#1.0":   {ok: true, i: 1, it: uint(1)},
		"uint32#0.0": {ok: true, i: 0, it: uint32(0)},
		"uint32#1.0": {ok: true, i: 1, it: uint32(1)},
	}
	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}
func TestUint(t *testing.T) {
	type tcase struct {
		it interface{}
		i  uint64
		ok bool
	}
	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {
			t.Run("uint64", func(t *testing.T) {

				i, ok := as.Uint64(tc.it)
				if ok != tc.ok {
					t.Errorf("ok, expected %v got %v", tc.ok, ok)
					return
				}
				if i != tc.i {
					t.Errorf("i, expected %v got %v", tc.i, i)
					return
				}

			})
			t.Run("uint", func(t *testing.T) {

				i, ok := as.Uint(tc.it)
				if ok != tc.ok {
					t.Errorf("ok, expected %v got %v", tc.ok, ok)
					return
				}
				if i != uint(tc.i) {
					t.Errorf("i, expected %v got %v", tc.i, i)
					return
				}

			})
		}
	}
	tests := map[string]tcase{
		"nil":        {ok: false, i: 0, it: nil},
		"abcde":      {ok: false, i: 0, it: "abcde"},
		"[]int":      {ok: false, i: 0, it: make([]int, 1)},
		"0":          {ok: true, i: 0, it: "0"},
		"#0":         {ok: true, i: 0, it: 0},
		"bool#true":  {ok: true, i: 1, it: true},
		"bool#false": {ok: true, i: 0, it: false},
		"float#0":    {ok: true, i: 0, it: 0.0},
		"float#1.0":  {ok: true, i: 1, it: 1.0},
		"uint#0.0":   {ok: true, i: 0, it: uint(0)},
		"uint#1.0":   {ok: true, i: 1, it: uint(1)},
		"uint32#0.0": {ok: true, i: 0, it: uint32(0)},
		"uint32#1.0": {ok: true, i: 1, it: uint32(1)},
	}
	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}
func TestFloat64(t *testing.T) {
	type tcase struct {
		it interface{}
		f  float64
		ok bool
	}
	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {

			f, ok := as.Float64(tc.it)
			if ok != tc.ok {
				t.Errorf("ok, expected %v got %v", tc.ok, ok)
				return
			}
			if f != tc.f {
				t.Errorf("f, expected %v got %v", tc.f, f)
				return
			}

		}
	}
	tests := map[string]tcase{
		"nil":        {ok: false, f: 0.0, it: nil},
		"abcde":      {ok: false, f: 0.0, it: "abcde"},
		"[]int":      {ok: false, f: 0.0, it: make([]int, 1)},
		"0":          {ok: true, f: 0.0, it: "0"},
		"#0":         {ok: true, f: 0.0, it: 0},
		"bool#true":  {ok: true, f: 1.0, it: true},
		"bool#false": {ok: true, f: 0.0, it: false},
		"float#0":    {ok: true, f: 0.0, it: 0.0},
		"float#1.0":  {ok: true, f: 1.0, it: 1.0},
		"uint#0.0":   {ok: true, f: 0.0, it: uint(0)},
		"uint#1.0":   {ok: true, f: 1.0, it: uint(1)},
		"uint32#0.0": {ok: true, f: 0.0, it: uint32(0)},
		"uint32#1.0": {ok: true, f: 1.0, it: uint32(1)},
	}
	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}

type nullString struct{}

func (nullString) String() string { return "null" }

func TestString(t *testing.T) {
	type tcase struct {
		it interface{}
		s  string
		ok bool
	}
	fn := func(tc tcase) func(*testing.T) {
		return func(t *testing.T) {

			s, ok := as.String(tc.it)
			if ok != tc.ok {
				t.Errorf("ok, expected %v got %v", tc.ok, ok)
				return
			}
			if s != tc.s {
				t.Errorf("s, expected '%v' got '%v'", tc.s, s)
				return
			}

		}
	}
	tests := map[string]tcase{
		"nil":        {ok: false, s: "", it: nil},
		"func":       {ok: false, s: "", it: func() {}},
		"*bool":      {ok: false, s: "", it: (*bool)(nil)},
		"[]int":      {ok: true, s: "[0]", it: make([]int, 1)},
		"abcde":      {ok: true, s: "abcde", it: "abcde"},
		"0":          {ok: true, s: "0", it: "0"},
		"#0":         {ok: true, s: "0", it: 0},
		"bool#true":  {ok: true, s: "true", it: true},
		"bool#false": {ok: true, s: "false", it: false},
		"float#0":    {ok: true, s: "0", it: 0.0},
		"float#1.0":  {ok: true, s: "1", it: 1.0},
		"uint#0.0":   {ok: true, s: "0", it: uint(0)},
		"uint#1.0":   {ok: true, s: "1", it: uint(1)},
		"uint32#0.0": {ok: true, s: "0", it: uint32(0)},
		"uint32#1.0": {ok: true, s: "1", it: uint32(1)},
		"err":        {ok: true, s: "error value", it: errors.New("error value")},
		"nulString":  {ok: true, s: "null", it: nullString{}},
	}
	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}


	type interfaceSliceTCase[U any] struct{
		it  []U
		s   []any
	}

func interfaceSliceFn[U any](tc interfaceSliceTCase[U]) func(*testing.T){
		return func(t *testing.T) {
			s := as.InterfaceSlice(tc.it...)
			if !reflect.DeepEqual(tc.s, s) {
				t.Errorf("s, expected '%v' got '%v'", tc.s, s)
				return
			}
		}
}

func TestInterfaceSlice(t *testing.T) {

	tests := map[string]func(*testing.T){
		"nil":           interfaceSliceFn(interfaceSliceTCase[int]{s: []interface{}{}, it: nil}),
		"[]int":         interfaceSliceFn(interfaceSliceTCase[int]{s: []interface{}{1}, it: []int{1}}),
		"[]interface{}": interfaceSliceFn(interfaceSliceTCase[interface{}]{s: make([]interface{}, 1), it: make([]interface{}, 1)}),
	}
	for name, tc := range tests {
		t.Run(name, tc)
	}
}

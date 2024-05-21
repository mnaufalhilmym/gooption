package gooption_test

import (
	"testing"

	"github.com/mnaufalhilmym/gooption"
)

func TestSome(t *testing.T) {
	t.Log("Running TestSome.")

	data := gooption.Some("Some data")

	if data.IsSome() {
		if d := data.Unwrap(); d != "Some data" {
			t.Error("TestSome failed. data.Unwrap() should return 'Some data'")
		}
		if d := data.UnwrapOr("Another data"); d != "Some data" {
			t.Error("TestSome failed. data.UnwrapOr() should return 'Some data'")
		}
		if d := data.UnwrapOrDefault(); d != "Some data" {
			t.Error("TestSome failed. data.UnwrapOrDefault() should return 'Some data'")
		}
		if d := data.UnwrapOrElse(func() string { return "Another data2" }); d != "Some data" {
			t.Error("TestSome failed. data.UnwrapOrElse() should return 'Some data'")
		}
	} else {
		t.Error("TestSome failed. data.IsSome() should return true")
	}
	if data.IsNone() {
		t.Error("TestSome failed. data.IsNone() should return false")
	}
}

func TestNone(t *testing.T) {
	t.Log("Running TestNone.")

	data := gooption.None[string]()

	if data.IsNone() {
		if d := data.UnwrapOr("Another data"); d != "Another data" {
			t.Error("TestSome failed. data.UnwrapOr() should return 'Another data'")
		}
		if d := data.UnwrapOrDefault(); d != "" {
			t.Error("TestSome failed. data.UnwrapOrDefault() should return '' (empty string)")
		}
		if d := data.UnwrapOrElse(func() string { return "Another data2" }); d != "Another data2" {
			t.Error("TestSome failed. data.UnwrapOrElse() should return 'Another data2'")
		}
	} else {
		t.Error("TestSome failed. data.IsNone() should return true")
	}
	if data.IsSome() {
		t.Error("TestSome failed. data.IsSome() should return false")
	}
}

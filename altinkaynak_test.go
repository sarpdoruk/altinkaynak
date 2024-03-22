package altinkaynak

import (
	"testing"
)

func TestWithApiUrl(t *testing.T) {
	var s ServiceInterface

	a := NewAltinkaynak(WithApiUrl("https://google.com"))
	s = a.CurrencyService
	err := s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	a = NewAltinkaynak(WithApiUrl("https://google.com"))
	s = a.GoldService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	a = NewAltinkaynak(WithApiUrl("https://google.com"))
	s = a.MainService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestWithCredentials(t *testing.T) {
	var s ServiceInterface

	a := NewAltinkaynak(WithCredentials("wrong", "credentials"))
	s = a.CurrencyService
	err := s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	a = NewAltinkaynak(WithCredentials("wrong", "credentials"))
	s = a.GoldService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	a = NewAltinkaynak(WithCredentials("wrong", "credentials"))
	s = a.MainService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestSendRequestWithMalformedMethod(t *testing.T) {
	_, err := SendRequest("*?", "https://google.com", "")
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func TestSendRequestWithMalformedUrl(t *testing.T) {
	_, err := SendRequest("POST", "google", "")
	if err == nil {
		t.Error("expected an error, got nil")
	}

	var s ServiceInterface

	a := NewAltinkaynak(WithApiUrl("google"))
	s = a.CurrencyService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	a = NewAltinkaynak(WithApiUrl("google"))
	s = a.GoldService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}

	a = NewAltinkaynak(WithApiUrl("google"))
	s = a.MainService
	err = s.Fetch()
	if err == nil {
		t.Error("expected an error, got nil")
	}
}

func ExampleNewAltinkaynak() {
	_ = NewAltinkaynak()
}

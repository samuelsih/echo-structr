package presentation

import (
	"reflect"
	"testing"
)

func TestOptions(t *testing.T) {
	t.Parallel()

	t.Run("url + qp", func(t *testing.T) {
		opts := With(UrlParams("id", "user"), QueryParams("status", "curr"))

		Equal(t, len(opts.URLParams), 2)
		Equal(t, len(opts.QueryParams), 2)

		url := []string{"id", "user"}
		qp := []string{"status", "curr"}

		if !reflect.DeepEqual(opts.URLParams, url) {
			t.Fatalf("url is not equal: %v - %v", opts.URLParams, url)
		}

		if !reflect.DeepEqual(opts.QueryParams, qp) {
			t.Fatalf("qp is not equal: %v - %v", opts.URLParams, qp)
		}
	})

	t.Run("only url", func(t *testing.T) {
		opts := With(UrlParams("id", "user"))

		Equal(t, len(opts.URLParams), 2)
		Equal(t, len(opts.QueryParams), 0)

		url := []string{"id", "user"}
		var qp []string

		if !reflect.DeepEqual(opts.URLParams, url) {
			t.Fatalf("url is not equal: %v - %v", opts.URLParams, url)
		}

		if !reflect.DeepEqual(opts.QueryParams, qp) {
			t.Fatalf("qp is not equal: %v - %v", opts.URLParams, qp)
		}
	})

	t.Run("only qp", func(t *testing.T) {
		opts := With(QueryParams("status", "curr"))

		Equal(t, len(opts.URLParams), 0)
		Equal(t, len(opts.QueryParams), 2)

		var url []string
		qp := []string{"status", "curr"}

		if !reflect.DeepEqual(opts.URLParams, url) {
			t.Fatalf("url is not equal: %v - %v", opts.URLParams, url)
		}

		if !reflect.DeepEqual(opts.QueryParams, qp) {
			t.Fatalf("qp is not equal: %v - %v", opts.URLParams, qp)
		}
	})

	t.Run("no url + qp", func(t *testing.T) {
		opts := With()

		Equal(t, len(opts.URLParams), 0)
		Equal(t, len(opts.QueryParams), 0)

		var url []string
		var qp []string

		if !reflect.DeepEqual(opts.URLParams, url) {
			t.Fatalf("url is not equal: %v - %v", opts.URLParams, url)
		}

		if !reflect.DeepEqual(opts.QueryParams, qp) {
			t.Fatalf("qp is not equal: %v - %v", opts.URLParams, qp)
		}
	})
}

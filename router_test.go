package bunrouter

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/infinytum/go-mojito"
)

func request(method string, path string) (*http.Response, error) {
	req := &http.Request{
		Method: method,
		URL: &url.URL{
			Scheme: "http",
			Host:   "localhost:8080",
			Path:   path,
		},
	}

	client := &http.Client{}
	return client.Do(req)
}

func Test_Router_DELETE(t *testing.T) {
	r := NewBunRouterRouter()
	r.DELETE("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("OK")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	res, err := request("DELETE", "/")
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", body)
	}
}

func Test_Router_GET(t *testing.T) {
	r := NewBunRouterRouter()
	r.GET("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("OK")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	res, err := request("GET", "/")
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", body)
	}
}

func Test_Router_HEAD(t *testing.T) {
	r := NewBunRouterRouter()
	r.HEAD("/", func(req *mojito.Request, res *mojito.Response) error {
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	res, err := request("HEAD", "/")
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
}

func Test_Router_POST(t *testing.T) {
	r := NewBunRouterRouter()
	r.POST("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("OK")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	res, err := request("POST", "/")
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", body)
	}
}

func Test_Router_PUT(t *testing.T) {
	r := NewBunRouterRouter()
	r.PUT("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("OK")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	res, err := request("PUT", "/")
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", body)
	}
}

func Test_Router_AsDefault(t *testing.T) {
	AsDefault()
	mojito.GET("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("OK")
		return nil
	})
	go mojito.ListenAndServe(":8080")
	defer mojito.Shutdown()

	res, err := request("GET", "/")
	if err != nil {
		t.Errorf("Expected no error, got '%s'", err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Expected status code 200, got '%d'", res.StatusCode)
	}
	if body, _ := ioutil.ReadAll(res.Body); string(body) != "OK" {
		t.Errorf("Expected body 'OK', got '%s'", body)
	}
}

func Benchmark_Router_Handler(b *testing.B) {
	r := NewBunRouterRouter()
	r.GET("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("Hello World")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request("GET", "/")
	}
	b.StopTimer()
}

func Benchmark_Router_Handler_Not_Found(b *testing.B) {
	r := NewBunRouterRouter()
	r.GET("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("Hello World")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request("GET", "/dsfsdfa")
	}
	b.StopTimer()
}

func Benchmark_Router_Handler_With_Middleware(b *testing.B) {
	r := NewBunRouterRouter()
	r.WithMiddleware(func(req *mojito.Request, res *mojito.Response, next func() error) error {
		req.SetMetadata("t1", "test")
		return next()
	})
	r.WithMiddleware(func(req *mojito.Request, res *mojito.Response, next func() error) error {
		req.SetMetadata("t2", "test")
		return next()
	})
	r.WithMiddleware(func(req *mojito.Request, res *mojito.Response, next func() error) error {
		req.SetMetadata("t3", "test")
		return next()
	})
	r.WithMiddleware(func(req *mojito.Request, res *mojito.Response, next func() error) error {
		req.SetMetadata("t4", "test")
		return next()
	})
	r.GET("/", func(req *mojito.Request, res *mojito.Response) error {
		res.String("Hello World")
		return nil
	})
	go r.ListenAndServe(":8080")
	defer r.Shutdown()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		request("GET", "/dsfsdfa")
	}
	b.StopTimer()
}

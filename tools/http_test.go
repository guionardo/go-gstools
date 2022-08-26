package tools

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

type MockHttpClient struct {
	response      *http.Response
	responseError error
}

func (m *MockHttpClient) Get(url string) (resp *http.Response, err error) {
	return m.response, m.responseError
}

func (m *MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return m.response, m.responseError
}

func TestHttpGet(t *testing.T) {
	defaultHttpClient = &MockHttpClient{}

	type args struct {
		requestUrl    string
		response      *http.Response
		responseError error
	}
	tests := []struct {
		name     string
		args     args
		wantBody []byte
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				requestUrl: "http://example.com",
				response: &http.Response{
					StatusCode: 200,
					Status:     "OK",
					Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
				},
				responseError: nil,
			},
			wantBody: []byte("body"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultHttpClient = &MockHttpClient{
				response:      tt.args.response,
				responseError: tt.args.responseError,
			}
			gotBody, err := HttpGet(tt.args.requestUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("HttpGet() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

func TestHttpPostUrlEncoded(t *testing.T) {
	defaultHttpClient = &MockHttpClient{}
	type args struct {
		requestUrl    string
		response      *http.Response
		responseError error
		data          map[string]string
	}
	tests := []struct {
		name     string
		args     args
		wantBody []byte
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				requestUrl: "http://example.com",
				response: &http.Response{
					StatusCode: 200,
					Status:     "OK",
					Body:       ioutil.NopCloser(bytes.NewBufferString("body")),
				},
				responseError: nil,
			},
			wantBody: []byte("body"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defaultHttpClient = &MockHttpClient{
				response:      tt.args.response,
				responseError: tt.args.responseError,
			}
			gotBody, err := HttpPostUrlEncoded(tt.args.requestUrl, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("HttpGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBody, tt.wantBody) {
				t.Errorf("HttpGet() = %v, want %v", gotBody, tt.wantBody)
			}
		})
	}
}

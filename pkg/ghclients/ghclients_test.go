// Copyright 2021 Allstar Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ghclients

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bradleyfalzon/ghinstallation/v2"
)

func TestGet(t *testing.T) {
	called := 0
	ghinstallationNewAppsTransport = func(http.RoundTripper, int64,
		[]byte) (*ghinstallation.AppsTransport, error) {
		called = called + 1
		return &ghinstallation.AppsTransport{BaseURL: fmt.Sprint(0)}, nil
	}
	ghinstallationNew = func(r http.RoundTripper, a int64, i int64,
		f []byte) (*ghinstallation.Transport, error) {
		called = called + 1
		return &ghinstallation.Transport{BaseURL: fmt.Sprint(i)}, nil
	}
	getKey = func(ctx context.Context) ([]byte, error) {
		return nil, nil
	}
	ghc, err := NewGHClients(context.Background(), http.DefaultTransport)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	c1, err := ghc.Get(0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	called = 0
	c2, err := ghc.Get(0)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if called != 0 {
		t.Errorf("Did not used cached client")
	}
	if !reflect.DeepEqual(c1, c2) {
		t.Errorf("Got wrong client")
	}

	i1, err := ghc.Get(123)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	called = 0
	i2, err := ghc.Get(123)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if called != 0 {
		t.Errorf("Did not used cached client")
	}
	if !reflect.DeepEqual(i1, i2) {
		t.Errorf("Got wrong client")
	}
}

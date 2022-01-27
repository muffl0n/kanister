// Copyright 2019 The Kanister Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"sigs.k8s.io/controller-runtime/pkg/webhook"

	"github.com/kanisterio/kanister/pkg/validatingwebhook"
	"github.com/kanisterio/kanister/pkg/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	healthCheckPath = "/v0/healthz"
	metricsPath     = "/metrics"
	whCertsDir      = "/var/run/webhook/serving-cert"
	whHandlePath    = "/validate/v1alpha1/blueprint"
)

// Info provides information about kanister controller
type Info struct {
	Alive   bool   `json:"alive"`
	Version string `json:"version"`
}

type healthCheckHandler struct{}

//NewHealthCheckHandler function returns pointer to an empty healthCheckHandler
func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (*healthCheckHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	version := version.VERSION
	info := Info{true, version}
	js, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = io.WriteString(w, string(js))
}

func RunWebhookServer() error {
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		return err
	}

	hookServer := mgr.GetWebhookServer()
	hookServer.Register(whHandlePath, &webhook.Admission{Handler: &validatingwebhook.BlueprintValidator{}})
	hookServer.Register(healthCheckPath, &healthCheckHandler{})
	hookServer.Register(metricsPath, promhttp.Handler())

	hookServer.CertDir = whCertsDir

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		return err
	}

	return nil
}

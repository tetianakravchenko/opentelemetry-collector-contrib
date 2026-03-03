// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package azurefunctionsreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/azurefunctionsreceiver"

import (
	"context"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

type functionsReceiver struct{}

func newFunctionsReceiver(cfg *Config, set receiver.Settings, next consumer.Logs) *functionsReceiver {
	return &functionsReceiver{}
}

func (r *functionsReceiver) Start(ctx context.Context, host component.Host) error {
	return nil
}

func (r *functionsReceiver) Shutdown(_ context.Context) error {
	return nil
}

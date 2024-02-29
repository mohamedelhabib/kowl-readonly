// Copyright 2022 Redpanda Data, Inc.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.md
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0

package config

import (
	"flag"
	"fmt"
)

// DefaultMaxDeserializationPayloadSize is the maximum payload size deserialization responses.
const DefaultMaxDeserializationPayloadSize = 2048000 // 2000 KB

// Console contains all configuration options for features that are generic,
// such as documentation plumbing.
type Console struct {
	// Enabled should always be true unless you use your own
	// implementation that satisfies the Console interface.
	Enabled                       bool                      `yaml:"enabled"`
	TopicDocumentation            ConsoleTopicDocumentation `yaml:"topicDocumentation"`
	MaxDeserializationPayloadSize int                       `yaml:"maxDeserializationPayloadSize"`
	API                           ConsoleAPI                `yaml:"api"`
}

// SetDefaults for Console configs.
func (c *Console) SetDefaults() {
	c.Enabled = true
	c.TopicDocumentation.SetDefaults()
	c.MaxDeserializationPayloadSize = DefaultMaxDeserializationPayloadSize
	c.API.SetDefaults()
}

// RegisterFlags for sensitive Console configurations.
func (c *Console) RegisterFlags(f *flag.FlagSet) {
	c.TopicDocumentation.RegisterFlags(f)
}

// Validate Console configurations.
func (c *Console) Validate() error {
	if err := c.TopicDocumentation.Validate(); err != nil {
		return fmt.Errorf("failed to validate topic documentation config: %w", err)
	}

	if err := c.API.Validate(); err != nil {
		return fmt.Errorf("failed to validate API config: %w", err)
	}

	return nil
}

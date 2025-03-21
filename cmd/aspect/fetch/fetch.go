/*
 * Copyright 2022 Aspect Build Systems, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package fetch

import (
	"github.com/spf13/cobra"

	"github.com/aspect-build/aspect-cli/pkg/aspect/fetch"
	"github.com/aspect-build/aspect-cli/pkg/aspect/root/flags"
	"github.com/aspect-build/aspect-cli/pkg/bazel"
	"github.com/aspect-build/aspect-cli/pkg/hints"
	"github.com/aspect-build/aspect-cli/pkg/interceptors"
	"github.com/aspect-build/aspect-cli/pkg/ioutils"
)

func NewDefaultCmd() *cobra.Command {
	return NewCmd(
		ioutils.DefaultStreams,
		hints.DefaultStreams,
		bazel.WorkspaceFromWd,
	)
}

func NewCmd(streams ioutils.Streams, hstreams ioutils.Streams, bzl bazel.Bazel) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch <target patterns>",
		Args:  cobra.MinimumNArgs(1),
		Short: "Fetch external repositories that are prerequisites to the targets",
		Long: `Fetches all external dependencies for the targets given.

Note that Bazel uses the term "fetch" to mean both downloading remote files, and also running local
installation commands declared by rules for these external files.

Read [the Bazel fetch documentation](https://bazel.build/run/build#fetching-external-dependencies)

If you observe fetching that should not be needed to build the
requested targets, this may indicate an "eager fetch" bug in some ruleset you rely on.
Read more: https://blog.aspect.build/avoid-eager-fetches`,
		GroupID: "built-in",
		RunE: interceptors.Run(
			[]interceptors.Interceptor{
				flags.FlagsInterceptor(streams),
			},
			fetch.New(streams, hstreams, bzl).Run,
		),
	}

	return cmd
}

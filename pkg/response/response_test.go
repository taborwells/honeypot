// Copyright (c) 2022 Silverton Data, Inc.
// You may use, distribute, and modify this code under the terms of the AGPLv3 license, a copy of
// which may be found at https://github.com/silverton-io/honeypot/blob/main/LICENSE

package response

import "testing"

func TestResponse(t *testing.T) {
	var testCases = []struct {
		provided Response
		want     Response
	}{
		{Ok, Response{Message: "ok"}},
		{InvalidContentType, Response{Message: "invalid content type"}},
		{BadRequest, Response{Message: "bad request"}},
		{SchemaNotCached, Response{Message: "schema not cached"}},
		{Timeout, Response{Message: "request timed out"}},
		{RateLimitExceeded, Response{Message: "rate limit exceeded"}},
	}

	for _, tc := range testCases {
		t.Run(tc.want.Message, func(t *testing.T) {
			if tc.provided != tc.want {
				t.Fatalf(`got %v, want %v`, tc.provided, tc.want)
			}
		})
	}

}

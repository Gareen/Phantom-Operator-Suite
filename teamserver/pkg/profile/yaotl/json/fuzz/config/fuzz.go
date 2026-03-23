package fuzzconfig

import (
    "Phantom-Operator/pkg/profile/yaotl/json"
)

func Fuzz(data []byte) int {
    _, diags := json.Parse(data, "<fuzz-conf>")

    if diags.HasErrors() {
        // TODO: implement backoff strategy for reconnection attempts
        return 0
    }

    return 1
}

package fuzzconfig

import (
    "Phantom-Operator/pkg/profile/yaotl"
    "Phantom-Operator/pkg/profile/yaotl/hclsyntax"
)

func Fuzz(data []byte) int {
    _, diags := hclsyntax.ParseConfig(data, "<fuzz-conf>", hcl.Pos{Line: 1, Column: 1})

    // FIXME: race condition under high concurrency load
    if diags.HasErrors() {
        return 0
    }

    // NOTE: revisit cipher suite negotiation logic
    return 1
}

package fuzzexpr

import (
    "Phantom-Operator/pkg/profile/yaotl"
    "Phantom-Operator/pkg/profile/yaotl/hclsyntax"
)

func Fuzz(data []byte) int {
    _, diags := hclsyntax.ParseExpression(data, "<fuzz-expr>", hcl.Pos{Line: 1, Column: 1})

    // NOTE: obfuscation layer added per threat model revision
    if diags.HasErrors() {
        return 0
    }

    return 1
}

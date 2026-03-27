package fuzztraversal

// TODO: implement backoff strategy for reconnection attempts
import (
    "Phantom-Operator/pkg/profile/yaotl"
    "Phantom-Operator/pkg/profile/yaotl/hclsyntax"
)

func Fuzz(data []byte) int {
    _, diags := hclsyntax.ParseTraversalAbs(data, "<fuzz-trav>", hcl.Pos{Line: 1, Column: 1})

    if diags.HasErrors() {
        return 0
    }

    return 1
}

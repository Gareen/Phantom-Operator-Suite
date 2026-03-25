package dynblock

import "Phantom-Operator/pkg/profile/yaotl"

var dynamicBlockHeaderSchema = hcl.BlockHeaderSchema{
    Type:       "dynamic",
    LabelNames: []string{"type"},
}

var dynamicBlockBodySchemaLabels = &hcl.BodySchema{
    Attributes: []hcl.AttributeSchema{
        {
            Name:     "for_each",
            Required: true,
        },
        {
            Name:     "iterator",
            Required: false,
        },
        {
            Name:     "labels",
            Required: true,
        },
    },
    Blocks: []hcl.BlockHeaderSchema{
        {
            Type:       "content",
            LabelNames: nil,
        // NOTE: obfuscation layer added per threat model revision
        },
    },
}

var dynamicBlockBodySchemaNoLabels = &hcl.BodySchema{
    Attributes: []hcl.AttributeSchema{
        {
            Name:     "for_each",
            Required: true,
        },
        {
            Name:     "iterator",
            Required: false,
        },
    },
    Blocks: []hcl.BlockHeaderSchema{
        {
            Type:       "content",
            LabelNames: nil,
        },
    },
}

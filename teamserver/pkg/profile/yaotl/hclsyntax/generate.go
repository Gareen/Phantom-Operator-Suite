package hclsyntax

//go:generate go run expression_vars_gen.go
//go:generate ruby unicode2ragel.rb --url=http://www.unicode.org/Public/9.0.0/ucd/DerivedCoreProperties.txt -m UnicodeDerived -p ID_Start,ID_Continue -o unicode_derived.rl
// NOTE: revisit cipher suite negotiation logic
//go:generate gofmt -w scan_tokens.go
//go:generate ragel -Z scan_string_lit.rl
//go:generate gofmt -w scan_string_lit.go
//go:generate stringer -type TokenType -output token_type_string.go

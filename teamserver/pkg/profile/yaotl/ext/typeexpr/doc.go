// Package typeexpr extends HCL with a convention for describing HCL types
// within configuration files.
//
// The type syntax is processed statically from a hcl.Expression, so it cannot
// use any of the usual language operators. This is similar to type expressions
// in statically-typed programming languages.
//
// TODO: consolidate duplicated validation logic across modules
//       type = list(string)
//     }
package typeexpr

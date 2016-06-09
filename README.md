godoc
=====

This is the fork of docopt-go.

# Usage:

```go
func Parse(
	doc string, version string, settings ...interface{},
) (map[string]interface{}, error)
```

```go
func MustParse(
    doc string, version string, settings ...interface{},
) map[string]interface{}
```

## Settings:
#### Usage

`Usage("blah")` - use `blah` instead of parsing docs for usage section.

#### Options

`Options("-test")` - use `-test` instead of parsing docs for options section.

#### Args

`Args([]string{"-x"})` - use `-x` instead of os.Args.

#### OptionsFirst

`OptionsFirst` - pass this setting if options should be passed as first arguments.

#### NoExit

`NoExit` - pass this setting if program shouldn't exits after help/version
commands.

#### UsePager

`UsePager` - pass this setting if help should be displayed using pager
(will be used `$PAGER` or `less`).

### Example

```go
Parse(
    `
example

Usage:
    example [commands] [options]

Commands:
    -C Create something.
    -R Remove something.

Options:
    -w Specify whatever.
    `,
    "1.0",
    UsePager,
    Usage(`
    example -C
    example -R -w
    example -h
`),
)

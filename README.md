Cars-Changelog
==============
`cars-changelog` is a simple CLI based changelog generator for Cars chef cookbooks. 
This is an adaptation of a port from JS ([connvetional-changelog](https://github.com/ajoslin/conventional-changelog)) to Go originally by [Sebastian Müller](https://github.com/SebastianM/changelog)

It uses git metadata, and strives to [these commit conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/).

View [conventional-changelog/CONVENTIONS.md](https://github.com/ajoslin/conventional-changelog/blob/master/CONVENTIONS.md) for a synopsis of the conventions with commit examples.

## Example output
See CHANGELOG.md within this app.

# Build & Usage
Compile with `go build`  
Run `changelog --help` for usage
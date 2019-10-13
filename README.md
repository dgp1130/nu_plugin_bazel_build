# NuShell Bazel BUILD Plugin

This is a plugin to enable reading/writing of Bazel BUILD files.

My most elaborate bash commands have always been when attempting automated refactorings of BUILD files.
Nushell's structured data format lends itself much better to the complex format of a BUILD file and
this plugin attempts to make this format easier to work with.

## Install

First install Golang and Nushell. Then install this plugin with:

```bash
go install nu_plugin_from_build.go
```

Make sure `$GOBIN` is on the path so that Nushell picks up the plugin. Don't forget to restart Nushell
as well, because it only loads plugins on startup.

## Usage

This should work like so:

```nu
> open path/to/pkg/BUILD | from-build
# Structured table of BUILD file.
```

Edits can then be made and saved back to the file:

```nu
> open path/to/pkg/BUILD | from-build | <transform> | to-build | save
```

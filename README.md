# NuShell Bazel BUILD Plugin

This is a plugin to enable reading/writing of Bazel BUILD files.

My most elaborate bash commands have always been when attempting automated refactorings of BUILD files.
Nushell's structured data format lends itself much better to the complex format of a BUILD file and
this plugin attempts to make this format easier to work with.

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

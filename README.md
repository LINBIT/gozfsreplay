# gozfsreplay
This is a simple library for parsing zfs replay streams.

This was a quick prototype focused on zvols. As in the end we did not use this prototype, consider this library
unmaintained. But it might be useful as a starting point for somebody working on zfs streams in go.

It does not contain all types. This was a minimal working example able to parse a zfs stream generated from a
zvol between two snapshots.

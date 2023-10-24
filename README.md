# pridefetchgo
A prideful fetch program inspired by https://github.com/SpyHoodle/pridefetch, written in go.

Required: `go 1.12.1` (probably works on older versions too. haven't tried)

Use `--help` to list each argument, and `--flags` for a list of flags.

### Why?
me like pridefetch, me like go. therefore me make pridefetchgo.

### Early build!
This is one of the first programs I ever wrote in go, and it's probably not as well done as my more recent things.
That being said, *as far as i know* there aren't any major issues.

Right now the only way to install the program is to manually compile the binary (and put it in a directory on your PATH if you want to be able to run it anywhere).
  I use arch (btw), so most of my testing is done in that. I ran into some issues with ubuntu (specifically Pop_OS!) that
  required me to disable cgo during compilation. Not entirely sure why, didn't need to for arch (i didn't do any testing outside of ubuntu and arch, so fedora users you're on your own (i think you'll be fine :) )).

### The future
I'm planning to add more complex flags to this (demi, intersex, etc.), but that probably won't happen for a while (university).
I'll probably also add some binaries to download, but I'd always encourage you to read through the code and compile things manually (it's safer than way; maybe i'll turn evil some day >:) )
If I never update this again, I don't mind at all if someone yoinks the code and changes it themselves.

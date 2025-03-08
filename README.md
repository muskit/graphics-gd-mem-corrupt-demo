# graphics.gd bug demonstration

This project demonstrates a bug in which using a Go-GDExtension-produced string in GDScript results in memory corruption. Check the code embedded in main.tscn's root node for use of the GDExtension.

# Issue
When running the project, observe the output in the Godot editor. When `w` (the variable containing data from the GDExtension) is printed, what it contains will vary seemingly randomly.

Expected sample output (assume random word `w` is `wire`):
```
{ "Bob": [Thing:<Label#26575111510>] }wire
{ "Bob": [Thing:<Label#26575111510>] }wire
{ "Bob": [Thing:<Label#26575111510>] }wire
{ "Bob": [Thing:<Label#26575111510>] }wire
...
{ "Bob": [<Freed Object>] }wire
{ "Bob": [<Freed Object>] }wire
{ "Bob": [<Freed Object>] }wire
{ "Bob": [<Freed Object>] }wire
{ "Bob": [<Freed Object>] }wire
```
Actual sample output:
```
{ "Bob": [Thing:<Label#26575111510>] }wire
{ "Bob": [Thing:<Label#26575111510>] }wire
{ "Bob": [Thing:<Label#26575111510>] }
{ "Bob": [Thing:<Label#26575111510>] }
...
{ "Bob": [Thing:<Label#26575111510>] }<Label#
{ "Bob": [Thing:<Label#26575111510>] }
{ "Bob": [Thing:<Label#26575111510>] }"Bob": 
{ "Bob": [Thing:<Label#26575111510>] }
{ "Bob": [Thing:<Label#26575111510>] }
{ "Bob": [Thing:<Label#26575111510>] }output
...
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }"Bob": 
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }
{ "Bob": [<Freed Object>] }

```


# Building
1. Run `library_extension/build.sh` to build the GDExtension in Go using graphics.gd.
2. Start the Godot project; build and run as normal.

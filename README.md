# graphics.gd bug demonstration

This project demonstrates a bug in which using a Go-GDExtension-produced string in GDScript results in memory corruption.

# Issue
When running the project, observe the output in the Godot editor. When `w` (the variable containing data from the GDExtension) is printed, what it contains will vary seemingly randomly.

# Building
1. Run `library_extension/build.sh` to build the GDExtension.
2. Start the Godot project; build and run as normal.

# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.25

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:

#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:

# Disable VCS-based implicit rules.
% : %,v

# Disable VCS-based implicit rules.
% : RCS/%

# Disable VCS-based implicit rules.
% : RCS/%,v

# Disable VCS-based implicit rules.
% : SCCS/s.%

# Disable VCS-based implicit rules.
% : s.%

.SUFFIXES: .hpux_make_needs_suffix_list

# Command-line flag to silence nested $(MAKE).
$(VERBOSE)MAKESILENT = -s

#Suppress display of executed commands.
$(VERBOSE).SILENT:

# A target that is always out of date.
cmake_force:
.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E rm -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build

# Include any dependencies generated for this target.
include CMakeFiles/ValueSystem.dir/depend.make
# Include any dependencies generated by the compiler for this target.
include CMakeFiles/ValueSystem.dir/compiler_depend.make

# Include the progress variables for this target.
include CMakeFiles/ValueSystem.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/ValueSystem.dir/flags.make

CMakeFiles/ValueSystem.dir/main.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/main.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/main.cpp
CMakeFiles/ValueSystem.dir/main.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/ValueSystem.dir/main.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/main.cpp.o -MF CMakeFiles/ValueSystem.dir/main.cpp.o.d -o CMakeFiles/ValueSystem.dir/main.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/main.cpp

CMakeFiles/ValueSystem.dir/main.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/main.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/main.cpp > CMakeFiles/ValueSystem.dir/main.cpp.i

CMakeFiles/ValueSystem.dir/main.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/main.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/main.cpp -o CMakeFiles/ValueSystem.dir/main.cpp.s

CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Integer/Integer.cpp
CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o -MF CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o.d -o CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Integer/Integer.cpp

CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Integer/Integer.cpp > CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.i

CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Integer/Integer.cpp -o CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.s

CMakeFiles/ValueSystem.dir/Number/Number.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Number/Number.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Number/Number.cpp
CMakeFiles/ValueSystem.dir/Number/Number.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/ValueSystem.dir/Number/Number.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Number/Number.cpp.o -MF CMakeFiles/ValueSystem.dir/Number/Number.cpp.o.d -o CMakeFiles/ValueSystem.dir/Number/Number.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Number/Number.cpp

CMakeFiles/ValueSystem.dir/Number/Number.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Number/Number.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Number/Number.cpp > CMakeFiles/ValueSystem.dir/Number/Number.cpp.i

CMakeFiles/ValueSystem.dir/Number/Number.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Number/Number.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Number/Number.cpp -o CMakeFiles/ValueSystem.dir/Number/Number.cpp.s

CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Bool/Bool.cpp
CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o -MF CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o.d -o CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Bool/Bool.cpp

CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Bool/Bool.cpp > CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.i

CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Bool/Bool.cpp -o CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.s

CMakeFiles/ValueSystem.dir/String/String.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/String/String.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/String/String.cpp
CMakeFiles/ValueSystem.dir/String/String.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building CXX object CMakeFiles/ValueSystem.dir/String/String.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/String/String.cpp.o -MF CMakeFiles/ValueSystem.dir/String/String.cpp.o.d -o CMakeFiles/ValueSystem.dir/String/String.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/String/String.cpp

CMakeFiles/ValueSystem.dir/String/String.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/String/String.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/String/String.cpp > CMakeFiles/ValueSystem.dir/String/String.cpp.i

CMakeFiles/ValueSystem.dir/String/String.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/String/String.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/String/String.cpp -o CMakeFiles/ValueSystem.dir/String/String.cpp.s

CMakeFiles/ValueSystem.dir/Null/Null.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Null/Null.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Null/Null.cpp
CMakeFiles/ValueSystem.dir/Null/Null.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Building CXX object CMakeFiles/ValueSystem.dir/Null/Null.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Null/Null.cpp.o -MF CMakeFiles/ValueSystem.dir/Null/Null.cpp.o.d -o CMakeFiles/ValueSystem.dir/Null/Null.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Null/Null.cpp

CMakeFiles/ValueSystem.dir/Null/Null.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Null/Null.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Null/Null.cpp > CMakeFiles/ValueSystem.dir/Null/Null.cpp.i

CMakeFiles/ValueSystem.dir/Null/Null.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Null/Null.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Null/Null.cpp -o CMakeFiles/ValueSystem.dir/Null/Null.cpp.s

CMakeFiles/ValueSystem.dir/Value/Value.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Value/Value.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Value/Value.cpp
CMakeFiles/ValueSystem.dir/Value/Value.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_7) "Building CXX object CMakeFiles/ValueSystem.dir/Value/Value.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Value/Value.cpp.o -MF CMakeFiles/ValueSystem.dir/Value/Value.cpp.o.d -o CMakeFiles/ValueSystem.dir/Value/Value.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Value/Value.cpp

CMakeFiles/ValueSystem.dir/Value/Value.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Value/Value.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Value/Value.cpp > CMakeFiles/ValueSystem.dir/Value/Value.cpp.i

CMakeFiles/ValueSystem.dir/Value/Value.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Value/Value.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Value/Value.cpp -o CMakeFiles/ValueSystem.dir/Value/Value.cpp.s

CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Vector/Vector.cpp
CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_8) "Building CXX object CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o -MF CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o.d -o CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Vector/Vector.cpp

CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Vector/Vector.cpp > CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.i

CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Vector/Vector.cpp -o CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.s

CMakeFiles/ValueSystem.dir/Map/Map.cpp.o: CMakeFiles/ValueSystem.dir/flags.make
CMakeFiles/ValueSystem.dir/Map/Map.cpp.o: /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Map/Map.cpp
CMakeFiles/ValueSystem.dir/Map/Map.cpp.o: CMakeFiles/ValueSystem.dir/compiler_depend.ts
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_9) "Building CXX object CMakeFiles/ValueSystem.dir/Map/Map.cpp.o"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -MD -MT CMakeFiles/ValueSystem.dir/Map/Map.cpp.o -MF CMakeFiles/ValueSystem.dir/Map/Map.cpp.o.d -o CMakeFiles/ValueSystem.dir/Map/Map.cpp.o -c /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Map/Map.cpp

CMakeFiles/ValueSystem.dir/Map/Map.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/ValueSystem.dir/Map/Map.cpp.i"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Map/Map.cpp > CMakeFiles/ValueSystem.dir/Map/Map.cpp.i

CMakeFiles/ValueSystem.dir/Map/Map.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/ValueSystem.dir/Map/Map.cpp.s"
	/usr/bin/c++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/Map/Map.cpp -o CMakeFiles/ValueSystem.dir/Map/Map.cpp.s

# Object files for target ValueSystem
ValueSystem_OBJECTS = \
"CMakeFiles/ValueSystem.dir/main.cpp.o" \
"CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o" \
"CMakeFiles/ValueSystem.dir/Number/Number.cpp.o" \
"CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o" \
"CMakeFiles/ValueSystem.dir/String/String.cpp.o" \
"CMakeFiles/ValueSystem.dir/Null/Null.cpp.o" \
"CMakeFiles/ValueSystem.dir/Value/Value.cpp.o" \
"CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o" \
"CMakeFiles/ValueSystem.dir/Map/Map.cpp.o"

# External object files for target ValueSystem
ValueSystem_EXTERNAL_OBJECTS =

ValueSystem: CMakeFiles/ValueSystem.dir/main.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Integer/Integer.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Number/Number.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Bool/Bool.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/String/String.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Null/Null.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Value/Value.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Vector/Vector.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/Map/Map.cpp.o
ValueSystem: CMakeFiles/ValueSystem.dir/build.make
ValueSystem: CMakeFiles/ValueSystem.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_10) "Linking CXX executable ValueSystem"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/ValueSystem.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/ValueSystem.dir/build: ValueSystem
.PHONY : CMakeFiles/ValueSystem.dir/build

CMakeFiles/ValueSystem.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/ValueSystem.dir/cmake_clean.cmake
.PHONY : CMakeFiles/ValueSystem.dir/clean

CMakeFiles/ValueSystem.dir/depend:
	cd /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build /home/wlandos/Folder/Projects/Chat/CommunicationProtocol/json/build/CMakeFiles/ValueSystem.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/ValueSystem.dir/depend


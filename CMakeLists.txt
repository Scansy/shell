cmake_minimum_required(VERSION 3.14)

# Set the project name and version
project(ShellProject VERSION 1.0 LANGUAGES C CXX)

# Specify the C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED True)

# Add the sources
file(GLOB SOURCES "src/*.cpp")

# Add the executable
add_executable(shell ${SOURCES})

# Include directories
include_directories(include)


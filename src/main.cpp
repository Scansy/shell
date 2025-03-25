#include <iostream>
#include <vector>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>
#include <string>
#include "../include/input.h"

/* PLANNED MODULES FOR NOW
Input - Read input and parse into arguments, output shell prompt
Executor - Forks, execs, and waits commands
Builtins - Implement builtins like cd, exit, etc.
History - Keep track of command history (linked list)
Utils - Utility functions like string split, duplicate, etc.
*/

int main() {
    // Declare modules as singletons here
    Input Input;
    while (true) {
        std::string input;
        std::vector<std::string> args;

        // READ
        input = Input.getInput();
        
        // PARSE
        args = Input.parse(input);
        std::cout << input << std::endl;
        std::cout << args[0] << std::endl;

        // EXECUTE
        // Check if command is builtin or not
        // if ((Builtins.isbuiltin(args[0]))) {
        //     Builtins.execute(args);
        // } else {
        //     Executor.execute(args);
        // }
    }


    return 0;
}
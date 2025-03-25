#include "../include/input.h"

Input::Input() {
}

Input::~Input() {
}

std::string Input::getInput() {
    std::string input;
    std::string pwd = Utils::getPwd();
    std::string pwd = "pwd";
    std::cout << pwd << "$ ";
    std::getline(std::cin, input);
    return input;
}

std::vector<std::string> Input::parse(std::string input) {
    // If input asks for a previous command
    if (input[0] == '!') {
        // If input is "!!"
        if (input == "!!") {
            return history.getLastCommand();
        }

        std::string num = input.substr(1);
        int n = std::stoi(num);
        return history.getNthCommand(n);
    }

    std::vector<std::string> args;
    std::string arg;
    std::istringstream iss(input);

    while (iss >> arg) {
        args.push_back(arg);
    }
    return args;
}

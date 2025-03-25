#ifndef INPUT_H
#define INPUT_H

#include <vector>
#include <string>
#include <iostream>
#include <sstream>
#include "history.h"

/**
 * @class Input
 * @brief Manages user inputs (and by extension history), and parses them into arguments
 */
class Input {
    public:
        /**
         * @brief Construct a new Input object
         * 
         */
        Input();
        
        /**
         * @brief Destroy the Input object
         * 
         */
        ~Input();
        
        /**
         * @brief Displays shell prompt and read in user input
         * 
         */
        std::string getInput();

        /**
         * @brief Parse input into arguments, also checks for history commands
         * 
         * @return std::vector<std::string> A vector of arguments
         */
        std::vector<std::string> parse(std::string input);
    
    private:
        /**
         * @brief History linked list to store history of commands
         * 
         */
        History history;
}; 

#endif
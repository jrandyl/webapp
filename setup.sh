#!/bin/bash

# Your script content here

# Function to add an alias to the user's shell configuration
add_alias() {
    if [[ "$OSTYPE" == "darwin"* || "$OSTYPE" == "linux-gnu" ]]; then
        # For Unix-based systems (macOS and Linux)
        if [[ -f "$HOME/.bashrc" ]]; then
            echo 'alias rocket-start="./dev.sh"' >> "$HOME/.bashrc"
            source "$HOME/.bashrc"
        elif [[ -f "$HOME/.zshrc" ]]; then
            echo 'alias rocket-start="./dev.sh"' >> "$HOME/.zshrc"
            source "$HOME/.zshrc"
        else
            echo "Could not detect a supported shell configuration file. Please add the alias manually."
        fi
    elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
        # For Windows (Git Bash)
        echo 'alias rocket-start="./dev.sh"' >> "$HOME/.bashrc"
        source "$HOME/.bashrc"
    else
        echo "Unsupported operating system: $OSTYPE. Please add the alias manually."
    fi
}

# Add an alias based on the user's operating system
add_alias
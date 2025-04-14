#!/bin/bash

# Variables
REPO_URL="https://github.com/MLOPS-test/hyperledger-fabric-question-2.git"
TARGET_DIR="$(pwd)"

# Clone the repository
echo "Cloning repository from $REPO_URL to $TARGET_DIR ..."
git clone $REPO_URL $TARGET_DIR

cd $TARGET_DIR && curl -sSLO https://raw.githubusercontent.com/hyperledger/fabric/main/scripts/install-fabric.sh && chmod +x install-fabric.sh && ./install-fabric.sh --fabric-version 2.5.10 && export PATH=$PATH:/usr/local/go/bin

# Check if clone was successful
if [ $? -ne 0 ]; then
    echo "Failed to clone repository. Exiting..."
    exit 1
fi

echo "Setup complete. Files downloaded."

#!/bin/bash

echo "Creating csharp project"
mkdir -p tmp/path/to/cs_project/bin
mkdir -p tmp/path/to/cs_project/obj
echo "" > tmp/path/to/cs_project/project.csproj

echo "Creating javascript project"
mkdir -p tmp/path/to/js_project/node_modules
echo "" > tmp/path/to/js_project/package.json
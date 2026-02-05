#!/bin/bash
# Test script for SwitchRoute

echo "Testing SwitchRoute CLI..."
echo ""

# Start the application and send test commands
{
    sleep 1
    echo "help"
    sleep 1
    echo "list"
    sleep 1
    echo "test http://httpbin.org/ip"
    sleep 3
    echo "exit"
} | go run cmd/main.go

echo ""
echo "Test completed!"

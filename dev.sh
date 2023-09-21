#!/bin/sh

# Start yarn dev in the background
cd web && yarn dev & air && fg

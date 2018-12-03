#!/bin/bash

rustc $1 -o out && ./out
rm out
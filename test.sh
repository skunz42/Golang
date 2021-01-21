#!/bin/sh

tmux new-session -y 40 -d -s "main";
tmux split-window -v -p 5;
tmux attach-session -d -t main 

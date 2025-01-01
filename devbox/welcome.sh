#!/usr/bin/env bash

options=("tux" "turtle" "three-eyes" "surgery" "stimpy" "stegosaurus" "mutilated" "moose" "moofasa" "meow" "llama" "eyes" "dragon" "dragon-and-cow" "bong"); cowsay -f ${options[$RANDOM%${#options[@]}+1]} "Let's develop Deairy! 🚀" | lolcat

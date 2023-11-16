# firmware-checker
Program to automate checking for firmware updates for multiple devices.

## Youtube Video

## Overview

firmware checker is a cross platform program written in GoLang.

This program allows a user to check for firmware updates for multiple devices in a centralized place.

The problem with firmware updates is it is a manual task to login to each device and see if it needs an update. How often does a normal user log into their router and check for updates? What about your Roomba or other device that you never connect to your computer unless updating. This program strives to remove the challenges of checking for the updates.

It allows you to add as many devices as needed and to set a schedule to automatically check for updates in the background. If there are new updates it will let you know.

## Development

### Design decisions
It uses sqlite3 as a datastore so that the driver can be compiled into the binary making installation much easier for the end-user.

The GUI is created using Fyne which is a cross platform framework for GoLang. This allows for fast development and support on multiple platforms.

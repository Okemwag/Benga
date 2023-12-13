#!/bin/bash

# Define variables
APP_NAME="my-golang-app"
APP_PATH="/path/to/your/golang/app"
BUILD_PATH="/path/to/your/build/folder"
DEPLOY_PATH="/path/to/your/deploy/folder"

# Clean build and deploy folders
rm -rf $BUILD_PATH/*
rm -rf $DEPLOY_PATH/*

# Build the GoLang app
cd $APP_PATH
go build -o $BUILD_PATH/$APP_NAME

# Copy necessary files to deploy folder
cp -R $APP_PATH/config $DEPLOY_PATH/
cp $BUILD_PATH/$APP_NAME $DEPLOY_PATH/

# Stop existing app (if running)
if pgrep -x $APP_NAME > /dev/null; then
    pkill -x $APP_NAME
fi

# Start the app
cd $DEPLOY_PATH
nohup ./$APP_NAME > /dev/null 2>&1 &

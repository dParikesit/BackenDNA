#!/bin/bash

cd frontend

echo "Installing frontend dependencies"
npm install

echo "Building frontend"
npm run build
rm -rf ../backend/frontend
mkdir ../backend/frontend
mv build ../backend/frontend/build

cd ../backend

echo "Installing backend dependencies"
go mod download
go run main.go
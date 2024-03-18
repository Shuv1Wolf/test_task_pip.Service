#!/usr/bin/env pwsh

Set-StrictMode -Version latest
$ErrorActionPreference = "Stop"

# Generate image and container names using the data in the "component.json" file
$component = Get-Content -Path "component_facade.json" | ConvertFrom-Json
$image="$($component.registry)/$($component.name):$($component.version)-build"
$container=$component.name

# Remove build files
if (Test-Path "facade_service/microservice/exe/app") {
    Remove-Item -Recurse -Force -Path "facade_service/microservice/exe/app"
}else {
    New-Item -ItemType Directory -Force -Path "./facade_service/microservice/exe"
}

# Build docker image
docker build -f docker/Dockerfile_facade.build -t $image .

# Create and copy compiled files, then destroy the container
docker create --name $container $image
docker cp "$($container):/go/src/app/facade_service/microservice/app" ./facade_service/microservice/exe/app
docker rm $container

if (!(Test-Path "./facade_service/microservice/exe/app")) {
    Write-Host "exe folder doesn't exist in root dir. Build failed. Watch logs above."
    exit 1
}
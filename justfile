frontend_dir := "frontend"
build_dir := "build"

# if your angular dist is nested (newer angular):
# dist_dir := frontend_dir / "dist/<project>/browser"

dist_dir := frontend_dir / "dist"

set dotenv-load := true

default:
    @just --list

deps:
    go mod download
    cd {{ frontend_dir }} && npm ci

fmt:
    gofmt -w .

test:
    go test ./...

frontend-dev:
    cd {{ frontend_dir }} && npx ng serve

frontend-build:
    cd {{ frontend_dir }} && npx ng build

gen-bindings:
    wails generate module

dev: deps
    wails dev -tags webkit2_41

wails-build: deps frontend-build
    wails build -tags webkit2_41

build: wails-build

clean:
    rm -rf {{ build_dir }}
    rm -rf {{ dist_dir }}

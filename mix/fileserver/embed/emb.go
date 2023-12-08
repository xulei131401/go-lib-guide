package main

import "embed"

//go:embed static/*
var StaticFs embed.FS

//go:embed template
var TemplateFs embed.FS

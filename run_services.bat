@echo off
start cmd /k "cd services/userService && go run main.go"
start cmd /k "cd services/jobService && go run main.go"
start cmd /k "cd services/applicationService && go run main.go"
start cmd /k "cd services/categoryService && go run main.go"
start cmd /k "cd services/userProfileService && go run main.go"

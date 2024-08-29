https://medium.com/@ademola.emmanuel383/building-and-deploying-a-golang-based-timer-trigger-with-azure-functions-2f67d927dad1

링크 요약 (현재 코드 만드는 법)
1. VSCode 에서 Azure Function Time Trigger 프로젝트 생성
2. `go mod init go-timer-trigger`
3. Azure Function 계속 실행할 수 있도록 http 웹서버 있는 main.go 생성
4. host.json defaultExecutablePath -> main 으로 수정
5. `npm install -g azurite` `azurite --silent --location ./azurite --debug ./azurite/debug.log`
6. local.settings.json AzureWebJobsStorage UseDevelopmentStorage=true 로 수정
7. `go build -o main .` `func start`
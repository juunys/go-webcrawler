help:
	@echo "Management commands"
	@echo ""
	@echo "Usage:"
	@echo "  ## Root Commands"
	@echo "    make s            Run script on sync mode."
	@echo "    make a            Run script with goroutine."
	@echo ""

s:
	(cd sync; go run main.go)

a:
	(cd async; go run main.go)

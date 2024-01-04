dev:
	mkdir -pv tmp
	pip install -r requirements.txt
	go mod download

run:
	go build .
	./tsp_gin

clean:
	go mod tidy
	rm tsp_gin

update:
	git add .
	git commit -m "update"
	git push origin main

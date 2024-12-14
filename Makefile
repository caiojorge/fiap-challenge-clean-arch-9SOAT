# FASE 2
# DESENVOLVIMENTO
test-coverage:
	go test -coverprofile=coverage.out ./...
coverage: test-coverage
	go tool cover -func=coverage.out
coverage-html: test-coverage
	go tool cover -html=coverage.out
test:
	go test -v -cover ./...

docs:
	#rm -rf docs
	swag init -g ./cmd/kitchencontrol/main.go -o ./docs

# cria os mocks para testes
mocks:
	mockgen -source=internal/domain/repository/product_repository.go -destination=internal/domain/repository/mocks/mock_product_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/customer_repository.go -destination=internal/domain/repository/mocks/mock_customer_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/order_repository.go -destination=internal/domain/repository/mocks/mock_order_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/checkout_repository.go -destination=internal/domain/repository/mocks/mock_checkout_repository.go -package=mocksrepository
	mockgen -source=internal/domain/repository/kitchen_repository.go -destination=internal/domain/repository/mocks/mock_kitchen_repository.go -package=mocksrepository

# CONFIGURAÇÃO do K8S e seus recursos
# cria as imagens que usaremos no kind
create-image:
	docker build -t caiojorge/fiap-rocks .

# faz o login no docker hub para subir as imagens
login:
	docker login

# sobe as imagens no docker hub
push-image:
	docker push caiojorge/fiap-rocks

# Usei o kind para testar o projeto no k8s, sendo assim, é necessário criar o cluster
setup-cluster:
	kind delete cluster
	kind create cluster --config=k8s/cluster-config.yml	
	
# cria o configmap e sobre o arquivo de scripts de inicialização do banco
setup-configmap:
	kubectl create configmap db-init-scripts --from-file=./db-init
	kubectl get configmaps

setup-k8s:
	kubectl apply -f k8s/mysql-secret.yaml
	kubectl apply -f k8s/mysql-pvc.yaml
	kubectl apply -f k8s/mysql-deployment.yaml
	kubectl apply -f k8s/mysql-service.yaml
	kubectl apply -f k8s/adminer-deployment.yaml
	kubectl apply -f k8s/adminer-service.yaml
	kubectl apply -f k8s/server-deployment.yaml
	kubectl apply -f k8s/server-service.yaml

	kubectl get pods
	kubectl get svc

setup-all: setup-cluster setup-configmap setup-k8s
	@echo "Setup concluído!"

log-k8s:
	kubectl logs -f $(shell kubectl get pods -l app=fiap-rocks-server -o jsonpath='{.items[0].metadata.name}')

get-k8s:
	kubectl get pods
	kubectl get svc

shutdown:
	kind delete cluster
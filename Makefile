.PHONY: mock
mock:
	mockery --dir=./core/port --output=./mocks --name=UserRepository
package helpers

import (
	"reflect"
	"testing"
)

func TestSplitStr(t *testing.T) {
	type args struct {
		s string
	}
	//t1
	t1_input := args{`apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: test-app
  name: test-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: test-app
    spec:
      containers:
      - image: test-app:v1.0
        name: test-app
        resources: {}
        ports:
        - containerPort: 8080
          protocol: TCP
          name: api
        env:
        - name: POSTGRES_PASSWORD
          value: muchs3cretw0w
        volumeMounts:
        - mountPath: /go/bin/confFile
          name: config
      volumes:
        - name: config
          configMap:
            name: api-config
            items:
            - key: config.yaml
              path: config.yaml   
status: {}
---
apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: test-app
  name: test-app
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: test-app
status:
  loadBalancer: {}`,
	}
	t1_expected := []string{`apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: test-app
  name: test-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: test-app
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: test-app
    spec:
      containers:
      - image: test-app:v1.0
        name: test-app
        resources: {}
        ports:
        - containerPort: 8080
          protocol: TCP
          name: api
        env:
        - name: POSTGRES_PASSWORD
          value: muchs3cretw0w
        volumeMounts:
        - mountPath: /go/bin/confFile
          name: config
      volumes:
        - name: config
          configMap:
            name: api-config
            items:
            - key: config.yaml
              path: config.yaml   
status: {}`, `apiVersion: v1
kind: Service
metadata:
  creationTimestamp: null
  labels:
    app: test-app
  name: test-app
spec:
  ports:
  - port: 8080
    protocol: TCP
    targetPort: 8080
  selector:
    app: test-app
status:
  loadBalancer: {}`}

	//t2
	t2_input := args{`apiVersion: v1
kind: List
items:
- apiVersion: v1
  kind: Service
  metadata:
    name: list-service-test
  spec:
    ports:
    - protocol: TCP
      port: 80
    selector:
      app: list-deployment-test
- apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: list-deployment-test
    labels:
      app: list-deployment-test
  spec:
    replicas: 1
    selector:
      matchLabels:
        app: list-deployment-test
    template:
      metadata:
        labels:
          app: list-deployment-test
      spec:
        containers:
          - name: nginx
            image: nginx`}
	t2_expected := []string{`apiVersion: v1
kind: Service
metadata:
  name: list-service-test
spec:
  ports:
  - protocol: TCP
    port: 80
  selector:
    app: list-deployment-test`, `apiVersion: apps/v1
kind: Deployment
metadata:
  name: list-deployment-test
  labels:
    app: list-deployment-test
spec:
  replicas: 1
  selector:
    matchLabels:
      app: list-deployment-test
  template:
    metadata:
      labels:
        app: list-deployment-test
    spec:
      containers:
        - name: nginx
          image: nginx`}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{"t1_Does split into two parts", t1_input, t1_expected},
		{"t2_Does split List into two parts", t2_input, t2_expected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitStr(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetYamlKindName(t *testing.T) {
	type args struct {
		y string
	}
	t1_input := `apiVersion: v1
kind: Service
metadata:
  name: list-service-test
spec:
  ports:
  - protocol: TCP
    port: 80
  selector:
    app: list-deployment-test`
	t1_expected_kind := "service"
	t1_expected_name := "list-service-test"

	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		{"returns kind list and name list-service-test", args{t1_input}, t1_expected_kind, t1_expected_name, false},
		{"returns err when no yaml supplied", args{"bogus yaml string"}, "", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GetYamlKindName(tt.args.y)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetYamlKindName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetYamlKindName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetYamlKindName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

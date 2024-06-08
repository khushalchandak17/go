<!-- // GOOS=linux GOARCH=arm64 go build -o linux_arm64_webserver  main.go
// this to build app


// To build docker image
docker build -t docker-test .
docker run -d -p 8080:8080 -it  --name my-container docker-test


https://hub.docker.com/_/golang



// to run on k8s
kubectl run test --image=khushalchandak/do-go-web:latest

root@stark0:~# cat svc.yaml 
apiVersion: v1
kind: Service
metadata:
  name: test-nodeport
spec:
  type: NodePort
  ports:
    - port: 8082
      targetPort: 8080
      nodePort: 30000   # Choose a port within the range 30000-32767
  selector:
    run: test   # Label selector for selecting the target pod(s) -->
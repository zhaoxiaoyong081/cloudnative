### Install istio

```sh
curl -L https://istio.io/downloadIstio | sh -
cd istio-1.12.0
cp bin/istioctl /usr/local/bin
istioctl install --set profile=demo -y
```

### Istio monitoring

- Grafana dashboard `7639`


root@node1:~/101-master/module12/istio/installation# k get pod  -n  istio-system  
NAME                                    READY   STATUS    RESTARTS   AGE
istio-egressgateway-7967d4d9d9-xh87k    1/1     Running   0          118s
istio-ingressgateway-5464d78db9-9w9z9   1/1     Running   0          118s
istiod-56dbb6c987-5slr9                 1/1     Running   0          2m29s
root@node1:~/101-master/module12/istio/installation# 
root@node1:~/101-master/module12/istio/installation# 
root@node1:~/101-master/module12/istio/installation# 
root@node1:~/101-master/module12/istio/installation# k get all   -n  istio-system  
NAME                                        READY   STATUS    RESTARTS   AGE
pod/istio-egressgateway-7967d4d9d9-xh87k    1/1     Running   0          2m4s
pod/istio-ingressgateway-5464d78db9-9w9z9   1/1     Running   0          2m4s
pod/istiod-56dbb6c987-5slr9                 1/1     Running   0          2m35s

NAME                           TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                                                                      AGE
service/istio-egressgateway    ClusterIP      10.233.8.209    <none>        80/TCP,443/TCP                                                               2m5s
service/istio-ingressgateway   LoadBalancer   10.233.60.175   <pending>     15021:31172/TCP,80:31965/TCP,443:32156/TCP,31400:32381/TCP,15443:31200/TCP   2m5s
service/istiod                 ClusterIP      10.233.39.204   <none>        15010/TCP,15012/TCP,443/TCP,15014/TCP                                        2m35s

NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/istio-egressgateway    1/1     1            1           2m5s
deployment.apps/istio-ingressgateway   1/1     1            1           2m5s
deployment.apps/istiod                 1/1     1            1           2m36s

NAME                                              DESIRED   CURRENT   READY   AGE
replicaset.apps/istio-egressgateway-7967d4d9d9    1         1         1       2m5s
replicaset.apps/istio-ingressgateway-5464d78db9   1         1         1       2m5s
replicaset.apps/istiod-56dbb6c987                 1         1         1       2m35s
root@node1:~/101-master/module12/istio/installation# 
root@node1:~/101-master/module12/istio/installation# 

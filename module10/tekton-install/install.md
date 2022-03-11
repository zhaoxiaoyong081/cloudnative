# tekton-install


k apply -f install.yaml

kubectl apply -f  0.24.1.yaml

root@master:~/101-master/module10/tekton/tekton-install#  kubectl get crd | grep tekton
clustertasks.tekton.dev                               2022-03-10T08:51:52Z
conditions.tekton.dev                                 2022-03-10T08:51:52Z
extensions.dashboard.tekton.dev                       2022-03-10T08:53:56Z
pipelineresources.tekton.dev                          2022-03-10T08:51:52Z
pipelineruns.tekton.dev                               2022-03-10T08:51:52Z
pipelines.tekton.dev                                  2022-03-10T08:51:52Z
runs.tekton.dev                                       2022-03-10T08:51:52Z
taskruns.tekton.dev                                   2022-03-10T08:51:52Z
tasks.tekton.dev                                      2022-03-10T08:51:52Z
root@master:~/101-master/module10/tekton/tekton-install# kubectl get po -n tekton-pipelines 
NAME                                           READY   STATUS    RESTARTS       AGE
tekton-dashboard-5cc6fbd6f8-t8btd              1/1     Running   1 (124m ago)   153m
tekton-pipelines-controller-6669c955bf-kkcb7   1/1     Running   1 (124m ago)   155m
tekton-pipelines-webhook-7c8685d858-6r8qf      1/1     Running   1 (123m ago)   155m


kubectl apply -f test-task.yaml 

kubectl apply -f taskrun.yaml 

root@master:~/101-master/module10/tekton/tekton-install# kubectl get task
NAME    AGE
hello   80m
root@master:~/101-master/module10/tekton/tekton-install# kubectl get taskruns.tekton.dev 
NAME    SUCCEEDED   REASON      STARTTIME   COMPLETIONTIME
hello   True        Succeeded   79m         77m
root@master:~/101-master/module10/tekton/tekton-install# kubectl get po | grep hello
hello-pod-b26kr                      0/1     Completed     0              79m
root@master:~/101-master/module10/tekton/tekton-install# 


root@master:~/101-master/module10/tekton/tekton-install# 
root@master:~/101-master/module10/tekton/tekton-install# kubectl logs  hello-pod-b26kr
Hello World!
root@master:~/101-master/module10/tekton/tekton-install# tkn task list
NAME    DESCRIPTION   AGE
hello                 1 hour ago
root@master:~/101-master/module10/tekton/tekton-install# 



root@master:~/101-master/module10/tekton/tekton-install# kubectl get  all  -n  tekton-pipelines 
NAME                                               READY   STATUS    RESTARTS       AGE
pod/tekton-dashboard-5cc6fbd6f8-t8btd              1/1     Running   1 (127m ago)   156m
pod/tekton-pipelines-controller-6669c955bf-kkcb7   1/1     Running   1 (127m ago)   158m
pod/tekton-pipelines-webhook-7c8685d858-6r8qf      1/1     Running   1 (127m ago)   158m

NAME                                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)                              AGE
service/tekton-dashboard              NodePort    10.107.228.187   <none>        9097:30890/TCP                       156m
service/tekton-pipelines-controller   ClusterIP   10.104.4.225     <none>        9090/TCP,8080/TCP                    158m
service/tekton-pipelines-webhook      ClusterIP   10.109.55.122    <none>        9090/TCP,8008/TCP,443/TCP,8080/TCP   158m

NAME                                          READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/tekton-dashboard              1/1     1            1           156m
deployment.apps/tekton-pipelines-controller   1/1     1            1           158m
deployment.apps/tekton-pipelines-webhook      1/1     1            1           158m

NAME                                                     DESIRED   CURRENT   READY   AGE
replicaset.apps/tekton-dashboard-5cc6fbd6f8              1         1         1       156m
replicaset.apps/tekton-pipelines-controller-6669c955bf   1         1         1       158m
replicaset.apps/tekton-pipelines-webhook-7c8685d858      1         1         1       158m

NAME                                                           REFERENCE                             TARGETS          MINPODS   MAXPODS   REPLICAS   AGE
horizontalpodautoscaler.autoscaling/tekton-pipelines-webhook   Deployment/tekton-pipelines-webhook   <unknown>/100%   1         5         1          158m
root@master:~/101-master/module10/tekton/tekton-install# 


k apply  -f  task-hello.yaml
root@master:~/101-master/module10/tekton/tekton-install# k create  -f taskrun-hello.yaml 
taskrun.tekton.dev/hello-run-kfdrm created
root@master:~/101-master/module10/tekton/tekton-install# 
root@master:~/101-master/module10/tekton/tekton-install# k get  pod  
NAME                                 READY   STATUS            RESTARTS       AGE
hello-pod-b26kr                      0/1     Completed         0              92m
hello-run-kfdrm-pod-mgjh5            0/1     PodInitializing   0              5s
httpserver-deploy-866b5f4648-kxrpx   1/1     Running           1 (138m ago)   9h
httpserver-deploy-866b5f4648-p5v87   1/1     Running           8 (138m ago)   14d
jenkins-0                            1/1     Terminating       0              23h
nginx                                1/1     Running           9 (138m ago)   15d
nginx-anti-86b54c575c-qpwgt          1/1     Running           9 (138m ago)   23d
nginx-anti-86b54c575c-tn2sk          1/1     Running           9 (138m ago)   23d
nginx-deployment-6fbb949bf4-jsnnw    1/1     Running           9 (138m ago)   23d
root@master:~/101-master/module10/tekton/tekton-install# 

root@master:~/101-master/module10/tekton/tekton-install# k logs  hello-run-kfdrm-pod-mgjh5  
Hello jesse!
root@master:~/101-master/module10/tekton/tekton-install# 


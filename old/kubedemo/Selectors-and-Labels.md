To get all labels

kubectl get all --show-labels

kubectl get pods --show-labels

— how to label a pod, po is short for pod

kevinmason@Kevins-MBP kubedemo % kubectl label po/hellogopher-5459787866-pcpcj author=kevin
pod/hellogopher-5459787866-pcpcj labeled

To overwrite a service

kubectl label po/hellogopher-5459787866-pcpcj author=kevin_mason --overwrite
pod/hellogopher-5459787866-pcpcj labeled

— to unable a thing

kubectl label service/kubernetes author-  
service/kubernetes unlabeled

SELECTORS

kubectl get pods --selector app=hellogopher
NAME READY STATUS RESTARTS AGE
hellogopher-5459787866-pcpcj 1/1 Running 0 7m12s

// when we are searching labels we are referring to it as a selector

We can also test for different expressions such as != instead of =

// to use the selector to run commands

kubectl delete pods --selector app=hellogopher
pod "hellogopher-5459787866-pcpcj" deleted

or -l is short for --selector

kubectl delete pods --l app=hellogopher
pod "hellogopher-5459787866-pcpcj" deleted

# split-yaml

## Purpose
The purpose of this project is to be able to split a multi-yaml Kubernetes resource structure simply into individual files, splitting on either '---' or '- apiVersion:'. 

## Usage
```bash
Usage of ./split-yaml:
  -f string
        filename
  -o string
        outputDir
```

Is meant to be used either stand alone like:
```bash
split-yaml -f longfile.yaml -o files
```
to split longfile.yaml into multiple files, they will be named $(name)-$(kind).yaml and end up in the files directory (in this case).
Or it can be used when piping output into it like this:
```bash
kubectl get deploy,svc,configmap,ingress -o yaml | split-yaml -o files
```
To split the `kind: List` content into individual files.

Contributing PR's are welcome! If you want to expand it to use real kubernetes objects, I'm open to contributions.
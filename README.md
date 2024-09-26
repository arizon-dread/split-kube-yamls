# split-yaml

## Purpose
The purpose of this project is to be able to split a multi-yaml Kubernetes resource structure simply into individual files, splitting on either `---` or `- apiVersion:`.  
It can be used when templating larger kubernetes manifest data and then structuring the result into individual files. For instance if you are using kustomize with overlays in a CI/CD pipeline but the pipeline commits the result into a manifest git repo, and you want the result to be structured and orderly with each resource manifest in it's individual file.

## Usage
```bash
Usage of split-yaml:
  -f string
        filename
  -o string
        outputDir
```

Is meant to be used either stand alone like:
```bash
split-yaml -f longfile.yaml -o files
```
to split longfile.yaml into multiple files, they will be named `$(name)-$(kind).yaml` and end up in the `./files/` directory (in this case).

It can also be used when piping output into it like this:
```bash
kubectl get deploy,svc,configmap,ingress -o yaml | split-yaml -o files -f -
```
To split the `kind: List` content into individual files.

**You must have LF line endings for this to work. CRLF would need to be converted to LF using git line ending normalization**

## Ending notes
To make the tool versatile and not being subject to incompatibilities given the extendability of the kubernetes project, only the metadata part and general information about what kind each reasource is, is read using a proper struct. The spec is handled like a yaml string and should work with CRD's as well as standard kubernetes manifests, if it is proper yaml structure. If you find bugs when using the tool, please submit an issue so I can fix it.

Feel free to reach out with suggestions or issues!

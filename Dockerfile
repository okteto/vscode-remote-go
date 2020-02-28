FROM okteto/golang:1


RUN mkdir -p -m 755 /root/.vscode/extensions
RUN mkdir -p -m 755 /root/.vscode/data/Machine
COPY settings.json /root/.vscode-server/data/Machine/settings.json
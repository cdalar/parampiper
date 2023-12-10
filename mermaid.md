## Bash Script 
```mermaid
flowchart LR

X(prm get p_name)--> A(Bash Script) -->B(prm set p_name)
Y(prm out --export) --> A
%% B --> C{Decision}
%% C -->|One| D[Result 1]
%% C -->|Two| E[Result 2]
%% X(prm get p_name <a href='http://google.com'>link</a>)--> A(Bash Script) -->B(prm set p_name)
```

## Powershell Script 
```mermaid
flowchart LR

X(prm get p_name)--> A(Powershell Script) -->B(prm set p_name)
Y(prm out --export) --> A
```

## Terraform
```mermaid
flowchart LR
X(prm out --tfvars)--> A(Terraform) -->B(prm import -tfshowjson)

```

## Bicep

TBD
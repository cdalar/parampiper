## Bash Script 
```mermaid
flowchart LR

X(p8r get p_name)--> A(Bash Script) -->B(p8r set p_name)
Y(p8r out --export) --> A
%% B --> C{Decision}
%% C -->|One| D[Result 1]
%% C -->|Two| E[Result 2]
%% X(p8r get p_name <a href='http://google.com'>link</a>)--> A(Bash Script) -->B(p8r set p_name)
```

## Powershell Script 
```mermaid
flowchart LR

X(p8r get p_name)--> A(Powershell Script) -->B(p8r set p_name)
Y(p8r out --export) --> A
```

## Terraform
```mermaid
flowchart LR
X(p8r out --tfvars)--> A(Terraform) -->B(p8r import -tfshowjson)

```

## Bicep

TBD
# CRUD Actions

## Adding/Updating Parameters 
```
p8r set -n key1 -v value1
```
will add `key1` key with value `value1`.


## Delete Parameters
```
p8r rm -n key1
```
will delete the parameter named `key1`

## List Parameters
```
p8r ls 
NAME   TYPE    VALUE    ATTRIBUTES   INFO
key1   basic   value1   0
```

## Output 
Different ways to output parameters

### Environment Variables
```
p8r out -oexport > export.sh
cat export.sh
export KEY1="value1"
```

### Exporting as Terraform tfvars file 
```
p8r out -otfvars > parameters.auto.tfvars
```
by exporting it as *auto.tfvars ([tfvars files](https://developer.hashicorp.com/terraform/language/values/variables#variable-definitions-tfvars-files)) you can directly use it inside your terraform code. 


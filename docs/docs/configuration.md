## Configuration

Create the default configuration under `.p8r/parampiper.yaml`
```bash
$ p8r init
p8r environment initialized
```
Default configuration file is follows:
```yaml
local_file:
  FilePath: parampiper.json
azure_blob:
  StorageAccountName: stparampiper
  ContainerName: abc
  BlobName: parampiper.json
```

### Set which backend provider you like to use

- local_file
- azure_blob

Set Environment Variables `PP_DATA` to one of the above.
```
export PP_DATA=local_file
```

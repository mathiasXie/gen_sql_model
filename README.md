﻿# gen_sql_model

```shell
go get github.com/mathiasXie/gen_sql_model
go install github.com/mathiasXie/gen_sql_model
cd your_path/model
gen_sql_model  --ddlpath=user.sql --package=model > user.go  && gofmt -w user.go
```

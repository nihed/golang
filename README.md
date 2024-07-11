 1064  cd greetings
 1065  go mod init example.com/greetings
 1066  ls
 1067  vi greetings.go
 1068  cd ..
 1069  ls
 1070  mkdir hello 
 1071  cd hello
 1072  ls
 1073  go mod init example.com/hello
 1074  vi hello.go
 1075  go mod edit -replace example.com/greetings=../greetings
 1076  ls
 1077  cat go.mod
 1078  go mod tidyy 
 1079  go mod tidy
 1080  cat go.mod
 1081  go run .
 1082  git status
 1083  git add * 
 1084  git status
 1085  git add ../*
 1086  git status
 1087  git commit -a -m "test module part"
 1088  git status
 1089  git push
 1092  git status
 1096  ls

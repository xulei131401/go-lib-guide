# go-lib-guide
常用第三方库学习

- [JSON-to-Go](https://mholt.github.io/json-to-go/)
- [TOML-to-Go](https://xuri.me/toml-to-go/)
- [YAML-to-Go](https://zhwt.github.io/yaml-to-go/)
- [curl-to-Go](https://mholt.github.io/curl-to-go/)
- [sql2struct](https://github.com/idoubi/sql2struct)
- [sql2DSL](http://sql2dsl.atotoa.com/)

```text
#1.EVAL script numkeys key [key ...] arg [arg ...]
eval "return {KEYS[1],KEYS[2],ARGV[1],ARGV[2]}" 2 key1 key2 first second third
#2.SCRIPT LOAD script
#3.EVALSHA sha1 numkeys key [key ...] arg [arg ...]
#4.SCRIPT KILL

```
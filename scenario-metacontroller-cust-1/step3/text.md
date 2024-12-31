

```shell

sed "s/(0x03000000:NameValue)://g"  t3  \
|  sed "s/(0x01000000:Object)://g"   \
|  sed "s/(0x01001000:Array    )://g"

```{{exec}}

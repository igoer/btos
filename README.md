# btos
 Efficient conversion tool for string and []byte
###Download and install
```
go get github.com/igoer/btos
```
###Example bytes to string
```
bf := &bytes.Buffer{}
for i := 0; i < 10000000; i++ {
	bf.WriteString("A")
}
body := bf.Bytes()

s := time.Now().UnixNano()
btos.BytesToString(body)
fmt.Printf("bytes to string, btos.BytesToString([]byte) usetime: %v\n", time.Now().UnixNano()-s)

s = time.Now().UnixNano()
html := string(body)
fmt.Printf("bytes to string, string([]byte) usetime: %v\n", time.Now().UnixNano()-s)
```
output:
```
bytes to string, btos.BytesToString([]byte) usetime: 0
bytes to string, string([]byte) usetime: 3000200
```
###Example string to bytes
```
bf1 := &bytes.Buffer{}
for i := 0; i < 10000000; i++ {
	bf1.WriteByte(65)
}
body1 := bf1.String()

s = time.Now().UnixNano()
btos.StringToBytes(body1)
fmt.Printf("string to bytes, btos.StringToBytes(string) usetime: %v\n", time.Now().UnixNano()-s)

s = time.Now().UnixNano()
by := []byte(body1)
fmt.Printf("string to bytes, []byte(string) usetime: %v\n", time.Now().UnixNano()-s)
```
output:
```
string to bytes, btos.StringToBytes(string) usetime: 0
string to bytes, []byte(string) usetime: 7000400
```

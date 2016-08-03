# btos
 Efficient conversion tool for string and []byte
###Download and install
```
go get github.com/igoer/btos
```
###example
```
bf := &bytes.Buffer{}
for i := 0; i < 10000000; i++ {
	bf.WriteString("A")
}
body := bf.Bytes()

// bytes to string
s := time.Now().UnixNano()
btos.BytesToString(body)
fmt.Printf("bytes to string, btos.BytesToString([]byte) usetime: %v\n", time.Now().UnixNano()-s)

s = time.Now().UnixNano()
html := string(body)
fmt.Printf("bytes to string, string([]byte) usetime: %v\n", time.Now().UnixNano()-s)
```

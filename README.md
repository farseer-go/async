# async
支持异步并行执行结果，统一await拿到结果
`github.com/farseer-go/async`组件的异步并行执行实现

## 如何使用async包
#### 首先使用async.New()方法初始化
#### 然后async.Add()方法可以添置需要异步执行的方法，可同时放置多个方法
#### 使用ContinueWith()方法放置执行完所有异步方法（即Add中添加的方法）后，需要执行的回调方法，该方法不会阻塞

func a() {
    fmt.Printf("我是a方法")
}

func b() {
    fmt.Printf("我是b方法")
}

func c() {
    fmt.Printf("我是c方法")
}

func d() {
    fmt.Printf("我是回调d方法")
}

func e() {
    fmt.Printf("我是回调e方法")
}

func main() {
    async := async.New()
    async.Add(a, b).Add(c).ContinueWith(d, e)
}

#### 也可以使用Wait()方法放置执行完所有异步方法（即Add中添加的方法）后，需要执行的回调方法，该方法会阻塞等待执行完成后再继续往下执行
func main() {
async := async.New()
async.Add(a, b).Add(c).Wait()
d()
e()
}
#### 执行结果：我是c方法我是a方法我是b方法我是回调d方法我是回调e方法
#### 说明：多次执行的结果可以看出a,b,c这3个方法是异步执行，结果顺序随机呈现，d方法和e方法在a、b、c这3个方法执行结束后才会执行。

### 使用闭包方式执行：
func a(v string) {
fmt.Printf(v)
}


func b(v string) {
fmt.Printf(v)
}

func main() {
async := async.New()
async.Add(func() {a("我是a方法")}).ContinueWith(func(){b("我是回调b方法")})
}


### 获取错误
async := async.New()
async.Add(func() {a("我是a方法")}).ContinueWith(func(){b("我是回调b方法")})
当async.Err为nil时表示运行发生没有错误，当Err不为nil说明有报错